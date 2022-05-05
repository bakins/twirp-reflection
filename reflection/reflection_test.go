package reflection_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/twitchtv/twirp/example"

	"github.com/bakins/twirp-reflection/reflection"
	twirp_reflection "github.com/bakins/twirp-reflection/v0"
)

func TestHandler(t *testing.T) {
	ts := example.NewHaberdasherServer(&nopHaberdasher{})

	s := reflection.NewServer()

	s.RegisterService(ts)
	s.RegisterService(s)
	mux := http.NewServeMux()

	mux.Handle(ts.PathPrefix(), ts)
	mux.Handle(s.PathPrefix(), s)

	svr := httptest.NewServer(mux)
	defer svr.Close()

	base := twirp_reflection.NewServerReflectionServiceProtobufClient(svr.URL, http.DefaultClient)
	client := reflection.NewClient(base)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	services, err := client.ListServices(ctx)
	require.NoError(t, err)
	require.Len(t, services, 2)
	require.Equal(
		t,
		[]string{
			"bakins.twirp.reflection.v0.ServerReflectionService",
			"twitch.twirp.example.Haberdasher",
		},
		services,
	)

	descriptor, err := client.GetServiceDescriptor(ctx, "twitch.twirp.example.Haberdasher")
	require.NoError(t, err)
	require.Equal(t, "twitch.twirp.example", descriptor.GetPackage())
	require.Len(t, descriptor.GetService(), 1)

	descriptors, err := client.GetSymbolDescriptor(ctx, "bakins.twirp.reflection.v0.ListServicesResponse")
	require.NoError(t, err)
	require.Len(t, descriptors, 1)
	require.Len(t, descriptors[0].GetService(), 1)
	require.Equal(t, "ServerReflectionService", descriptors[0].GetService()[0].GetName())
}

type nopHaberdasher struct{}

func (h *nopHaberdasher) MakeHat(ctx context.Context, size *example.Size) (*example.Hat, error) {
	return nil, nil
}
