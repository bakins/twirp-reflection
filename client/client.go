package client

import (
	"context"
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	reflection "github.com/bakins/twirp-reflection"
)

type Handler struct {
	client reflection.ServerReflectionService
}

func New(base reflection.ServerReflectionService) *Handler {
	h := Handler{
		client: base,
	}

	return &h
}

func (h *Handler) ListServices(ctx context.Context) ([]string, error) {
	var req reflection.ListServicesRequest

	resp, err := h.client.ListServices(ctx, &req)
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

func (h *Handler) GetServiceDescriptor(ctx context.Context, name string) (*descriptorpb.FileDescriptorProto, error) {
	req := reflection.GetServiceDescriptorRequest{
		Name: name,
	}

	fmt.Println(&req)

	resp, err := h.client.GetServiceDescriptor(ctx, &req)
	if err != nil {
		return nil, err
	}

	var descriptor descriptorpb.FileDescriptorProto
	if err := proto.Unmarshal(resp.FileDescriptor, &descriptor); err != nil {
		return nil, err
	}

	return &descriptor, nil
}

func (h *Handler) GetSymbolDescriptor(ctx context.Context, symbol string) ([]*descriptorpb.FileDescriptorProto, error) {
	req := reflection.GetSymbolDescriptorRequest{
		Symbol: symbol,
	}

	resp, err := h.client.GetSymbolDescriptor(ctx, &req)
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
