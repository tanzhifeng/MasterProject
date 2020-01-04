// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pt_login/room.proto

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

type S2CNotifyAppendRoom struct {
	Item                 *pt_com.RoomItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *S2CNotifyAppendRoom) Reset()         { *m = S2CNotifyAppendRoom{} }
func (m *S2CNotifyAppendRoom) String() string { return proto.CompactTextString(m) }
func (*S2CNotifyAppendRoom) ProtoMessage()    {}
func (*S2CNotifyAppendRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ea27b92c163e15, []int{0}
}

func (m *S2CNotifyAppendRoom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CNotifyAppendRoom.Unmarshal(m, b)
}
func (m *S2CNotifyAppendRoom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CNotifyAppendRoom.Marshal(b, m, deterministic)
}
func (m *S2CNotifyAppendRoom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CNotifyAppendRoom.Merge(m, src)
}
func (m *S2CNotifyAppendRoom) XXX_Size() int {
	return xxx_messageInfo_S2CNotifyAppendRoom.Size(m)
}
func (m *S2CNotifyAppendRoom) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CNotifyAppendRoom.DiscardUnknown(m)
}

var xxx_messageInfo_S2CNotifyAppendRoom proto.InternalMessageInfo

func (m *S2CNotifyAppendRoom) GetItem() *pt_com.RoomItem {
	if m != nil {
		return m.Item
	}
	return nil
}

type S2CNotifyRemoveRoom struct {
	Roomid               uint64   `protobuf:"varint,1,opt,name=roomid,proto3" json:"roomid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2CNotifyRemoveRoom) Reset()         { *m = S2CNotifyRemoveRoom{} }
func (m *S2CNotifyRemoveRoom) String() string { return proto.CompactTextString(m) }
func (*S2CNotifyRemoveRoom) ProtoMessage()    {}
func (*S2CNotifyRemoveRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ea27b92c163e15, []int{1}
}

func (m *S2CNotifyRemoveRoom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CNotifyRemoveRoom.Unmarshal(m, b)
}
func (m *S2CNotifyRemoveRoom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CNotifyRemoveRoom.Marshal(b, m, deterministic)
}
func (m *S2CNotifyRemoveRoom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CNotifyRemoveRoom.Merge(m, src)
}
func (m *S2CNotifyRemoveRoom) XXX_Size() int {
	return xxx_messageInfo_S2CNotifyRemoveRoom.Size(m)
}
func (m *S2CNotifyRemoveRoom) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CNotifyRemoveRoom.DiscardUnknown(m)
}

var xxx_messageInfo_S2CNotifyRemoveRoom proto.InternalMessageInfo

func (m *S2CNotifyRemoveRoom) GetRoomid() uint64 {
	if m != nil {
		return m.Roomid
	}
	return 0
}

type S2CNotifyAppendTable struct {
	Item                 *pt_com.RoomTableItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *S2CNotifyAppendTable) Reset()         { *m = S2CNotifyAppendTable{} }
func (m *S2CNotifyAppendTable) String() string { return proto.CompactTextString(m) }
func (*S2CNotifyAppendTable) ProtoMessage()    {}
func (*S2CNotifyAppendTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ea27b92c163e15, []int{2}
}

func (m *S2CNotifyAppendTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CNotifyAppendTable.Unmarshal(m, b)
}
func (m *S2CNotifyAppendTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CNotifyAppendTable.Marshal(b, m, deterministic)
}
func (m *S2CNotifyAppendTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CNotifyAppendTable.Merge(m, src)
}
func (m *S2CNotifyAppendTable) XXX_Size() int {
	return xxx_messageInfo_S2CNotifyAppendTable.Size(m)
}
func (m *S2CNotifyAppendTable) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CNotifyAppendTable.DiscardUnknown(m)
}

var xxx_messageInfo_S2CNotifyAppendTable proto.InternalMessageInfo

func (m *S2CNotifyAppendTable) GetItem() *pt_com.RoomTableItem {
	if m != nil {
		return m.Item
	}
	return nil
}

type S2CNotifyRemoveTable struct {
	Tableid              uint32   `protobuf:"varint,1,opt,name=tableid,proto3" json:"tableid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2CNotifyRemoveTable) Reset()         { *m = S2CNotifyRemoveTable{} }
func (m *S2CNotifyRemoveTable) String() string { return proto.CompactTextString(m) }
func (*S2CNotifyRemoveTable) ProtoMessage()    {}
func (*S2CNotifyRemoveTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ea27b92c163e15, []int{3}
}

func (m *S2CNotifyRemoveTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CNotifyRemoveTable.Unmarshal(m, b)
}
func (m *S2CNotifyRemoveTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CNotifyRemoveTable.Marshal(b, m, deterministic)
}
func (m *S2CNotifyRemoveTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CNotifyRemoveTable.Merge(m, src)
}
func (m *S2CNotifyRemoveTable) XXX_Size() int {
	return xxx_messageInfo_S2CNotifyRemoveTable.Size(m)
}
func (m *S2CNotifyRemoveTable) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CNotifyRemoveTable.DiscardUnknown(m)
}

var xxx_messageInfo_S2CNotifyRemoveTable proto.InternalMessageInfo

func (m *S2CNotifyRemoveTable) GetTableid() uint32 {
	if m != nil {
		return m.Tableid
	}
	return 0
}

type C2SRoomUserEnter struct {
	Roomid               uint64   `protobuf:"varint,1,opt,name=roomid,proto3" json:"roomid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2SRoomUserEnter) Reset()         { *m = C2SRoomUserEnter{} }
func (m *C2SRoomUserEnter) String() string { return proto.CompactTextString(m) }
func (*C2SRoomUserEnter) ProtoMessage()    {}
func (*C2SRoomUserEnter) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ea27b92c163e15, []int{4}
}

func (m *C2SRoomUserEnter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2SRoomUserEnter.Unmarshal(m, b)
}
func (m *C2SRoomUserEnter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2SRoomUserEnter.Marshal(b, m, deterministic)
}
func (m *C2SRoomUserEnter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2SRoomUserEnter.Merge(m, src)
}
func (m *C2SRoomUserEnter) XXX_Size() int {
	return xxx_messageInfo_C2SRoomUserEnter.Size(m)
}
func (m *C2SRoomUserEnter) XXX_DiscardUnknown() {
	xxx_messageInfo_C2SRoomUserEnter.DiscardUnknown(m)
}

var xxx_messageInfo_C2SRoomUserEnter proto.InternalMessageInfo

func (m *C2SRoomUserEnter) GetRoomid() uint64 {
	if m != nil {
		return m.Roomid
	}
	return 0
}

type S2CRoomUserEnterResult struct {
	Code                 int32                            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Roomid               uint64                           `protobuf:"varint,2,opt,name=roomid,proto3" json:"roomid,omitempty"`
	Users                map[uint64]*pt_com.RoomUserItem  `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Tables               map[uint32]*pt_com.RoomTableItem `protobuf:"bytes,4,rep,name=tables,proto3" json:"tables,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *S2CRoomUserEnterResult) Reset()         { *m = S2CRoomUserEnterResult{} }
func (m *S2CRoomUserEnterResult) String() string { return proto.CompactTextString(m) }
func (*S2CRoomUserEnterResult) ProtoMessage()    {}
func (*S2CRoomUserEnterResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ea27b92c163e15, []int{5}
}

func (m *S2CRoomUserEnterResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CRoomUserEnterResult.Unmarshal(m, b)
}
func (m *S2CRoomUserEnterResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CRoomUserEnterResult.Marshal(b, m, deterministic)
}
func (m *S2CRoomUserEnterResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CRoomUserEnterResult.Merge(m, src)
}
func (m *S2CRoomUserEnterResult) XXX_Size() int {
	return xxx_messageInfo_S2CRoomUserEnterResult.Size(m)
}
func (m *S2CRoomUserEnterResult) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CRoomUserEnterResult.DiscardUnknown(m)
}

var xxx_messageInfo_S2CRoomUserEnterResult proto.InternalMessageInfo

func (m *S2CRoomUserEnterResult) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *S2CRoomUserEnterResult) GetRoomid() uint64 {
	if m != nil {
		return m.Roomid
	}
	return 0
}

func (m *S2CRoomUserEnterResult) GetUsers() map[uint64]*pt_com.RoomUserItem {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *S2CRoomUserEnterResult) GetTables() map[uint32]*pt_com.RoomTableItem {
	if m != nil {
		return m.Tables
	}
	return nil
}

type S2CRoomUserLeaveResult struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Roomid               uint64   `protobuf:"varint,2,opt,name=roomid,proto3" json:"roomid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2CRoomUserLeaveResult) Reset()         { *m = S2CRoomUserLeaveResult{} }
func (m *S2CRoomUserLeaveResult) String() string { return proto.CompactTextString(m) }
func (*S2CRoomUserLeaveResult) ProtoMessage()    {}
func (*S2CRoomUserLeaveResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ea27b92c163e15, []int{6}
}

func (m *S2CRoomUserLeaveResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CRoomUserLeaveResult.Unmarshal(m, b)
}
func (m *S2CRoomUserLeaveResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CRoomUserLeaveResult.Marshal(b, m, deterministic)
}
func (m *S2CRoomUserLeaveResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CRoomUserLeaveResult.Merge(m, src)
}
func (m *S2CRoomUserLeaveResult) XXX_Size() int {
	return xxx_messageInfo_S2CRoomUserLeaveResult.Size(m)
}
func (m *S2CRoomUserLeaveResult) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CRoomUserLeaveResult.DiscardUnknown(m)
}

var xxx_messageInfo_S2CRoomUserLeaveResult proto.InternalMessageInfo

func (m *S2CRoomUserLeaveResult) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *S2CRoomUserLeaveResult) GetRoomid() uint64 {
	if m != nil {
		return m.Roomid
	}
	return 0
}

type S2CRoomNotifyAppendUser struct {
	Item                 *pt_com.RoomUserItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *S2CRoomNotifyAppendUser) Reset()         { *m = S2CRoomNotifyAppendUser{} }
func (m *S2CRoomNotifyAppendUser) String() string { return proto.CompactTextString(m) }
func (*S2CRoomNotifyAppendUser) ProtoMessage()    {}
func (*S2CRoomNotifyAppendUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ea27b92c163e15, []int{7}
}

func (m *S2CRoomNotifyAppendUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CRoomNotifyAppendUser.Unmarshal(m, b)
}
func (m *S2CRoomNotifyAppendUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CRoomNotifyAppendUser.Marshal(b, m, deterministic)
}
func (m *S2CRoomNotifyAppendUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CRoomNotifyAppendUser.Merge(m, src)
}
func (m *S2CRoomNotifyAppendUser) XXX_Size() int {
	return xxx_messageInfo_S2CRoomNotifyAppendUser.Size(m)
}
func (m *S2CRoomNotifyAppendUser) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CRoomNotifyAppendUser.DiscardUnknown(m)
}

var xxx_messageInfo_S2CRoomNotifyAppendUser proto.InternalMessageInfo

func (m *S2CRoomNotifyAppendUser) GetItem() *pt_com.RoomUserItem {
	if m != nil {
		return m.Item
	}
	return nil
}

type S2CRoomNotifyRemoveUser struct {
	Userid               uint64   `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2CRoomNotifyRemoveUser) Reset()         { *m = S2CRoomNotifyRemoveUser{} }
func (m *S2CRoomNotifyRemoveUser) String() string { return proto.CompactTextString(m) }
func (*S2CRoomNotifyRemoveUser) ProtoMessage()    {}
func (*S2CRoomNotifyRemoveUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ea27b92c163e15, []int{8}
}

func (m *S2CRoomNotifyRemoveUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CRoomNotifyRemoveUser.Unmarshal(m, b)
}
func (m *S2CRoomNotifyRemoveUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CRoomNotifyRemoveUser.Marshal(b, m, deterministic)
}
func (m *S2CRoomNotifyRemoveUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CRoomNotifyRemoveUser.Merge(m, src)
}
func (m *S2CRoomNotifyRemoveUser) XXX_Size() int {
	return xxx_messageInfo_S2CRoomNotifyRemoveUser.Size(m)
}
func (m *S2CRoomNotifyRemoveUser) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CRoomNotifyRemoveUser.DiscardUnknown(m)
}

var xxx_messageInfo_S2CRoomNotifyRemoveUser proto.InternalMessageInfo

func (m *S2CRoomNotifyRemoveUser) GetUserid() uint64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

type C2SRoomUserSitdown struct {
	Tableid              uint32   `protobuf:"varint,1,opt,name=tableid,proto3" json:"tableid,omitempty"`
	Chairid              uint32   `protobuf:"varint,2,opt,name=chairid,proto3" json:"chairid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2SRoomUserSitdown) Reset()         { *m = C2SRoomUserSitdown{} }
func (m *C2SRoomUserSitdown) String() string { return proto.CompactTextString(m) }
func (*C2SRoomUserSitdown) ProtoMessage()    {}
func (*C2SRoomUserSitdown) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ea27b92c163e15, []int{9}
}

func (m *C2SRoomUserSitdown) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2SRoomUserSitdown.Unmarshal(m, b)
}
func (m *C2SRoomUserSitdown) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2SRoomUserSitdown.Marshal(b, m, deterministic)
}
func (m *C2SRoomUserSitdown) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2SRoomUserSitdown.Merge(m, src)
}
func (m *C2SRoomUserSitdown) XXX_Size() int {
	return xxx_messageInfo_C2SRoomUserSitdown.Size(m)
}
func (m *C2SRoomUserSitdown) XXX_DiscardUnknown() {
	xxx_messageInfo_C2SRoomUserSitdown.DiscardUnknown(m)
}

var xxx_messageInfo_C2SRoomUserSitdown proto.InternalMessageInfo

func (m *C2SRoomUserSitdown) GetTableid() uint32 {
	if m != nil {
		return m.Tableid
	}
	return 0
}

func (m *C2SRoomUserSitdown) GetChairid() uint32 {
	if m != nil {
		return m.Chairid
	}
	return 0
}

type S2CRoomUserSitdownResult struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Tableid              uint32   `protobuf:"varint,2,opt,name=tableid,proto3" json:"tableid,omitempty"`
	Chairid              uint32   `protobuf:"varint,3,opt,name=chairid,proto3" json:"chairid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2CRoomUserSitdownResult) Reset()         { *m = S2CRoomUserSitdownResult{} }
func (m *S2CRoomUserSitdownResult) String() string { return proto.CompactTextString(m) }
func (*S2CRoomUserSitdownResult) ProtoMessage()    {}
func (*S2CRoomUserSitdownResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ea27b92c163e15, []int{10}
}

func (m *S2CRoomUserSitdownResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CRoomUserSitdownResult.Unmarshal(m, b)
}
func (m *S2CRoomUserSitdownResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CRoomUserSitdownResult.Marshal(b, m, deterministic)
}
func (m *S2CRoomUserSitdownResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CRoomUserSitdownResult.Merge(m, src)
}
func (m *S2CRoomUserSitdownResult) XXX_Size() int {
	return xxx_messageInfo_S2CRoomUserSitdownResult.Size(m)
}
func (m *S2CRoomUserSitdownResult) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CRoomUserSitdownResult.DiscardUnknown(m)
}

var xxx_messageInfo_S2CRoomUserSitdownResult proto.InternalMessageInfo

func (m *S2CRoomUserSitdownResult) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *S2CRoomUserSitdownResult) GetTableid() uint32 {
	if m != nil {
		return m.Tableid
	}
	return 0
}

func (m *S2CRoomUserSitdownResult) GetChairid() uint32 {
	if m != nil {
		return m.Chairid
	}
	return 0
}

type S2CRoomUserStandupResult struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2CRoomUserStandupResult) Reset()         { *m = S2CRoomUserStandupResult{} }
func (m *S2CRoomUserStandupResult) String() string { return proto.CompactTextString(m) }
func (*S2CRoomUserStandupResult) ProtoMessage()    {}
func (*S2CRoomUserStandupResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ea27b92c163e15, []int{11}
}

func (m *S2CRoomUserStandupResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CRoomUserStandupResult.Unmarshal(m, b)
}
func (m *S2CRoomUserStandupResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CRoomUserStandupResult.Marshal(b, m, deterministic)
}
func (m *S2CRoomUserStandupResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CRoomUserStandupResult.Merge(m, src)
}
func (m *S2CRoomUserStandupResult) XXX_Size() int {
	return xxx_messageInfo_S2CRoomUserStandupResult.Size(m)
}
func (m *S2CRoomUserStandupResult) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CRoomUserStandupResult.DiscardUnknown(m)
}

var xxx_messageInfo_S2CRoomUserStandupResult proto.InternalMessageInfo

func (m *S2CRoomUserStandupResult) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type S2CRoomNotifyUserSitdown struct {
	Userid               uint64   `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Tableid              uint32   `protobuf:"varint,2,opt,name=tableid,proto3" json:"tableid,omitempty"`
	Chairid              uint32   `protobuf:"varint,3,opt,name=chairid,proto3" json:"chairid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2CRoomNotifyUserSitdown) Reset()         { *m = S2CRoomNotifyUserSitdown{} }
func (m *S2CRoomNotifyUserSitdown) String() string { return proto.CompactTextString(m) }
func (*S2CRoomNotifyUserSitdown) ProtoMessage()    {}
func (*S2CRoomNotifyUserSitdown) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ea27b92c163e15, []int{12}
}

func (m *S2CRoomNotifyUserSitdown) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CRoomNotifyUserSitdown.Unmarshal(m, b)
}
func (m *S2CRoomNotifyUserSitdown) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CRoomNotifyUserSitdown.Marshal(b, m, deterministic)
}
func (m *S2CRoomNotifyUserSitdown) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CRoomNotifyUserSitdown.Merge(m, src)
}
func (m *S2CRoomNotifyUserSitdown) XXX_Size() int {
	return xxx_messageInfo_S2CRoomNotifyUserSitdown.Size(m)
}
func (m *S2CRoomNotifyUserSitdown) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CRoomNotifyUserSitdown.DiscardUnknown(m)
}

var xxx_messageInfo_S2CRoomNotifyUserSitdown proto.InternalMessageInfo

func (m *S2CRoomNotifyUserSitdown) GetUserid() uint64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *S2CRoomNotifyUserSitdown) GetTableid() uint32 {
	if m != nil {
		return m.Tableid
	}
	return 0
}

func (m *S2CRoomNotifyUserSitdown) GetChairid() uint32 {
	if m != nil {
		return m.Chairid
	}
	return 0
}

type S2CRoomNotifyUserStandup struct {
	Userid               uint64   `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2CRoomNotifyUserStandup) Reset()         { *m = S2CRoomNotifyUserStandup{} }
func (m *S2CRoomNotifyUserStandup) String() string { return proto.CompactTextString(m) }
func (*S2CRoomNotifyUserStandup) ProtoMessage()    {}
func (*S2CRoomNotifyUserStandup) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ea27b92c163e15, []int{13}
}

func (m *S2CRoomNotifyUserStandup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CRoomNotifyUserStandup.Unmarshal(m, b)
}
func (m *S2CRoomNotifyUserStandup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CRoomNotifyUserStandup.Marshal(b, m, deterministic)
}
func (m *S2CRoomNotifyUserStandup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CRoomNotifyUserStandup.Merge(m, src)
}
func (m *S2CRoomNotifyUserStandup) XXX_Size() int {
	return xxx_messageInfo_S2CRoomNotifyUserStandup.Size(m)
}
func (m *S2CRoomNotifyUserStandup) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CRoomNotifyUserStandup.DiscardUnknown(m)
}

var xxx_messageInfo_S2CRoomNotifyUserStandup proto.InternalMessageInfo

func (m *S2CRoomNotifyUserStandup) GetUserid() uint64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func init() {
	proto.RegisterType((*S2CNotifyAppendRoom)(nil), "pt_login.s2c_notify_append_room")
	proto.RegisterType((*S2CNotifyRemoveRoom)(nil), "pt_login.s2c_notify_remove_room")
	proto.RegisterType((*S2CNotifyAppendTable)(nil), "pt_login.s2c_notify_append_table")
	proto.RegisterType((*S2CNotifyRemoveTable)(nil), "pt_login.s2c_notify_remove_table")
	proto.RegisterType((*C2SRoomUserEnter)(nil), "pt_login.c2s_room_user_enter")
	proto.RegisterType((*S2CRoomUserEnterResult)(nil), "pt_login.s2c_room_user_enter_result")
	proto.RegisterMapType((map[uint32]*pt_com.RoomTableItem)(nil), "pt_login.s2c_room_user_enter_result.TablesEntry")
	proto.RegisterMapType((map[uint64]*pt_com.RoomUserItem)(nil), "pt_login.s2c_room_user_enter_result.UsersEntry")
	proto.RegisterType((*S2CRoomUserLeaveResult)(nil), "pt_login.s2c_room_user_leave_result")
	proto.RegisterType((*S2CRoomNotifyAppendUser)(nil), "pt_login.s2c_room_notify_append_user")
	proto.RegisterType((*S2CRoomNotifyRemoveUser)(nil), "pt_login.s2c_room_notify_remove_user")
	proto.RegisterType((*C2SRoomUserSitdown)(nil), "pt_login.c2s_room_user_sitdown")
	proto.RegisterType((*S2CRoomUserSitdownResult)(nil), "pt_login.s2c_room_user_sitdown_result")
	proto.RegisterType((*S2CRoomUserStandupResult)(nil), "pt_login.s2c_room_user_standup_result")
	proto.RegisterType((*S2CRoomNotifyUserSitdown)(nil), "pt_login.s2c_room_notify_user_sitdown")
	proto.RegisterType((*S2CRoomNotifyUserStandup)(nil), "pt_login.s2c_room_notify_user_standup")
}

func init() { proto.RegisterFile("pt_login/room.proto", fileDescriptor_21ea27b92c163e15) }

var fileDescriptor_21ea27b92c163e15 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcd, 0x8b, 0xd4, 0x30,
	0x14, 0xc0, 0x99, 0xaf, 0x2a, 0xaf, 0x0c, 0x68, 0x16, 0xbb, 0xa5, 0x7a, 0x58, 0x02, 0xc2, 0xa2,
	0x6e, 0x67, 0xe9, 0xa2, 0x88, 0x17, 0x4f, 0x2b, 0x2b, 0x5e, 0xa4, 0xe8, 0x39, 0x74, 0xdb, 0xac,
	0x56, 0xdb, 0xa4, 0xa4, 0xe9, 0xca, 0xfc, 0xf5, 0x4a, 0x3e, 0xea, 0x34, 0x3b, 0x64, 0xd0, 0x3d,
	0x35, 0x99, 0x97, 0xf7, 0xfb, 0xe5, 0xe5, 0x3d, 0x06, 0x8e, 0x3a, 0x49, 0x1a, 0xfe, 0xad, 0x66,
	0x1b, 0xc1, 0x79, 0x9b, 0x76, 0x82, 0x4b, 0x8e, 0x1e, 0x8e, 0x3f, 0x26, 0x2a, 0x5c, 0xf2, 0x76,
	0x53, 0xf2, 0xb6, 0xe5, 0xcc, 0x84, 0xf1, 0x7b, 0x88, 0xfa, 0xac, 0x24, 0x8c, 0xcb, 0xfa, 0x66,
	0x4b, 0x8a, 0xae, 0xa3, 0xac, 0x22, 0x2a, 0x1d, 0x3d, 0x87, 0x65, 0x2d, 0x69, 0x1b, 0xcf, 0x4e,
	0x66, 0xa7, 0x61, 0xf6, 0x38, 0x35, 0xd9, 0xa9, 0x8a, 0x11, 0x15, 0xc8, 0x75, 0x18, 0x9f, 0x3b,
	0x00, 0x41, 0x5b, 0x7e, 0x4b, 0x0d, 0x20, 0x82, 0x40, 0x7d, 0xeb, 0x4a, 0x23, 0x96, 0xb9, 0xdd,
	0xe1, 0x0f, 0x70, 0xbc, 0xaf, 0x94, 0xc5, 0x75, 0x43, 0xd1, 0x4b, 0xc7, 0x79, 0xec, 0x38, 0xf5,
	0x89, 0xa9, 0xf9, 0xc2, 0xe1, 0x58, 0xb3, 0xe1, 0xc4, 0xf0, 0x40, 0x2f, 0xac, 0x7b, 0x9d, 0x8f,
	0x5b, 0x7c, 0x06, 0x47, 0x65, 0xd6, 0xeb, 0x0b, 0x92, 0xa1, 0xa7, 0x82, 0x50, 0x26, 0xa9, 0xf0,
	0xde, 0xf5, 0xf7, 0x1c, 0x12, 0x25, 0xb9, 0x73, 0x9e, 0x08, 0xda, 0x0f, 0x8d, 0x44, 0x08, 0x96,
	0x25, 0xaf, 0xa8, 0x4e, 0x5a, 0xe5, 0x7a, 0x3d, 0x41, 0xcd, 0xa7, 0x28, 0x74, 0x09, 0x2b, 0x05,
	0xe8, 0xe3, 0xc5, 0xc9, 0xe2, 0x34, 0xcc, 0x36, 0xe9, 0xd8, 0x98, 0xd4, 0x2f, 0x48, 0xbf, 0xaa,
	0x8c, 0x4b, 0x26, 0xc5, 0x36, 0x37, 0xd9, 0xe8, 0x0a, 0x02, 0x5d, 0x4b, 0x1f, 0x2f, 0x35, 0xe7,
	0xfc, 0x9f, 0x38, 0x5f, 0x74, 0x8a, 0x01, 0xd9, 0xfc, 0xe4, 0x33, 0xc0, 0x0e, 0x8f, 0x1e, 0xc1,
	0xe2, 0x27, 0xdd, 0xda, 0xf2, 0xd5, 0x12, 0xbd, 0x82, 0xd5, 0x6d, 0xd1, 0x0c, 0x54, 0xd7, 0x11,
	0x66, 0x91, 0xd3, 0x0d, 0xad, 0xd0, 0xcd, 0x30, 0x87, 0xde, 0xcd, 0xdf, 0xce, 0x92, 0x1c, 0xc2,
	0x89, 0x68, 0x8a, 0x5c, 0x1b, 0xe4, 0x99, 0x8b, 0xf4, 0x36, 0x78, 0xc7, 0xc4, 0x57, 0x77, 0x1b,
	0xd0, 0xd0, 0x42, 0x4d, 0xd8, 0x7f, 0x37, 0x00, 0x7f, 0x84, 0xa7, 0x7f, 0x49, 0xee, 0xf0, 0x29,
	0x2e, 0x7a, 0xe1, 0xcc, 0x9e, 0xaf, 0x5a, 0x33, 0x7a, 0xaf, 0xf7, 0x51, 0x76, 0xfe, 0x34, 0x2a,
	0x82, 0x40, 0x7d, 0x77, 0xd3, 0x64, 0x76, 0xf8, 0x13, 0x3c, 0x71, 0x87, 0xaf, 0xaf, 0x65, 0xc5,
	0x7f, 0x31, 0xff, 0xbc, 0xaa, 0x48, 0xf9, 0xbd, 0xa8, 0x85, 0xad, 0x66, 0x9d, 0x8f, 0x5b, 0x7c,
	0x03, 0xcf, 0xdc, 0x87, 0xb1, 0xb0, 0x43, 0x4f, 0x33, 0xf1, 0xcc, 0xbd, 0x9e, 0x85, 0xeb, 0xc9,
	0xf6, 0x3c, 0xb2, 0x60, 0xd5, 0xd0, 0x1d, 0xf0, 0xe0, 0x1f, 0x93, 0x1c, 0xfb, 0x3e, 0x4e, 0xbd,
	0x9e, 0x07, 0xba, 0xd7, 0xfd, 0xde, 0xf8, 0x5c, 0xe6, 0x9a, 0x3e, 0xd7, 0x75, 0xa0, 0xff, 0x00,
	0x2f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x96, 0x84, 0x9d, 0x36, 0x05, 0x00, 0x00,
}
