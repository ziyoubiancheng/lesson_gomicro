// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rate.proto

package rate

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

type Request struct {
	HotelIds             []string `protobuf:"bytes,1,rep,name=hotelIds,proto3" json:"hotelIds,omitempty"`
	InDate               string   `protobuf:"bytes,2,opt,name=inDate,proto3" json:"inDate,omitempty"`
	OutDate              string   `protobuf:"bytes,3,opt,name=outDate,proto3" json:"outDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_426335fde4cae2d1, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetHotelIds() []string {
	if m != nil {
		return m.HotelIds
	}
	return nil
}

func (m *Request) GetInDate() string {
	if m != nil {
		return m.InDate
	}
	return ""
}

func (m *Request) GetOutDate() string {
	if m != nil {
		return m.OutDate
	}
	return ""
}

type Result struct {
	RatePlans            []*RatePlan `protobuf:"bytes,1,rep,name=ratePlans,proto3" json:"ratePlans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_426335fde4cae2d1, []int{1}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetRatePlans() []*RatePlan {
	if m != nil {
		return m.RatePlans
	}
	return nil
}

type RatePlan struct {
	HotelId              string    `protobuf:"bytes,1,opt,name=hotelId,proto3" json:"hotelId,omitempty"`
	Code                 string    `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	InDate               string    `protobuf:"bytes,3,opt,name=inDate,proto3" json:"inDate,omitempty"`
	OutDate              string    `protobuf:"bytes,4,opt,name=outDate,proto3" json:"outDate,omitempty"`
	RoomType             *RoomType `protobuf:"bytes,5,opt,name=roomType,proto3" json:"roomType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RatePlan) Reset()         { *m = RatePlan{} }
func (m *RatePlan) String() string { return proto.CompactTextString(m) }
func (*RatePlan) ProtoMessage()    {}
func (*RatePlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_426335fde4cae2d1, []int{2}
}

func (m *RatePlan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RatePlan.Unmarshal(m, b)
}
func (m *RatePlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RatePlan.Marshal(b, m, deterministic)
}
func (m *RatePlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RatePlan.Merge(m, src)
}
func (m *RatePlan) XXX_Size() int {
	return xxx_messageInfo_RatePlan.Size(m)
}
func (m *RatePlan) XXX_DiscardUnknown() {
	xxx_messageInfo_RatePlan.DiscardUnknown(m)
}

var xxx_messageInfo_RatePlan proto.InternalMessageInfo

func (m *RatePlan) GetHotelId() string {
	if m != nil {
		return m.HotelId
	}
	return ""
}

func (m *RatePlan) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *RatePlan) GetInDate() string {
	if m != nil {
		return m.InDate
	}
	return ""
}

func (m *RatePlan) GetOutDate() string {
	if m != nil {
		return m.OutDate
	}
	return ""
}

func (m *RatePlan) GetRoomType() *RoomType {
	if m != nil {
		return m.RoomType
	}
	return nil
}

type RoomType struct {
	BookableRate         float64  `protobuf:"fixed64,1,opt,name=bookableRate,proto3" json:"bookableRate,omitempty"`
	TotalRate            float64  `protobuf:"fixed64,2,opt,name=totalRate,proto3" json:"totalRate,omitempty"`
	TotalRateInclusive   float64  `protobuf:"fixed64,3,opt,name=totalRateInclusive,proto3" json:"totalRateInclusive,omitempty"`
	Code                 string   `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Currency             string   `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	RoomDescription      string   `protobuf:"bytes,6,opt,name=roomDescription,proto3" json:"roomDescription,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomType) Reset()         { *m = RoomType{} }
func (m *RoomType) String() string { return proto.CompactTextString(m) }
func (*RoomType) ProtoMessage()    {}
func (*RoomType) Descriptor() ([]byte, []int) {
	return fileDescriptor_426335fde4cae2d1, []int{3}
}

func (m *RoomType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomType.Unmarshal(m, b)
}
func (m *RoomType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomType.Marshal(b, m, deterministic)
}
func (m *RoomType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomType.Merge(m, src)
}
func (m *RoomType) XXX_Size() int {
	return xxx_messageInfo_RoomType.Size(m)
}
func (m *RoomType) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomType.DiscardUnknown(m)
}

var xxx_messageInfo_RoomType proto.InternalMessageInfo

func (m *RoomType) GetBookableRate() float64 {
	if m != nil {
		return m.BookableRate
	}
	return 0
}

func (m *RoomType) GetTotalRate() float64 {
	if m != nil {
		return m.TotalRate
	}
	return 0
}

func (m *RoomType) GetTotalRateInclusive() float64 {
	if m != nil {
		return m.TotalRateInclusive
	}
	return 0
}

func (m *RoomType) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *RoomType) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *RoomType) GetRoomDescription() string {
	if m != nil {
		return m.RoomDescription
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "rate.Request")
	proto.RegisterType((*Result)(nil), "rate.Result")
	proto.RegisterType((*RatePlan)(nil), "rate.RatePlan")
	proto.RegisterType((*RoomType)(nil), "rate.RoomType")
}

func init() { proto.RegisterFile("rate.proto", fileDescriptor_426335fde4cae2d1) }

var fileDescriptor_426335fde4cae2d1 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xd1, 0x4a, 0xf3, 0x30,
	0x14, 0xc7, 0xc9, 0xb7, 0x7e, 0x5d, 0x7b, 0x9c, 0x0a, 0xe7, 0x42, 0xc2, 0xf0, 0x62, 0xf4, 0xc6,
	0x21, 0x32, 0x61, 0x82, 0x4f, 0x30, 0x90, 0xdd, 0xc9, 0x41, 0xf0, 0xba, 0xed, 0x02, 0x16, 0x63,
	0x53, 0x93, 0x54, 0xd8, 0x8b, 0xf8, 0x68, 0x3e, 0x8f, 0x24, 0x4d, 0xbb, 0x4e, 0xf4, 0xee, 0xfc,
	0xfe, 0x27, 0x9c, 0xfe, 0x92, 0x1e, 0x00, 0x9d, 0x5b, 0xb1, 0x6a, 0xb4, 0xb2, 0x0a, 0x23, 0x57,
	0x67, 0xcf, 0x30, 0x25, 0xf1, 0xde, 0x0a, 0x63, 0x71, 0x0e, 0xc9, 0x8b, 0xb2, 0x42, 0x6e, 0x77,
	0x86, 0xb3, 0xc5, 0x64, 0x99, 0xd2, 0xc0, 0x78, 0x01, 0x71, 0x55, 0x6f, 0x72, 0x2b, 0xf8, 0xbf,
	0x05, 0x5b, 0xa6, 0x14, 0x08, 0x39, 0x4c, 0x55, 0x6b, 0x7d, 0x63, 0xe2, 0x1b, 0x3d, 0x66, 0xf7,
	0x10, 0x93, 0x30, 0xad, 0xb4, 0x78, 0x03, 0xa9, 0xfb, 0xd4, 0xa3, 0xcc, 0xeb, 0x6e, 0xf0, 0xc9,
	0xfa, 0x6c, 0xe5, 0x45, 0x28, 0xc4, 0x74, 0x38, 0x90, 0x7d, 0x32, 0x48, 0xfa, 0xdc, 0x8d, 0x0f,
	0x0a, 0x9c, 0x75, 0xe3, 0x03, 0x22, 0x42, 0x54, 0xaa, 0x5d, 0xaf, 0xe3, 0xeb, 0x91, 0xe4, 0xe4,
	0x2f, 0xc9, 0xe8, 0x48, 0x12, 0xaf, 0x21, 0xd1, 0x4a, 0xbd, 0x3d, 0xed, 0x1b, 0xc1, 0xff, 0x2f,
	0xd8, 0xc8, 0x2c, 0xa4, 0x34, 0xf4, 0xb3, 0x2f, 0x27, 0x16, 0x00, 0x33, 0x98, 0x15, 0x4a, 0xbd,
	0xe6, 0x85, 0x14, 0x4e, 0xd6, 0xdb, 0x31, 0x3a, 0xca, 0xf0, 0x12, 0x52, 0xab, 0x6c, 0x2e, 0xa9,
	0x7f, 0x36, 0x46, 0x87, 0x00, 0x57, 0x80, 0x03, 0x6c, 0xeb, 0x52, 0xb6, 0xa6, 0xfa, 0xe8, 0xc4,
	0x19, 0xfd, 0xd2, 0x19, 0x2e, 0x1c, 0x8d, 0x2e, 0x3c, 0x87, 0xa4, 0x6c, 0xb5, 0x16, 0x75, 0xb9,
	0xf7, 0xfa, 0x29, 0x0d, 0x8c, 0x4b, 0x38, 0x77, 0xea, 0x1b, 0x61, 0x4a, 0x5d, 0x35, 0xb6, 0x52,
	0x35, 0x8f, 0xfd, 0x91, 0x9f, 0xf1, 0xfa, 0x16, 0x22, 0x6f, 0x74, 0x05, 0xc9, 0x83, 0xb0, 0xae,
	0x34, 0x78, 0x1a, 0x9e, 0xa1, 0x5b, 0x8d, 0xf9, 0xac, 0x47, 0xf7, 0x43, 0x8b, 0xd8, 0x2f, 0xd0,
	0xdd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0x81, 0x29, 0x45, 0x4e, 0x02, 0x00, 0x00,
}
