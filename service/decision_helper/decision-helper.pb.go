// Code generated by protoc-gen-go. DO NOT EDIT.
// source: decision-helper.proto

package decision_helper

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Topic struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Topic) Reset()         { *m = Topic{} }
func (m *Topic) String() string { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()    {}
func (*Topic) Descriptor() ([]byte, []int) {
	return fileDescriptor_30e41342fbb1d60d, []int{0}
}

func (m *Topic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Topic.Unmarshal(m, b)
}
func (m *Topic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Topic.Marshal(b, m, deterministic)
}
func (m *Topic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topic.Merge(m, src)
}
func (m *Topic) XXX_Size() int {
	return xxx_messageInfo_Topic.Size(m)
}
func (m *Topic) XXX_DiscardUnknown() {
	xxx_messageInfo_Topic.DiscardUnknown(m)
}

var xxx_messageInfo_Topic proto.InternalMessageInfo

func (m *Topic) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Topic) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type Decision struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TopicId              int32    `protobuf:"varint,2,opt,name=topicId,proto3" json:"topicId,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Hots                 int32    `protobuf:"varint,4,opt,name=hots,proto3" json:"hots,omitempty"`
	Nots                 int32    `protobuf:"varint,5,opt,name=nots,proto3" json:"nots,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Decision) Reset()         { *m = Decision{} }
func (m *Decision) String() string { return proto.CompactTextString(m) }
func (*Decision) ProtoMessage()    {}
func (*Decision) Descriptor() ([]byte, []int) {
	return fileDescriptor_30e41342fbb1d60d, []int{1}
}

func (m *Decision) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Decision.Unmarshal(m, b)
}
func (m *Decision) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Decision.Marshal(b, m, deterministic)
}
func (m *Decision) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Decision.Merge(m, src)
}
func (m *Decision) XXX_Size() int {
	return xxx_messageInfo_Decision.Size(m)
}
func (m *Decision) XXX_DiscardUnknown() {
	xxx_messageInfo_Decision.DiscardUnknown(m)
}

var xxx_messageInfo_Decision proto.InternalMessageInfo

func (m *Decision) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Decision) GetTopicId() int32 {
	if m != nil {
		return m.TopicId
	}
	return 0
}

func (m *Decision) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Decision) GetHots() int32 {
	if m != nil {
		return m.Hots
	}
	return 0
}

func (m *Decision) GetNots() int32 {
	if m != nil {
		return m.Nots
	}
	return 0
}

type TopicsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicsRequest) Reset()         { *m = TopicsRequest{} }
func (m *TopicsRequest) String() string { return proto.CompactTextString(m) }
func (*TopicsRequest) ProtoMessage()    {}
func (*TopicsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30e41342fbb1d60d, []int{2}
}

func (m *TopicsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicsRequest.Unmarshal(m, b)
}
func (m *TopicsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicsRequest.Marshal(b, m, deterministic)
}
func (m *TopicsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicsRequest.Merge(m, src)
}
func (m *TopicsRequest) XXX_Size() int {
	return xxx_messageInfo_TopicsRequest.Size(m)
}
func (m *TopicsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TopicsRequest proto.InternalMessageInfo

type TopicsReply struct {
	Topics               []*Topic `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicsReply) Reset()         { *m = TopicsReply{} }
func (m *TopicsReply) String() string { return proto.CompactTextString(m) }
func (*TopicsReply) ProtoMessage()    {}
func (*TopicsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_30e41342fbb1d60d, []int{3}
}

func (m *TopicsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicsReply.Unmarshal(m, b)
}
func (m *TopicsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicsReply.Marshal(b, m, deterministic)
}
func (m *TopicsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicsReply.Merge(m, src)
}
func (m *TopicsReply) XXX_Size() int {
	return xxx_messageInfo_TopicsReply.Size(m)
}
func (m *TopicsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicsReply.DiscardUnknown(m)
}

var xxx_messageInfo_TopicsReply proto.InternalMessageInfo

func (m *TopicsReply) GetTopics() []*Topic {
	if m != nil {
		return m.Topics
	}
	return nil
}

type DecisionsRequest struct {
	TopicId              int32    `protobuf:"varint,1,opt,name=topicId,proto3" json:"topicId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecisionsRequest) Reset()         { *m = DecisionsRequest{} }
func (m *DecisionsRequest) String() string { return proto.CompactTextString(m) }
func (*DecisionsRequest) ProtoMessage()    {}
func (*DecisionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30e41342fbb1d60d, []int{4}
}

func (m *DecisionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecisionsRequest.Unmarshal(m, b)
}
func (m *DecisionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecisionsRequest.Marshal(b, m, deterministic)
}
func (m *DecisionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecisionsRequest.Merge(m, src)
}
func (m *DecisionsRequest) XXX_Size() int {
	return xxx_messageInfo_DecisionsRequest.Size(m)
}
func (m *DecisionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecisionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecisionsRequest proto.InternalMessageInfo

func (m *DecisionsRequest) GetTopicId() int32 {
	if m != nil {
		return m.TopicId
	}
	return 0
}

type DecisionsReply struct {
	Decisions            []*Decision `protobuf:"bytes,1,rep,name=decisions,proto3" json:"decisions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DecisionsReply) Reset()         { *m = DecisionsReply{} }
func (m *DecisionsReply) String() string { return proto.CompactTextString(m) }
func (*DecisionsReply) ProtoMessage()    {}
func (*DecisionsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_30e41342fbb1d60d, []int{5}
}

func (m *DecisionsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecisionsReply.Unmarshal(m, b)
}
func (m *DecisionsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecisionsReply.Marshal(b, m, deterministic)
}
func (m *DecisionsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecisionsReply.Merge(m, src)
}
func (m *DecisionsReply) XXX_Size() int {
	return xxx_messageInfo_DecisionsReply.Size(m)
}
func (m *DecisionsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DecisionsReply.DiscardUnknown(m)
}

var xxx_messageInfo_DecisionsReply proto.InternalMessageInfo

func (m *DecisionsReply) GetDecisions() []*Decision {
	if m != nil {
		return m.Decisions
	}
	return nil
}

type VoteRequest struct {
	TopicId              int32    `protobuf:"varint,1,opt,name=topicId,proto3" json:"topicId,omitempty"`
	DecisionId           int32    `protobuf:"varint,2,opt,name=decisionId,proto3" json:"decisionId,omitempty"`
	Hot                  bool     `protobuf:"varint,3,opt,name=hot,proto3" json:"hot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoteRequest) Reset()         { *m = VoteRequest{} }
func (m *VoteRequest) String() string { return proto.CompactTextString(m) }
func (*VoteRequest) ProtoMessage()    {}
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30e41342fbb1d60d, []int{6}
}

func (m *VoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteRequest.Unmarshal(m, b)
}
func (m *VoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteRequest.Marshal(b, m, deterministic)
}
func (m *VoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteRequest.Merge(m, src)
}
func (m *VoteRequest) XXX_Size() int {
	return xxx_messageInfo_VoteRequest.Size(m)
}
func (m *VoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VoteRequest proto.InternalMessageInfo

func (m *VoteRequest) GetTopicId() int32 {
	if m != nil {
		return m.TopicId
	}
	return 0
}

func (m *VoteRequest) GetDecisionId() int32 {
	if m != nil {
		return m.DecisionId
	}
	return 0
}

func (m *VoteRequest) GetHot() bool {
	if m != nil {
		return m.Hot
	}
	return false
}

type VoteReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoteReply) Reset()         { *m = VoteReply{} }
func (m *VoteReply) String() string { return proto.CompactTextString(m) }
func (*VoteReply) ProtoMessage()    {}
func (*VoteReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_30e41342fbb1d60d, []int{7}
}

func (m *VoteReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteReply.Unmarshal(m, b)
}
func (m *VoteReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteReply.Marshal(b, m, deterministic)
}
func (m *VoteReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteReply.Merge(m, src)
}
func (m *VoteReply) XXX_Size() int {
	return xxx_messageInfo_VoteReply.Size(m)
}
func (m *VoteReply) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteReply.DiscardUnknown(m)
}

var xxx_messageInfo_VoteReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Topic)(nil), "decision_helper.Topic")
	proto.RegisterType((*Decision)(nil), "decision_helper.Decision")
	proto.RegisterType((*TopicsRequest)(nil), "decision_helper.TopicsRequest")
	proto.RegisterType((*TopicsReply)(nil), "decision_helper.TopicsReply")
	proto.RegisterType((*DecisionsRequest)(nil), "decision_helper.DecisionsRequest")
	proto.RegisterType((*DecisionsReply)(nil), "decision_helper.DecisionsReply")
	proto.RegisterType((*VoteRequest)(nil), "decision_helper.VoteRequest")
	proto.RegisterType((*VoteReply)(nil), "decision_helper.VoteReply")
}

func init() { proto.RegisterFile("decision-helper.proto", fileDescriptor_30e41342fbb1d60d) }

var fileDescriptor_30e41342fbb1d60d = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x41, 0x4b, 0xfb, 0x30,
	0x1c, 0xfd, 0xb7, 0x5b, 0xf7, 0x5f, 0x7f, 0xc5, 0x6d, 0xc4, 0x29, 0x75, 0xca, 0x9c, 0x39, 0xed,
	0xe0, 0x7a, 0x98, 0x82, 0x27, 0x2f, 0x43, 0x36, 0x87, 0x08, 0xa3, 0x88, 0x20, 0x08, 0xa2, 0x6b,
	0xb0, 0x85, 0xb8, 0xc4, 0x36, 0xa2, 0xfb, 0x0a, 0x7e, 0x6a, 0x49, 0xb2, 0x74, 0x9d, 0x73, 0x7a,
	0x4b, 0x5e, 0xde, 0xef, 0xbd, 0x97, 0x47, 0x02, 0x3b, 0x11, 0x99, 0x26, 0x59, 0xc2, 0x66, 0xbd,
	0x98, 0x50, 0x4e, 0xd2, 0x80, 0xa7, 0x4c, 0x30, 0x54, 0x37, 0xf0, 0x83, 0x86, 0x71, 0x0f, 0x9c,
	0x1b, 0xc6, 0x93, 0x29, 0xaa, 0x81, 0x9d, 0x44, 0xbe, 0xd5, 0xb1, 0xba, 0x4e, 0x68, 0x27, 0x11,
	0x6a, 0x82, 0x23, 0x12, 0x41, 0x89, 0x6f, 0x77, 0xac, 0xae, 0x1b, 0xea, 0x0d, 0x4e, 0xa1, 0x7a,
	0xb1, 0x50, 0x58, 0x9b, 0xf0, 0xe1, 0xbf, 0x90, 0x52, 0xe3, 0x48, 0xcd, 0x38, 0xa1, 0xd9, 0x2e,
	0xb5, 0x4a, 0x05, 0x2d, 0x84, 0xa0, 0x1c, 0x33, 0x91, 0xf9, 0x65, 0x45, 0x56, 0x6b, 0x89, 0xcd,
	0x24, 0xe6, 0x68, 0x4c, 0xae, 0x71, 0x1d, 0xb6, 0x54, 0xc4, 0x2c, 0x24, 0xaf, 0x6f, 0x24, 0x13,
	0xf8, 0x1c, 0x3c, 0x03, 0x70, 0x3a, 0x47, 0x01, 0x54, 0x94, 0x51, 0xe6, 0x5b, 0x9d, 0x52, 0xd7,
	0xeb, 0xef, 0x06, 0xdf, 0x2e, 0x19, 0x28, 0x76, 0xb8, 0x60, 0xe1, 0x63, 0x68, 0x98, 0x3b, 0x18,
	0xc9, 0x62, 0x76, 0x6b, 0x25, 0x3b, 0x1e, 0x43, 0xad, 0xc0, 0x96, 0x7e, 0x67, 0xe0, 0x1a, 0x03,
	0x63, 0xb9, 0xb7, 0x66, 0x69, 0x66, 0xc2, 0x25, 0x17, 0xdf, 0x81, 0x77, 0xcb, 0x04, 0xf9, 0xd3,
	0x13, 0xb5, 0x01, 0xcc, 0x54, 0x5e, 0x66, 0x01, 0x41, 0x0d, 0x28, 0xc5, 0x4c, 0xa8, 0x36, 0xab,
	0xa1, 0x5c, 0x62, 0x0f, 0x5c, 0x2d, 0xcd, 0xe9, 0xbc, 0xff, 0x69, 0x2f, 0x33, 0x5f, 0xaa, 0x38,
	0xe8, 0x0a, 0xdc, 0x11, 0x11, 0xba, 0x35, 0xd4, 0xfe, 0xb9, 0x20, 0x53, 0x46, 0xeb, 0x60, 0xe3,
	0x39, 0xa7, 0x73, 0xfc, 0x0f, 0xdd, 0x43, 0x73, 0x44, 0x44, 0xde, 0xca, 0x90, 0xa5, 0xfa, 0x09,
	0x1d, 0x6d, 0x6c, 0x21, 0x97, 0x3e, 0xfc, 0x8d, 0xa2, 0xd5, 0xaf, 0xa1, 0x2e, 0xaf, 0x32, 0x64,
	0x69, 0xfe, 0xd2, 0xd6, 0x03, 0x15, 0x7a, 0x6c, 0xb5, 0x36, 0x9c, 0x2a, 0xb9, 0xc1, 0x29, 0xec,
	0x27, 0x2c, 0x78, 0x4e, 0xf9, 0x34, 0x20, 0x1f, 0x8f, 0x2f, 0x9c, 0x92, 0x2c, 0x88, 0x09, 0xa5,
	0xec, 0x9d, 0xa5, 0x34, 0x1a, 0x6c, 0xaf, 0x16, 0x35, 0x91, 0xbf, 0x64, 0x62, 0x3d, 0x55, 0xd4,
	0x77, 0x39, 0xf9, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xee, 0xd3, 0x34, 0x47, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DecisionHelperClient is the client API for DecisionHelper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DecisionHelperClient interface {
	// Returns topics
	GetTopics(ctx context.Context, in *TopicsRequest, opts ...grpc.CallOption) (*TopicsReply, error)
	GetDecisionsForTopic(ctx context.Context, in *DecisionsRequest, opts ...grpc.CallOption) (*DecisionsReply, error)
	VoteForDecision(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteReply, error)
}

type decisionHelperClient struct {
	cc *grpc.ClientConn
}

func NewDecisionHelperClient(cc *grpc.ClientConn) DecisionHelperClient {
	return &decisionHelperClient{cc}
}

func (c *decisionHelperClient) GetTopics(ctx context.Context, in *TopicsRequest, opts ...grpc.CallOption) (*TopicsReply, error) {
	out := new(TopicsReply)
	err := c.cc.Invoke(ctx, "/decision_helper.DecisionHelper/GetTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *decisionHelperClient) GetDecisionsForTopic(ctx context.Context, in *DecisionsRequest, opts ...grpc.CallOption) (*DecisionsReply, error) {
	out := new(DecisionsReply)
	err := c.cc.Invoke(ctx, "/decision_helper.DecisionHelper/GetDecisionsForTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *decisionHelperClient) VoteForDecision(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteReply, error) {
	out := new(VoteReply)
	err := c.cc.Invoke(ctx, "/decision_helper.DecisionHelper/VoteForDecision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DecisionHelperServer is the server API for DecisionHelper service.
type DecisionHelperServer interface {
	// Returns topics
	GetTopics(context.Context, *TopicsRequest) (*TopicsReply, error)
	GetDecisionsForTopic(context.Context, *DecisionsRequest) (*DecisionsReply, error)
	VoteForDecision(context.Context, *VoteRequest) (*VoteReply, error)
}

// UnimplementedDecisionHelperServer can be embedded to have forward compatible implementations.
type UnimplementedDecisionHelperServer struct {
}

func (*UnimplementedDecisionHelperServer) GetTopics(ctx context.Context, req *TopicsRequest) (*TopicsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopics not implemented")
}
func (*UnimplementedDecisionHelperServer) GetDecisionsForTopic(ctx context.Context, req *DecisionsRequest) (*DecisionsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDecisionsForTopic not implemented")
}
func (*UnimplementedDecisionHelperServer) VoteForDecision(ctx context.Context, req *VoteRequest) (*VoteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteForDecision not implemented")
}

func RegisterDecisionHelperServer(s *grpc.Server, srv DecisionHelperServer) {
	s.RegisterService(&_DecisionHelper_serviceDesc, srv)
}

func _DecisionHelper_GetTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecisionHelperServer).GetTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/decision_helper.DecisionHelper/GetTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecisionHelperServer).GetTopics(ctx, req.(*TopicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DecisionHelper_GetDecisionsForTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecisionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecisionHelperServer).GetDecisionsForTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/decision_helper.DecisionHelper/GetDecisionsForTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecisionHelperServer).GetDecisionsForTopic(ctx, req.(*DecisionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DecisionHelper_VoteForDecision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecisionHelperServer).VoteForDecision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/decision_helper.DecisionHelper/VoteForDecision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecisionHelperServer).VoteForDecision(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DecisionHelper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "decision_helper.DecisionHelper",
	HandlerType: (*DecisionHelperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTopics",
			Handler:    _DecisionHelper_GetTopics_Handler,
		},
		{
			MethodName: "GetDecisionsForTopic",
			Handler:    _DecisionHelper_GetDecisionsForTopic_Handler,
		},
		{
			MethodName: "VoteForDecision",
			Handler:    _DecisionHelper_VoteForDecision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "decision-helper.proto",
}
