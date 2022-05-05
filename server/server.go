package server

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/twitchtv/twirp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	reflection "github.com/bakins/twirp-reflection"
)

type Handler struct {
	lock               sync.Mutex
	services           map[string]TwirpServer
	descriptorResolver protodesc.Resolver
}

var _ reflection.ServerReflectionService = &Handler{}

type Option interface {
	apply(*Handler)
}

func New(options ...Option) *Handler {
	h := Handler{
		services:           make(map[string]TwirpServer),
		descriptorResolver: protoregistry.GlobalFiles,
	}

	return &h
}

func (h *Handler) Register(server TwirpServer) {
	name := parseTwirpServicePath(server.PathPrefix())

	h.lock.Lock()
	defer h.lock.Unlock()

	h.services[name] = server
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

func (h *Handler) ListServices(_ context.Context, _ *reflection.ListServicesRequest) (*reflection.ListServicesResponse, error) {
	var resp reflection.ListServicesResponse

	h.lock.Lock()
	defer h.lock.Unlock()

	for name := range h.services {
		r := reflection.ListServiceResponse{
			Name: name,
		}

		resp.Services = append(resp.Services, &r)
	}

	return &resp, nil
}

func (h *Handler) GetServiceDescriptor(_ context.Context, req *reflection.GetServiceDescriptorRequest) (*reflection.GetServiceDescriptorResponse, error) {
	fmt.Println("GetServiceDescriptor", req.Name)

	svc := h.getServer(req.Name)
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

func (h *Handler) getServer(name string) TwirpServer {
	h.lock.Lock()
	defer h.lock.Unlock()

	svc, ok := h.services[name]
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

func (h *Handler) GetSymbolDescriptor(_ context.Context, req *reflection.GetSymbolDescriptorRequest) (*reflection.GetSymbolDescriptorResponse, error) {
	// sentFileDescriptors := map[string]bool{}
	fmt.Println("I am here")

	return &reflection.GetSymbolDescriptorResponse{}, nil
	/*
		descriptors, err := h.fileDescEncodingContainingSymbol(req.Symbol, sentFileDescriptors)
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
	*/
}

func (h *Handler) fileDescEncodingContainingSymbol(name string, sentFileDescriptors map[string]bool) ([][]byte, error) {
	fmt.Println(name, protoreflect.FullName(name))

	d, err := h.descriptorResolver.FindDescriptorByName(protoreflect.FullName(name))
	if err != nil {
		return nil, err
	}

	fmt.Println(name, d.ParentFile())

	return h.fileDescWithDependencies(d.ParentFile(), sentFileDescriptors)
}

func (h *Handler) fileDescWithDependencies(fd protoreflect.FileDescriptor, sentFileDescriptors map[string]bool) ([][]byte, error) {
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
