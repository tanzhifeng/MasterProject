// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pt_logical/login.proto

package pt_logical

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

type S2LLoginUser struct {
	Loginsocket          uint32   `protobuf:"varint,1,opt,name=loginsocket,proto3" json:"loginsocket,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Account              string   `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2LLoginUser) Reset()         { *m = S2LLoginUser{} }
func (m *S2LLoginUser) String() string { return proto.CompactTextString(m) }
func (*S2LLoginUser) ProtoMessage()    {}
func (*S2LLoginUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_1686212162667a44, []int{0}
}

func (m *S2LLoginUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2LLoginUser.Unmarshal(m, b)
}
func (m *S2LLoginUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2LLoginUser.Marshal(b, m, deterministic)
}
func (m *S2LLoginUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2LLoginUser.Merge(m, src)
}
func (m *S2LLoginUser) XXX_Size() int {
	return xxx_messageInfo_S2LLoginUser.Size(m)
}
func (m *S2LLoginUser) XXX_DiscardUnknown() {
	xxx_messageInfo_S2LLoginUser.DiscardUnknown(m)
}

var xxx_messageInfo_S2LLoginUser proto.InternalMessageInfo

func (m *S2LLoginUser) GetLoginsocket() uint32 {
	if m != nil {
		return m.Loginsocket
	}
	return 0
}

func (m *S2LLoginUser) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *S2LLoginUser) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *S2LLoginUser) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type L2SLoginUserResult struct {
	Loginsocket          uint32                      `protobuf:"varint,1,opt,name=loginsocket,proto3" json:"loginsocket,omitempty"`
	Address              string                      `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Code                 int32                       `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	Userid               uint64                      `protobuf:"varint,4,opt,name=userid,proto3" json:"userid,omitempty"`
	Account              string                      `protobuf:"bytes,5,opt,name=account,proto3" json:"account,omitempty"`
	Password             string                      `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	Gold                 uint64                      `protobuf:"varint,7,opt,name=gold,proto3" json:"gold,omitempty"`
	Rooms                map[uint64]*pt_com.RoomItem `protobuf:"bytes,8,rep,name=rooms,proto3" json:"rooms,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *L2SLoginUserResult) Reset()         { *m = L2SLoginUserResult{} }
func (m *L2SLoginUserResult) String() string { return proto.CompactTextString(m) }
func (*L2SLoginUserResult) ProtoMessage()    {}
func (*L2SLoginUserResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_1686212162667a44, []int{1}
}

func (m *L2SLoginUserResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2SLoginUserResult.Unmarshal(m, b)
}
func (m *L2SLoginUserResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2SLoginUserResult.Marshal(b, m, deterministic)
}
func (m *L2SLoginUserResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2SLoginUserResult.Merge(m, src)
}
func (m *L2SLoginUserResult) XXX_Size() int {
	return xxx_messageInfo_L2SLoginUserResult.Size(m)
}
func (m *L2SLoginUserResult) XXX_DiscardUnknown() {
	xxx_messageInfo_L2SLoginUserResult.DiscardUnknown(m)
}

var xxx_messageInfo_L2SLoginUserResult proto.InternalMessageInfo

func (m *L2SLoginUserResult) GetLoginsocket() uint32 {
	if m != nil {
		return m.Loginsocket
	}
	return 0
}

func (m *L2SLoginUserResult) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *L2SLoginUserResult) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *L2SLoginUserResult) GetUserid() uint64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *L2SLoginUserResult) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *L2SLoginUserResult) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *L2SLoginUserResult) GetGold() uint64 {
	if m != nil {
		return m.Gold
	}
	return 0
}

func (m *L2SLoginUserResult) GetRooms() map[uint64]*pt_com.RoomItem {
	if m != nil {
		return m.Rooms
	}
	return nil
}

type S2LUserLost struct {
	Loginsocket          uint32   `protobuf:"varint,1,opt,name=loginsocket,proto3" json:"loginsocket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2LUserLost) Reset()         { *m = S2LUserLost{} }
func (m *S2LUserLost) String() string { return proto.CompactTextString(m) }
func (*S2LUserLost) ProtoMessage()    {}
func (*S2LUserLost) Descriptor() ([]byte, []int) {
	return fileDescriptor_1686212162667a44, []int{2}
}

func (m *S2LUserLost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2LUserLost.Unmarshal(m, b)
}
func (m *S2LUserLost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2LUserLost.Marshal(b, m, deterministic)
}
func (m *S2LUserLost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2LUserLost.Merge(m, src)
}
func (m *S2LUserLost) XXX_Size() int {
	return xxx_messageInfo_S2LUserLost.Size(m)
}
func (m *S2LUserLost) XXX_DiscardUnknown() {
	xxx_messageInfo_S2LUserLost.DiscardUnknown(m)
}

var xxx_messageInfo_S2LUserLost proto.InternalMessageInfo

func (m *S2LUserLost) GetLoginsocket() uint32 {
	if m != nil {
		return m.Loginsocket
	}
	return 0
}

type S2LRegisterUser struct {
	Loginsocket          uint32   `protobuf:"varint,1,opt,name=loginsocket,proto3" json:"loginsocket,omitempty"`
	Account              string   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2LRegisterUser) Reset()         { *m = S2LRegisterUser{} }
func (m *S2LRegisterUser) String() string { return proto.CompactTextString(m) }
func (*S2LRegisterUser) ProtoMessage()    {}
func (*S2LRegisterUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_1686212162667a44, []int{3}
}

func (m *S2LRegisterUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2LRegisterUser.Unmarshal(m, b)
}
func (m *S2LRegisterUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2LRegisterUser.Marshal(b, m, deterministic)
}
func (m *S2LRegisterUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2LRegisterUser.Merge(m, src)
}
func (m *S2LRegisterUser) XXX_Size() int {
	return xxx_messageInfo_S2LRegisterUser.Size(m)
}
func (m *S2LRegisterUser) XXX_DiscardUnknown() {
	xxx_messageInfo_S2LRegisterUser.DiscardUnknown(m)
}

var xxx_messageInfo_S2LRegisterUser proto.InternalMessageInfo

func (m *S2LRegisterUser) GetLoginsocket() uint32 {
	if m != nil {
		return m.Loginsocket
	}
	return 0
}

func (m *S2LRegisterUser) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *S2LRegisterUser) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type L2SRegisterUserResult struct {
	Loginsocket          uint32   `protobuf:"varint,1,opt,name=loginsocket,proto3" json:"loginsocket,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Account              string   `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2SRegisterUserResult) Reset()         { *m = L2SRegisterUserResult{} }
func (m *L2SRegisterUserResult) String() string { return proto.CompactTextString(m) }
func (*L2SRegisterUserResult) ProtoMessage()    {}
func (*L2SRegisterUserResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_1686212162667a44, []int{4}
}

func (m *L2SRegisterUserResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2SRegisterUserResult.Unmarshal(m, b)
}
func (m *L2SRegisterUserResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2SRegisterUserResult.Marshal(b, m, deterministic)
}
func (m *L2SRegisterUserResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2SRegisterUserResult.Merge(m, src)
}
func (m *L2SRegisterUserResult) XXX_Size() int {
	return xxx_messageInfo_L2SRegisterUserResult.Size(m)
}
func (m *L2SRegisterUserResult) XXX_DiscardUnknown() {
	xxx_messageInfo_L2SRegisterUserResult.DiscardUnknown(m)
}

var xxx_messageInfo_L2SRegisterUserResult proto.InternalMessageInfo

func (m *L2SRegisterUserResult) GetLoginsocket() uint32 {
	if m != nil {
		return m.Loginsocket
	}
	return 0
}

func (m *L2SRegisterUserResult) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *L2SRegisterUserResult) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *L2SRegisterUserResult) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*S2LLoginUser)(nil), "pt_logical.s2l_login_user")
	proto.RegisterType((*L2SLoginUserResult)(nil), "pt_logical.l2s_login_user_result")
	proto.RegisterMapType((map[uint64]*pt_com.RoomItem)(nil), "pt_logical.l2s_login_user_result.RoomsEntry")
	proto.RegisterType((*S2LUserLost)(nil), "pt_logical.s2l_user_lost")
	proto.RegisterType((*S2LRegisterUser)(nil), "pt_logical.s2l_register_user")
	proto.RegisterType((*L2SRegisterUserResult)(nil), "pt_logical.l2s_register_user_result")
}

func init() { proto.RegisterFile("pt_logical/login.proto", fileDescriptor_1686212162667a44) }

var fileDescriptor_1686212162667a44 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcd, 0x6a, 0xe3, 0x30,
	0x10, 0xc6, 0x7f, 0x49, 0x76, 0x42, 0x96, 0x8d, 0x96, 0x0d, 0x22, 0x27, 0xe3, 0xcb, 0xe6, 0x50,
	0x1c, 0xea, 0x5e, 0x4a, 0x8f, 0x85, 0x9e, 0x7a, 0xd3, 0x0b, 0x18, 0xd7, 0x16, 0xc1, 0x44, 0xf6,
	0x18, 0x49, 0x6e, 0xc9, 0xb5, 0x87, 0x3e, 0x64, 0x9f, 0xa6, 0x48, 0x76, 0x1a, 0xa7, 0x34, 0x90,
	0xb6, 0x27, 0xcf, 0x37, 0x3f, 0x9a, 0x6f, 0xe6, 0x1b, 0xc3, 0xa2, 0xd1, 0xa9, 0xc0, 0x4d, 0x99,
	0x67, 0x62, 0x6d, 0xbe, 0x75, 0xdc, 0x48, 0xd4, 0x48, 0xe0, 0xe0, 0x5f, 0xfe, 0x6d, 0x74, 0x9a,
	0x63, 0xb5, 0xce, 0xb1, 0xaa, 0xb0, 0x4f, 0x88, 0x9e, 0x1d, 0xf8, 0xad, 0x12, 0x61, 0x93, 0xea,
	0xb4, 0x55, 0x5c, 0x92, 0x10, 0xa6, 0x16, 0x29, 0xcc, 0xb7, 0x5c, 0x53, 0x27, 0x74, 0x56, 0x33,
	0x36, 0x74, 0x11, 0x0a, 0xe3, 0xac, 0x28, 0x24, 0x57, 0x8a, 0xba, 0xa1, 0xb3, 0xfa, 0xc5, 0xf6,
	0xd0, 0x46, 0xf2, 0x1c, 0xdb, 0x5a, 0x53, 0xaf, 0x8f, 0x74, 0x90, 0x2c, 0x61, 0xd2, 0x64, 0x4a,
	0x3d, 0xa1, 0x2c, 0xa8, 0x6f, 0x43, 0xef, 0x38, 0x7a, 0x75, 0xe1, 0x9f, 0x48, 0xd4, 0x80, 0x44,
	0x2a, 0xb9, 0x6a, 0x85, 0xfe, 0x11, 0x17, 0x02, 0x7e, 0x8e, 0x05, 0xb7, 0x44, 0x02, 0x66, 0x6d,
	0xb2, 0x80, 0x91, 0x79, 0xbe, 0xec, 0x38, 0xf8, 0xac, 0x47, 0x43, 0xde, 0xc1, 0x69, 0xde, 0xa3,
	0x63, 0xde, 0xa6, 0xc3, 0x06, 0x45, 0x41, 0xc7, 0xf6, 0x2d, 0x6b, 0x93, 0x5b, 0x08, 0x24, 0x62,
	0xa5, 0xe8, 0x24, 0xf4, 0x56, 0xd3, 0xe4, 0x22, 0x3e, 0x28, 0x10, 0x7f, 0x3a, 0x63, 0xcc, 0x4c,
	0xfa, 0x5d, 0xad, 0xe5, 0x8e, 0x75, 0xa5, 0xcb, 0x7b, 0x80, 0x83, 0x93, 0xfc, 0x01, 0x6f, 0xcb,
	0x77, 0x76, 0x76, 0x9f, 0x19, 0x93, 0xfc, 0x87, 0xe0, 0x31, 0x13, 0x2d, 0xb7, 0x13, 0x4f, 0x93,
	0x79, 0xdc, 0x29, 0x1b, 0x9b, 0xea, 0xb4, 0xd4, 0xbc, 0x62, 0x5d, 0xfc, 0xc6, 0xbd, 0x76, 0xa2,
	0x4b, 0x98, 0x19, 0x81, 0x6d, 0x47, 0x81, 0xea, 0x8c, 0x9d, 0x46, 0x5b, 0x98, 0x9b, 0x12, 0xc9,
	0x37, 0xa5, 0xd2, 0x5c, 0x7e, 0xe5, 0x2c, 0xfa, 0x25, 0xba, 0xa7, 0x97, 0xe8, 0x7d, 0x10, 0xff,
	0xc5, 0x01, 0x6a, 0x16, 0x73, 0xd4, 0xed, 0x7c, 0xfd, 0xf7, 0x2a, 0xbb, 0x03, 0x95, 0xbf, 0x75,
	0x85, 0x0f, 0x23, 0xfb, 0x47, 0x5c, 0xbd, 0x05, 0x00, 0x00, 0xff, 0xff, 0x32, 0x63, 0x9d, 0xee,
	0x4c, 0x03, 0x00, 0x00,
}