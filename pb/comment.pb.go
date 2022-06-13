// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/comment.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CommentActionRequest struct {
	VideoID              int64    `protobuf:"varint,1,opt,name=videoID,proto3" json:"videoID,omitempty"`
	UserID               int64    `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	ActionType           int64    `protobuf:"varint,3,opt,name=ActionType,proto3" json:"ActionType,omitempty"`
	CommentText          string   `protobuf:"bytes,4,opt,name=CommentText,proto3" json:"CommentText,omitempty"`
	CommentID            int64    `protobuf:"varint,5,opt,name=CommentID,proto3" json:"CommentID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommentActionRequest) Reset()         { *m = CommentActionRequest{} }
func (m *CommentActionRequest) String() string { return proto.CompactTextString(m) }
func (*CommentActionRequest) ProtoMessage()    {}
func (*CommentActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1662d110fc169bd4, []int{0}
}

func (m *CommentActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommentActionRequest.Unmarshal(m, b)
}
func (m *CommentActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommentActionRequest.Marshal(b, m, deterministic)
}
func (m *CommentActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentActionRequest.Merge(m, src)
}
func (m *CommentActionRequest) XXX_Size() int {
	return xxx_messageInfo_CommentActionRequest.Size(m)
}
func (m *CommentActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommentActionRequest proto.InternalMessageInfo

func (m *CommentActionRequest) GetVideoID() int64 {
	if m != nil {
		return m.VideoID
	}
	return 0
}

func (m *CommentActionRequest) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *CommentActionRequest) GetActionType() int64 {
	if m != nil {
		return m.ActionType
	}
	return 0
}

func (m *CommentActionRequest) GetCommentText() string {
	if m != nil {
		return m.CommentText
	}
	return ""
}

func (m *CommentActionRequest) GetCommentID() int64 {
	if m != nil {
		return m.CommentID
	}
	return 0
}

type CommentActionResponse struct {
	CommentID            int64    `protobuf:"varint,1,opt,name=CommentID,proto3" json:"CommentID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommentActionResponse) Reset()         { *m = CommentActionResponse{} }
func (m *CommentActionResponse) String() string { return proto.CompactTextString(m) }
func (*CommentActionResponse) ProtoMessage()    {}
func (*CommentActionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1662d110fc169bd4, []int{1}
}

func (m *CommentActionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommentActionResponse.Unmarshal(m, b)
}
func (m *CommentActionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommentActionResponse.Marshal(b, m, deterministic)
}
func (m *CommentActionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentActionResponse.Merge(m, src)
}
func (m *CommentActionResponse) XXX_Size() int {
	return xxx_messageInfo_CommentActionResponse.Size(m)
}
func (m *CommentActionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentActionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommentActionResponse proto.InternalMessageInfo

func (m *CommentActionResponse) GetCommentID() int64 {
	if m != nil {
		return m.CommentID
	}
	return 0
}

type CommentListRequest struct {
	VideoID              int64    `protobuf:"varint,1,opt,name=videoID,proto3" json:"videoID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommentListRequest) Reset()         { *m = CommentListRequest{} }
func (m *CommentListRequest) String() string { return proto.CompactTextString(m) }
func (*CommentListRequest) ProtoMessage()    {}
func (*CommentListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1662d110fc169bd4, []int{2}
}

func (m *CommentListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommentListRequest.Unmarshal(m, b)
}
func (m *CommentListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommentListRequest.Marshal(b, m, deterministic)
}
func (m *CommentListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentListRequest.Merge(m, src)
}
func (m *CommentListRequest) XXX_Size() int {
	return xxx_messageInfo_CommentListRequest.Size(m)
}
func (m *CommentListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommentListRequest proto.InternalMessageInfo

func (m *CommentListRequest) GetVideoID() int64 {
	if m != nil {
		return m.VideoID
	}
	return 0
}

type CommentListResponse struct {
	List                 []*Comment `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CommentListResponse) Reset()         { *m = CommentListResponse{} }
func (m *CommentListResponse) String() string { return proto.CompactTextString(m) }
func (*CommentListResponse) ProtoMessage()    {}
func (*CommentListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1662d110fc169bd4, []int{3}
}

func (m *CommentListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommentListResponse.Unmarshal(m, b)
}
func (m *CommentListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommentListResponse.Marshal(b, m, deterministic)
}
func (m *CommentListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentListResponse.Merge(m, src)
}
func (m *CommentListResponse) XXX_Size() int {
	return xxx_messageInfo_CommentListResponse.Size(m)
}
func (m *CommentListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommentListResponse proto.InternalMessageInfo

func (m *CommentListResponse) GetList() []*Comment {
	if m != nil {
		return m.List
	}
	return nil
}

type Comment struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	VideoId              int64    `protobuf:"varint,3,opt,name=VideoId,proto3" json:"VideoId,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	CreateDate           string   `protobuf:"bytes,5,opt,name=createDate,proto3" json:"createDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}
func (*Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor_1662d110fc169bd4, []int{4}
}

func (m *Comment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Comment.Unmarshal(m, b)
}
func (m *Comment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Comment.Marshal(b, m, deterministic)
}
func (m *Comment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment.Merge(m, src)
}
func (m *Comment) XXX_Size() int {
	return xxx_messageInfo_Comment.Size(m)
}
func (m *Comment) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment.DiscardUnknown(m)
}

var xxx_messageInfo_Comment proto.InternalMessageInfo

func (m *Comment) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Comment) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Comment) GetVideoId() int64 {
	if m != nil {
		return m.VideoId
	}
	return 0
}

func (m *Comment) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Comment) GetCreateDate() string {
	if m != nil {
		return m.CreateDate
	}
	return ""
}

func init() {
	proto.RegisterType((*CommentActionRequest)(nil), "pb.CommentActionRequest")
	proto.RegisterType((*CommentActionResponse)(nil), "pb.CommentActionResponse")
	proto.RegisterType((*CommentListRequest)(nil), "pb.CommentListRequest")
	proto.RegisterType((*CommentListResponse)(nil), "pb.CommentListResponse")
	proto.RegisterType((*Comment)(nil), "pb.Comment")
}

func init() {
	proto.RegisterFile("pb/comment.proto", fileDescriptor_1662d110fc169bd4)
}

var fileDescriptor_1662d110fc169bd4 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x99, 0x34, 0x5f, 0x4a, 0x6f, 0xa1, 0x7c, 0x8c, 0x5a, 0xc7, 0x22, 0x1a, 0xb2, 0xea,
	0x2a, 0x42, 0x45, 0x5d, 0x88, 0x0b, 0x35, 0x9b, 0x80, 0xab, 0x58, 0x5d, 0xb8, 0x6b, 0x92, 0xbb,
	0x08, 0xd8, 0xcc, 0x98, 0x99, 0x16, 0x7d, 0x00, 0xc1, 0x47, 0xf1, 0x31, 0x65, 0x26, 0x13, 0x9c,
	0x56, 0xc1, 0xe5, 0x39, 0xe7, 0xfe, 0x9b, 0x5f, 0x02, 0xff, 0x45, 0x7e, 0x52, 0xf0, 0xe5, 0x12,
	0x6b, 0x15, 0x8b, 0x86, 0x2b, 0x4e, 0x3d, 0x91, 0x47, 0x9f, 0x04, 0x76, 0x6f, 0x5b, 0xf7, 0xba,
	0x50, 0x15, 0xaf, 0x33, 0x7c, 0x59, 0xa1, 0x54, 0x94, 0x41, 0x7f, 0x5d, 0x95, 0xc8, 0xd3, 0x84,
	0x91, 0x90, 0x4c, 0x7b, 0x59, 0x27, 0xe9, 0x18, 0x82, 0x95, 0xc4, 0x26, 0x4d, 0x98, 0x67, 0x02,
	0xab, 0xe8, 0x11, 0x40, 0x3b, 0x62, 0xfe, 0x26, 0x90, 0xf5, 0x4c, 0xe6, 0x38, 0x34, 0x84, 0xa1,
	0xdd, 0x34, 0xc7, 0x57, 0xc5, 0xfc, 0x90, 0x4c, 0x07, 0x99, 0x6b, 0xd1, 0x43, 0x18, 0x58, 0x99,
	0x26, 0xec, 0x9f, 0x19, 0xf0, 0x6d, 0x44, 0x67, 0xb0, 0xb7, 0x75, 0xa9, 0x14, 0xbc, 0x96, 0xb8,
	0xd9, 0x46, 0xb6, 0xdb, 0x62, 0xa0, 0x56, 0xdc, 0x55, 0x52, 0xfd, 0xf9, 0xbc, 0xe8, 0x1c, 0x76,
	0x36, 0xea, 0xed, 0x92, 0x63, 0xf0, 0x9f, 0x2b, 0xa9, 0x18, 0x09, 0x7b, 0xd3, 0xe1, 0x6c, 0x18,
	0x8b, 0x3c, 0xb6, 0x65, 0x99, 0x09, 0xa2, 0x77, 0x02, 0x7d, 0xeb, 0xd0, 0x11, 0x78, 0x69, 0x69,
	0x07, 0x7b, 0x69, 0xa9, 0x91, 0x3d, 0x68, 0x48, 0x65, 0x87, 0xac, 0x55, 0xfa, 0x8a, 0x47, 0xb3,
	0xb6, 0xb4, 0xbc, 0x3a, 0xa9, 0x93, 0x82, 0xd7, 0x0a, 0xeb, 0x0e, 0x54, 0x27, 0x35, 0xe6, 0xa2,
	0xc1, 0x85, 0xc2, 0x64, 0xa1, 0xd0, 0x50, 0x1a, 0x64, 0x8e, 0x33, 0xfb, 0x20, 0x30, 0xb2, 0x77,
	0xdc, 0x63, 0xb3, 0xae, 0x0a, 0xa4, 0x57, 0x10, 0xb4, 0xc8, 0x28, 0x73, 0xee, 0xde, 0xf8, 0xde,
	0x93, 0x83, 0x5f, 0x12, 0xfb, 0xf4, 0x0b, 0xf0, 0x35, 0x0a, 0x3a, 0x76, 0x4a, 0x1c, 0x96, 0x93,
	0xfd, 0x1f, 0x7e, 0xdb, 0x78, 0x13, 0x3c, 0xf9, 0xf1, 0xa5, 0xc8, 0xf3, 0xc0, 0xfc, 0x6f, 0xa7,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xa1, 0x98, 0xdc, 0x83, 0x02, 0x00, 0x00,
}
