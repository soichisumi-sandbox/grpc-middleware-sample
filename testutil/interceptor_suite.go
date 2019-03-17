// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package testing

import (
	"github.com/soichisumi-sandbox/grpc-middleware-sample/api-pb"
	"net"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// InterceptorTestSuite is a testify/Suite that starts a gRPC PingService server and a client.
type InterceptorTestSuite struct {
	suite.Suite

	TestService apipb.UserServiceServer
	ServerOpts  []grpc.ServerOption
	ClientOpts  []grpc.DialOption

	serverAddr     string
	ServerListener net.Listener
	Server         *grpc.Server
	clientConn     *grpc.ClientConn
	Client         apipb.UserServiceClient

	restartServerWithDelayedStart chan time.Duration
	serverRunning                 chan bool
}

func (s *InterceptorTestSuite) SetupSuite() {
	s.restartServerWithDelayedStart = make(chan time.Duration)
	s.serverRunning = make(chan bool)

	s.serverAddr = "127.0.0.1:0"

	go func() {
		for {
			var err error
			s.ServerListener, err = net.Listen("tcp", s.serverAddr)
			s.serverAddr = s.ServerListener.Addr().String()
			require.NoError(s.T(), err, "must be able to allocate a port for serverListener")
			// This is the point where we hook up the interceptor
			s.Server = grpc.NewServer(s.ServerOpts...)
			// Crete a service of the instantiator hasn't provided one.
			if s.TestService == nil {
				s.TestService = &TestUserService{T: s.T()}
			}
			apipb.RegisterUserServiceServer(s.Server, s.TestService)

			go func() {
				s.Server.Serve(s.ServerListener)
			}()
			if s.Client == nil {
				s.Client = s.NewClient(s.ClientOpts...)
			}

			s.serverRunning <- true

			d := <-s.restartServerWithDelayedStart
			s.Server.Stop()
			time.Sleep(d)
		}
	}()

	<-s.serverRunning
}

func (s *InterceptorTestSuite) RestartServer(delayedStart time.Duration) <-chan bool {
	s.restartServerWithDelayedStart <- delayedStart
	time.Sleep(10 * time.Millisecond)
	return s.serverRunning
}

func (s *InterceptorTestSuite) NewClient(dialOpts ...grpc.DialOption) apipb.UserServiceClient {
	newDialOpts := append(dialOpts, grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
	newDialOpts = append(newDialOpts, grpc.WithInsecure())

	clientConn, err := grpc.Dial(s.ServerAddr(), newDialOpts...)
	require.NoError(s.T(), err, "must not error on client Dial")
	return apipb.NewUserServiceClient(clientConn)
}

func (s *InterceptorTestSuite) ServerAddr() string {
	return s.serverAddr
}

func (s *InterceptorTestSuite) SimpleCtx() context.Context {
	ctx, _ := context.WithTimeout(context.TODO(), 2*time.Second)
	return ctx
}

func (s *InterceptorTestSuite) DeadlineCtx(deadline time.Time) context.Context {
	ctx, _ := context.WithDeadline(context.TODO(), deadline)
	return ctx
}

func (s *InterceptorTestSuite) TearDownSuite() {
	time.Sleep(10 * time.Millisecond)
	if s.ServerListener != nil {
		s.Server.GracefulStop()
		s.T().Logf("stopped grpc.Server at: %v", s.ServerAddr())
		s.ServerListener.Close()

	}
	if s.clientConn != nil {
		s.clientConn.Close()
	}
}
