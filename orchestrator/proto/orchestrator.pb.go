// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/orchestrator.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StartSagaRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartSagaRequest) Reset() {
	*x = StartSagaRequest{}
	mi := &file_proto_orchestrator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartSagaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartSagaRequest) ProtoMessage() {}

func (x *StartSagaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orchestrator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartSagaRequest.ProtoReflect.Descriptor instead.
func (*StartSagaRequest) Descriptor() ([]byte, []int) {
	return file_proto_orchestrator_proto_rawDescGZIP(), []int{0}
}

func (x *StartSagaRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type StartSagaResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartSagaResponse) Reset() {
	*x = StartSagaResponse{}
	mi := &file_proto_orchestrator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartSagaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartSagaResponse) ProtoMessage() {}

func (x *StartSagaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orchestrator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartSagaResponse.ProtoReflect.Descriptor instead.
func (*StartSagaResponse) Descriptor() ([]byte, []int) {
	return file_proto_orchestrator_proto_rawDescGZIP(), []int{1}
}

func (x *StartSagaResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *StartSagaResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_orchestrator_proto protoreflect.FileDescriptor

const file_proto_orchestrator_proto_rawDesc = "" +
	"\n" +
	"\x18proto/orchestrator.proto\x12\forchestrator\"-\n" +
	"\x10StartSagaRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\"G\n" +
	"\x11StartSagaResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage2c\n" +
	"\x13OrchestratorService\x12L\n" +
	"\tStartSaga\x12\x1e.orchestrator.StartSagaRequest\x1a\x1f.orchestrator.StartSagaResponseB&Z$github.com/roybas/orchestrator/protob\x06proto3"

var (
	file_proto_orchestrator_proto_rawDescOnce sync.Once
	file_proto_orchestrator_proto_rawDescData []byte
)

func file_proto_orchestrator_proto_rawDescGZIP() []byte {
	file_proto_orchestrator_proto_rawDescOnce.Do(func() {
		file_proto_orchestrator_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_orchestrator_proto_rawDesc), len(file_proto_orchestrator_proto_rawDesc)))
	})
	return file_proto_orchestrator_proto_rawDescData
}

var file_proto_orchestrator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_orchestrator_proto_goTypes = []any{
	(*StartSagaRequest)(nil),  // 0: orchestrator.StartSagaRequest
	(*StartSagaResponse)(nil), // 1: orchestrator.StartSagaResponse
}
var file_proto_orchestrator_proto_depIdxs = []int32{
	0, // 0: orchestrator.OrchestratorService.StartSaga:input_type -> orchestrator.StartSagaRequest
	1, // 1: orchestrator.OrchestratorService.StartSaga:output_type -> orchestrator.StartSagaResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_orchestrator_proto_init() }
func file_proto_orchestrator_proto_init() {
	if File_proto_orchestrator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_orchestrator_proto_rawDesc), len(file_proto_orchestrator_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_orchestrator_proto_goTypes,
		DependencyIndexes: file_proto_orchestrator_proto_depIdxs,
		MessageInfos:      file_proto_orchestrator_proto_msgTypes,
	}.Build()
	File_proto_orchestrator_proto = out.File
	file_proto_orchestrator_proto_goTypes = nil
	file_proto_orchestrator_proto_depIdxs = nil
}
