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

// LiderClient is the client API for Lider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LiderClient interface {
	// Petición de unirse al juego
	UnirseAJuego(ctx context.Context, in *UnirseReq, opts ...grpc.CallOption) (Lider_UnirseAJuegoClient, error)
	// Recibe jugada de un jugador
	EnviarJugada(ctx context.Context, in *EnviarJugadaReq, opts ...grpc.CallOption) (*EnviarJugadaResp, error)
	// Recibe el estado de los jugadores después de que haya finalizado una etapa
	EnviarEstado(ctx context.Context, in *EnviarEstadoReq, opts ...grpc.CallOption) (*EnviarEstadoResp, error)
	//	Recibe la petición de ver el pozo por parte del jugador
	PedirPozo(ctx context.Context, in *PedirPozoReq, opts ...grpc.CallOption) (*PedirPozoResp, error)
}

type liderClient struct {
	cc grpc.ClientConnInterface
}

func NewLiderClient(cc grpc.ClientConnInterface) LiderClient {
	return &liderClient{cc}
}

func (c *liderClient) UnirseAJuego(ctx context.Context, in *UnirseReq, opts ...grpc.CallOption) (Lider_UnirseAJuegoClient, error) {
	stream, err := c.cc.NewStream(ctx, &Lider_ServiceDesc.Streams[0], "/grpc.Lider/UnirseAJuego", opts...)
	if err != nil {
		return nil, err
	}
	x := &liderUnirseAJuegoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Lider_UnirseAJuegoClient interface {
	Recv() (*UnirseResp, error)
	grpc.ClientStream
}

type liderUnirseAJuegoClient struct {
	grpc.ClientStream
}

func (x *liderUnirseAJuegoClient) Recv() (*UnirseResp, error) {
	m := new(UnirseResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *liderClient) EnviarJugada(ctx context.Context, in *EnviarJugadaReq, opts ...grpc.CallOption) (*EnviarJugadaResp, error) {
	out := new(EnviarJugadaResp)
	err := c.cc.Invoke(ctx, "/grpc.Lider/EnviarJugada", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liderClient) EnviarEstado(ctx context.Context, in *EnviarEstadoReq, opts ...grpc.CallOption) (*EnviarEstadoResp, error) {
	out := new(EnviarEstadoResp)
	err := c.cc.Invoke(ctx, "/grpc.Lider/EnviarEstado", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liderClient) PedirPozo(ctx context.Context, in *PedirPozoReq, opts ...grpc.CallOption) (*PedirPozoResp, error) {
	out := new(PedirPozoResp)
	err := c.cc.Invoke(ctx, "/grpc.Lider/PedirPozo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiderServer is the server API for Lider service.
// All implementations must embed UnimplementedLiderServer
// for forward compatibility
type LiderServer interface {
	// Petición de unirse al juego
	UnirseAJuego(*UnirseReq, Lider_UnirseAJuegoServer) error
	// Recibe jugada de un jugador
	EnviarJugada(context.Context, *EnviarJugadaReq) (*EnviarJugadaResp, error)
	// Recibe el estado de los jugadores después de que haya finalizado una etapa
	EnviarEstado(context.Context, *EnviarEstadoReq) (*EnviarEstadoResp, error)
	//	Recibe la petición de ver el pozo por parte del jugador
	PedirPozo(context.Context, *PedirPozoReq) (*PedirPozoResp, error)
	mustEmbedUnimplementedLiderServer()
}

// UnimplementedLiderServer must be embedded to have forward compatible implementations.
type UnimplementedLiderServer struct {
}

func (UnimplementedLiderServer) UnirseAJuego(*UnirseReq, Lider_UnirseAJuegoServer) error {
	return status.Errorf(codes.Unimplemented, "method UnirseAJuego not implemented")
}
func (UnimplementedLiderServer) EnviarJugada(context.Context, *EnviarJugadaReq) (*EnviarJugadaResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarJugada not implemented")
}
func (UnimplementedLiderServer) EnviarEstado(context.Context, *EnviarEstadoReq) (*EnviarEstadoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarEstado not implemented")
}
func (UnimplementedLiderServer) PedirPozo(context.Context, *PedirPozoReq) (*PedirPozoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PedirPozo not implemented")
}
func (UnimplementedLiderServer) mustEmbedUnimplementedLiderServer() {}

// UnsafeLiderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiderServer will
// result in compilation errors.
type UnsafeLiderServer interface {
	mustEmbedUnimplementedLiderServer()
}

func RegisterLiderServer(s grpc.ServiceRegistrar, srv LiderServer) {
	s.RegisterService(&Lider_ServiceDesc, srv)
}

func _Lider_UnirseAJuego_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UnirseReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LiderServer).UnirseAJuego(m, &liderUnirseAJuegoServer{stream})
}

type Lider_UnirseAJuegoServer interface {
	Send(*UnirseResp) error
	grpc.ServerStream
}

type liderUnirseAJuegoServer struct {
	grpc.ServerStream
}

func (x *liderUnirseAJuegoServer) Send(m *UnirseResp) error {
	return x.ServerStream.SendMsg(m)
}

func _Lider_EnviarJugada_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnviarJugadaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServer).EnviarJugada(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Lider/EnviarJugada",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServer).EnviarJugada(ctx, req.(*EnviarJugadaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lider_EnviarEstado_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnviarEstadoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServer).EnviarEstado(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Lider/EnviarEstado",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServer).EnviarEstado(ctx, req.(*EnviarEstadoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lider_PedirPozo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PedirPozoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServer).PedirPozo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Lider/PedirPozo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServer).PedirPozo(ctx, req.(*PedirPozoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Lider_ServiceDesc is the grpc.ServiceDesc for Lider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Lider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Lider",
	HandlerType: (*LiderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnviarJugada",
			Handler:    _Lider_EnviarJugada_Handler,
		},
		{
			MethodName: "EnviarEstado",
			Handler:    _Lider_EnviarEstado_Handler,
		},
		{
			MethodName: "PedirPozo",
			Handler:    _Lider_PedirPozo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UnirseAJuego",
			Handler:       _Lider_UnirseAJuego_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "lider.proto",
}
