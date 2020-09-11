package svid

import (
	"context"
	"crypto/x509"
	"net/url"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/imkira/go-observer"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/spiffe/spire/pkg/agent/manager/cache"
	"github.com/spiffe/spire/pkg/agent/plugin/keymanager/memory"
	"github.com/spiffe/spire/pkg/common/telemetry"
	"github.com/spiffe/spire/proto/spire/api/node"
	"github.com/spiffe/spire/test/clock"
	"github.com/spiffe/spire/test/fakes/fakeagentcatalog"
	mock_client "github.com/spiffe/spire/test/mock/agent/client"
	"github.com/spiffe/spire/test/util"
	"github.com/stretchr/testify/suite"
	tomb "gopkg.in/tomb.v2"
)

func TestRotator(t *testing.T) {
	suite.Run(t, new(RotatorTestSuite))
}

type RotatorTestSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	client *mock_client.MockClient

	bundle observer.Property

	r *rotator

	mockClock *clock.Mock
}

func (s *RotatorTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())

	s.client = mock_client.NewMockClient(s.ctrl)

	b, err := util.LoadBundleFixture()
	s.Require().NoError(err)
	s.bundle = observer.NewProperty(b)

	cat := fakeagentcatalog.New()
	cat.SetKeyManager(fakeagentcatalog.KeyManager(memory.New()))

	s.mockClock = clock.NewMock(s.T())
	s.mockClock.Set(time.Now())
	log, _ := test.NewNullLogger()
	td := url.URL{
		Scheme: "spiffe",
		Host:   "example.org",
	}
	c := &RotatorConfig{
		Catalog:      cat,
		Log:          log,
		Metrics:      telemetry.Blackhole{},
		TrustDomain:  td,
		BundleStream: cache.NewBundleStream(s.bundle.Observe()),
		Clk:          s.mockClock,
	}
	s.r, _ = newRotator(c)
	s.r.client = s.client
}

func (s *RotatorTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *RotatorTestSuite) TestRun() {
	cert, key, err := util.LoadSVIDFixture()
	s.Require().NoError(err)
	s.expectSVIDRotation(cert)

	state := State{
		SVID: []*x509.Certificate{cert},
		Key:  key,
	}
	s.r.state = observer.NewProperty(state)

	stream := s.r.Subscribe()

	ctx, cancel := context.WithCancel(context.Background())
	t := new(tomb.Tomb)
	t.Go(func() error {
		return s.r.Run(ctx)
	})

	// We should have the latest state
	s.Assert().False(stream.HasNext())

	// Should be equal to the fixture
	state = stream.Value().(State)
	s.Require().Len(state.SVID, 1)
	s.Assert().Equal(cert, state.SVID[0])
	s.Assert().Equal(key, state.Key)

	cancel()
	s.Require().Equal(context.Canceled, t.Wait())
}

func (s *RotatorTestSuite) TestRunWithUpdates() {
	// Cert that's valid for 1hr
	temp, err := util.NewSVIDTemplate(s.mockClock, "spiffe://example.org/test")
	s.Require().NoError(err)
	goodCert, _, err := util.SelfSign(temp)
	s.Require().NoError(err)

	// Cert that's expiring
	temp.NotBefore = s.mockClock.Now().Add(-1 * time.Hour)
	temp.NotAfter = s.mockClock.Now()
	badCert, _, err := util.SelfSign(temp)
	s.Require().NoError(err)

	state := State{
		SVID: []*x509.Certificate{badCert},
	}
	s.r.state = observer.NewProperty(state)

	s.expectSVIDRotation(goodCert)

	stream := s.r.Subscribe()

	ctx, cancel := context.WithCancel(context.Background())
	t := new(tomb.Tomb)
	t.Go(func() error {
		return s.r.Run(ctx)
	})

	s.mockClock.Add(s.r.c.Interval)
	err = s.r.rotateSVID(context.Background())
	s.Require().NoError(err)

	select {
	case <-time.After(time.Second):
		s.T().Error("timed out while waiting for expected SVID rotation")
	case <-stream.Changes():
		state = stream.Next().(State)
		s.Require().Len(state.SVID, 1)
		s.Assert().Equal(goodCert, state.SVID[0])
	}

	cancel()
	s.Require().Equal(context.Canceled, t.Wait())
}

func (s *RotatorTestSuite) TestRotateSVID() {
	// Cert that's valid for 1hr
	temp, err := util.NewSVIDTemplate(s.mockClock, "spiffe://example.org/test")
	s.Require().NoError(err)
	goodCert, _, err := util.SelfSign(temp)
	s.Require().NoError(err)

	// Cert that's expiring
	temp.NotBefore = s.mockClock.Now().Add(-1 * time.Hour)
	temp.NotAfter = s.mockClock.Now()
	badCert, _, err := util.SelfSign(temp)
	s.Require().NoError(err)

	state := State{
		SVID: []*x509.Certificate{badCert},
	}
	s.r.state = observer.NewProperty(state)

	stream := s.r.Subscribe()
	s.expectSVIDRotation(goodCert)
	err = s.r.rotateSVID(context.Background())
	s.Assert().NoError(err)
	s.Require().True(stream.HasNext())

	state = stream.Next().(State)
	s.Require().Len(state.SVID, 1)
	s.Assert().True(goodCert.Equal(state.SVID[0]))
}

// expectSVIDRotation sets the appropriate expectations for an SVID rotation, and returns
// the the provided certificate to the client.Client caller.
func (s *RotatorTestSuite) expectSVIDRotation(cert *x509.Certificate) {
	s.client.EXPECT().
		RenewSVID(gomock.Any(), gomock.Any()).
		Return(&node.X509SVID{
			CertChain: cert.Raw,
		}, nil)
	s.client.EXPECT().Release().MaxTimes(2)
}
