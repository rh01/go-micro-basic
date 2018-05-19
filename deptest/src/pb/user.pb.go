// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=Address" json:"Address,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=Phone" json:"Phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d97a7407b988ae67, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type ModifyUserReq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=Address" json:"Address,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=Phone" json:"Phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyUserReq) Reset()         { *m = ModifyUserReq{} }
func (m *ModifyUserReq) String() string { return proto.CompactTextString(m) }
func (*ModifyUserReq) ProtoMessage()    {}
func (*ModifyUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d97a7407b988ae67, []int{1}
}
func (m *ModifyUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyUserReq.Unmarshal(m, b)
}
func (m *ModifyUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyUserReq.Marshal(b, m, deterministic)
}
func (dst *ModifyUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyUserReq.Merge(dst, src)
}
func (m *ModifyUserReq) XXX_Size() int {
	return xxx_messageInfo_ModifyUserReq.Size(m)
}
func (m *ModifyUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyUserReq proto.InternalMessageInfo

func (m *ModifyUserReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ModifyUserReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModifyUserReq) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ModifyUserReq) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type ModifyUserResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyUserResp) Reset()         { *m = ModifyUserResp{} }
func (m *ModifyUserResp) String() string { return proto.CompactTextString(m) }
func (*ModifyUserResp) ProtoMessage()    {}
func (*ModifyUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d97a7407b988ae67, []int{2}
}
func (m *ModifyUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyUserResp.Unmarshal(m, b)
}
func (m *ModifyUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyUserResp.Marshal(b, m, deterministic)
}
func (dst *ModifyUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyUserResp.Merge(dst, src)
}
func (m *ModifyUserResp) XXX_Size() int {
	return xxx_messageInfo_ModifyUserResp.Size(m)
}
func (m *ModifyUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyUserResp proto.InternalMessageInfo

type SelectUserReq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SelectUserReq) Reset()         { *m = SelectUserReq{} }
func (m *SelectUserReq) String() string { return proto.CompactTextString(m) }
func (*SelectUserReq) ProtoMessage()    {}
func (*SelectUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d97a7407b988ae67, []int{3}
}
func (m *SelectUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelectUserReq.Unmarshal(m, b)
}
func (m *SelectUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelectUserReq.Marshal(b, m, deterministic)
}
func (dst *SelectUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelectUserReq.Merge(dst, src)
}
func (m *SelectUserReq) XXX_Size() int {
	return xxx_messageInfo_SelectUserReq.Size(m)
}
func (m *SelectUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SelectUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_SelectUserReq proto.InternalMessageInfo

func (m *SelectUserReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SelectUserResp struct {
	Users                *User    `protobuf:"bytes,1,opt,name=users" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SelectUserResp) Reset()         { *m = SelectUserResp{} }
func (m *SelectUserResp) String() string { return proto.CompactTextString(m) }
func (*SelectUserResp) ProtoMessage()    {}
func (*SelectUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d97a7407b988ae67, []int{4}
}
func (m *SelectUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelectUserResp.Unmarshal(m, b)
}
func (m *SelectUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelectUserResp.Marshal(b, m, deterministic)
}
func (dst *SelectUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelectUserResp.Merge(dst, src)
}
func (m *SelectUserResp) XXX_Size() int {
	return xxx_messageInfo_SelectUserResp.Size(m)
}
func (m *SelectUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SelectUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_SelectUserResp proto.InternalMessageInfo

func (m *SelectUserResp) GetUsers() *User {
	if m != nil {
		return m.Users
	}
	return nil
}

type InsertUserReq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=Address" json:"Address,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=Phone" json:"Phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InsertUserReq) Reset()         { *m = InsertUserReq{} }
func (m *InsertUserReq) String() string { return proto.CompactTextString(m) }
func (*InsertUserReq) ProtoMessage()    {}
func (*InsertUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d97a7407b988ae67, []int{5}
}
func (m *InsertUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertUserReq.Unmarshal(m, b)
}
func (m *InsertUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertUserReq.Marshal(b, m, deterministic)
}
func (dst *InsertUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertUserReq.Merge(dst, src)
}
func (m *InsertUserReq) XXX_Size() int {
	return xxx_messageInfo_InsertUserReq.Size(m)
}
func (m *InsertUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_InsertUserReq proto.InternalMessageInfo

func (m *InsertUserReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *InsertUserReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InsertUserReq) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *InsertUserReq) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type InsertUserResp struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=Address" json:"Address,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=Phone" json:"Phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InsertUserResp) Reset()         { *m = InsertUserResp{} }
func (m *InsertUserResp) String() string { return proto.CompactTextString(m) }
func (*InsertUserResp) ProtoMessage()    {}
func (*InsertUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d97a7407b988ae67, []int{6}
}
func (m *InsertUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertUserResp.Unmarshal(m, b)
}
func (m *InsertUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertUserResp.Marshal(b, m, deterministic)
}
func (dst *InsertUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertUserResp.Merge(dst, src)
}
func (m *InsertUserResp) XXX_Size() int {
	return xxx_messageInfo_InsertUserResp.Size(m)
}
func (m *InsertUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_InsertUserResp proto.InternalMessageInfo

func (m *InsertUserResp) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *InsertUserResp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InsertUserResp) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *InsertUserResp) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type DeleteUserReq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserReq) Reset()         { *m = DeleteUserReq{} }
func (m *DeleteUserReq) String() string { return proto.CompactTextString(m) }
func (*DeleteUserReq) ProtoMessage()    {}
func (*DeleteUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d97a7407b988ae67, []int{7}
}
func (m *DeleteUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserReq.Unmarshal(m, b)
}
func (m *DeleteUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserReq.Marshal(b, m, deterministic)
}
func (dst *DeleteUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserReq.Merge(dst, src)
}
func (m *DeleteUserReq) XXX_Size() int {
	return xxx_messageInfo_DeleteUserReq.Size(m)
}
func (m *DeleteUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserReq proto.InternalMessageInfo

func (m *DeleteUserReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteUserResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserResp) Reset()         { *m = DeleteUserResp{} }
func (m *DeleteUserResp) String() string { return proto.CompactTextString(m) }
func (*DeleteUserResp) ProtoMessage()    {}
func (*DeleteUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d97a7407b988ae67, []int{8}
}
func (m *DeleteUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserResp.Unmarshal(m, b)
}
func (m *DeleteUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserResp.Marshal(b, m, deterministic)
}
func (dst *DeleteUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserResp.Merge(dst, src)
}
func (m *DeleteUserResp) XXX_Size() int {
	return xxx_messageInfo_DeleteUserResp.Size(m)
}
func (m *DeleteUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*ModifyUserReq)(nil), "pb.ModifyUserReq")
	proto.RegisterType((*ModifyUserResp)(nil), "pb.ModifyUserResp")
	proto.RegisterType((*SelectUserReq)(nil), "pb.SelectUserReq")
	proto.RegisterType((*SelectUserResp)(nil), "pb.SelectUserResp")
	proto.RegisterType((*InsertUserReq)(nil), "pb.InsertUserReq")
	proto.RegisterType((*InsertUserResp)(nil), "pb.InsertUserResp")
	proto.RegisterType((*DeleteUserReq)(nil), "pb.DeleteUserReq")
	proto.RegisterType((*DeleteUserResp)(nil), "pb.DeleteUserResp")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_d97a7407b988ae67) }

var fileDescriptor_user_d97a7407b988ae67 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x9b, 0x98, 0xf8, 0x67, 0x4a, 0x42, 0x1d, 0x3c, 0x04, 0x0f, 0xb6, 0xe4, 0xd4, 0x53,
	0x90, 0x8a, 0x0f, 0x20, 0x78, 0xf1, 0x20, 0x48, 0x8a, 0x17, 0x6f, 0x26, 0x3b, 0x62, 0xa0, 0x26,
	0xeb, 0x4e, 0x14, 0x7c, 0x64, 0xdf, 0x42, 0x76, 0x07, 0x49, 0xb6, 0xd0, 0x63, 0x6e, 0x99, 0x2f,
	0xf3, 0xcb, 0xb7, 0xfb, 0xcd, 0x04, 0xe0, 0x8b, 0xc9, 0x14, 0xda, 0x74, 0x7d, 0x87, 0xa1, 0xae,
	0xf2, 0x17, 0x88, 0x9e, 0x99, 0x0c, 0xa6, 0x10, 0x36, 0x2a, 0x0b, 0x56, 0xc1, 0x3a, 0x2e, 0xc3,
	0x46, 0x21, 0x42, 0xd4, 0xbe, 0x7e, 0x50, 0x16, 0xae, 0x82, 0xf5, 0x59, 0xe9, 0x9e, 0x31, 0x83,
	0x93, 0x3b, 0xa5, 0x0c, 0x31, 0x67, 0x47, 0x4e, 0xfe, 0x2f, 0xf1, 0x02, 0xe2, 0xa7, 0xf7, 0xae,
	0xa5, 0x2c, 0x72, 0xba, 0x14, 0x79, 0x0d, 0xc9, 0x63, 0xa7, 0x9a, 0xb7, 0x1f, 0xeb, 0x50, 0xd2,
	0xe7, 0x24, 0x26, 0x0b, 0x48, 0xc7, 0x26, 0xac, 0xf3, 0x25, 0x24, 0x5b, 0xda, 0x51, 0xdd, 0x1f,
	0xb0, 0xcd, 0xaf, 0x21, 0x1d, 0x37, 0xb0, 0xc6, 0x2b, 0x88, 0x6d, 0x2e, 0xec, 0x9a, 0xe6, 0x9b,
	0xd3, 0x42, 0x57, 0x85, 0x7b, 0x29, 0xb2, 0xbd, 0xc9, 0x43, 0xcb, 0x64, 0xfa, 0x29, 0x6f, 0xa2,
	0x20, 0x1d, 0x9b, 0xb0, 0x9e, 0xc4, 0x65, 0x09, 0xc9, 0x3d, 0xed, 0xa8, 0xa7, 0x43, 0xe9, 0x2c,
	0x20, 0x1d, 0x37, 0xb0, 0xde, 0xfc, 0x06, 0x30, 0xb7, 0xc5, 0x96, 0xcc, 0x77, 0x53, 0x13, 0xde,
	0x02, 0x0c, 0x07, 0xc5, 0x73, 0x1b, 0x96, 0x97, 0xce, 0x25, 0xee, 0x4b, 0xac, 0xf3, 0x99, 0xc5,
	0x86, 0x0f, 0x0b, 0xe6, 0x9d, 0x44, 0x30, 0xdf, 0x5b, 0xb0, 0x61, 0x5a, 0x82, 0x79, 0xe3, 0x15,
	0xcc, 0x1f, 0xa8, 0x60, 0xc3, 0x5e, 0x08, 0xe6, 0x2d, 0xa3, 0x60, 0x7b, 0xab, 0x33, 0xab, 0x8e,
	0xdd, 0xaf, 0x71, 0xf3, 0x17, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xe8, 0xf2, 0xcc, 0x28, 0x03, 0x00,
	0x00,
}
