// +build unit

package grpc

import (
	"context"
	"net"
	"testing"
	"time"

	traefikstatic "github.com/containous/traefik/v2/pkg/config/static"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockserver "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/grpc/server/mock"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/metrics/generic"
	"google.golang.org/grpc"
)

func TestEntryPoint(t *testing.T) {
	ctrlr := gomock.NewController(t)
	defer ctrlr.Finish()

	cfg := &traefikstatic.EntryPoint{
		Address: "127.0.0.1:0",
		Transport: &traefikstatic.EntryPointsTransport{
			RespondingTimeouts: &traefikstatic.RespondingTimeouts{},
			LifeCycle:          &traefikstatic.LifeCycle{},
		},
	}

	builder := mockserver.NewMockBuilder(ctrlr)
	ep := NewEntryPoint("", cfg, builder, generic.NewTCP())

	builder.EXPECT().Build(gomock.Any(), gomock.Any(), gomock.Any()).Return(grpc.NewServer(), nil)
	_ = ep.BuildServer(context.Background(), nil)

	done := make(chan struct{})
	go func() {
		_ = ep.ListenAndServe(context.Background())
		close(done)
	}()

	// Wait a few millisecond for server to start
	time.Sleep(500 * time.Millisecond)

	_, err := net.Dial("tcp", ep.Addr())
	require.NoError(t, err, "Dial should not error")

	_ = ep.Shutdown(context.Background())
	<-done
	_ = ep.Close()
}
