package testproto

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

type Role int32

const (
	Role_UNKNOWN Role = 0
	Role_REGULAR Role = 1
	Role_ADMIN   Role = 2
)

var Role_name = map[int32]string{
	0: "UNKNOWN",
	1: "REGULAR",
	2: "ADMIN",
}

var Role_value = map[string]int32{
	"UNKNOWN": 0,
	"REGULAR": 1,
	"ADMIN":   2,
}

func (x Role) String() string {
	return proto.EnumName(Role_name, int32(x))
}

func (Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

type Permission int32

const (
	Permission_READ    Permission = 0
	Permission_WRITE   Permission = 1
	Permission_EXECUTE Permission = 2
)

var Permission_name = map[int32]string{
	0: "READ",
	1: "WRITE",
	2: "EXECUTE",
}

var Permission_value = map[string]int32{
	"READ":    0,
	"WRITE":   1,
	"EXECUTE": 2,
}

func (x Permission) String() string {
	return proto.EnumName(Permission_name, int32(x))
}

func (Permission) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

type Image struct {
	OriginalUrl          string   `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
	ResizedUrl           string   `protobuf:"bytes,2,opt,name=resized_url,json=resizedUrl,proto3" json:"resized_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetOriginalUrl() string {
	if m != nil {
		return m.OriginalUrl
	}
	return ""
}

func (m *Image) GetResizedUrl() string {
	if m != nil {
		return m.ResizedUrl
	}
	return ""
}

type Metrics struct {
	Height               uint32   `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Weight               uint32   `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metrics) Reset()         { *m = Metrics{} }
func (m *Metrics) String() string { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()    {}
func (*Metrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *Metrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metrics.Unmarshal(m, b)
}
func (m *Metrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metrics.Marshal(b, m, deterministic)
}
func (m *Metrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metrics.Merge(m, src)
}
func (m *Metrics) XXX_Size() int {
	return xxx_messageInfo_Metrics.Size(m)
}
func (m *Metrics) XXX_DiscardUnknown() {
	xxx_messageInfo_Metrics.DiscardUnknown(m)
}

var xxx_messageInfo_Metrics proto.InternalMessageInfo

func (m *Metrics) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Metrics) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

type User struct {
	Id          uint32            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username    string            `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Role        Role              `protobuf:"varint,3,opt,name=role,proto3,enum=Role" json:"role,omitempty"`
	Meta        map[string]string `protobuf:"bytes,4,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Deactivated bool              `protobuf:"varint,5,opt,name=deactivated,proto3" json:"deactivated,omitempty"`
	Permissions []Permission      `protobuf:"varint,6,rep,packed,name=permissions,proto3,enum=Permission" json:"permissions,omitempty"`
	// Types that are valid to be assigned to Name:
	//	*User_MaleName
	//	*User_FemaleName
	Name                 isUser_Name `protobuf_oneof:"name"`
	Details              []*any.Any  `protobuf:"bytes,9,rep,name=details,proto3" json:"details,omitempty"`
	Images               []*Image    `protobuf:"bytes,10,rep,name=images,proto3" json:"images,omitempty"`
	Avatar               *Image      `protobuf:"bytes,11,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Tags                 []string    `protobuf:"bytes,12,rep,name=tags,proto3" json:"tags,omitempty"`
	Friends              []*User     `protobuf:"bytes,13,rep,name=friends,proto3" json:"friends,omitempty"`
	ExtraUser            *any.Any    `protobuf:"bytes,14,opt,name=extra_user,json=extraUser,proto3" json:"extra_user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{2}
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

func (m *User) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_UNKNOWN
}

func (m *User) GetMeta() map[string]string {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *User) GetDeactivated() bool {
	if m != nil {
		return m.Deactivated
	}
	return false
}

func (m *User) GetPermissions() []Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type isUser_Name interface {
	isUser_Name()
}

type User_MaleName struct {
	MaleName string `protobuf:"bytes,7,opt,name=male_name,json=maleName,proto3,oneof"`
}

type User_FemaleName struct {
	FemaleName string `protobuf:"bytes,8,opt,name=female_name,json=femaleName,proto3,oneof"`
}

func (*User_MaleName) isUser_Name() {}

func (*User_FemaleName) isUser_Name() {}

func (m *User) GetName() isUser_Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *User) GetMaleName() string {
	if x, ok := m.GetName().(*User_MaleName); ok {
		return x.MaleName
	}
	return ""
}

func (m *User) GetFemaleName() string {
	if x, ok := m.GetName().(*User_FemaleName); ok {
		return x.FemaleName
	}
	return ""
}

func (m *User) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *User) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *User) GetAvatar() *Image {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *User) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *User) GetFriends() []*User {
	if m != nil {
		return m.Friends
	}
	return nil
}

func (m *User) GetExtraUser() *any.Any {
	if m != nil {
		return m.ExtraUser
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*User) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*User_MaleName)(nil),
		(*User_FemaleName)(nil),
	}
}

type UpdateUserRequest struct {
	User                 *User                 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	FieldMask            *field_mask.FieldMask `protobuf:"bytes,2,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{3}
}

func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UpdateUserRequest) GetFieldMask() *field_mask.FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func init() {
	proto.RegisterEnum("Role", Role_name, Role_value)
	proto.RegisterEnum("Permission", Permission_name, Permission_value)
	proto.RegisterType((*Image)(nil), "Image")
	proto.RegisterType((*Metrics)(nil), "Metrics")
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterMapType((map[string]string)(nil), "User.MetaEntry")
	proto.RegisterType((*UpdateUserRequest)(nil), "UpdateUserRequest")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 613 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x5d, 0x4f, 0xdb, 0x4a,
	0x10, 0x86, 0x71, 0xe2, 0x7c, 0x78, 0x0c, 0x39, 0x39, 0x2b, 0x74, 0x64, 0x22, 0x9d, 0x62, 0xd2,
	0x9b, 0x88, 0x0a, 0x47, 0x0a, 0x17, 0x2d, 0xbd, 0x0b, 0xc5, 0x6d, 0x11, 0x4d, 0x5a, 0xad, 0xb0,
	0xa8, 0x7a, 0x13, 0x2d, 0x78, 0x62, 0x56, 0xf8, 0x83, 0xee, 0xae, 0xd3, 0xa6, 0xbf, 0xa6, 0x3f,
	0xb5, 0xda, 0xb5, 0x03, 0xa8, 0xed, 0x95, 0x77, 0x9e, 0x79, 0x77, 0xfc, 0x6a, 0x66, 0x16, 0x40,
	0xa1, 0x54, 0xc1, 0xbd, 0x28, 0x54, 0x31, 0xd8, 0x4b, 0x8a, 0x22, 0x49, 0x71, 0x6c, 0xa2, 0xeb,
	0x72, 0x39, 0x66, 0xf9, 0xba, 0x4e, 0xf9, 0xbf, 0xa7, 0x96, 0x1c, 0xd3, 0x78, 0x91, 0x31, 0x79,
	0x57, 0x29, 0x86, 0x17, 0xd0, 0x3a, 0xcf, 0x58, 0x82, 0xe4, 0x00, 0xb6, 0x0b, 0xc1, 0x13, 0x9e,
	0xb3, 0x74, 0x51, 0x8a, 0xd4, 0xb3, 0x7c, 0x6b, 0xe4, 0x50, 0x77, 0xc3, 0x22, 0x91, 0x92, 0x7d,
	0x70, 0x05, 0x4a, 0xfe, 0x03, 0x63, 0xa3, 0x68, 0x18, 0x05, 0xd4, 0x28, 0x12, 0xe9, 0xf0, 0x04,
	0x3a, 0x33, 0x54, 0x82, 0xdf, 0x48, 0xf2, 0x1f, 0xb4, 0x6f, 0x91, 0x27, 0xb7, 0xca, 0x14, 0xda,
	0xa1, 0x75, 0xa4, 0xf9, 0xb7, 0x8a, 0x37, 0x2a, 0x5e, 0x45, 0xc3, 0x9f, 0x36, 0xd8, 0x91, 0x44,
	0x41, 0x7a, 0xd0, 0xe0, 0x71, 0x7d, 0xa9, 0xc1, 0x63, 0x32, 0x80, 0x6e, 0x29, 0x51, 0xe4, 0x2c,
	0xc3, 0xfa, 0x8f, 0x0f, 0x31, 0xd9, 0x03, 0x5b, 0x14, 0x29, 0x7a, 0x4d, 0xdf, 0x1a, 0xf5, 0x26,
	0xad, 0x80, 0x16, 0x29, 0x52, 0x83, 0xc8, 0x73, 0xb0, 0x33, 0x54, 0xcc, 0xb3, 0xfd, 0xe6, 0xc8,
	0x9d, 0xfc, 0x13, 0xe8, 0xda, 0xc1, 0x0c, 0x15, 0x0b, 0x73, 0x25, 0xd6, 0xd4, 0x24, 0x89, 0x0f,
	0x6e, 0x8c, 0xec, 0x46, 0xf1, 0x15, 0x53, 0x18, 0x7b, 0x2d, 0xdf, 0x1a, 0x75, 0xe9, 0x53, 0x44,
	0x8e, 0xc0, 0xbd, 0x47, 0x91, 0x71, 0x29, 0x79, 0x91, 0x4b, 0xaf, 0xed, 0x37, 0x47, 0xbd, 0x89,
	0x1b, 0x7c, 0x7a, 0x60, 0xf4, 0x69, 0x9e, 0xfc, 0x0f, 0x4e, 0xc6, 0x52, 0x5c, 0x18, 0xb7, 0x1d,
	0xed, 0xf6, 0xfd, 0x16, 0xed, 0x6a, 0x34, 0xd7, 0x7e, 0x0f, 0xc0, 0x5d, 0xe2, 0xa3, 0xa0, 0x5b,
	0x0b, 0xa0, 0x82, 0x46, 0x12, 0x40, 0x27, 0x46, 0xc5, 0x78, 0x2a, 0x3d, 0xc7, 0x58, 0xdf, 0x0d,
	0xaa, 0x19, 0x06, 0x9b, 0x19, 0x06, 0xd3, 0x7c, 0x4d, 0x37, 0x22, 0xf2, 0x0c, 0xda, 0x5c, 0xcf,
	0x4f, 0x7a, 0x60, 0xe4, 0xed, 0xc0, 0x8c, 0x93, 0xd6, 0x54, 0xe7, 0xd9, 0x8a, 0x29, 0x26, 0x3c,
	0xd7, 0xb7, 0x9e, 0xe6, 0x2b, 0x4a, 0x08, 0xd8, 0x8a, 0x25, 0xd2, 0xdb, 0xf6, 0x9b, 0x23, 0x87,
	0x9a, 0x33, 0xd9, 0x87, 0xce, 0x52, 0x70, 0xcc, 0x63, 0xe9, 0xed, 0x98, 0xa2, 0x2d, 0xd3, 0x3e,
	0xba, 0xa1, 0xe4, 0x18, 0x00, 0xbf, 0x2b, 0xc1, 0x16, 0x7a, 0x12, 0x5e, 0xcf, 0x14, 0xfe, 0xbb,
	0x4f, 0xc7, 0xe8, 0xf4, 0xed, 0xc1, 0x4b, 0x70, 0x1e, 0xfa, 0x4f, 0xfa, 0xd0, 0xbc, 0xc3, 0x75,
	0xbd, 0x64, 0xfa, 0x48, 0x76, 0xa1, 0xb5, 0x62, 0x69, 0xb9, 0x19, 0x72, 0x15, 0xbc, 0x6e, 0xbc,
	0xb2, 0x4e, 0xdb, 0x60, 0xeb, 0x76, 0x0d, 0x39, 0xfc, 0x1b, 0xdd, 0xc7, 0x4c, 0xa1, 0x31, 0x83,
	0x5f, 0x4b, 0x94, 0x4a, 0xaf, 0x80, 0x31, 0x61, 0x19, 0x13, 0xb5, 0x51, 0x83, 0xc8, 0x09, 0xc0,
	0xe3, 0xba, 0x9b, 0xb2, 0xee, 0x64, 0xf0, 0x87, 0xcb, 0xb7, 0x5a, 0x32, 0x63, 0xf2, 0x8e, 0x3a,
	0xcb, 0xcd, 0xf1, 0xf0, 0x05, 0xd8, 0x7a, 0x97, 0x88, 0x0b, 0x9d, 0x68, 0x7e, 0x31, 0xff, 0x78,
	0x35, 0xef, 0x6f, 0xe9, 0x80, 0x86, 0xef, 0xa2, 0x0f, 0x53, 0xda, 0xb7, 0x88, 0x03, 0xad, 0xe9,
	0xd9, 0xec, 0x7c, 0xde, 0x6f, 0x1c, 0x06, 0x00, 0x8f, 0xfb, 0x40, 0xba, 0x60, 0xd3, 0x70, 0x7a,
	0xd6, 0xdf, 0xd2, 0x92, 0x2b, 0x7a, 0x7e, 0x19, 0xf6, 0x2d, 0x7d, 0x35, 0xfc, 0x1c, 0xbe, 0x89,
	0x2e, 0xc3, 0x7e, 0xe3, 0x74, 0xfc, 0xe5, 0x28, 0xe1, 0xea, 0xb6, 0xbc, 0x0e, 0x6e, 0x8a, 0x6c,
	0x9c, 0x61, 0x9e, 0xb3, 0xbc, 0x58, 0x55, 0x4f, 0x53, 0x5b, 0x3d, 0x2a, 0x15, 0x4f, 0xe5, 0x58,
	0xbf, 0xf0, 0xca, 0x65, 0xdb, 0x7c, 0x8e, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x86, 0x73, 0xb5,
	0x6f, 0xf5, 0x03, 0x00, 0x00,
}
