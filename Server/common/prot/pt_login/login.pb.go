// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pt_login/login.proto

package pt_login

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	pt_com "../pt_com"
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

type C2SLoginUser struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2SLoginUser) Reset()         { *m = C2SLoginUser{} }
func (m *C2SLoginUser) String() string { return proto.CompactTextString(m) }
func (*C2SLoginUser) ProtoMessage()    {}
func (*C2SLoginUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fb07df2e2d485ed, []int{0}
}

func (m *C2SLoginUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2SLoginUser.Unmarshal(m, b)
}
func (m *C2SLoginUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2SLoginUser.Marshal(b, m, deterministic)
}
func (m *C2SLoginUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2SLoginUser.Merge(m, src)
}
func (m *C2SLoginUser) XXX_Size() int {
	return xxx_messageInfo_C2SLoginUser.Size(m)
}
func (m *C2SLoginUser) XXX_DiscardUnknown() {
	xxx_messageInfo_C2SLoginUser.DiscardUnknown(m)
}

var xxx_messageInfo_C2SLoginUser proto.InternalMessageInfo

func (m *C2SLoginUser) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *C2SLoginUser) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type S2CLoginUserResult struct {
	Code                 int32                       `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Time                 int64                       `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Userid               uint64                      `protobuf:"varint,3,opt,name=userid,proto3" json:"userid,omitempty"`
	Account              string                      `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
	Password             string                      `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Gold                 uint64                      `protobuf:"varint,6,opt,name=gold,proto3" json:"gold,omitempty"`
	Rooms                map[uint64]*pt_com.RoomItem `protobuf:"bytes,7,rep,name=rooms,proto3" json:"rooms,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *S2CLoginUserResult) Reset()         { *m = S2CLoginUserResult{} }
func (m *S2CLoginUserResult) String() string { return proto.CompactTextString(m) }
func (*S2CLoginUserResult) ProtoMessage()    {}
func (*S2CLoginUserResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fb07df2e2d485ed, []int{1}
}

func (m *S2CLoginUserResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CLoginUserResult.Unmarshal(m, b)
}
func (m *S2CLoginUserResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CLoginUserResult.Marshal(b, m, deterministic)
}
func (m *S2CLoginUserResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CLoginUserResult.Merge(m, src)
}
func (m *S2CLoginUserResult) XXX_Size() int {
	return xxx_messageInfo_S2CLoginUserResult.Size(m)
}
func (m *S2CLoginUserResult) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CLoginUserResult.DiscardUnknown(m)
}

var xxx_messageInfo_S2CLoginUserResult proto.InternalMessageInfo

func (m *S2CLoginUserResult) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *S2CLoginUserResult) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *S2CLoginUserResult) GetUserid() uint64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *S2CLoginUserResult) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *S2CLoginUserResult) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *S2CLoginUserResult) GetGold() uint64 {
	if m != nil {
		return m.Gold
	}
	return 0
}

func (m *S2CLoginUserResult) GetRooms() map[uint64]*pt_com.RoomItem {
	if m != nil {
		return m.Rooms
	}
	return nil
}

type C2SRegisterUser struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2SRegisterUser) Reset()         { *m = C2SRegisterUser{} }
func (m *C2SRegisterUser) String() string { return proto.CompactTextString(m) }
func (*C2SRegisterUser) ProtoMessage()    {}
func (*C2SRegisterUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fb07df2e2d485ed, []int{2}
}

func (m *C2SRegisterUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2SRegisterUser.Unmarshal(m, b)
}
func (m *C2SRegisterUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2SRegisterUser.Marshal(b, m, deterministic)
}
func (m *C2SRegisterUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2SRegisterUser.Merge(m, src)
}
func (m *C2SRegisterUser) XXX_Size() int {
	return xxx_messageInfo_C2SRegisterUser.Size(m)
}
func (m *C2SRegisterUser) XXX_DiscardUnknown() {
	xxx_messageInfo_C2SRegisterUser.DiscardUnknown(m)
}

var xxx_messageInfo_C2SRegisterUser proto.InternalMessageInfo

func (m *C2SRegisterUser) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *C2SRegisterUser) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type S2CRegisterUserResult struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Account              string   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2CRegisterUserResult) Reset()         { *m = S2CRegisterUserResult{} }
func (m *S2CRegisterUserResult) String() string { return proto.CompactTextString(m) }
func (*S2CRegisterUserResult) ProtoMessage()    {}
func (*S2CRegisterUserResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fb07df2e2d485ed, []int{3}
}

func (m *S2CRegisterUserResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CRegisterUserResult.Unmarshal(m, b)
}
func (m *S2CRegisterUserResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CRegisterUserResult.Marshal(b, m, deterministic)
}
func (m *S2CRegisterUserResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CRegisterUserResult.Merge(m, src)
}
func (m *S2CRegisterUserResult) XXX_Size() int {
	return xxx_messageInfo_S2CRegisterUserResult.Size(m)
}
func (m *S2CRegisterUserResult) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CRegisterUserResult.DiscardUnknown(m)
}

var xxx_messageInfo_S2CRegisterUserResult proto.InternalMessageInfo

func (m *S2CRegisterUserResult) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *S2CRegisterUserResult) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *S2CRegisterUserResult) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*C2SLoginUser)(nil), "pt_login.c2s_login_user")
	proto.RegisterType((*S2CLoginUserResult)(nil), "pt_login.s2c_login_user_result")
	proto.RegisterMapType((map[uint64]*pt_com.RoomItem)(nil), "pt_login.s2c_login_user_result.RoomsEntry")
	proto.RegisterType((*C2SRegisterUser)(nil), "pt_login.c2s_register_user")
	proto.RegisterType((*S2CRegisterUserResult)(nil), "pt_login.s2c_register_user_result")
}

func init() { proto.RegisterFile("pt_login/login.proto", fileDescriptor_2fb07df2e2d485ed) }

var fileDescriptor_2fb07df2e2d485ed = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x25, 0x5f, 0x6d, 0x9d, 0x82, 0xd8, 0xf5, 0x83, 0xa5, 0xa7, 0xd2, 0x8b, 0xc5, 0xc3, 0x16,
	0xe2, 0x45, 0x3c, 0x79, 0x51, 0x10, 0x6f, 0xfb, 0x07, 0x42, 0xdd, 0x2c, 0x21, 0x98, 0xcd, 0x84,
	0xdd, 0x8d, 0xd2, 0x7f, 0xe5, 0x4f, 0x94, 0xdd, 0x34, 0x9a, 0x80, 0xf5, 0xe2, 0x25, 0xbc, 0x99,
	0xd9, 0x79, 0xef, 0xcd, 0x23, 0x70, 0xd1, 0xd8, 0xac, 0xc2, 0xa2, 0xac, 0xb7, 0xfe, 0xcb, 0x1a,
	0x8d, 0x16, 0xc9, 0xac, 0xef, 0x2e, 0xcf, 0x1b, 0x9b, 0x09, 0x54, 0x5b, 0x81, 0x4a, 0xe1, 0x61,
	0xbc, 0x7e, 0x82, 0x53, 0x91, 0x9a, 0xee, 0x45, 0xd6, 0x1a, 0xa9, 0x09, 0x85, 0xe9, 0x4e, 0x08,
	0x6c, 0x6b, 0x4b, 0x83, 0x55, 0xb0, 0x39, 0xe1, 0x7d, 0x49, 0x96, 0x30, 0x6b, 0x76, 0xc6, 0x7c,
	0xa0, 0xce, 0x69, 0xe8, 0x47, 0xdf, 0xf5, 0xfa, 0x33, 0x84, 0x4b, 0x93, 0x8a, 0x01, 0x51, 0xa6,
	0xa5, 0x69, 0x2b, 0x4b, 0x08, 0xc4, 0x02, 0x73, 0xe9, 0xc9, 0x12, 0xee, 0xb1, 0xeb, 0xd9, 0x52,
	0x49, 0xcf, 0x12, 0x71, 0x8f, 0xc9, 0x15, 0x4c, 0xdc, 0x5a, 0x99, 0xd3, 0x68, 0x15, 0x6c, 0x62,
	0x7e, 0xa8, 0x86, 0x7e, 0xe2, 0xe3, 0x7e, 0x92, 0xb1, 0x1f, 0xa7, 0x50, 0x60, 0x95, 0xd3, 0x89,
	0xe7, 0xf2, 0x98, 0x3c, 0x40, 0xa2, 0x11, 0x95, 0xa1, 0xd3, 0x55, 0xb4, 0x99, 0xa7, 0x37, 0xac,
	0x8f, 0x86, 0xfd, 0xea, 0x9c, 0x71, 0xf7, 0xf8, 0xb1, 0xb6, 0x7a, 0xcf, 0xbb, 0xc5, 0xe5, 0x0b,
	0xc0, 0x4f, 0x93, 0x9c, 0x41, 0xf4, 0x26, 0xf7, 0xfe, 0xb0, 0x98, 0x3b, 0x48, 0xae, 0x21, 0x79,
	0xdf, 0x55, 0x6d, 0x77, 0xd8, 0x3c, 0x5d, 0xb0, 0x2e, 0x72, 0xe6, 0xb6, 0xb3, 0xd2, 0x4a, 0xc5,
	0xbb, 0xf9, 0x7d, 0x78, 0x17, 0xac, 0x9f, 0x61, 0xe1, 0xa2, 0xd7, 0xb2, 0x28, 0x8d, 0x95, 0xfa,
	0x3f, 0xe9, 0xe7, 0x40, 0xdd, 0x09, 0x23, 0xaa, 0xbf, 0xf2, 0x1f, 0xa8, 0x84, 0xc7, 0x55, 0xa2,
	0xb1, 0xca, 0xeb, 0xc4, 0xff, 0x32, 0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x31, 0xc2,
	0x15, 0x69, 0x02, 0x00, 0x00,
}
