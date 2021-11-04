// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PozoClient is the client API for Pozo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PozoClient interface {
	// Envía el monto total acumulado
	Amount(ctx context.Context, in *AmountReq, opts ...grpc.CallOption) (*AmountResp, error)
}

type pozoClient struct {
	cc grpc.ClientConnInterface
}

func NewPozoClient(cc grpc.ClientConnInterface) PozoClient {
	return &pozoClient{cc}
}

func (c *pozoClient) Amount(ctx context.Context, in *AmountReq, opts ...grpc.CallOption) (*AmountResp, error) {
	out := new(AmountResp)
	err := c.cc.Invoke(ctx, "/grpc.Pozo/Amount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PozoServer is the server API for Pozo service.
// All implementations must embed UnimplementedPozoServer
// for forward compatibility
type PozoServer interface {
	// Envía el monto total acumulado
	Amount(context.Context, *AmountReq) (*AmountResp, error)
	mustEmbedUnimplementedPozoServer()
}

// UnimplementedPozoServer must be embedded to have forward compatible implementations.
type UnimplementedPozoServer struct {
}

func (UnimplementedPozoServer) Amount(context.Context, *AmountReq) (*AmountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Amount not implemented")
}
func (UnimplementedPozoServer) mustEmbedUnimplementedPozoServer() {}

// UnsafePozoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PozoServer will
// result in compilation errors.
type UnsafePozoServer interface {
	mustEmbedUnimplementedPozoServer()
}

func RegisterPozoServer(s grpc.ServiceRegistrar, srv PozoServer) {
	s.RegisterService(&Pozo_ServiceDesc, srv)
}

func _Pozo_Amount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AmountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PozoServer).Amount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Pozo/Amount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PozoServer).Amount(ctx, req.(*AmountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Pozo_ServiceDesc is the grpc.ServiceDesc for Pozo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pozo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Pozo",
	HandlerType: (*PozoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Amount",
			Handler:    _Pozo_Amount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pozo.proto",
}