// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fk-module.proto

/*
Package fk_module is a generated protocol buffer package.

It is generated from these files:
	fk-module.proto

It has these top-level messages:
	QueryCapabilities
	QuerySensorCapabilities
	Capabilities
	SensorCapabilities
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
	QueryType_QUERY_SENSOR_CAPABILITIES QueryType = 2
	QueryType_QUERY_BEGIN_TAKE_READINGS QueryType = 3
	QueryType_QUERY_READING_STATUS      QueryType = 4
)

var QueryType_name = map[int32]string{
	0: "QUERY_NONE",
	1: "QUERY_CAPABILITIES",
	2: "QUERY_SENSOR_CAPABILITIES",
	3: "QUERY_BEGIN_TAKE_READINGS",
	4: "QUERY_READING_STATUS",
}
var QueryType_value = map[string]int32{
	"QUERY_NONE":                0,
	"QUERY_CAPABILITIES":        1,
	"QUERY_SENSOR_CAPABILITIES": 2,
	"QUERY_BEGIN_TAKE_READINGS": 3,
	"QUERY_READING_STATUS":      4,
}

func (x QueryType) String() string {
	return proto.EnumName(QueryType_name, int32(x))
}
func (QueryType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ReplyType int32

const (
	ReplyType_REPLY_NONE                ReplyType = 0
	ReplyType_REPLY_RETRY               ReplyType = 1
	ReplyType_REPLY_ERROR               ReplyType = 2
	ReplyType_REPLY_CAPABILITIES        ReplyType = 3
	ReplyType_REPLY_SENSOR_CAPABILITIES ReplyType = 4
	ReplyType_REPLY_READING_STATUS      ReplyType = 5
)

var ReplyType_name = map[int32]string{
	0: "REPLY_NONE",
	1: "REPLY_RETRY",
	2: "REPLY_ERROR",
	3: "REPLY_CAPABILITIES",
	4: "REPLY_SENSOR_CAPABILITIES",
	5: "REPLY_READING_STATUS",
}
var ReplyType_value = map[string]int32{
	"REPLY_NONE":                0,
	"REPLY_RETRY":               1,
	"REPLY_ERROR":               2,
	"REPLY_CAPABILITIES":        3,
	"REPLY_SENSOR_CAPABILITIES": 4,
	"REPLY_READING_STATUS":      5,
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
	Version    int32  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	CallerTime uint32 `protobuf:"varint,2,opt,name=callerTime" json:"callerTime,omitempty"`
}

func (m *QueryCapabilities) Reset()                    { *m = QueryCapabilities{} }
func (m *QueryCapabilities) String() string            { return proto.CompactTextString(m) }
func (*QueryCapabilities) ProtoMessage()               {}
func (*QueryCapabilities) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *QueryCapabilities) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *QueryCapabilities) GetCallerTime() uint32 {
	if m != nil {
		return m.CallerTime
	}
	return 0
}

type QuerySensorCapabilities struct {
	Sensor int32 `protobuf:"varint,1,opt,name=sensor" json:"sensor,omitempty"`
}

func (m *QuerySensorCapabilities) Reset()                    { *m = QuerySensorCapabilities{} }
func (m *QuerySensorCapabilities) String() string            { return proto.CompactTextString(m) }
func (*QuerySensorCapabilities) ProtoMessage()               {}
func (*QuerySensorCapabilities) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *QuerySensorCapabilities) GetSensor() int32 {
	if m != nil {
		return m.Sensor
	}
	return 0
}

type Capabilities struct {
	Version         int32      `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Type            ModuleType `protobuf:"varint,2,opt,name=type,enum=fk_module.ModuleType" json:"type,omitempty"`
	Name            string     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	NumberOfSensors int32      `protobuf:"varint,4,opt,name=numberOfSensors" json:"numberOfSensors,omitempty"`
}

func (m *Capabilities) Reset()                    { *m = Capabilities{} }
func (m *Capabilities) String() string            { return proto.CompactTextString(m) }
func (*Capabilities) ProtoMessage()               {}
func (*Capabilities) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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

func (m *Capabilities) GetNumberOfSensors() int32 {
	if m != nil {
		return m.NumberOfSensors
	}
	return 0
}

type SensorCapabilities struct {
	Id            int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	UnitOfMeasure string `protobuf:"bytes,3,opt,name=unitOfMeasure" json:"unitOfMeasure,omitempty"`
}

func (m *SensorCapabilities) Reset()                    { *m = SensorCapabilities{} }
func (m *SensorCapabilities) String() string            { return proto.CompactTextString(m) }
func (*SensorCapabilities) ProtoMessage()               {}
func (*SensorCapabilities) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SensorCapabilities) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SensorCapabilities) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SensorCapabilities) GetUnitOfMeasure() string {
	if m != nil {
		return m.UnitOfMeasure
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
}

func (m *SensorReadings) Reset()                    { *m = SensorReadings{} }
func (m *SensorReadings) String() string            { return proto.CompactTextString(m) }
func (*SensorReadings) ProtoMessage()               {}
func (*SensorReadings) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type SensorReading struct {
	Sensor int32   `protobuf:"varint,1,opt,name=sensor" json:"sensor,omitempty"`
	Time   uint32  `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	Value  float32 `protobuf:"fixed32,3,opt,name=value" json:"value,omitempty"`
}

func (m *SensorReading) Reset()                    { *m = SensorReading{} }
func (m *SensorReading) String() string            { return proto.CompactTextString(m) }
func (*SensorReading) ProtoMessage()               {}
func (*SensorReading) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SensorReading) GetSensor() int32 {
	if m != nil {
		return m.Sensor
	}
	return 0
}

func (m *SensorReading) GetTime() uint32 {
	if m != nil {
		return m.Time
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
	Type                    QueryType                `protobuf:"varint,1,opt,name=type,enum=fk_module.QueryType" json:"type,omitempty"`
	QueryCapabilities       *QueryCapabilities       `protobuf:"bytes,2,opt,name=queryCapabilities" json:"queryCapabilities,omitempty"`
	QuerySensorCapabilities *QuerySensorCapabilities `protobuf:"bytes,3,opt,name=querySensorCapabilities" json:"querySensorCapabilities,omitempty"`
	BeginTakeReadings       *BeginTakeReadings       `protobuf:"bytes,4,opt,name=beginTakeReadings" json:"beginTakeReadings,omitempty"`
	QueryReadingStatus      *QueryReadingStatus      `protobuf:"bytes,5,opt,name=queryReadingStatus" json:"queryReadingStatus,omitempty"`
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

func (m *WireMessageQuery) GetQuerySensorCapabilities() *QuerySensorCapabilities {
	if m != nil {
		return m.QuerySensorCapabilities
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
	SensorReading      *SensorReading      `protobuf:"bytes,6,opt,name=sensorReading" json:"sensorReading,omitempty"`
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

func (m *WireMessageReply) GetSensorReading() *SensorReading {
	if m != nil {
		return m.SensorReading
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryCapabilities)(nil), "fk_module.QueryCapabilities")
	proto.RegisterType((*QuerySensorCapabilities)(nil), "fk_module.QuerySensorCapabilities")
	proto.RegisterType((*Capabilities)(nil), "fk_module.Capabilities")
	proto.RegisterType((*SensorCapabilities)(nil), "fk_module.SensorCapabilities")
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
	// 782 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x5b, 0x6b, 0xe3, 0x46,
	0x14, 0x8e, 0x64, 0xcb, 0xa9, 0x8f, 0x63, 0x47, 0x19, 0xdc, 0x44, 0x85, 0xa6, 0x04, 0x51, 0x8a,
	0x63, 0x9a, 0x40, 0xdd, 0x87, 0x3e, 0x14, 0x02, 0xbe, 0x0c, 0x41, 0x8d, 0x2d, 0x27, 0x23, 0x99,
	0x36, 0x50, 0x6a, 0xe4, 0x78, 0x1c, 0x44, 0x6c, 0xd9, 0xd1, 0x25, 0x34, 0x8f, 0xfb, 0xbc, 0xb0,
	0xb0, 0xff, 0x74, 0x7f, 0xc2, 0x32, 0x33, 0xb2, 0x2d, 0x59, 0x4a, 0xd8, 0x27, 0xcf, 0xf9, 0xce,
	0xed, 0x3b, 0xe7, 0x9b, 0x91, 0xe1, 0x70, 0xf6, 0x74, 0xb1, 0x58, 0x4e, 0xa3, 0x39, 0xbd, 0x5c,
	0xf9, 0xcb, 0x70, 0x89, 0xca, 0xb3, 0xa7, 0xb1, 0x00, 0xf4, 0x01, 0x1c, 0xdd, 0x45, 0xd4, 0x7f,
	0xed, 0x3a, 0x2b, 0x67, 0xe2, 0xce, 0xdd, 0xd0, 0xa5, 0x01, 0xd2, 0x60, 0xff, 0x85, 0xfa, 0x81,
	0xbb, 0xf4, 0x34, 0xe9, 0x4c, 0x6a, 0x28, 0x64, 0x6d, 0xa2, 0x9f, 0x00, 0x1e, 0x9c, 0xf9, 0x9c,
	0xfa, 0xb6, 0xbb, 0xa0, 0x9a, 0x7c, 0x26, 0x35, 0xaa, 0x24, 0x81, 0xe8, 0xbf, 0xc1, 0x09, 0x2f,
	0x67, 0x51, 0x2f, 0x58, 0xfa, 0xa9, 0xa2, 0xc7, 0x50, 0x0a, 0x38, 0x1a, 0xd7, 0x8c, 0x2d, 0xfd,
	0xb3, 0x04, 0x07, 0xdf, 0xd8, 0xfd, 0x1c, 0x8a, 0xe1, 0xeb, 0x4a, 0xf4, 0xad, 0xb5, 0xbe, 0xbf,
	0xdc, 0x8c, 0x71, 0x39, 0xe0, 0x3f, 0xf6, 0xeb, 0x8a, 0x12, 0x1e, 0x82, 0x10, 0x14, 0x3d, 0x67,
	0x41, 0xb5, 0xc2, 0x99, 0xd4, 0x28, 0x13, 0x7e, 0x46, 0x0d, 0x38, 0xf4, 0xa2, 0xc5, 0x84, 0xfa,
	0xc3, 0x99, 0xe0, 0x17, 0x68, 0x45, 0xde, 0x60, 0x17, 0xd6, 0xff, 0x03, 0x94, 0x33, 0x41, 0x0d,
	0x64, 0x77, 0x1a, 0x73, 0x92, 0xdd, 0xe9, 0xa6, 0x87, 0x9c, 0xe8, 0xf1, 0x33, 0x54, 0x23, 0xcf,
	0x0d, 0x87, 0xb3, 0x01, 0x75, 0x82, 0xc8, 0x5f, 0x13, 0x48, 0x83, 0xfa, 0x39, 0x1c, 0x75, 0xe8,
	0xa3, 0xeb, 0xd9, 0xce, 0x13, 0x25, 0xd4, 0x99, 0xba, 0xde, 0x63, 0x80, 0xea, 0xa0, 0xb8, 0xde,
	0x94, 0xfe, 0x1f, 0x77, 0x10, 0x86, 0x5e, 0x07, 0xc4, 0x37, 0x1a, 0x87, 0x59, 0xa1, 0x13, 0x46,
	0x81, 0x7e, 0x05, 0xd5, 0x14, 0x80, 0x2e, 0x40, 0x09, 0x42, 0x27, 0xa4, 0x3c, 0xb9, 0xd6, 0x3a,
	0x49, 0xec, 0x26, 0x11, 0x48, 0x89, 0x88, 0xd2, 0x55, 0xa8, 0x89, 0x01, 0xd7, 0xdd, 0xf5, 0x3b,
	0xa8, 0xa6, 0x90, 0xb7, 0xf4, 0x62, 0x53, 0x87, 0x5b, 0xf1, 0xf9, 0x99, 0x51, 0x7f, 0x71, 0xe6,
	0x91, 0x98, 0x56, 0x26, 0xc2, 0xd0, 0x3f, 0x14, 0x40, 0xfd, 0xdb, 0xf5, 0xe9, 0x80, 0x06, 0x81,
	0xf3, 0x48, 0xf9, 0x18, 0xa8, 0x11, 0x6b, 0x28, 0x78, 0xd6, 0x13, 0x3c, 0xb9, 0x3f, 0x21, 0xe1,
	0x5f, 0x70, 0xf4, 0xbc, 0x7b, 0x35, 0x79, 0xd7, 0x4a, 0xeb, 0xc7, 0xdd, 0xb4, 0x64, 0x0c, 0xc9,
	0xa6, 0xa1, 0x7f, 0xe1, 0xe4, 0x39, 0xff, 0x5e, 0x72, 0xca, 0x95, 0x96, 0xbe, 0x5b, 0x31, 0x1b,
	0x49, 0xde, 0x2a, 0xc1, 0x98, 0x4e, 0x76, 0xe5, 0xe4, 0x57, 0x2b, 0xcd, 0x34, 0x23, 0x39, 0xc9,
	0xa6, 0xa1, 0x01, 0xa0, 0xe7, 0x8c, 0xde, 0x9a, 0xc2, 0x8b, 0x9d, 0xee, 0x92, 0x4c, 0x05, 0x91,
	0x9c, 0x44, 0xfd, 0x06, 0x14, 0xec, 0xfb, 0x4b, 0xff, 0x9d, 0xbd, 0x73, 0x7f, 0x62, 0xef, 0x1a,
	0xec, 0x2f, 0x84, 0x62, 0xf1, 0xcd, 0x5e, 0x9b, 0xfa, 0x17, 0x39, 0x25, 0x28, 0xa1, 0xab, 0xf9,
	0x7b, 0x82, 0x72, 0x7f, 0xa2, 0xf0, 0x2f, 0xa0, 0x50, 0xd6, 0x2b, 0x16, 0x51, 0xdd, 0xe5, 0x40,
	0x84, 0x1b, 0xfd, 0x09, 0x07, 0x0f, 0x59, 0x85, 0x92, 0x57, 0x3a, 0x25, 0x4b, 0x2a, 0x98, 0xed,
	0x2f, 0xc8, 0x8a, 0x5c, 0xcc, 0xec, 0x2f, 0x47, 0xdf, 0x9c, 0x44, 0x74, 0x05, 0x55, 0x3f, 0x47,
	0x09, 0x2d, 0xff, 0x7d, 0x45, 0x01, 0x49, 0x87, 0xb3, 0xfc, 0x20, 0xf9, 0xac, 0xb4, 0x52, 0x26,
	0x3f, 0xf5, 0xec, 0x48, 0x3a, 0xbc, 0xf9, 0x51, 0x82, 0xf2, 0xe6, 0x61, 0xa0, 0x1a, 0xc0, 0xdd,
	0x08, 0x93, 0xfb, 0xb1, 0x39, 0x34, 0xb1, 0xba, 0x87, 0x8e, 0x01, 0x09, 0xbb, 0xdb, 0xbe, 0x6d,
	0x77, 0x8c, 0xbe, 0x61, 0x1b, 0xd8, 0x52, 0x25, 0x74, 0x0a, 0x3f, 0x08, 0xdc, 0xc2, 0xa6, 0x35,
	0x24, 0x69, 0xb7, 0xbc, 0x75, 0x77, 0xf0, 0xb5, 0x61, 0x8e, 0xed, 0xf6, 0x0d, 0x1e, 0x13, 0xdc,
	0xee, 0x19, 0xe6, 0xb5, 0xa5, 0x16, 0x90, 0x06, 0x75, 0xe1, 0x8e, 0xb1, 0xb1, 0x65, 0xb7, 0xed,
	0x91, 0xa5, 0x16, 0x9b, 0x9f, 0x24, 0x28, 0x6f, 0x54, 0x65, 0x6c, 0x08, 0xbe, 0xed, 0x6f, 0xd8,
	0x1c, 0x42, 0x45, 0xd8, 0x04, 0xdb, 0xe4, 0x5e, 0x95, 0xb6, 0x00, 0x26, 0x64, 0x48, 0x54, 0x99,
	0xf1, 0x15, 0x40, 0x8a, 0x50, 0x81, 0x11, 0x12, 0x78, 0x1e, 0xdf, 0x22, 0x23, 0xb4, 0x2e, 0x9c,
	0x22, 0xa4, 0x34, 0x7f, 0x05, 0xd8, 0x7e, 0xfa, 0x11, 0x40, 0x49, 0x14, 0x50, 0xf7, 0x10, 0x82,
	0x5a, 0x77, 0x38, 0x18, 0x8c, 0x4c, 0xa3, 0xdb, 0xb6, 0x8d, 0xa1, 0x69, 0xa9, 0x52, 0xf3, 0x0f,
	0x38, 0x48, 0x7e, 0x0c, 0xd1, 0x77, 0x50, 0x34, 0x7a, 0x7d, 0x46, 0xbd, 0x0c, 0x0a, 0xdf, 0x85,
	0x2a, 0x31, 0xb0, 0x33, 0xb2, 0xee, 0x55, 0x99, 0x9d, 0x7a, 0x6c, 0xb2, 0x42, 0xb3, 0x05, 0xe5,
	0xcd, 0x2b, 0x61, 0x70, 0x3c, 0x70, 0x05, 0xf6, 0xaf, 0xb1, 0x89, 0x49, 0xbb, 0xaf, 0x4a, 0x6c,
	0x1b, 0x23, 0x13, 0xff, 0x73, 0x8b, 0xbb, 0x36, 0xee, 0xa9, 0xf2, 0xa4, 0xc4, 0xff, 0x6b, 0x7f,
	0xff, 0x1a, 0x00, 0x00, 0xff, 0xff, 0xac, 0xc8, 0x78, 0xd5, 0x7e, 0x07, 0x00, 0x00,
}
