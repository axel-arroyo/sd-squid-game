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

// DataNodeClient is the client API for DataNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataNodeClient interface {
	// Registra las jugadas de un jugador
	RegistrarJugadaJug(ctx context.Context, in *RegistrarJugadaJugReq, opts ...grpc.CallOption) (*RegistrarJugadaJugResp, error)
}

type dataNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewDataNodeClient(cc grpc.ClientConnInterface) DataNodeClient {
	return &dataNodeClient{cc}
}

func (c *dataNodeClient) RegistrarJugadaJug(ctx context.Context, in *RegistrarJugadaJugReq, opts ...grpc.CallOption) (*RegistrarJugadaJugResp, error) {
	out := new(RegistrarJugadaJugResp)
	err := c.cc.Invoke(ctx, "/grpc.DataNode/RegistrarJugadaJug", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataNodeServer is the server API for DataNode service.
// All implementations must embed UnimplementedDataNodeServer
// for forward compatibility
type DataNodeServer interface {
	// Registra las jugadas de un jugador
	RegistrarJugadaJug(context.Context, *RegistrarJugadaJugReq) (*RegistrarJugadaJugResp, error)
	mustEmbedUnimplementedDataNodeServer()
}

// UnimplementedDataNodeServer must be embedded to have forward compatible implementations.
type UnimplementedDataNodeServer struct {
}

func (UnimplementedDataNodeServer) RegistrarJugadaJug(context.Context, *RegistrarJugadaJugReq) (*RegistrarJugadaJugResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegistrarJugadaJug not implemented")
}
func (UnimplementedDataNodeServer) mustEmbedUnimplementedDataNodeServer() {}

// UnsafeDataNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataNodeServer will
// result in compilation errors.
type UnsafeDataNodeServer interface {
	mustEmbedUnimplementedDataNodeServer()
}

func RegisterDataNodeServer(s grpc.ServiceRegistrar, srv DataNodeServer) {
	s.RegisterService(&DataNode_ServiceDesc, srv)
}

func _DataNode_RegistrarJugadaJug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrarJugadaJugReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServer).RegistrarJugadaJug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.DataNode/RegistrarJugadaJug",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServer).RegistrarJugadaJug(ctx, req.(*RegistrarJugadaJugReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DataNode_ServiceDesc is the grpc.ServiceDesc for DataNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.DataNode",
	HandlerType: (*DataNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegistrarJugadaJug",
			Handler:    _DataNode_RegistrarJugadaJug_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "datanode.proto",
}
