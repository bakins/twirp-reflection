package server_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/twitchtv/twirp/example"

	reflection "github.com/bakins/twirp-reflection"
	"github.com/bakins/twirp-reflection/client"
	"github.com/bakins/twirp-reflection/server"
)

func TestHandler(t *testing.T) {
	ts := example.NewHaberdasherServer(&nopHaberdasher{})

	h := server.New()

	h.Register(ts)

	rfl := reflection.NewServerReflectionServiceServer(h)

	h.Register(rfl)

	mux := http.NewServeMux()

	mux.Handle(ts.PathPrefix(), ts)
	mux.Handle(rfl.PathPrefix(), rfl)

	svr := httptest.NewServer(mux)
	defer svr.Close()

	base := reflection.NewServerReflectionServiceProtobufClient(svr.URL, http.DefaultClient)
	client := client.New(base)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	services, err := client.ListServices(ctx)
	require.NoError(t, err)
	require.Len(t, services, 2)
	require.Equal(
		t,
		[]string{
			"bakins.twirp.reflection.v1.ServerReflectionService",
			"twitch.twirp.example.Haberdasher",
		},
		services,
	)

	descriptor, err := client.GetServiceDescriptor(ctx, "twitch.twirp.example.Haberdasher")
	require.NoError(t, err)
	require.Equal(t, "twitch.twirp.example", descriptor.GetPackage())
	require.Len(t, descriptor.GetService(), 1)
}

type nopHaberdasher struct{}

func (h *nopHaberdasher) MakeHat(ctx context.Context, size *example.Size) (*example.Hat, error) {
	return nil, nil
}
