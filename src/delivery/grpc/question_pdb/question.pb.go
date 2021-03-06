// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: src/delivery/grpc/proto/question.proto

package question_pdb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionText string `protobuf:"bytes,1,opt,name=question_text,json=questionText,proto3" json:"question_text,omitempty"`
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_delivery_grpc_proto_question_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_src_delivery_grpc_proto_question_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_src_delivery_grpc_proto_question_proto_rawDescGZIP(), []int{0}
}

func (x *Question) GetQuestionText() string {
	if x != nil {
		return x.QuestionText
	}
	return ""
}

type QuestionListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionVersion int32 `protobuf:"varint,1,opt,name=question_version,json=questionVersion,proto3" json:"question_version,omitempty"`
}

func (x *QuestionListRequest) Reset() {
	*x = QuestionListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_delivery_grpc_proto_question_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionListRequest) ProtoMessage() {}

func (x *QuestionListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_delivery_grpc_proto_question_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionListRequest.ProtoReflect.Descriptor instead.
func (*QuestionListRequest) Descriptor() ([]byte, []int) {
	return file_src_delivery_grpc_proto_question_proto_rawDescGZIP(), []int{1}
}

func (x *QuestionListRequest) GetQuestionVersion() int32 {
	if x != nil {
		return x.QuestionVersion
	}
	return 0
}

type QuestionListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Questions []*Question `protobuf:"bytes,1,rep,name=questions,proto3" json:"questions,omitempty"`
}

func (x *QuestionListResponse) Reset() {
	*x = QuestionListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_delivery_grpc_proto_question_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionListResponse) ProtoMessage() {}

func (x *QuestionListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_delivery_grpc_proto_question_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionListResponse.ProtoReflect.Descriptor instead.
func (*QuestionListResponse) Descriptor() ([]byte, []int) {
	return file_src_delivery_grpc_proto_question_proto_rawDescGZIP(), []int{2}
}

func (x *QuestionListResponse) GetQuestions() []*Question {
	if x != nil {
		return x.Questions
	}
	return nil
}

var File_src_delivery_grpc_proto_question_proto protoreflect.FileDescriptor

var file_src_delivery_grpc_proto_question_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x72, 0x63, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x70, 0x64, 0x62, 0x22, 0x2f, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x22, 0x40, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29,
	0x0a, 0x10, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x4c, 0x0a, 0x14, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x34, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x64, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x68, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x64, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x64, 0x62, 0x2e, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_delivery_grpc_proto_question_proto_rawDescOnce sync.Once
	file_src_delivery_grpc_proto_question_proto_rawDescData = file_src_delivery_grpc_proto_question_proto_rawDesc
)

func file_src_delivery_grpc_proto_question_proto_rawDescGZIP() []byte {
	file_src_delivery_grpc_proto_question_proto_rawDescOnce.Do(func() {
		file_src_delivery_grpc_proto_question_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_delivery_grpc_proto_question_proto_rawDescData)
	})
	return file_src_delivery_grpc_proto_question_proto_rawDescData
}

var file_src_delivery_grpc_proto_question_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_src_delivery_grpc_proto_question_proto_goTypes = []interface{}{
	(*Question)(nil),             // 0: question_pdb.Question
	(*QuestionListRequest)(nil),  // 1: question_pdb.QuestionListRequest
	(*QuestionListResponse)(nil), // 2: question_pdb.QuestionListResponse
}
var file_src_delivery_grpc_proto_question_proto_depIdxs = []int32{
	0, // 0: question_pdb.QuestionListResponse.questions:type_name -> question_pdb.Question
	1, // 1: question_pdb.QuestionService.ListQuestion:input_type -> question_pdb.QuestionListRequest
	2, // 2: question_pdb.QuestionService.ListQuestion:output_type -> question_pdb.QuestionListResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_src_delivery_grpc_proto_question_proto_init() }
func file_src_delivery_grpc_proto_question_proto_init() {
	if File_src_delivery_grpc_proto_question_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_delivery_grpc_proto_question_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Question); i {
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
		file_src_delivery_grpc_proto_question_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionListRequest); i {
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
		file_src_delivery_grpc_proto_question_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionListResponse); i {
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
			RawDescriptor: file_src_delivery_grpc_proto_question_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_delivery_grpc_proto_question_proto_goTypes,
		DependencyIndexes: file_src_delivery_grpc_proto_question_proto_depIdxs,
		MessageInfos:      file_src_delivery_grpc_proto_question_proto_msgTypes,
	}.Build()
	File_src_delivery_grpc_proto_question_proto = out.File
	file_src_delivery_grpc_proto_question_proto_rawDesc = nil
	file_src_delivery_grpc_proto_question_proto_goTypes = nil
	file_src_delivery_grpc_proto_question_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// QuestionServiceClient is the client API for QuestionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QuestionServiceClient interface {
	ListQuestion(ctx context.Context, in *QuestionListRequest, opts ...grpc.CallOption) (*QuestionListResponse, error)
}

type questionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuestionServiceClient(cc grpc.ClientConnInterface) QuestionServiceClient {
	return &questionServiceClient{cc}
}

func (c *questionServiceClient) ListQuestion(ctx context.Context, in *QuestionListRequest, opts ...grpc.CallOption) (*QuestionListResponse, error) {
	out := new(QuestionListResponse)
	err := c.cc.Invoke(ctx, "/question_pdb.QuestionService/ListQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuestionServiceServer is the server API for QuestionService service.
type QuestionServiceServer interface {
	ListQuestion(context.Context, *QuestionListRequest) (*QuestionListResponse, error)
}

// UnimplementedQuestionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQuestionServiceServer struct {
}

func (*UnimplementedQuestionServiceServer) ListQuestion(context.Context, *QuestionListRequest) (*QuestionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQuestion not implemented")
}

func RegisterQuestionServiceServer(s *grpc.Server, srv QuestionServiceServer) {
	s.RegisterService(&_QuestionService_serviceDesc, srv)
}

func _QuestionService_ListQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).ListQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/question_pdb.QuestionService/ListQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).ListQuestion(ctx, req.(*QuestionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QuestionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "question_pdb.QuestionService",
	HandlerType: (*QuestionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListQuestion",
			Handler:    _QuestionService_ListQuestion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/delivery/grpc/proto/question.proto",
}
