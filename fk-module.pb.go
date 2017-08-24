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
func (ModuleType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ReadingState int32

const (
	ReadingState_IDLE ReadingState = 0
	ReadingState_BUSY ReadingState = 1
	ReadingState_DONE ReadingState = 2
)

var ReadingState_name = map[int32]string{
	0: "IDLE",
	1: "BUSY",
	2: "DONE",
}
var ReadingState_value = map[string]int32{
	"IDLE": 0,
	"BUSY": 1,
	"DONE": 2,
}

func (x ReadingState) String() string {
	return proto.EnumName(ReadingState_name, int32(x))
}
func (ReadingState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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
	QueryCapabilities  *QueryCapabilities  `protobuf:"bytes,1,opt,name=queryCapabilities" json:"queryCapabilities,omitempty"`
	BeginTakeReadings  *BeginTakeReadings  `protobuf:"bytes,2,opt,name=beginTakeReadings" json:"beginTakeReadings,omitempty"`
	QueryReadingStatus *QueryReadingStatus `protobuf:"bytes,3,opt,name=queryReadingStatus" json:"queryReadingStatus,omitempty"`
}

func (m *WireMessageQuery) Reset()                    { *m = WireMessageQuery{} }
func (m *WireMessageQuery) String() string            { return proto.CompactTextString(m) }
func (*WireMessageQuery) ProtoMessage()               {}
func (*WireMessageQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

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

type WireMessageReply struct {
	Capabilities       *Capabilities       `protobuf:"bytes,1,opt,name=capabilities" json:"capabilities,omitempty"`
	SensorCapabilities *SensorCapabilities `protobuf:"bytes,2,opt,name=sensorCapabilities" json:"sensorCapabilities,omitempty"`
	ReadingStatus      *ReadingStatus      `protobuf:"bytes,3,opt,name=readingStatus" json:"readingStatus,omitempty"`
	SensorReadings     *SensorReadings     `protobuf:"bytes,4,opt,name=sensorReadings" json:"sensorReadings,omitempty"`
}

func (m *WireMessageReply) Reset()                    { *m = WireMessageReply{} }
func (m *WireMessageReply) String() string            { return proto.CompactTextString(m) }
func (*WireMessageReply) ProtoMessage()               {}
func (*WireMessageReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

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
	proto.RegisterType((*WireMessageReply)(nil), "fk_module.WireMessageReply")
	proto.RegisterEnum("fk_module.ModuleType", ModuleType_name, ModuleType_value)
	proto.RegisterEnum("fk_module.ReadingState", ReadingState_name, ReadingState_value)
}

func init() { proto.RegisterFile("fk-module.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xc1, 0x6e, 0xda, 0x40,
	0x10, 0x8d, 0x0d, 0xa4, 0xc9, 0x04, 0xa8, 0x19, 0x51, 0xd5, 0x95, 0x1a, 0x09, 0xf9, 0x44, 0x50,
	0xe0, 0x40, 0x7b, 0xa8, 0x14, 0x35, 0x52, 0x42, 0x38, 0x50, 0xd5, 0xa0, 0xae, 0x89, 0xaa, 0xf6,
	0x52, 0x2d, 0x65, 0x40, 0x16, 0x60, 0x88, 0xd7, 0x8e, 0x4a, 0xff, 0xa6, 0x9f, 0xd9, 0x5b, 0xc5,
	0xda, 0xa1, 0x5e, 0x16, 0x72, 0x62, 0x76, 0xfd, 0x66, 0xe6, 0xcd, 0x9b, 0xc7, 0xc2, 0xcb, 0xc9,
	0xac, 0xb9, 0x58, 0x8e, 0xe3, 0x39, 0xb5, 0x56, 0xe1, 0x32, 0x5a, 0xe2, 0xe9, 0x64, 0xf6, 0x23,
	0xb9, 0x70, 0x9a, 0x50, 0xf9, 0x12, 0x53, 0xb8, 0xee, 0xf0, 0x15, 0x1f, 0xf9, 0x73, 0x3f, 0xf2,
	0x49, 0xa0, 0x0d, 0x2f, 0x1e, 0x29, 0x14, 0xfe, 0x32, 0xb0, 0x8d, 0x9a, 0x51, 0xcf, 0xb1, 0xa7,
	0xa3, 0x33, 0x85, 0xe2, 0x73, 0xc8, 0xc2, 0x16, 0x89, 0x17, 0x90, 0x8f, 0xd6, 0x2b, 0xb2, 0xcd,
	0x9a, 0x51, 0x2f, 0xb7, 0x5f, 0xb5, 0xb6, 0x2d, 0x5b, 0xae, 0xfc, 0x19, 0xae, 0x57, 0xc4, 0x24,
	0x04, 0x11, 0xf2, 0x01, 0x5f, 0x90, 0x9d, 0xab, 0x19, 0xf5, 0x53, 0x26, 0x63, 0xe7, 0x37, 0xa0,
	0x47, 0x81, 0x58, 0x86, 0x4a, 0xbb, 0x06, 0x58, 0x41, 0xbc, 0x18, 0x51, 0x38, 0x98, 0x30, 0xe2,
	0x63, 0x3f, 0x98, 0x8a, 0xb4, 0xaf, 0x76, 0x8f, 0x1f, 0xe0, 0x24, 0x7c, 0xc2, 0x98, 0xb5, 0x5c,
	0xfd, 0xac, 0xfd, 0x36, 0x43, 0x22, 0x29, 0x9e, 0x82, 0x5d, 0x8a, 0x38, 0xdb, 0xa2, 0x9d, 0x8f,
	0x50, 0xd1, 0x3e, 0x63, 0x15, 0x0a, 0x7e, 0x30, 0xa6, 0x5f, 0x69, 0xbf, 0xe4, 0xb0, 0xa5, 0x6e,
	0x66, 0xa8, 0x5f, 0x40, 0xe5, 0x96, 0xa6, 0x7e, 0x30, 0xe4, 0x33, 0xda, 0xb2, 0xd9, 0x9b, 0xee,
	0x54, 0x01, 0xa5, 0xfa, 0x29, 0xcc, 0x8b, 0x78, 0x14, 0x0b, 0xe7, 0x1a, 0x4a, 0xca, 0x05, 0x36,
	0xa1, 0x20, 0x22, 0x1e, 0x91, 0x4c, 0x2e, 0xb7, 0x5f, 0x67, 0xe6, 0xc8, 0x00, 0x89, 0x25, 0x28,
	0xe7, 0x3b, 0x94, 0x15, 0xfe, 0x62, 0x43, 0x33, 0xf2, 0x17, 0x94, 0x6e, 0x53, 0xc6, 0xf8, 0x5e,
	0xd3, 0xc7, 0x3e, 0xa4, 0x4f, 0x46, 0x9b, 0x2b, 0x28, 0x29, 0x9f, 0x0e, 0xe8, 0x52, 0x85, 0xc2,
	0x23, 0x9f, 0xc7, 0x89, 0x30, 0x26, 0x4b, 0x0e, 0xce, 0x5f, 0x03, 0xac, 0xaf, 0x7e, 0x48, 0x2e,
	0x09, 0xc1, 0xa7, 0x24, 0x47, 0xc7, 0x4f, 0x50, 0x79, 0xd8, 0x75, 0xa0, 0x2c, 0xa6, 0x2e, 0x4c,
	0x73, 0x29, 0xd3, 0xd3, 0x36, 0xb5, 0x46, 0xbb, 0xd2, 0x4b, 0x0a, 0x6a, 0x2d, 0x6d, 0x3d, 0x4c,
	0x4f, 0x43, 0x17, 0xf0, 0x41, 0xdb, 0x8d, 0xf4, 0xe8, 0x59, 0xfb, 0x7c, 0x97, 0x98, 0x02, 0x62,
	0x7b, 0x12, 0x9d, 0x3f, 0xa6, 0x32, 0x3b, 0xa3, 0xd5, 0x7c, 0x8d, 0x57, 0x50, 0xfc, 0xa9, 0x8f,
	0x9d, 0xdd, 0xaf, 0x32, 0xb1, 0x02, 0xde, 0x10, 0x14, 0xda, 0x5f, 0x24, 0x9d, 0xf6, 0x5c, 0x5b,
	0xa5, 0x52, 0x68, 0x4f, 0x22, 0x5e, 0x43, 0x29, 0xdc, 0x33, 0xaa, 0xbd, 0xdf, 0x6c, 0xb1, 0x60,
	0x2a, 0x1c, 0x6f, 0xa0, 0x2c, 0x14, 0xd7, 0xd9, 0x79, 0x59, 0xe0, 0xcd, 0x21, 0x57, 0x09, 0xb6,
	0x93, 0xd0, 0xb8, 0x04, 0xf8, 0xff, 0x38, 0x20, 0xc0, 0xb1, 0xd7, 0xed, 0x7b, 0x03, 0x66, 0x1d,
	0x21, 0x42, 0xb9, 0x33, 0x70, 0xdd, 0xfb, 0x7e, 0xaf, 0x73, 0x33, 0xec, 0x0d, 0xfa, 0x9e, 0x65,
	0x34, 0x2e, 0xa1, 0x98, 0x75, 0x3f, 0x9e, 0x40, 0xbe, 0x77, 0xf7, 0xb9, 0x6b, 0x1d, 0x6d, 0xa2,
	0xdb, 0x7b, 0xef, 0x9b, 0x65, 0x6c, 0xa2, 0xbb, 0x41, 0xbf, 0x6b, 0x99, 0xa3, 0x63, 0xf9, 0xf4,
	0xbd, 0xfb, 0x17, 0x00, 0x00, 0xff, 0xff, 0x37, 0x07, 0xed, 0x94, 0x0d, 0x05, 0x00, 0x00,
}
