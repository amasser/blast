// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/index/index.proto

package index

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	raft "github.com/mosuka/blast/protobuf/raft"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IndexCommand_Type int32

const (
	IndexCommand_UNKNOWN_COMMAND IndexCommand_Type = 0
	IndexCommand_SET_METADATA    IndexCommand_Type = 1
	IndexCommand_DELETE_METADATA IndexCommand_Type = 2
	IndexCommand_INDEX_DOCUMENT  IndexCommand_Type = 3
	IndexCommand_DELETE_DOCUMENT IndexCommand_Type = 4
)

var IndexCommand_Type_name = map[int32]string{
	0: "UNKNOWN_COMMAND",
	1: "SET_METADATA",
	2: "DELETE_METADATA",
	3: "INDEX_DOCUMENT",
	4: "DELETE_DOCUMENT",
}

var IndexCommand_Type_value = map[string]int32{
	"UNKNOWN_COMMAND": 0,
	"SET_METADATA":    1,
	"DELETE_METADATA": 2,
	"INDEX_DOCUMENT":  3,
	"DELETE_DOCUMENT": 4,
}

func (x IndexCommand_Type) String() string {
	return proto.EnumName(IndexCommand_Type_name, int32(x))
}

func (IndexCommand_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7b2daf652facb3ae, []int{7, 0}
}

type Document struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Fields               *any.Any `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Document) Reset()         { *m = Document{} }
func (m *Document) String() string { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()    {}
func (*Document) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2daf652facb3ae, []int{0}
}

func (m *Document) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Document.Unmarshal(m, b)
}
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Document.Marshal(b, m, deterministic)
}
func (m *Document) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document.Merge(m, src)
}
func (m *Document) XXX_Size() int {
	return xxx_messageInfo_Document.Size(m)
}
func (m *Document) XXX_DiscardUnknown() {
	xxx_messageInfo_Document.DiscardUnknown(m)
}

var xxx_messageInfo_Document proto.InternalMessageInfo

func (m *Document) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Document) GetFields() *any.Any {
	if m != nil {
		return m.Fields
	}
	return nil
}

type GetRequest struct {
	Document             *Document `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2daf652facb3ae, []int{1}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetDocument() *Document {
	if m != nil {
		return m.Document
	}
	return nil
}

type GetResponse struct {
	Document             *Document `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2daf652facb3ae, []int{2}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetDocument() *Document {
	if m != nil {
		return m.Document
	}
	return nil
}

type SearchRequest struct {
	SearchRequest        *any.Any `protobuf:"bytes,1,opt,name=search_request,json=searchRequest,proto3" json:"search_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2daf652facb3ae, []int{3}
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

func (m *SearchRequest) GetSearchRequest() *any.Any {
	if m != nil {
		return m.SearchRequest
	}
	return nil
}

type SearchResponse struct {
	SearchResult         *any.Any `protobuf:"bytes,1,opt,name=search_result,json=searchResult,proto3" json:"search_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2daf652facb3ae, []int{4}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetSearchResult() *any.Any {
	if m != nil {
		return m.SearchResult
	}
	return nil
}

type IndexRequest struct {
	Document             *Document `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *IndexRequest) Reset()         { *m = IndexRequest{} }
func (m *IndexRequest) String() string { return proto.CompactTextString(m) }
func (*IndexRequest) ProtoMessage()    {}
func (*IndexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2daf652facb3ae, []int{5}
}

func (m *IndexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexRequest.Unmarshal(m, b)
}
func (m *IndexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexRequest.Marshal(b, m, deterministic)
}
func (m *IndexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexRequest.Merge(m, src)
}
func (m *IndexRequest) XXX_Size() int {
	return xxx_messageInfo_IndexRequest.Size(m)
}
func (m *IndexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IndexRequest proto.InternalMessageInfo

func (m *IndexRequest) GetDocument() *Document {
	if m != nil {
		return m.Document
	}
	return nil
}

type DeleteRequest struct {
	Document             *Document `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2daf652facb3ae, []int{6}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetDocument() *Document {
	if m != nil {
		return m.Document
	}
	return nil
}

type IndexCommand struct {
	Type                 IndexCommand_Type `protobuf:"varint,1,opt,name=type,proto3,enum=index.IndexCommand_Type" json:"type,omitempty"`
	Data                 *any.Any          `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *IndexCommand) Reset()         { *m = IndexCommand{} }
func (m *IndexCommand) String() string { return proto.CompactTextString(m) }
func (*IndexCommand) ProtoMessage()    {}
func (*IndexCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2daf652facb3ae, []int{7}
}

func (m *IndexCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexCommand.Unmarshal(m, b)
}
func (m *IndexCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexCommand.Marshal(b, m, deterministic)
}
func (m *IndexCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexCommand.Merge(m, src)
}
func (m *IndexCommand) XXX_Size() int {
	return xxx_messageInfo_IndexCommand.Size(m)
}
func (m *IndexCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexCommand.DiscardUnknown(m)
}

var xxx_messageInfo_IndexCommand proto.InternalMessageInfo

func (m *IndexCommand) GetType() IndexCommand_Type {
	if m != nil {
		return m.Type
	}
	return IndexCommand_UNKNOWN_COMMAND
}

func (m *IndexCommand) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("index.IndexCommand_Type", IndexCommand_Type_name, IndexCommand_Type_value)
	proto.RegisterType((*Document)(nil), "index.Document")
	proto.RegisterType((*GetRequest)(nil), "index.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "index.GetResponse")
	proto.RegisterType((*SearchRequest)(nil), "index.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "index.SearchResponse")
	proto.RegisterType((*IndexRequest)(nil), "index.IndexRequest")
	proto.RegisterType((*DeleteRequest)(nil), "index.DeleteRequest")
	proto.RegisterType((*IndexCommand)(nil), "index.IndexCommand")
}

func init() { proto.RegisterFile("protobuf/index/index.proto", fileDescriptor_7b2daf652facb3ae) }

var fileDescriptor_7b2daf652facb3ae = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdb, 0x6e, 0xda, 0x40,
	0x10, 0x0d, 0xc4, 0x50, 0x3a, 0x5c, 0x42, 0x36, 0x49, 0x45, 0xdd, 0x97, 0xc8, 0x0f, 0x15, 0x52,
	0x2b, 0x23, 0x11, 0xf5, 0x92, 0xcb, 0x43, 0x29, 0xb6, 0x68, 0x1a, 0x30, 0x12, 0x38, 0x6a, 0xd5,
	0x17, 0x64, 0xf0, 0x02, 0x56, 0x6c, 0xaf, 0xcb, 0xae, 0xab, 0xf2, 0x0b, 0xfd, 0xbd, 0xfe, 0x50,
	0xe5, 0xf5, 0xda, 0x31, 0x69, 0x49, 0xd4, 0xbc, 0xac, 0xd8, 0x33, 0xe7, 0x9c, 0x19, 0x76, 0x66,
	0x0c, 0x72, 0xb0, 0x22, 0x8c, 0x4c, 0xc3, 0x79, 0xcb, 0xf1, 0x6d, 0xfc, 0x33, 0x3e, 0x55, 0x0e,
	0xa2, 0x02, 0xbf, 0xc8, 0xcf, 0x17, 0x84, 0x2c, 0x5c, 0xdc, 0x4a, 0x99, 0x96, 0xbf, 0x8e, 0x19,
	0xf2, 0x8b, 0xbb, 0x21, 0xec, 0x05, 0x2c, 0x09, 0x36, 0x52, 0x74, 0x65, 0xcd, 0x19, 0x3f, 0xe2,
	0x88, 0xf2, 0x09, 0x4a, 0x1a, 0x99, 0x85, 0x1e, 0xf6, 0x19, 0xaa, 0x41, 0xde, 0xb1, 0x1b, 0xb9,
	0xe3, 0x5c, 0xf3, 0xe9, 0x28, 0xef, 0xd8, 0xe8, 0x35, 0x14, 0xe7, 0x0e, 0x76, 0x6d, 0xda, 0xc8,
	0x1f, 0xe7, 0x9a, 0xe5, 0xf6, 0xa1, 0x1a, 0xe7, 0x50, 0x13, 0x37, 0xb5, 0xe3, 0xaf, 0x47, 0x82,
	0xa3, 0x9c, 0x02, 0xf4, 0x30, 0x1b, 0xe1, 0xef, 0x21, 0xa6, 0x0c, 0xbd, 0x82, 0x92, 0x2d, 0x7c,
	0xb9, 0x63, 0xb9, 0xbd, 0xa7, 0xc6, 0x7f, 0x28, 0x49, 0x37, 0x4a, 0x09, 0xca, 0x19, 0x94, 0xb9,
	0x94, 0x06, 0xc4, 0xa7, 0xf8, 0xff, 0xb4, 0x7d, 0xa8, 0x8e, 0xb1, 0xb5, 0x9a, 0x2d, 0x93, 0xcc,
	0xe7, 0x50, 0xa3, 0x1c, 0x98, 0xac, 0x62, 0x44, 0x78, 0xfc, 0xbb, 0xfa, 0x2a, 0xcd, 0x8a, 0x95,
	0x2b, 0xa8, 0x25, 0x6e, 0xa2, 0x98, 0x53, 0xa8, 0xa6, 0x76, 0x34, 0x74, 0xef, 0x77, 0xab, 0x24,
	0x6e, 0x11, 0x53, 0x39, 0x87, 0xca, 0x65, 0x54, 0xf6, 0xa3, 0xde, 0xe4, 0x02, 0xaa, 0x1a, 0x76,
	0x31, 0xc3, 0x8f, 0x52, 0xff, 0xce, 0x89, 0xdc, 0x5d, 0xe2, 0x79, 0x96, 0x1f, 0xf5, 0x52, 0x62,
	0xeb, 0x00, 0x73, 0x65, 0xad, 0xdd, 0x10, 0xca, 0x2c, 0x45, 0x35, 0xd7, 0x01, 0x1e, 0x71, 0x16,
	0x6a, 0x82, 0x64, 0x5b, 0xcc, 0xba, 0xb7, 0xef, 0x9c, 0xa1, 0xdc, 0x80, 0x14, 0xe9, 0xd0, 0x01,
	0xec, 0x5d, 0x1b, 0x57, 0xc6, 0xf0, 0x8b, 0x31, 0xe9, 0x0e, 0x07, 0x83, 0x8e, 0xa1, 0xd5, 0x77,
	0x50, 0x1d, 0x2a, 0x63, 0xdd, 0x9c, 0x0c, 0x74, 0xb3, 0xa3, 0x75, 0xcc, 0x4e, 0x3d, 0x17, 0xd1,
	0x34, 0xbd, 0xaf, 0x9b, 0xfa, 0x2d, 0x98, 0x47, 0x08, 0x6a, 0x97, 0x86, 0xa6, 0x7f, 0x9d, 0x68,
	0xc3, 0xee, 0xf5, 0x40, 0x37, 0xcc, 0xfa, 0x6e, 0x86, 0x98, 0x82, 0x52, 0xfb, 0x97, 0x04, 0x05,
	0x5e, 0x32, 0x3a, 0x01, 0xe9, 0x33, 0x71, 0x7c, 0xb4, 0xaf, 0xf2, 0x59, 0x8e, 0x7e, 0x8b, 0x77,
	0x92, 0x9f, 0xfd, 0x55, 0xad, 0x1e, 0x6d, 0x82, 0xb2, 0x83, 0xde, 0x40, 0xa1, 0x8f, 0xad, 0x1f,
	0x18, 0xa1, 0x58, 0xc5, 0x2f, 0x0f, 0xcb, 0xce, 0xe0, 0x49, 0x0f, 0x33, 0x83, 0xd8, 0x18, 0x6d,
	0x21, 0xc9, 0x47, 0xb1, 0xa1, 0xa0, 0x25, 0xb3, 0xa3, 0xec, 0xa0, 0x0f, 0x7c, 0x29, 0xba, 0x6e,
	0x48, 0x19, 0x5e, 0x6d, 0x95, 0x37, 0x52, 0xb9, 0x60, 0x66, 0x1c, 0x2e, 0xa0, 0x34, 0xf6, 0xad,
	0x80, 0x2e, 0x09, 0xdb, 0xaa, 0xdf, 0x5e, 0xbb, 0x0a, 0xbb, 0x3d, 0xcc, 0xd0, 0xbe, 0xe8, 0xf7,
	0xed, 0x82, 0xca, 0x28, 0x0b, 0xa5, 0xd9, 0xde, 0x41, 0x31, 0x9e, 0x7f, 0x74, 0x28, 0xe2, 0x1b,
	0xcb, 0x25, 0x1f, 0xdd, 0x41, 0x53, 0xe1, 0xdb, 0xa4, 0x33, 0x07, 0xd9, 0xd1, 0x7a, 0xf8, 0x71,
	0xdf, 0x43, 0x31, 0x1e, 0xf3, 0x34, 0xe1, 0xc6, 0xd4, 0x6f, 0x57, 0x7e, 0x6c, 0x7e, 0x7b, 0xb9,
	0x70, 0xd8, 0x32, 0x9c, 0xaa, 0x33, 0xe2, 0xb5, 0x3c, 0x42, 0xc3, 0x1b, 0xab, 0x35, 0x75, 0x2d,
	0xca, 0x5a, 0x9b, 0x1f, 0xd2, 0x69, 0x91, 0xdf, 0x4f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9f,
	0x80, 0xa6, 0x0d, 0x61, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IndexClient is the client API for Index service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IndexClient interface {
	Join(ctx context.Context, in *raft.JoinRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Leave(ctx context.Context, in *raft.LeaveRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetNode(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*raft.GetNodeResponse, error)
	GetCluster(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*raft.GetClusterResponse, error)
	Snapshot(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type indexClient struct {
	cc *grpc.ClientConn
}

func NewIndexClient(cc *grpc.ClientConn) IndexClient {
	return &indexClient{cc}
}

func (c *indexClient) Join(ctx context.Context, in *raft.JoinRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/index.Index/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) Leave(ctx context.Context, in *raft.LeaveRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/index.Index/Leave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) GetNode(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*raft.GetNodeResponse, error) {
	out := new(raft.GetNodeResponse)
	err := c.cc.Invoke(ctx, "/index.Index/GetNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) GetCluster(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*raft.GetClusterResponse, error) {
	out := new(raft.GetClusterResponse)
	err := c.cc.Invoke(ctx, "/index.Index/GetCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) Snapshot(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/index.Index/Snapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/index.Index/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/index.Index/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/index.Index/Index", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/index.Index/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexServer is the server API for Index service.
type IndexServer interface {
	Join(context.Context, *raft.JoinRequest) (*empty.Empty, error)
	Leave(context.Context, *raft.LeaveRequest) (*empty.Empty, error)
	GetNode(context.Context, *empty.Empty) (*raft.GetNodeResponse, error)
	GetCluster(context.Context, *empty.Empty) (*raft.GetClusterResponse, error)
	Snapshot(context.Context, *empty.Empty) (*empty.Empty, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	Index(context.Context, *IndexRequest) (*empty.Empty, error)
	Delete(context.Context, *DeleteRequest) (*empty.Empty, error)
}

func RegisterIndexServer(s *grpc.Server, srv IndexServer) {
	s.RegisterService(&_Index_serviceDesc, srv)
}

func _Index_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).Join(ctx, req.(*raft.JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.LeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).Leave(ctx, req.(*raft.LeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/GetNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).GetNode(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/GetCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).GetCluster(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_Snapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).Snapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/Snapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).Snapshot(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).Index(ctx, req.(*IndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Index_serviceDesc = grpc.ServiceDesc{
	ServiceName: "index.Index",
	HandlerType: (*IndexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _Index_Join_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _Index_Leave_Handler,
		},
		{
			MethodName: "GetNode",
			Handler:    _Index_GetNode_Handler,
		},
		{
			MethodName: "GetCluster",
			Handler:    _Index_GetCluster_Handler,
		},
		{
			MethodName: "Snapshot",
			Handler:    _Index_Snapshot_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Index_Get_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Index_Search_Handler,
		},
		{
			MethodName: "Index",
			Handler:    _Index_Index_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Index_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/index/index.proto",
}
