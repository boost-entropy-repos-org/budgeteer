// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/budgeteer/budget_service.proto

package budgeteerpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Purchase struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount               int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Purchase) Reset()         { *m = Purchase{} }
func (m *Purchase) String() string { return proto.CompactTextString(m) }
func (*Purchase) ProtoMessage()    {}
func (*Purchase) Descriptor() ([]byte, []int) {
	return fileDescriptor_budget_service_cc71af05fa2a8eaa, []int{0}
}
func (m *Purchase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Purchase.Unmarshal(m, b)
}
func (m *Purchase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Purchase.Marshal(b, m, deterministic)
}
func (dst *Purchase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Purchase.Merge(dst, src)
}
func (m *Purchase) XXX_Size() int {
	return xxx_messageInfo_Purchase.Size(m)
}
func (m *Purchase) XXX_DiscardUnknown() {
	xxx_messageInfo_Purchase.DiscardUnknown(m)
}

var xxx_messageInfo_Purchase proto.InternalMessageInfo

func (m *Purchase) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Purchase) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type GetPurchasesRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPurchasesRequest) Reset()         { *m = GetPurchasesRequest{} }
func (m *GetPurchasesRequest) String() string { return proto.CompactTextString(m) }
func (*GetPurchasesRequest) ProtoMessage()    {}
func (*GetPurchasesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_budget_service_cc71af05fa2a8eaa, []int{1}
}
func (m *GetPurchasesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPurchasesRequest.Unmarshal(m, b)
}
func (m *GetPurchasesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPurchasesRequest.Marshal(b, m, deterministic)
}
func (dst *GetPurchasesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPurchasesRequest.Merge(dst, src)
}
func (m *GetPurchasesRequest) XXX_Size() int {
	return xxx_messageInfo_GetPurchasesRequest.Size(m)
}
func (m *GetPurchasesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPurchasesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPurchasesRequest proto.InternalMessageInfo

func (m *GetPurchasesRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Purchase)(nil), "budgeteer.Purchase")
	proto.RegisterType((*GetPurchasesRequest)(nil), "budgeteer.GetPurchasesRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BudgetServiceClient is the client API for BudgetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BudgetServiceClient interface {
	GetPurchase(ctx context.Context, in *GetPurchasesRequest, opts ...grpc.CallOption) (*Purchase, error)
}

type budgetServiceClient struct {
	cc *grpc.ClientConn
}

func NewBudgetServiceClient(cc *grpc.ClientConn) BudgetServiceClient {
	return &budgetServiceClient{cc}
}

func (c *budgetServiceClient) GetPurchase(ctx context.Context, in *GetPurchasesRequest, opts ...grpc.CallOption) (*Purchase, error) {
	out := new(Purchase)
	err := c.cc.Invoke(ctx, "/budgeteer.BudgetService/GetPurchase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BudgetServiceServer is the server API for BudgetService service.
type BudgetServiceServer interface {
	GetPurchase(context.Context, *GetPurchasesRequest) (*Purchase, error)
}

func RegisterBudgetServiceServer(s *grpc.Server, srv BudgetServiceServer) {
	s.RegisterService(&_BudgetService_serviceDesc, srv)
}

func _BudgetService_GetPurchase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPurchasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).GetPurchase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeteer.BudgetService/GetPurchase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).GetPurchase(ctx, req.(*GetPurchasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BudgetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "budgeteer.BudgetService",
	HandlerType: (*BudgetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPurchase",
			Handler:    _BudgetService_GetPurchase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/budgeteer/budget_service.proto",
}

func init() {
	proto.RegisterFile("proto/budgeteer/budget_service.proto", fileDescriptor_budget_service_cc71af05fa2a8eaa)
}

var fileDescriptor_budget_service_cc71af05fa2a8eaa = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2a, 0x4d, 0x49, 0x4f, 0x2d, 0x49, 0x4d, 0x2d, 0x82, 0xb2, 0xe2, 0x8b, 0x53,
	0x8b, 0xca, 0x32, 0x93, 0x53, 0xf5, 0xc0, 0xd2, 0x42, 0x9c, 0x70, 0x79, 0x25, 0x23, 0x2e, 0x8e,
	0x80, 0xd2, 0xa2, 0xe4, 0x8c, 0xc4, 0xe2, 0x54, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x31, 0x2e, 0xb6, 0xc4, 0xdc, 0xfc, 0xd2,
	0xbc, 0x12, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x28, 0x4f, 0x49, 0x95, 0x4b, 0xd8, 0x3d,
	0xb5, 0x04, 0xa6, 0xad, 0x38, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x04, 0x5d, 0xbb, 0x51, 0x28,
	0x17, 0xaf, 0x13, 0xd8, 0x9e, 0x60, 0x88, 0xe5, 0x42, 0x2e, 0x5c, 0xdc, 0x48, 0xfa, 0x84, 0xe4,
	0xf4, 0xe0, 0xce, 0xd0, 0xc3, 0x62, 0x9e, 0x94, 0x30, 0x92, 0x3c, 0x4c, 0x52, 0x89, 0xc1, 0x89,
	0x37, 0x8a, 0x1b, 0x2e, 0x5e, 0x90, 0x94, 0xc4, 0x06, 0xf6, 0x92, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x09, 0x4a, 0x33, 0x35, 0xfa, 0x00, 0x00, 0x00,
}