// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fk-module.proto

/*
Package fk_module is a generated protocol buffer package.

It is generated from these files:
	fk-module.proto

It has these top-level messages:
	QueryCapabilities
	Capabilities
	SensorCapabilities
	SensorReadingMeta
	BeginTakeReadings
	QueryReadingStatus
	ReadingStatus
	SensorReadings
	SensorReading
	WireMessageQuery
	Error
	WireMessageReply
*/
package fk_module

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

type QueryType int32

const (
	QueryType_QUERY_NONE                QueryType = 0
	QueryType_QUERY_CAPABILITIES        QueryType = 1
	QueryType_QUERY_BEGIN_TAKE_READINGS QueryType = 2
	QueryType_QUERY_READING_STATUS      QueryType = 3
)

var QueryType_name = map[int32]string{
	0: "QUERY_NONE",
	1: "QUERY_CAPABILITIES",
	2: "QUERY_BEGIN_TAKE_READINGS",
	3: "QUERY_READING_STATUS",
}
var QueryType_value = map[string]int32{
	"QUERY_NONE":                0,
	"QUERY_CAPABILITIES":        1,
	"QUERY_BEGIN_TAKE_READINGS": 2,
	"QUERY_READING_STATUS":      3,
}

func (x QueryType) String() string {
	return proto.EnumName(QueryType_name, int32(x))
}
func (QueryType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ReplyType int32

const (
	ReplyType_REPLY_NONE           ReplyType = 0
	ReplyType_REPLY_RETRY          ReplyType = 1
	ReplyType_REPLY_ERROR          ReplyType = 2
	ReplyType_REPLY_CAPABILITIES   ReplyType = 3
	ReplyType_REPLY_READING_STATUS ReplyType = 4
)

var ReplyType_name = map[int32]string{
	0: "REPLY_NONE",
	1: "REPLY_RETRY",
	2: "REPLY_ERROR",
	3: "REPLY_CAPABILITIES",
	4: "REPLY_READING_STATUS",
}
var ReplyType_value = map[string]int32{
	"REPLY_NONE":           0,
	"REPLY_RETRY":          1,
	"REPLY_ERROR":          2,
	"REPLY_CAPABILITIES":   3,
	"REPLY_READING_STATUS": 4,
}

func (x ReplyType) String() string {
	return proto.EnumName(ReplyType_name, int32(x))
}
func (ReplyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ModuleType int32

const (
	ModuleType_SENSOR         ModuleType = 0
	ModuleType_COMMUNICATIONS ModuleType = 1
)

var ModuleType_name = map[int32]string{
	0: "SENSOR",
	1: "COMMUNICATIONS",
}
var ModuleType_value = map[string]int32{
	"SENSOR":         0,
	"COMMUNICATIONS": 1,
}

func (x ModuleType) String() string {
	return proto.EnumName(ModuleType_name, int32(x))
}
func (ModuleType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ReadingState int32

const (
	ReadingState_IDLE  ReadingState = 0
	ReadingState_BEGIN ReadingState = 1
	ReadingState_BUSY  ReadingState = 2
	ReadingState_DONE  ReadingState = 3
)

var ReadingState_name = map[int32]string{
	0: "IDLE",
	1: "BEGIN",
	2: "BUSY",
	3: "DONE",
}
var ReadingState_value = map[string]int32{
	"IDLE":  0,
	"BEGIN": 1,
	"BUSY":  2,
	"DONE":  3,
}

func (x ReadingState) String() string {
	return proto.EnumName(ReadingState_name, int32(x))
}
func (ReadingState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ErrorType int32

const (
	ErrorType_NONE       ErrorType = 0
	ErrorType_GENERAL    ErrorType = 1
	ErrorType_UNEXPECTED ErrorType = 2
)

var ErrorType_name = map[int32]string{
	0: "NONE",
	1: "GENERAL",
	2: "UNEXPECTED",
}
var ErrorType_value = map[string]int32{
	"NONE":       0,
	"GENERAL":    1,
	"UNEXPECTED": 2,
}

func (x ErrorType) String() string {
	return proto.EnumName(ErrorType_name, int32(x))
}
func (ErrorType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type QueryCapabilities struct {
	Version int64 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
}

func (m *QueryCapabilities) Reset()                    { *m = QueryCapabilities{} }
func (m *QueryCapabilities) String() string            { return proto.CompactTextString(m) }
func (*QueryCapabilities) ProtoMessage()               {}
func (*QueryCapabilities) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *QueryCapabilities) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type Capabilities struct {
	Version int32      `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Type    ModuleType `protobuf:"varint,2,opt,name=type,enum=fk_module.ModuleType" json:"type,omitempty"`
	Name    string     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *Capabilities) Reset()                    { *m = Capabilities{} }
func (m *Capabilities) String() string            { return proto.CompactTextString(m) }
func (*Capabilities) ProtoMessage()               {}
func (*Capabilities) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Capabilities) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Capabilities) GetType() ModuleType {
	if m != nil {
		return m.Type
	}
	return ModuleType_SENSOR
}

func (m *Capabilities) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SensorCapabilities struct {
	NumberOfReadings int32                `protobuf:"varint,1,opt,name=numberOfReadings" json:"numberOfReadings,omitempty"`
	Readings         []*SensorReadingMeta `protobuf:"bytes,2,rep,name=readings" json:"readings,omitempty"`
}

func (m *SensorCapabilities) Reset()                    { *m = SensorCapabilities{} }
func (m *SensorCapabilities) String() string            { return proto.CompactTextString(m) }
func (*SensorCapabilities) ProtoMessage()               {}
func (*SensorCapabilities) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SensorCapabilities) GetNumberOfReadings() int32 {
	if m != nil {
		return m.NumberOfReadings
	}
	return 0
}

func (m *SensorCapabilities) GetReadings() []*SensorReadingMeta {
	if m != nil {
		return m.Readings
	}
	return nil
}

type SensorReadingMeta struct {
	Index int32  `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *SensorReadingMeta) Reset()                    { *m = SensorReadingMeta{} }
func (m *SensorReadingMeta) String() string            { return proto.CompactTextString(m) }
func (*SensorReadingMeta) ProtoMessage()               {}
func (*SensorReadingMeta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SensorReadingMeta) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *SensorReadingMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type BeginTakeReadings struct {
	Index int32 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
}

func (m *BeginTakeReadings) Reset()                    { *m = BeginTakeReadings{} }
func (m *BeginTakeReadings) String() string            { return proto.CompactTextString(m) }
func (*BeginTakeReadings) ProtoMessage()               {}
func (*BeginTakeReadings) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BeginTakeReadings) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type QueryReadingStatus struct {
}

func (m *QueryReadingStatus) Reset()                    { *m = QueryReadingStatus{} }
func (m *QueryReadingStatus) String() string            { return proto.CompactTextString(m) }
func (*QueryReadingStatus) ProtoMessage()               {}
func (*QueryReadingStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ReadingStatus struct {
	State ReadingState `protobuf:"varint,1,opt,name=state,enum=fk_module.ReadingState" json:"state,omitempty"`
}

func (m *ReadingStatus) Reset()                    { *m = ReadingStatus{} }
func (m *ReadingStatus) String() string            { return proto.CompactTextString(m) }
func (*ReadingStatus) ProtoMessage()               {}
func (*ReadingStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReadingStatus) GetState() ReadingState {
	if m != nil {
		return m.State
	}
	return ReadingState_IDLE
}

type SensorReadings struct {
	Time     int64            `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	Readings []*SensorReading `protobuf:"bytes,2,rep,name=readings" json:"readings,omitempty"`
}

func (m *SensorReadings) Reset()                    { *m = SensorReadings{} }
func (m *SensorReadings) String() string            { return proto.CompactTextString(m) }
func (*SensorReadings) ProtoMessage()               {}
func (*SensorReadings) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SensorReadings) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *SensorReadings) GetReadings() []*SensorReading {
	if m != nil {
		return m.Readings
	}
	return nil
}

type SensorReading struct {
	Index int32   `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	Value float32 `protobuf:"fixed32,2,opt,name=value" json:"value,omitempty"`
}

func (m *SensorReading) Reset()                    { *m = SensorReading{} }
func (m *SensorReading) String() string            { return proto.CompactTextString(m) }
func (*SensorReading) ProtoMessage()               {}
func (*SensorReading) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SensorReading) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *SensorReading) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type WireMessageQuery struct {
	Type               QueryType           `protobuf:"varint,1,opt,name=type,enum=fk_module.QueryType" json:"type,omitempty"`
	QueryCapabilities  *QueryCapabilities  `protobuf:"bytes,2,opt,name=queryCapabilities" json:"queryCapabilities,omitempty"`
	BeginTakeReadings  *BeginTakeReadings  `protobuf:"bytes,3,opt,name=beginTakeReadings" json:"beginTakeReadings,omitempty"`
	QueryReadingStatus *QueryReadingStatus `protobuf:"bytes,4,opt,name=queryReadingStatus" json:"queryReadingStatus,omitempty"`
}

func (m *WireMessageQuery) Reset()                    { *m = WireMessageQuery{} }
func (m *WireMessageQuery) String() string            { return proto.CompactTextString(m) }
func (*WireMessageQuery) ProtoMessage()               {}
func (*WireMessageQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *WireMessageQuery) GetType() QueryType {
	if m != nil {
		return m.Type
	}
	return QueryType_QUERY_NONE
}

func (m *WireMessageQuery) GetQueryCapabilities() *QueryCapabilities {
	if m != nil {
		return m.QueryCapabilities
	}
	return nil
}

func (m *WireMessageQuery) GetBeginTakeReadings() *BeginTakeReadings {
	if m != nil {
		return m.BeginTakeReadings
	}
	return nil
}

func (m *WireMessageQuery) GetQueryReadingStatus() *QueryReadingStatus {
	if m != nil {
		return m.QueryReadingStatus
	}
	return nil
}

type Error struct {
	Type    ErrorType `protobuf:"varint,1,opt,name=type,enum=fk_module.ErrorType" json:"type,omitempty"`
	Message string    `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Error) GetType() ErrorType {
	if m != nil {
		return m.Type
	}
	return ErrorType_NONE
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type WireMessageReply struct {
	Type               ReplyType           `protobuf:"varint,1,opt,name=type,enum=fk_module.ReplyType" json:"type,omitempty"`
	Error              *Error              `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Capabilities       *Capabilities       `protobuf:"bytes,3,opt,name=capabilities" json:"capabilities,omitempty"`
	SensorCapabilities *SensorCapabilities `protobuf:"bytes,4,opt,name=sensorCapabilities" json:"sensorCapabilities,omitempty"`
	ReadingStatus      *ReadingStatus      `protobuf:"bytes,5,opt,name=readingStatus" json:"readingStatus,omitempty"`
	SensorReadings     *SensorReadings     `protobuf:"bytes,6,opt,name=sensorReadings" json:"sensorReadings,omitempty"`
}

func (m *WireMessageReply) Reset()                    { *m = WireMessageReply{} }
func (m *WireMessageReply) String() string            { return proto.CompactTextString(m) }
func (*WireMessageReply) ProtoMessage()               {}
func (*WireMessageReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *WireMessageReply) GetType() ReplyType {
	if m != nil {
		return m.Type
	}
	return ReplyType_REPLY_NONE
}

func (m *WireMessageReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *WireMessageReply) GetCapabilities() *Capabilities {
	if m != nil {
		return m.Capabilities
	}
	return nil
}

func (m *WireMessageReply) GetSensorCapabilities() *SensorCapabilities {
	if m != nil {
		return m.SensorCapabilities
	}
	return nil
}

func (m *WireMessageReply) GetReadingStatus() *ReadingStatus {
	if m != nil {
		return m.ReadingStatus
	}
	return nil
}

func (m *WireMessageReply) GetSensorReadings() *SensorReadings {
	if m != nil {
		return m.SensorReadings
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryCapabilities)(nil), "fk_module.QueryCapabilities")
	proto.RegisterType((*Capabilities)(nil), "fk_module.Capabilities")
	proto.RegisterType((*SensorCapabilities)(nil), "fk_module.SensorCapabilities")
	proto.RegisterType((*SensorReadingMeta)(nil), "fk_module.SensorReadingMeta")
	proto.RegisterType((*BeginTakeReadings)(nil), "fk_module.BeginTakeReadings")
	proto.RegisterType((*QueryReadingStatus)(nil), "fk_module.QueryReadingStatus")
	proto.RegisterType((*ReadingStatus)(nil), "fk_module.ReadingStatus")
	proto.RegisterType((*SensorReadings)(nil), "fk_module.SensorReadings")
	proto.RegisterType((*SensorReading)(nil), "fk_module.SensorReading")
	proto.RegisterType((*WireMessageQuery)(nil), "fk_module.WireMessageQuery")
	proto.RegisterType((*Error)(nil), "fk_module.Error")
	proto.RegisterType((*WireMessageReply)(nil), "fk_module.WireMessageReply")
	proto.RegisterEnum("fk_module.QueryType", QueryType_name, QueryType_value)
	proto.RegisterEnum("fk_module.ReplyType", ReplyType_name, ReplyType_value)
	proto.RegisterEnum("fk_module.ModuleType", ModuleType_name, ModuleType_value)
	proto.RegisterEnum("fk_module.ReadingState", ReadingState_name, ReadingState_value)
	proto.RegisterEnum("fk_module.ErrorType", ErrorType_name, ErrorType_value)
}

func init() { proto.RegisterFile("fk-module.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 735 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0xdd, 0x4e, 0xdb, 0x4a,
	0x10, 0xc6, 0x4e, 0x0c, 0x64, 0x80, 0xb0, 0x59, 0xe5, 0x9c, 0x63, 0xa4, 0x83, 0x14, 0xf9, 0xe2,
	0x28, 0x44, 0x07, 0x2e, 0xd2, 0x4a, 0xad, 0x84, 0x8a, 0xe4, 0x24, 0x2b, 0xe4, 0x92, 0x38, 0xb0,
	0x76, 0xd4, 0xd2, 0x9b, 0xc8, 0x29, 0x4b, 0x64, 0x91, 0x3f, 0x6c, 0x07, 0x95, 0xbe, 0x4e, 0x1f,
	0xb1, 0x2f, 0x50, 0x79, 0xd7, 0x31, 0xde, 0x38, 0x70, 0x15, 0xef, 0xcc, 0x37, 0x33, 0xdf, 0x7c,
	0x33, 0xbb, 0x81, 0xc3, 0xfb, 0x87, 0xd3, 0xe9, 0xfc, 0x6e, 0x39, 0x61, 0x67, 0x8b, 0x60, 0x1e,
	0xcd, 0x71, 0xe9, 0xfe, 0x61, 0x28, 0x0c, 0xc6, 0x29, 0x54, 0x6e, 0x96, 0x2c, 0x78, 0x6e, 0x7b,
	0x0b, 0x6f, 0xe4, 0x4f, 0xfc, 0xc8, 0x67, 0x21, 0xd6, 0x61, 0xe7, 0x89, 0x05, 0xa1, 0x3f, 0x9f,
	0xe9, 0x4a, 0x4d, 0xa9, 0x17, 0xe8, 0xea, 0x68, 0x8c, 0x61, 0xff, 0x2d, 0xa4, 0x96, 0x22, 0xf1,
	0x09, 0x14, 0xa3, 0xe7, 0x05, 0xd3, 0xd5, 0x9a, 0x52, 0x2f, 0x37, 0xff, 0x3a, 0x4b, 0x4b, 0x9e,
	0xf5, 0xf8, 0x8f, 0xfb, 0xbc, 0x60, 0x94, 0x43, 0x30, 0x86, 0xe2, 0xcc, 0x9b, 0x32, 0xbd, 0x50,
	0x53, 0xea, 0x25, 0xca, 0xbf, 0x8d, 0x9f, 0x80, 0x1d, 0x36, 0x0b, 0xe7, 0x81, 0x54, 0xae, 0x01,
	0x68, 0xb6, 0x9c, 0x8e, 0x58, 0xd0, 0xbf, 0xa7, 0xcc, 0xbb, 0xf3, 0x67, 0xe3, 0x30, 0xa9, 0x9b,
	0xb3, 0xe3, 0x8f, 0xb0, 0x1b, 0xac, 0x30, 0x6a, 0xad, 0x50, 0xdf, 0x6b, 0xfe, 0x9b, 0x21, 0x21,
	0x92, 0x27, 0xe0, 0x1e, 0x8b, 0x3c, 0x9a, 0xa2, 0x8d, 0x4f, 0x50, 0xc9, 0xb9, 0x71, 0x15, 0x34,
	0x7f, 0x76, 0xc7, 0x7e, 0x24, 0xf5, 0xc4, 0x21, 0xa5, 0xae, 0x66, 0xa8, 0x9f, 0x40, 0xa5, 0xc5,
	0xc6, 0xfe, 0xcc, 0xf5, 0x1e, 0x58, 0xca, 0x66, 0x63, 0xb8, 0x51, 0x05, 0xcc, 0xd5, 0x4f, 0x60,
	0x4e, 0xe4, 0x45, 0xcb, 0xd0, 0xb8, 0x80, 0x03, 0xc9, 0x80, 0x4f, 0x41, 0x0b, 0x23, 0x2f, 0x62,
	0x3c, 0xb8, 0xdc, 0xfc, 0x27, 0xd3, 0x47, 0x06, 0xc8, 0xa8, 0x40, 0x19, 0xdf, 0xa0, 0x2c, 0xf1,
	0x0f, 0x63, 0x9a, 0x91, 0x3f, 0x65, 0xc9, 0x34, 0xf9, 0x37, 0x7e, 0x9f, 0xd3, 0x47, 0x7f, 0x4d,
	0x9f, 0x8c, 0x36, 0xe7, 0x70, 0x20, 0xb9, 0x5e, 0xd1, 0xa5, 0x0a, 0xda, 0x93, 0x37, 0x59, 0x0a,
	0x61, 0x54, 0x2a, 0x0e, 0xc6, 0x2f, 0x15, 0xd0, 0x17, 0x3f, 0x60, 0x3d, 0x16, 0x86, 0xde, 0x98,
	0xf1, 0xd6, 0x71, 0x3d, 0x59, 0x14, 0xd1, 0x5b, 0x35, 0xc3, 0x81, 0xfb, 0x33, 0x7b, 0xf2, 0x19,
	0x2a, 0x8f, 0xeb, 0xbb, 0xca, 0x0b, 0xc8, 0xa3, 0xcd, 0xed, 0x33, 0xcd, 0x87, 0xc5, 0xb9, 0x46,
	0xeb, 0x43, 0xe2, 0x0b, 0x28, 0xe7, 0xca, 0x0d, 0x92, 0xe6, 0xc3, 0x70, 0x0f, 0xf0, 0x63, 0x6e,
	0x8a, 0x7a, 0x91, 0x27, 0x3b, 0x5e, 0x27, 0x26, 0x81, 0xe8, 0x86, 0x40, 0xe3, 0x0a, 0x34, 0x12,
	0x04, 0xf3, 0xe0, 0x0d, 0x65, 0xb8, 0x3f, 0xa3, 0x8c, 0x0e, 0x3b, 0x53, 0xa1, 0x69, 0xb2, 0x89,
	0xab, 0xa3, 0xf1, 0x5b, 0x96, 0x9c, 0xb2, 0xc5, 0xe4, 0x2d, 0xc9, 0xb9, 0x3f, 0x93, 0xf8, 0x3f,
	0xd0, 0x58, 0x5c, 0x2b, 0x91, 0x19, 0xad, 0x73, 0xa0, 0xc2, 0x8d, 0xcf, 0x61, 0xff, 0x7b, 0x76,
	0x2a, 0x42, 0xc9, 0xec, 0xa2, 0x4a, 0x03, 0x91, 0xc0, 0xb1, 0x7e, 0x61, 0xee, 0xae, 0x6f, 0xd0,
	0x2f, 0xff, 0x20, 0xd0, 0x0d, 0x81, 0xf8, 0x02, 0x0e, 0x02, 0x69, 0x12, 0x1a, 0xcf, 0xa4, 0x6f,
	0xbe, 0x35, 0xcb, 0x90, 0xca, 0x70, 0x6c, 0x42, 0x39, 0x94, 0xae, 0x8f, 0xbe, 0xcd, 0x13, 0x1c,
	0xbd, 0x76, 0x3d, 0x42, 0xba, 0x16, 0xd0, 0x98, 0x40, 0x29, 0x5d, 0x5e, 0x5c, 0x06, 0xb8, 0x19,
	0x10, 0x7a, 0x3b, 0xb4, 0xfb, 0x36, 0x41, 0x5b, 0xf8, 0x6f, 0xc0, 0xe2, 0xdc, 0x36, 0xaf, 0xcd,
	0x96, 0xd5, 0xb5, 0x5c, 0x8b, 0x38, 0x48, 0xc1, 0xc7, 0x70, 0x24, 0xec, 0x2d, 0x72, 0x69, 0xd9,
	0x43, 0xd7, 0xbc, 0x22, 0x43, 0x4a, 0xcc, 0x8e, 0x65, 0x5f, 0x3a, 0x48, 0xc5, 0x3a, 0x54, 0x85,
	0x3b, 0xb1, 0x0d, 0x1d, 0xd7, 0x74, 0x07, 0x0e, 0x2a, 0x34, 0xe6, 0x50, 0x4a, 0xe7, 0x16, 0x57,
	0xa3, 0xe4, 0xba, 0x9b, 0x56, 0x3b, 0x84, 0x3d, 0x71, 0xa6, 0xc4, 0xa5, 0xb7, 0x48, 0x79, 0x31,
	0x10, 0x4a, 0xfb, 0x14, 0xa9, 0x31, 0x1f, 0x61, 0x90, 0xf8, 0x14, 0xe2, 0x82, 0xab, 0x48, 0xa9,
	0x60, 0xb1, 0xf1, 0x3f, 0xc0, 0xcb, 0x23, 0x8e, 0x01, 0xb6, 0x1d, 0x62, 0x3b, 0x7d, 0x8a, 0xb6,
	0x30, 0x86, 0x72, 0xbb, 0xdf, 0xeb, 0x0d, 0x6c, 0xab, 0x6d, 0xba, 0x56, 0xdf, 0x76, 0x90, 0xd2,
	0xf8, 0x00, 0xfb, 0xd9, 0x57, 0x0a, 0xef, 0x42, 0xd1, 0xea, 0x74, 0x63, 0x6e, 0x25, 0xd0, 0x78,
	0xaf, 0x48, 0x89, 0x8d, 0xad, 0x81, 0x73, 0x8b, 0xd4, 0xf8, 0xab, 0x13, 0x53, 0x2f, 0x34, 0x9a,
	0x50, 0x4a, 0x17, 0x3d, 0x36, 0x27, 0x1d, 0xed, 0xc1, 0xce, 0x25, 0xb1, 0x09, 0x35, 0xbb, 0x48,
	0x89, 0xdb, 0x1d, 0xd8, 0xe4, 0xeb, 0x35, 0x69, 0xbb, 0xa4, 0x83, 0xd4, 0xd1, 0x36, 0xff, 0x87,
	0x7b, 0xf7, 0x27, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x34, 0xd9, 0xd4, 0xf4, 0x06, 0x00, 0x00,
}
