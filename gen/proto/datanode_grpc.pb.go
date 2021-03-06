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

// DatanodeClient is the client API for Datanode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatanodeClient interface {
	// Registra las jugadas de un jugador
	GuardarJugada(ctx context.Context, in *GuardarJugadaReq, opts ...grpc.CallOption) (*GuardarJugadaResp, error)
	// Devuelve las jugadas de un jugador
	ObtenerJugadas(ctx context.Context, in *ObtenerJugadasReq, opts ...grpc.CallOption) (*ObtenerJugadasResp, error)
}

type datanodeClient struct {
	cc grpc.ClientConnInterface
}

func NewDatanodeClient(cc grpc.ClientConnInterface) DatanodeClient {
	return &datanodeClient{cc}
}

func (c *datanodeClient) GuardarJugada(ctx context.Context, in *GuardarJugadaReq, opts ...grpc.CallOption) (*GuardarJugadaResp, error) {
	out := new(GuardarJugadaResp)
	err := c.cc.Invoke(ctx, "/grpc.Datanode/GuardarJugada", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeClient) ObtenerJugadas(ctx context.Context, in *ObtenerJugadasReq, opts ...grpc.CallOption) (*ObtenerJugadasResp, error) {
	out := new(ObtenerJugadasResp)
	err := c.cc.Invoke(ctx, "/grpc.Datanode/ObtenerJugadas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatanodeServer is the server API for Datanode service.
// All implementations must embed UnimplementedDatanodeServer
// for forward compatibility
type DatanodeServer interface {
	// Registra las jugadas de un jugador
	GuardarJugada(context.Context, *GuardarJugadaReq) (*GuardarJugadaResp, error)
	// Devuelve las jugadas de un jugador
	ObtenerJugadas(context.Context, *ObtenerJugadasReq) (*ObtenerJugadasResp, error)
	mustEmbedUnimplementedDatanodeServer()
}

// UnimplementedDatanodeServer must be embedded to have forward compatible implementations.
type UnimplementedDatanodeServer struct {
}

func (UnimplementedDatanodeServer) GuardarJugada(context.Context, *GuardarJugadaReq) (*GuardarJugadaResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GuardarJugada not implemented")
}
func (UnimplementedDatanodeServer) ObtenerJugadas(context.Context, *ObtenerJugadasReq) (*ObtenerJugadasResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObtenerJugadas not implemented")
}
func (UnimplementedDatanodeServer) mustEmbedUnimplementedDatanodeServer() {}

// UnsafeDatanodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatanodeServer will
// result in compilation errors.
type UnsafeDatanodeServer interface {
	mustEmbedUnimplementedDatanodeServer()
}

func RegisterDatanodeServer(s grpc.ServiceRegistrar, srv DatanodeServer) {
	s.RegisterService(&Datanode_ServiceDesc, srv)
}

func _Datanode_GuardarJugada_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuardarJugadaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatanodeServer).GuardarJugada(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Datanode/GuardarJugada",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatanodeServer).GuardarJugada(ctx, req.(*GuardarJugadaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datanode_ObtenerJugadas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObtenerJugadasReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatanodeServer).ObtenerJugadas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Datanode/ObtenerJugadas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatanodeServer).ObtenerJugadas(ctx, req.(*ObtenerJugadasReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Datanode_ServiceDesc is the grpc.ServiceDesc for Datanode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Datanode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Datanode",
	HandlerType: (*DatanodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GuardarJugada",
			Handler:    _Datanode_GuardarJugada_Handler,
		},
		{
			MethodName: "ObtenerJugadas",
			Handler:    _Datanode_ObtenerJugadas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "datanode.proto",
}
