// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: lider.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PedirPozoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PedirPozoReq) Reset() {
	*x = PedirPozoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PedirPozoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PedirPozoReq) ProtoMessage() {}

func (x *PedirPozoReq) ProtoReflect() protoreflect.Message {
	mi := &file_lider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PedirPozoReq.ProtoReflect.Descriptor instead.
func (*PedirPozoReq) Descriptor() ([]byte, []int) {
	return file_lider_proto_rawDescGZIP(), []int{0}
}

type PedirPozoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valor int32 `protobuf:"varint,1,opt,name=valor,proto3" json:"valor,omitempty"`
}

func (x *PedirPozoResp) Reset() {
	*x = PedirPozoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PedirPozoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PedirPozoResp) ProtoMessage() {}

func (x *PedirPozoResp) ProtoReflect() protoreflect.Message {
	mi := &file_lider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PedirPozoResp.ProtoReflect.Descriptor instead.
func (*PedirPozoResp) Descriptor() ([]byte, []int) {
	return file_lider_proto_rawDescGZIP(), []int{1}
}

func (x *PedirPozoResp) GetValor() int32 {
	if x != nil {
		return x.Valor
	}
	return 0
}

type EnviarEstadoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumJugador int32 `protobuf:"varint,1,opt,name=numJugador,proto3" json:"numJugador,omitempty"`
	Estado     bool  `protobuf:"varint,2,opt,name=estado,proto3" json:"estado,omitempty"`
	Etapa      int32 `protobuf:"varint,3,opt,name=etapa,proto3" json:"etapa,omitempty"`
}

func (x *EnviarEstadoReq) Reset() {
	*x = EnviarEstadoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnviarEstadoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnviarEstadoReq) ProtoMessage() {}

func (x *EnviarEstadoReq) ProtoReflect() protoreflect.Message {
	mi := &file_lider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnviarEstadoReq.ProtoReflect.Descriptor instead.
func (*EnviarEstadoReq) Descriptor() ([]byte, []int) {
	return file_lider_proto_rawDescGZIP(), []int{2}
}

func (x *EnviarEstadoReq) GetNumJugador() int32 {
	if x != nil {
		return x.NumJugador
	}
	return 0
}

func (x *EnviarEstadoReq) GetEstado() bool {
	if x != nil {
		return x.Estado
	}
	return false
}

func (x *EnviarEstadoReq) GetEtapa() int32 {
	if x != nil {
		return x.Etapa
	}
	return 0
}

type EnviarEstadoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *EnviarEstadoResp) Reset() {
	*x = EnviarEstadoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnviarEstadoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnviarEstadoResp) ProtoMessage() {}

func (x *EnviarEstadoResp) ProtoReflect() protoreflect.Message {
	mi := &file_lider_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnviarEstadoResp.ProtoReflect.Descriptor instead.
func (*EnviarEstadoResp) Descriptor() ([]byte, []int) {
	return file_lider_proto_rawDescGZIP(), []int{3}
}

func (x *EnviarEstadoResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UnirseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumJugador int32  `protobuf:"varint,1,opt,name=numJugador,proto3" json:"numJugador,omitempty"`
	Msg        string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *UnirseReq) Reset() {
	*x = UnirseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lider_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnirseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnirseReq) ProtoMessage() {}

func (x *UnirseReq) ProtoReflect() protoreflect.Message {
	mi := &file_lider_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnirseReq.ProtoReflect.Descriptor instead.
func (*UnirseReq) Descriptor() ([]byte, []int) {
	return file_lider_proto_rawDescGZIP(), []int{4}
}

func (x *UnirseReq) GetNumJugador() int32 {
	if x != nil {
		return x.NumJugador
	}
	return 0
}

func (x *UnirseReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UnirseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *UnirseResp) Reset() {
	*x = UnirseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lider_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnirseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnirseResp) ProtoMessage() {}

func (x *UnirseResp) ProtoReflect() protoreflect.Message {
	mi := &file_lider_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnirseResp.ProtoReflect.Descriptor instead.
func (*UnirseResp) Descriptor() ([]byte, []int) {
	return file_lider_proto_rawDescGZIP(), []int{5}
}

func (x *UnirseResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type EnviarJugadaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumJugador int32 `protobuf:"varint,1,opt,name=numJugador,proto3" json:"numJugador,omitempty"`
	Etapa      int32 `protobuf:"varint,2,opt,name=etapa,proto3" json:"etapa,omitempty"`
	Ronda      int32 `protobuf:"varint,3,opt,name=ronda,proto3" json:"ronda,omitempty"`
	Jugada     int32 `protobuf:"varint,4,opt,name=jugada,proto3" json:"jugada,omitempty"`
}

func (x *EnviarJugadaReq) Reset() {
	*x = EnviarJugadaReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lider_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnviarJugadaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnviarJugadaReq) ProtoMessage() {}

func (x *EnviarJugadaReq) ProtoReflect() protoreflect.Message {
	mi := &file_lider_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnviarJugadaReq.ProtoReflect.Descriptor instead.
func (*EnviarJugadaReq) Descriptor() ([]byte, []int) {
	return file_lider_proto_rawDescGZIP(), []int{6}
}

func (x *EnviarJugadaReq) GetNumJugador() int32 {
	if x != nil {
		return x.NumJugador
	}
	return 0
}

func (x *EnviarJugadaReq) GetEtapa() int32 {
	if x != nil {
		return x.Etapa
	}
	return 0
}

func (x *EnviarJugadaReq) GetRonda() int32 {
	if x != nil {
		return x.Ronda
	}
	return 0
}

func (x *EnviarJugadaReq) GetJugada() int32 {
	if x != nil {
		return x.Jugada
	}
	return 0
}

type EnviarJugadaResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg       string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Eliminado bool   `protobuf:"varint,2,opt,name=eliminado,proto3" json:"eliminado,omitempty"`
}

func (x *EnviarJugadaResp) Reset() {
	*x = EnviarJugadaResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lider_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnviarJugadaResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnviarJugadaResp) ProtoMessage() {}

func (x *EnviarJugadaResp) ProtoReflect() protoreflect.Message {
	mi := &file_lider_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnviarJugadaResp.ProtoReflect.Descriptor instead.
func (*EnviarJugadaResp) Descriptor() ([]byte, []int) {
	return file_lider_proto_rawDescGZIP(), []int{7}
}

func (x *EnviarJugadaResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *EnviarJugadaResp) GetEliminado() bool {
	if x != nil {
		return x.Eliminado
	}
	return false
}

var File_lider_proto protoreflect.FileDescriptor

var file_lider_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67,
	0x72, 0x70, 0x63, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x65, 0x64, 0x69, 0x72, 0x50, 0x6f, 0x7a, 0x6f,
	0x52, 0x65, 0x71, 0x22, 0x25, 0x0a, 0x0d, 0x50, 0x65, 0x64, 0x69, 0x72, 0x50, 0x6f, 0x7a, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x22, 0x5f, 0x0a, 0x0f, 0x45, 0x6e,
	0x76, 0x69, 0x61, 0x72, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a,
	0x0a, 0x6e, 0x75, 0x6d, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65,
	0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x74, 0x61, 0x70, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x74, 0x61, 0x70, 0x61, 0x22, 0x24, 0x0a, 0x10, 0x45,
	0x6e, 0x76, 0x69, 0x61, 0x72, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x3d, 0x0a, 0x09, 0x55, 0x6e, 0x69, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1e,
	0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x22, 0x1e, 0x0a, 0x0a, 0x55, 0x6e, 0x69, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x22, 0x75, 0x0a, 0x0f, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61,
	0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x4a, 0x75, 0x67, 0x61,
	0x64, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x74, 0x61, 0x70, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x65, 0x74, 0x61, 0x70, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6e,
	0x64, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x6f, 0x6e, 0x64, 0x61, 0x12,
	0x16, 0x0a, 0x06, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x22, 0x42, 0x0a, 0x10, 0x45, 0x6e, 0x76, 0x69, 0x61,
	0x72, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1c, 0x0a,
	0x09, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x6e, 0x61, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x6e, 0x61, 0x64, 0x6f, 0x32, 0xf0, 0x01, 0x0a, 0x05,
	0x4c, 0x69, 0x64, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x0c, 0x55, 0x6e, 0x69, 0x72, 0x73, 0x65, 0x41,
	0x4a, 0x75, 0x65, 0x67, 0x6f, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x6e, 0x69,
	0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x6e,
	0x69, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x30, 0x01, 0x12, 0x3d, 0x0a, 0x0c, 0x45, 0x6e,
	0x76, 0x69, 0x61, 0x72, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x52, 0x65,
	0x71, 0x1a, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x4a,
	0x75, 0x67, 0x61, 0x64, 0x61, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3d, 0x0a, 0x0c, 0x45, 0x6e, 0x76,
	0x69, 0x61, 0x72, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x52, 0x65, 0x71,
	0x1a, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x45, 0x73,
	0x74, 0x61, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x09, 0x50, 0x65, 0x64, 0x69,
	0x72, 0x50, 0x6f, 0x7a, 0x6f, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x64,
	0x69, 0x72, 0x50, 0x6f, 0x7a, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x50, 0x65, 0x64, 0x69, 0x72, 0x50, 0x6f, 0x7a, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_lider_proto_rawDescOnce sync.Once
	file_lider_proto_rawDescData = file_lider_proto_rawDesc
)

func file_lider_proto_rawDescGZIP() []byte {
	file_lider_proto_rawDescOnce.Do(func() {
		file_lider_proto_rawDescData = protoimpl.X.CompressGZIP(file_lider_proto_rawDescData)
	})
	return file_lider_proto_rawDescData
}

var file_lider_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_lider_proto_goTypes = []interface{}{
	(*PedirPozoReq)(nil),     // 0: grpc.PedirPozoReq
	(*PedirPozoResp)(nil),    // 1: grpc.PedirPozoResp
	(*EnviarEstadoReq)(nil),  // 2: grpc.EnviarEstadoReq
	(*EnviarEstadoResp)(nil), // 3: grpc.EnviarEstadoResp
	(*UnirseReq)(nil),        // 4: grpc.UnirseReq
	(*UnirseResp)(nil),       // 5: grpc.UnirseResp
	(*EnviarJugadaReq)(nil),  // 6: grpc.EnviarJugadaReq
	(*EnviarJugadaResp)(nil), // 7: grpc.EnviarJugadaResp
}
var file_lider_proto_depIdxs = []int32{
	4, // 0: grpc.Lider.UnirseAJuego:input_type -> grpc.UnirseReq
	6, // 1: grpc.Lider.EnviarJugada:input_type -> grpc.EnviarJugadaReq
	2, // 2: grpc.Lider.EnviarEstado:input_type -> grpc.EnviarEstadoReq
	0, // 3: grpc.Lider.PedirPozo:input_type -> grpc.PedirPozoReq
	5, // 4: grpc.Lider.UnirseAJuego:output_type -> grpc.UnirseResp
	7, // 5: grpc.Lider.EnviarJugada:output_type -> grpc.EnviarJugadaResp
	3, // 6: grpc.Lider.EnviarEstado:output_type -> grpc.EnviarEstadoResp
	1, // 7: grpc.Lider.PedirPozo:output_type -> grpc.PedirPozoResp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lider_proto_init() }
func file_lider_proto_init() {
	if File_lider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PedirPozoReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PedirPozoResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnviarEstadoReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnviarEstadoResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lider_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnirseReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lider_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnirseResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lider_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnviarJugadaReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lider_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnviarJugadaResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lider_proto_goTypes,
		DependencyIndexes: file_lider_proto_depIdxs,
		MessageInfos:      file_lider_proto_msgTypes,
	}.Build()
	File_lider_proto = out.File
	file_lider_proto_rawDesc = nil
	file_lider_proto_goTypes = nil
	file_lider_proto_depIdxs = nil
}
