// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/pbfile/hello.proto

package pbfile

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

type SearchRequest struct {
	Condition            string                   `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	PersonNum            map[string]*SearchResult `protobuf:"bytes,2,rep,name=person_num,json=personNum,proto3" json:"person_num,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db567b686b10bd0b, []int{0}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetCondition() string {
	if m != nil {
		return m.Condition
	}
	return ""
}

func (m *SearchRequest) GetPersonNum() map[string]*SearchResult {
	if m != nil {
		return m.PersonNum
	}
	return nil
}

type SearchResult struct {
	Userinfo             string   `protobuf:"bytes,1,opt,name=userinfo,proto3" json:"userinfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchResult) Reset()         { *m = SearchResult{} }
func (m *SearchResult) String() string { return proto.CompactTextString(m) }
func (*SearchResult) ProtoMessage()    {}
func (*SearchResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_db567b686b10bd0b, []int{1}
}

func (m *SearchResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResult.Unmarshal(m, b)
}
func (m *SearchResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResult.Marshal(b, m, deterministic)
}
func (m *SearchResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResult.Merge(m, src)
}
func (m *SearchResult) XXX_Size() int {
	return xxx_messageInfo_SearchResult.Size(m)
}
func (m *SearchResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResult.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResult proto.InternalMessageInfo

func (m *SearchResult) GetUserinfo() string {
	if m != nil {
		return m.Userinfo
	}
	return ""
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "pbfile.SearchRequest")
	proto.RegisterMapType((map[string]*SearchResult)(nil), "pbfile.SearchRequest.PersonNumEntry")
	proto.RegisterType((*SearchResult)(nil), "pbfile.SearchResult")
}

func init() { proto.RegisterFile("grpc/pbfile/hello.proto", fileDescriptor_db567b686b10bd0b) }

var fileDescriptor_db567b686b10bd0b = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x2f, 0x2a, 0x48,
	0xd6, 0x2f, 0x48, 0x4a, 0xcb, 0xcc, 0x49, 0xd5, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0x88, 0x29, 0x1d, 0x63, 0xe4, 0xe2, 0x0d, 0x4e, 0x4d, 0x2c,
	0x4a, 0xce, 0x08, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe1, 0xe2, 0x4c, 0xce, 0xcf,
	0x4b, 0xc9, 0x2c, 0xc9, 0xcc, 0xcf, 0x93, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x08, 0x08,
	0x39, 0x73, 0x71, 0x15, 0xa4, 0x16, 0x15, 0xe7, 0xe7, 0xc5, 0xe7, 0x95, 0xe6, 0x4a, 0x30, 0x29,
	0x30, 0x6b, 0x70, 0x1b, 0xa9, 0xe8, 0x41, 0x0c, 0xd3, 0x43, 0x31, 0x48, 0x2f, 0x00, 0xac, 0xce,
	0xaf, 0x34, 0xd7, 0x35, 0xaf, 0xa4, 0xa8, 0x32, 0x88, 0xb3, 0x00, 0xc6, 0x97, 0x0a, 0xe2, 0xe2,
	0x43, 0x95, 0x14, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x84, 0x5a, 0x07, 0x62, 0x0a, 0x69, 0x71,
	0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x1b, 0x89, 0xa0, 0xdb,
	0x51, 0x5c, 0x9a, 0x53, 0x12, 0x04, 0x51, 0x62, 0xc5, 0x64, 0xc1, 0xa8, 0xa4, 0xc5, 0xc5, 0x83,
	0x2c, 0x25, 0x24, 0xc5, 0xc5, 0x51, 0x5a, 0x9c, 0x5a, 0x94, 0x99, 0x97, 0x96, 0x0f, 0x35, 0x16,
	0xce, 0x37, 0x72, 0xe5, 0xe2, 0x82, 0xa8, 0x0d, 0x2d, 0x4e, 0x2d, 0x12, 0x32, 0xe7, 0x62, 0x83,
	0xf0, 0x84, 0x44, 0xb1, 0x7a, 0x44, 0x0a, 0xab, 0xdd, 0x4a, 0x0c, 0x49, 0x6c, 0xe0, 0xa0, 0x34,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xf0, 0x2c, 0xf0, 0x65, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SearchUserClient is the client API for SearchUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchUserClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResult, error)
}

type searchUserClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchUserClient(cc grpc.ClientConnInterface) SearchUserClient {
	return &searchUserClient{cc}
}

func (c *searchUserClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResult, error) {
	out := new(SearchResult)
	err := c.cc.Invoke(ctx, "/pbfile.SearchUser/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchUserServer is the server API for SearchUser service.
type SearchUserServer interface {
	Search(context.Context, *SearchRequest) (*SearchResult, error)
}

// UnimplementedSearchUserServer can be embedded to have forward compatible implementations.
type UnimplementedSearchUserServer struct {
}

func (*UnimplementedSearchUserServer) Search(ctx context.Context, req *SearchRequest) (*SearchResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

func RegisterSearchUserServer(s *grpc.Server, srv SearchUserServer) {
	s.RegisterService(&_SearchUser_serviceDesc, srv)
}

func _SearchUser_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchUserServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbfile.SearchUser/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchUserServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchUser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbfile.SearchUser",
	HandlerType: (*SearchUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SearchUser_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/pbfile/hello.proto",
}
