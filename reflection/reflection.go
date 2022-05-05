package reflection

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"sort"
	"strings"
	"sync"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/twitchtv/twirp"

	reflection "github.com/bakins/twirp-reflection/v0"
)

type Client struct {
	client reflection.ServerReflectionService
}

func NewClient(base reflection.ServerReflectionService) *Client {
	h := Client{
		client: base,
	}

	return &h
}

func (c *Client) ListServices(ctx context.Context) ([]string, error) {
	var req reflection.ListServicesRequest

	resp, err := c.client.ListServices(ctx, &req)
	if err != nil {
		return nil, err
	}

	out := make([]string, 0, len(resp.Services))

	for _, s := range resp.Services {
		out = append(out, s.Name)
	}

	sort.Strings(out)

	return out, nil
}

func (c *Client) GetServiceDescriptor(ctx context.Context, name string) (*descriptorpb.FileDescriptorProto, error) {
	req := reflection.GetServiceDescriptorRequest{
		Name: name,
	}

	resp, err := c.client.GetServiceDescriptor(ctx, &req)
	if err != nil {
		return nil, err
	}

	var descriptor descriptorpb.FileDescriptorProto
	if err := proto.Unmarshal(resp.FileDescriptor, &descriptor); err != nil {
		return nil, err
	}

	return &descriptor, nil
}

func (c *Client) GetSymbolDescriptor(ctx context.Context, symbol string) ([]*descriptorpb.FileDescriptorProto, error) {
	req := reflection.GetSymbolDescriptorRequest{
		Symbol: symbol,
	}

	resp, err := c.client.GetSymbolDescriptor(ctx, &req)
	if err != nil {
		return nil, err
	}

	out := make([]*descriptorpb.FileDescriptorProto, 0, len(resp.FileDescriptors))

	for _, d := range resp.FileDescriptors {
		var descriptor descriptorpb.FileDescriptorProto
		if err := proto.Unmarshal(d, &descriptor); err != nil {
			return nil, err
		}

		out = append(out, &descriptor)
	}

	return out, nil
}

type Server struct {
	lock               sync.Mutex
	services           map[string]TwirpServer
	descriptorResolver protodesc.Resolver
	reflection.TwirpServer
}

var _ reflection.ServerReflectionService = &Server{}

type ServerOption interface {
	apply(*Server)
}

func NewServer(options ...ServerOption) *Server {
	s := Server{
		services:           make(map[string]TwirpServer),
		descriptorResolver: protoregistry.GlobalFiles,
	}

	s.TwirpServer = reflection.NewServerReflectionServiceServer(&s)

	return &s
}

func (s *Server) RegisterService(server TwirpServer) {
	name := parseTwirpServicePath(server.PathPrefix())

	s.lock.Lock()
	defer s.lock.Unlock()

	s.services[name] = server
}

// parseTwirpServicePath extracts path components from PathPrefix()
// Expected format: "/[<prefix>]/<package>.<Service>/"
func parseTwirpServicePath(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		// assume entire thing is package
		return path
	}

	return parts[len(parts)-2]
}

func (s *Server) ListServices(_ context.Context, _ *reflection.ListServicesRequest) (*reflection.ListServicesResponse, error) {
	var resp reflection.ListServicesResponse

	s.lock.Lock()
	defer s.lock.Unlock()

	for name := range s.services {
		r := reflection.ListServiceResponse{
			Name: name,
		}

		resp.Services = append(resp.Services, &r)
	}

	return &resp, nil
}

func (s *Server) GetServiceDescriptor(_ context.Context, req *reflection.GetServiceDescriptorRequest) (*reflection.GetServiceDescriptorResponse, error) {
	svc := s.getServer(req.Name)
	if svc == nil {
		return nil, twirp.NotFound.Errorf("service %q not found", req.Name)
	}

	descriptor, _ := svc.ServiceDescriptor()

	data, err := gunzip(descriptor)
	if err != nil {
		return nil, twirp.WrapError(twirp.InternalError("failed to uncompress descriptor"), err)
	}

	resp := reflection.GetServiceDescriptorResponse{
		FileDescriptor: data,
	}

	return &resp, nil
}

func (s *Server) getServer(name string) TwirpServer {
	s.lock.Lock()
	defer s.lock.Unlock()

	svc, ok := s.services[name]
	if !ok {
		return nil
	}

	return svc
}

func gunzip(data []byte) ([]byte, error) {
	rdr, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	if _, err := io.Copy(&out, rdr); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func (s *Server) GetSymbolDescriptor(_ context.Context, req *reflection.GetSymbolDescriptorRequest) (*reflection.GetSymbolDescriptorResponse, error) {
	sentFileDescriptors := map[string]bool{}
	descriptors, err := s.fileDescEncodingContainingSymbol(req.Symbol, sentFileDescriptors)
	if err != nil {
		return nil, twirp.WrapError(
			twirp.NotFound.Errorf("symbol %q not found", req.Symbol),
			err,
		)
	}

	resp := reflection.GetSymbolDescriptorResponse{
		FileDescriptors: descriptors,
	}

	return &resp, nil
}

func (s *Server) fileDescEncodingContainingSymbol(name string, sentFileDescriptors map[string]bool) ([][]byte, error) {
	d, err := s.descriptorResolver.FindDescriptorByName(protoreflect.FullName(name))
	if err != nil {
		return nil, err
	}

	return s.fileDescWithDependencies(d.ParentFile(), sentFileDescriptors)
}

func (s *Server) fileDescWithDependencies(fd protoreflect.FileDescriptor, sentFileDescriptors map[string]bool) ([][]byte, error) {
	var r [][]byte
	queue := []protoreflect.FileDescriptor{fd}
	for len(queue) > 0 {
		currentfd := queue[0]
		queue = queue[1:]
		if sent := sentFileDescriptors[currentfd.Path()]; len(r) == 0 || !sent {
			sentFileDescriptors[currentfd.Path()] = true
			fdProto := protodesc.ToFileDescriptorProto(currentfd)
			currentfdEncoded, err := proto.Marshal(fdProto)
			if err != nil {
				return nil, err
			}
			r = append(r, currentfdEncoded)
		}
		for i := 0; i < currentfd.Imports().Len(); i++ {
			queue = append(queue, currentfd.Imports().Get(i))
		}
	}
	return r, nil
}

// TwirpServer is the interface generated server structs will support.
type TwirpServer interface {
	// ServiceDescriptor returns gzipped bytes describing the .proto file that
	// this service was generated from. Once unzipped, the bytes can be
	// unmarshalled as a
	// google.golang.org/protobuf/types/descriptorpb.FileDescriptorProto.
	//
	// The returned integer is the index of this particular service within that
	// FileDescriptorProto's 'Service' slice of ServiceDescriptorProtos. This is a
	// low-level field, expected to be used for reflection.
	ServiceDescriptor() ([]byte, int)

	// PathPrefix returns the HTTP URL path prefix for all methods handled by this
	// service. This can be used with an HTTP mux to route Twirp requests.
	// The path prefix is in the form: "/<prefix>/<package>.<Service>/"
	// that is, everything in a Twirp route except for the <Method> at the end.
	PathPrefix() string
}
