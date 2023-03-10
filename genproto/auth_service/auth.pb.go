// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

//

package auth_service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/struct"
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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Phone                string   `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Username             string   `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	CreatedAt            string   `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *User) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type HasAccessModel struct {
	UserId               string   `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HasAccessModel) Reset()         { *m = HasAccessModel{} }
func (m *HasAccessModel) String() string { return proto.CompactTextString(m) }
func (*HasAccessModel) ProtoMessage()    {}
func (*HasAccessModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *HasAccessModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasAccessModel.Unmarshal(m, b)
}
func (m *HasAccessModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasAccessModel.Marshal(b, m, deterministic)
}
func (m *HasAccessModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasAccessModel.Merge(m, src)
}
func (m *HasAccessModel) XXX_Size() int {
	return xxx_messageInfo_HasAccessModel.Size(m)
}
func (m *HasAccessModel) XXX_DiscardUnknown() {
	xxx_messageInfo_HasAccessModel.DiscardUnknown(m)
}

var xxx_messageInfo_HasAccessModel proto.InternalMessageInfo

func (m *HasAccessModel) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "auth_service.User")
	proto.RegisterType((*HasAccessModel)(nil), "auth_service.HasAccessModel")
}

func init() {
	proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874)
}

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd0, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x07, 0x70, 0x5a, 0xb7, 0xae, 0x7d, 0xc8, 0x0e, 0x41, 0x59, 0x98, 0x0a, 0xb2, 0x93, 0x5e,
	0xec, 0xc1, 0x4f, 0x50, 0x4f, 0x7a, 0xd0, 0x83, 0xe0, 0xc5, 0x4b, 0xc9, 0x92, 0xb7, 0x2e, 0xd0,
	0x35, 0x25, 0x2f, 0xd1, 0x6f, 0xeb, 0x67, 0x91, 0x24, 0x2d, 0xec, 0xd6, 0xff, 0xff, 0x57, 0xf8,
	0x93, 0x07, 0x20, 0xbc, 0x3b, 0x3e, 0x8d, 0xd6, 0x38, 0xc3, 0x2e, 0xc3, 0x77, 0x4b, 0x68, 0x7f,
	0xb4, 0xc4, 0xed, 0x6d, 0x67, 0x4c, 0xd7, 0x63, 0x1d, 0x6d, 0xef, 0x0f, 0x35, 0x39, 0xeb, 0xa5,
	0x4b, 0xff, 0xee, 0xfe, 0x32, 0x58, 0x7c, 0x11, 0x5a, 0xb6, 0x86, 0x5c, 0x2b, 0x9e, 0xdd, 0x67,
	0x0f, 0xd5, 0x67, 0xae, 0x15, 0xbb, 0x82, 0xe5, 0x78, 0x34, 0x03, 0xf2, 0x3c, 0x56, 0x29, 0xb0,
	0x3b, 0x80, 0x83, 0xb6, 0xe4, 0xda, 0x41, 0x9c, 0x90, 0x5f, 0x44, 0xaa, 0x62, 0xf3, 0x21, 0x4e,
	0xc8, 0x6e, 0xa0, 0xea, 0xc5, 0xac, 0x8b, 0xa8, 0x65, 0x28, 0x22, 0x6e, 0xa1, 0xf4, 0x84, 0x36,
	0xda, 0x32, 0xd9, 0x9c, 0x83, 0x8d, 0x82, 0xe8, 0xd7, 0x58, 0xc5, 0x8b, 0x64, 0x73, 0x0e, 0x9b,
	0xd2, 0xa2, 0x70, 0xa8, 0x5a, 0xe1, 0xf8, 0x2a, 0x6d, 0x4e, 0x4d, 0xe3, 0x02, 0xfb, 0x51, 0xcd,
	0x5c, 0x26, 0x9e, 0x9a, 0xc6, 0xed, 0x1e, 0x61, 0xfd, 0x2a, 0xa8, 0x91, 0x12, 0x89, 0xde, 0x8d,
	0xc2, 0x9e, 0x6d, 0x60, 0x15, 0x76, 0x5b, 0xad, 0xa6, 0x07, 0x14, 0x21, 0xbe, 0xa9, 0x97, 0xcd,
	0xf7, 0x75, 0x87, 0x43, 0xbc, 0x4b, 0x7d, 0x7e, 0xc2, 0x7d, 0x11, 0xbb, 0xe7, 0xff, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x5d, 0xf3, 0xea, 0x18, 0x65, 0x01, 0x00, 0x00,
}
