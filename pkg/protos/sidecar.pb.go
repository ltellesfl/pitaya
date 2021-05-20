// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: sidecar.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StartPitayaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *StartPitayaRequest_ServerConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *StartPitayaRequest) Reset() {
	*x = StartPitayaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sidecar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartPitayaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartPitayaRequest) ProtoMessage() {}

func (x *StartPitayaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sidecar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartPitayaRequest.ProtoReflect.Descriptor instead.
func (*StartPitayaRequest) Descriptor() ([]byte, []int) {
	return file_sidecar_proto_rawDescGZIP(), []int{0}
}

func (x *StartPitayaRequest) GetConfig() *StartPitayaRequest_ServerConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type RequestTo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerID string `protobuf:"bytes,1,opt,name=serverID,proto3" json:"serverID,omitempty"`
	Msg      *Msg   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *RequestTo) Reset() {
	*x = RequestTo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sidecar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestTo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestTo) ProtoMessage() {}

func (x *RequestTo) ProtoReflect() protoreflect.Message {
	mi := &file_sidecar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestTo.ProtoReflect.Descriptor instead.
func (*RequestTo) Descriptor() ([]byte, []int) {
	return file_sidecar_proto_rawDescGZIP(), []int{1}
}

func (x *RequestTo) GetServerID() string {
	if x != nil {
		return x.ServerID
	}
	return ""
}

func (x *RequestTo) GetMsg() *Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

type RPCResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqId uint64    `protobuf:"varint,1,opt,name=reqId,proto3" json:"reqId,omitempty"`
	Res   *Response `protobuf:"bytes,2,opt,name=res,proto3" json:"res,omitempty"`
	Err   *Error    `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *RPCResponse) Reset() {
	*x = RPCResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sidecar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCResponse) ProtoMessage() {}

func (x *RPCResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sidecar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCResponse.ProtoReflect.Descriptor instead.
func (*RPCResponse) Descriptor() ([]byte, []int) {
	return file_sidecar_proto_rawDescGZIP(), []int{2}
}

func (x *RPCResponse) GetReqId() uint64 {
	if x != nil {
		return x.ReqId
	}
	return 0
}

func (x *RPCResponse) GetRes() *Response {
	if x != nil {
		return x.Res
	}
	return nil
}

func (x *RPCResponse) GetErr() *Error {
	if x != nil {
		return x.Err
	}
	return nil
}

type SidecarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqId uint64   `protobuf:"varint,1,opt,name=reqId,proto3" json:"reqId,omitempty"`
	Req   *Request `protobuf:"bytes,2,opt,name=req,proto3" json:"req,omitempty"`
}

func (x *SidecarRequest) Reset() {
	*x = SidecarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sidecar_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SidecarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SidecarRequest) ProtoMessage() {}

func (x *SidecarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sidecar_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SidecarRequest.ProtoReflect.Descriptor instead.
func (*SidecarRequest) Descriptor() ([]byte, []int) {
	return file_sidecar_proto_rawDescGZIP(), []int{3}
}

func (x *SidecarRequest) GetReqId() uint64 {
	if x != nil {
		return x.ReqId
	}
	return 0
}

func (x *SidecarRequest) GetReq() *Request {
	if x != nil {
		return x.Req
	}
	return nil
}

type StartPitayaRequest_ServerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IsFrontend             bool              `protobuf:"varint,2,opt,name=isFrontend,proto3" json:"isFrontend,omitempty"`
	ServerType             string            `protobuf:"bytes,3,opt,name=serverType,proto3" json:"serverType,omitempty"`
	Metadata               map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ShouldCompressMessages bool              `protobuf:"varint,5,opt,name=shouldCompressMessages,proto3" json:"shouldCompressMessages,omitempty"`
	DebugLog               bool              `protobuf:"varint,6,opt,name=debugLog,proto3" json:"debugLog,omitempty"`
}

func (x *StartPitayaRequest_ServerConfig) Reset() {
	*x = StartPitayaRequest_ServerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sidecar_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartPitayaRequest_ServerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartPitayaRequest_ServerConfig) ProtoMessage() {}

func (x *StartPitayaRequest_ServerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_sidecar_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartPitayaRequest_ServerConfig.ProtoReflect.Descriptor instead.
func (*StartPitayaRequest_ServerConfig) Descriptor() ([]byte, []int) {
	return file_sidecar_proto_rawDescGZIP(), []int{0, 0}
}

func (x *StartPitayaRequest_ServerConfig) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StartPitayaRequest_ServerConfig) GetIsFrontend() bool {
	if x != nil {
		return x.IsFrontend
	}
	return false
}

func (x *StartPitayaRequest_ServerConfig) GetServerType() string {
	if x != nil {
		return x.ServerType
	}
	return ""
}

func (x *StartPitayaRequest_ServerConfig) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *StartPitayaRequest_ServerConfig) GetShouldCompressMessages() bool {
	if x != nil {
		return x.ShouldCompressMessages
	}
	return false
}

func (x *StartPitayaRequest_ServerConfig) GetDebugLog() bool {
	if x != nil {
		return x.DebugLog
	}
	return false
}

var File_sidecar_proto protoreflect.FileDescriptor

var file_sidecar_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x03, 0x0a, 0x12,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x69, 0x74, 0x61, 0x79, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x50, 0x69, 0x74, 0x61, 0x79, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x1a, 0xc2, 0x02, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x72, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x51, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x69, 0x74, 0x61, 0x79, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x16, 0x73, 0x68, 0x6f, 0x75, 0x6c,
	0x64, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x43,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x65, 0x62, 0x75, 0x67, 0x4c, 0x6f, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x64, 0x65, 0x62, 0x75, 0x67, 0x4c, 0x6f, 0x67, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x46, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x54, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x1d, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x22, 0x68, 0x0a, 0x0b, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x72, 0x65, 0x71, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03, 0x72, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x03, 0x65, 0x72, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x49, 0x0a, 0x0e, 0x53, 0x69,
	0x64, 0x65, 0x63, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x65, 0x71, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x72, 0x65, 0x71,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x03, 0x72, 0x65, 0x71, 0x32, 0x92, 0x02, 0x0a, 0x07, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61,
	0x72, 0x12, 0x3c, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x52, 0x50, 0x43, 0x12, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x69, 0x64,
	0x65, 0x63, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x28, 0x01, 0x30, 0x01, 0x12,
	0x28, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x50, 0x43, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x50, 0x43, 0x54, 0x6f, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x50, 0x69, 0x74, 0x61, 0x79, 0x61, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x69, 0x74, 0x61, 0x79, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x33, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x70, 0x50, 0x69, 0x74,
	0x61, 0x79, 0x61, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x3c, 0x5a, 0x29, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x70, 0x66, 0x72, 0x65, 0x65,
	0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x70, 0x69, 0x74, 0x61, 0x79, 0x61, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0xaa, 0x02, 0x0e, 0x4e, 0x50, 0x69, 0x74, 0x61, 0x79,
	0x61, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sidecar_proto_rawDescOnce sync.Once
	file_sidecar_proto_rawDescData = file_sidecar_proto_rawDesc
)

func file_sidecar_proto_rawDescGZIP() []byte {
	file_sidecar_proto_rawDescOnce.Do(func() {
		file_sidecar_proto_rawDescData = protoimpl.X.CompressGZIP(file_sidecar_proto_rawDescData)
	})
	return file_sidecar_proto_rawDescData
}

var file_sidecar_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sidecar_proto_goTypes = []interface{}{
	(*StartPitayaRequest)(nil),              // 0: protos.StartPitayaRequest
	(*RequestTo)(nil),                       // 1: protos.RequestTo
	(*RPCResponse)(nil),                     // 2: protos.RPCResponse
	(*SidecarRequest)(nil),                  // 3: protos.SidecarRequest
	(*StartPitayaRequest_ServerConfig)(nil), // 4: protos.StartPitayaRequest.ServerConfig
	nil,                                     // 5: protos.StartPitayaRequest.ServerConfig.MetadataEntry
	(*Msg)(nil),                             // 6: protos.Msg
	(*Response)(nil),                        // 7: protos.Response
	(*Error)(nil),                           // 8: protos.Error
	(*Request)(nil),                         // 9: protos.Request
	(*emptypb.Empty)(nil),                   // 10: google.protobuf.Empty
}
var file_sidecar_proto_depIdxs = []int32{
	4,  // 0: protos.StartPitayaRequest.config:type_name -> protos.StartPitayaRequest.ServerConfig
	6,  // 1: protos.RequestTo.msg:type_name -> protos.Msg
	7,  // 2: protos.RPCResponse.res:type_name -> protos.Response
	8,  // 3: protos.RPCResponse.err:type_name -> protos.Error
	9,  // 4: protos.SidecarRequest.req:type_name -> protos.Request
	5,  // 5: protos.StartPitayaRequest.ServerConfig.metadata:type_name -> protos.StartPitayaRequest.ServerConfig.MetadataEntry
	2,  // 6: protos.Sidecar.ListenRPC:input_type -> protos.RPCResponse
	6,  // 7: protos.Sidecar.SendRPC:input_type -> protos.Msg
	1,  // 8: protos.Sidecar.SendRPCTo:input_type -> protos.RequestTo
	0,  // 9: protos.Sidecar.StartPitaya:input_type -> protos.StartPitayaRequest
	10, // 10: protos.Sidecar.StopPitaya:input_type -> google.protobuf.Empty
	3,  // 11: protos.Sidecar.ListenRPC:output_type -> protos.SidecarRequest
	7,  // 12: protos.Sidecar.SendRPC:output_type -> protos.Response
	7,  // 13: protos.Sidecar.SendRPCTo:output_type -> protos.Response
	8,  // 14: protos.Sidecar.StartPitaya:output_type -> protos.Error
	8,  // 15: protos.Sidecar.StopPitaya:output_type -> protos.Error
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_sidecar_proto_init() }
func file_sidecar_proto_init() {
	if File_sidecar_proto != nil {
		return
	}
	file_error_proto_init()
	file_msg_proto_init()
	file_response_proto_init()
	file_request_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sidecar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartPitayaRequest); i {
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
		file_sidecar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestTo); i {
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
		file_sidecar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCResponse); i {
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
		file_sidecar_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SidecarRequest); i {
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
		file_sidecar_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartPitayaRequest_ServerConfig); i {
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
			RawDescriptor: file_sidecar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sidecar_proto_goTypes,
		DependencyIndexes: file_sidecar_proto_depIdxs,
		MessageInfos:      file_sidecar_proto_msgTypes,
	}.Build()
	File_sidecar_proto = out.File
	file_sidecar_proto_rawDesc = nil
	file_sidecar_proto_goTypes = nil
	file_sidecar_proto_depIdxs = nil
}
