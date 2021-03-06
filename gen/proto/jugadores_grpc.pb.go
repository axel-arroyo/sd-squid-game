// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// JugadorClient is the client API for Jugador service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JugadorClient interface {
}

type jugadorClient struct {
	cc grpc.ClientConnInterface
}

func NewJugadorClient(cc grpc.ClientConnInterface) JugadorClient {
	return &jugadorClient{cc}
}

// JugadorServer is the server API for Jugador service.
// All implementations must embed UnimplementedJugadorServer
// for forward compatibility
type JugadorServer interface {
	mustEmbedUnimplementedJugadorServer()
}

// UnimplementedJugadorServer must be embedded to have forward compatible implementations.
type UnimplementedJugadorServer struct {
}

func (UnimplementedJugadorServer) mustEmbedUnimplementedJugadorServer() {}

// UnsafeJugadorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JugadorServer will
// result in compilation errors.
type UnsafeJugadorServer interface {
	mustEmbedUnimplementedJugadorServer()
}

func RegisterJugadorServer(s grpc.ServiceRegistrar, srv JugadorServer) {
	s.RegisterService(&Jugador_ServiceDesc, srv)
}

// Jugador_ServiceDesc is the grpc.ServiceDesc for Jugador service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Jugador_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Jugador",
	HandlerType: (*JugadorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "jugadores.proto",
}
