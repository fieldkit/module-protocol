// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fk-module.proto

/*
Package fk_module is a generated protocol buffer package.

It is generated from these files:
	fk-module.proto

It has these top-level messages:
	QueryCapabilities
	QuerySensorCapabilities
	Firmware
	Capabilities
	SensorCapabilities
	BeginTakeReadings
	QueryReadingStatus
	ReadingStatus
	SensorReadings
	SensorReading
	Custom
	Data
	DataStatus
	BeginTransmission
	QueryTransmissionStatus
	TransmissionStatus
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
	QueryType_QUERY_CUSTOM              QueryType = 5
	QueryType_QUERY_DATA_APPEND         QueryType = 6
	QueryType_QUERY_DATA_CLEAR          QueryType = 7
	QueryType_QUERY_BEGIN_TRANSMISSION  QueryType = 8
	QueryType_QUERY_TRANSMISSION_STATUS QueryType = 9
	QueryType_QUERY_DATA_PREPARE        QueryType = 10
)

var QueryType_name = map[int32]string{
	0:  "QUERY_NONE",
	1:  "QUERY_CAPABILITIES",
	2:  "QUERY_SENSOR_CAPABILITIES",
	3:  "QUERY_BEGIN_TAKE_READINGS",
	4:  "QUERY_READING_STATUS",
	5:  "QUERY_CUSTOM",
	6:  "QUERY_DATA_APPEND",
	7:  "QUERY_DATA_CLEAR",
	8:  "QUERY_BEGIN_TRANSMISSION",
	9:  "QUERY_TRANSMISSION_STATUS",
	10: "QUERY_DATA_PREPARE",
}
var QueryType_value = map[string]int32{
	"QUERY_NONE":                0,
	"QUERY_CAPABILITIES":        1,
	"QUERY_SENSOR_CAPABILITIES": 2,
	"QUERY_BEGIN_TAKE_READINGS": 3,
	"QUERY_READING_STATUS":      4,
	"QUERY_CUSTOM":              5,
	"QUERY_DATA_APPEND":         6,
	"QUERY_DATA_CLEAR":          7,
	"QUERY_BEGIN_TRANSMISSION":  8,
	"QUERY_TRANSMISSION_STATUS": 9,
	"QUERY_DATA_PREPARE":        10,
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
	ReplyType_REPLY_CUSTOM              ReplyType = 6
	ReplyType_REPLY_DATA                ReplyType = 7
	ReplyType_REPLY_TRANSMISSION_STATUS ReplyType = 8
)

var ReplyType_name = map[int32]string{
	0: "REPLY_NONE",
	1: "REPLY_RETRY",
	2: "REPLY_ERROR",
	3: "REPLY_CAPABILITIES",
	4: "REPLY_SENSOR_CAPABILITIES",
	5: "REPLY_READING_STATUS",
	6: "REPLY_CUSTOM",
	7: "REPLY_DATA",
	8: "REPLY_TRANSMISSION_STATUS",
}
var ReplyType_value = map[string]int32{
	"REPLY_NONE":                0,
	"REPLY_RETRY":               1,
	"REPLY_ERROR":               2,
	"REPLY_CAPABILITIES":        3,
	"REPLY_SENSOR_CAPABILITIES": 4,
	"REPLY_READING_STATUS":      5,
	"REPLY_CUSTOM":              6,
	"REPLY_DATA":                7,
	"REPLY_TRANSMISSION_STATUS": 8,
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

type TransmissionState int32

const (
	TransmissionState_TRANSMISSION_IDLE  TransmissionState = 0
	TransmissionState_TRANSMISSION_BEGIN TransmissionState = 1
	TransmissionState_TRANSMISSION_BUSY  TransmissionState = 2
	TransmissionState_TRANSMISSION_DONE  TransmissionState = 3
)

var TransmissionState_name = map[int32]string{
	0: "TRANSMISSION_IDLE",
	1: "TRANSMISSION_BEGIN",
	2: "TRANSMISSION_BUSY",
	3: "TRANSMISSION_DONE",
}
var TransmissionState_value = map[string]int32{
	"TRANSMISSION_IDLE":  0,
	"TRANSMISSION_BEGIN": 1,
	"TRANSMISSION_BUSY":  2,
	"TRANSMISSION_DONE":  3,
}

func (x TransmissionState) String() string {
	return proto.EnumName(TransmissionState_name, int32(x))
}
func (TransmissionState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
func (ErrorType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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

type Firmware struct {
	Git   string `protobuf:"bytes,1,opt,name=git" json:"git,omitempty"`
	Build string `protobuf:"bytes,2,opt,name=build" json:"build,omitempty"`
}

func (m *Firmware) Reset()                    { *m = Firmware{} }
func (m *Firmware) String() string            { return proto.CompactTextString(m) }
func (*Firmware) ProtoMessage()               {}
func (*Firmware) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Firmware) GetGit() string {
	if m != nil {
		return m.Git
	}
	return ""
}

func (m *Firmware) GetBuild() string {
	if m != nil {
		return m.Build
	}
	return ""
}

type Capabilities struct {
	Version                 int32      `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Type                    ModuleType `protobuf:"varint,2,opt,name=type,enum=fk_module.ModuleType" json:"type,omitempty"`
	Name                    string     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Module                  string     `protobuf:"bytes,7,opt,name=module" json:"module,omitempty"`
	NumberOfSensors         int32      `protobuf:"varint,4,opt,name=numberOfSensors" json:"numberOfSensors,omitempty"`
	Firmware                *Firmware  `protobuf:"bytes,5,opt,name=firmware" json:"firmware,omitempty"`
	MinimumNumberOfReadings int32      `protobuf:"varint,6,opt,name=minimumNumberOfReadings" json:"minimumNumberOfReadings,omitempty"`
}

func (m *Capabilities) Reset()                    { *m = Capabilities{} }
func (m *Capabilities) String() string            { return proto.CompactTextString(m) }
func (*Capabilities) ProtoMessage()               {}
func (*Capabilities) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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

func (m *Capabilities) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *Capabilities) GetNumberOfSensors() int32 {
	if m != nil {
		return m.NumberOfSensors
	}
	return 0
}

func (m *Capabilities) GetFirmware() *Firmware {
	if m != nil {
		return m.Firmware
	}
	return nil
}

func (m *Capabilities) GetMinimumNumberOfReadings() int32 {
	if m != nil {
		return m.MinimumNumberOfReadings
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
func (*SensorCapabilities) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
	Index      int32  `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	CallerTime uint32 `protobuf:"varint,2,opt,name=callerTime" json:"callerTime,omitempty"`
	Number     uint32 `protobuf:"varint,3,opt,name=number" json:"number,omitempty"`
}

func (m *BeginTakeReadings) Reset()                    { *m = BeginTakeReadings{} }
func (m *BeginTakeReadings) String() string            { return proto.CompactTextString(m) }
func (*BeginTakeReadings) ProtoMessage()               {}
func (*BeginTakeReadings) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *BeginTakeReadings) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *BeginTakeReadings) GetCallerTime() uint32 {
	if m != nil {
		return m.CallerTime
	}
	return 0
}

func (m *BeginTakeReadings) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type QueryReadingStatus struct {
	Sleep uint32 `protobuf:"varint,1,opt,name=sleep" json:"sleep,omitempty"`
}

func (m *QueryReadingStatus) Reset()                    { *m = QueryReadingStatus{} }
func (m *QueryReadingStatus) String() string            { return proto.CompactTextString(m) }
func (*QueryReadingStatus) ProtoMessage()               {}
func (*QueryReadingStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *QueryReadingStatus) GetSleep() uint32 {
	if m != nil {
		return m.Sleep
	}
	return 0
}

type ReadingStatus struct {
	State   ReadingState `protobuf:"varint,1,opt,name=state,enum=fk_module.ReadingState" json:"state,omitempty"`
	Backoff uint32       `protobuf:"varint,2,opt,name=backoff" json:"backoff,omitempty"`
	Elapsed uint32       `protobuf:"varint,3,opt,name=elapsed" json:"elapsed,omitempty"`
}

func (m *ReadingStatus) Reset()                    { *m = ReadingStatus{} }
func (m *ReadingStatus) String() string            { return proto.CompactTextString(m) }
func (*ReadingStatus) ProtoMessage()               {}
func (*ReadingStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReadingStatus) GetState() ReadingState {
	if m != nil {
		return m.State
	}
	return ReadingState_IDLE
}

func (m *ReadingStatus) GetBackoff() uint32 {
	if m != nil {
		return m.Backoff
	}
	return 0
}

func (m *ReadingStatus) GetElapsed() uint32 {
	if m != nil {
		return m.Elapsed
	}
	return 0
}

type SensorReadings struct {
}

func (m *SensorReadings) Reset()                    { *m = SensorReadings{} }
func (m *SensorReadings) String() string            { return proto.CompactTextString(m) }
func (*SensorReadings) ProtoMessage()               {}
func (*SensorReadings) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type SensorReading struct {
	Sensor int32   `protobuf:"varint,1,opt,name=sensor" json:"sensor,omitempty"`
	Time   uint32  `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	Value  float32 `protobuf:"fixed32,3,opt,name=value" json:"value,omitempty"`
}

func (m *SensorReading) Reset()                    { *m = SensorReading{} }
func (m *SensorReading) String() string            { return proto.CompactTextString(m) }
func (*SensorReading) ProtoMessage()               {}
func (*SensorReading) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

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

type Custom struct {
	Message []byte `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *Custom) Reset()                    { *m = Custom{} }
func (m *Custom) String() string            { return proto.CompactTextString(m) }
func (*Custom) ProtoMessage()               {}
func (*Custom) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Custom) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type Data struct {
	Size int32  `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Kind int32  `protobuf:"varint,3,opt,name=kind" json:"kind,omitempty"`
}

func (m *Data) Reset()                    { *m = Data{} }
func (m *Data) String() string            { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()               {}
func (*Data) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Data) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Data) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Data) GetKind() int32 {
	if m != nil {
		return m.Kind
	}
	return 0
}

type DataStatus struct {
	Size int32 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
}

func (m *DataStatus) Reset()                    { *m = DataStatus{} }
func (m *DataStatus) String() string            { return proto.CompactTextString(m) }
func (*DataStatus) ProtoMessage()               {}
func (*DataStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *DataStatus) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type BeginTransmission struct {
	Kind int32 `protobuf:"varint,1,opt,name=kind" json:"kind,omitempty"`
}

func (m *BeginTransmission) Reset()                    { *m = BeginTransmission{} }
func (m *BeginTransmission) String() string            { return proto.CompactTextString(m) }
func (*BeginTransmission) ProtoMessage()               {}
func (*BeginTransmission) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *BeginTransmission) GetKind() int32 {
	if m != nil {
		return m.Kind
	}
	return 0
}

type QueryTransmissionStatus struct {
}

func (m *QueryTransmissionStatus) Reset()                    { *m = QueryTransmissionStatus{} }
func (m *QueryTransmissionStatus) String() string            { return proto.CompactTextString(m) }
func (*QueryTransmissionStatus) ProtoMessage()               {}
func (*QueryTransmissionStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type TransmissionStatus struct {
	State TransmissionState `protobuf:"varint,1,opt,name=state,enum=fk_module.TransmissionState" json:"state,omitempty"`
}

func (m *TransmissionStatus) Reset()                    { *m = TransmissionStatus{} }
func (m *TransmissionStatus) String() string            { return proto.CompactTextString(m) }
func (*TransmissionStatus) ProtoMessage()               {}
func (*TransmissionStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *TransmissionStatus) GetState() TransmissionState {
	if m != nil {
		return m.State
	}
	return TransmissionState_TRANSMISSION_IDLE
}

type WireMessageQuery struct {
	Type                    QueryType                `protobuf:"varint,1,opt,name=type,enum=fk_module.QueryType" json:"type,omitempty"`
	QueryCapabilities       *QueryCapabilities       `protobuf:"bytes,2,opt,name=queryCapabilities" json:"queryCapabilities,omitempty"`
	QuerySensorCapabilities *QuerySensorCapabilities `protobuf:"bytes,3,opt,name=querySensorCapabilities" json:"querySensorCapabilities,omitempty"`
	BeginTakeReadings       *BeginTakeReadings       `protobuf:"bytes,4,opt,name=beginTakeReadings" json:"beginTakeReadings,omitempty"`
	QueryReadingStatus      *QueryReadingStatus      `protobuf:"bytes,5,opt,name=queryReadingStatus" json:"queryReadingStatus,omitempty"`
	Custom                  *Custom                  `protobuf:"bytes,6,opt,name=custom" json:"custom,omitempty"`
	Data                    *Data                    `protobuf:"bytes,7,opt,name=data" json:"data,omitempty"`
	BeginTransmission       *BeginTransmission       `protobuf:"bytes,8,opt,name=beginTransmission" json:"beginTransmission,omitempty"`
	QueryTransmissionStatus *QueryTransmissionStatus `protobuf:"bytes,9,opt,name=queryTransmissionStatus" json:"queryTransmissionStatus,omitempty"`
}

func (m *WireMessageQuery) Reset()                    { *m = WireMessageQuery{} }
func (m *WireMessageQuery) String() string            { return proto.CompactTextString(m) }
func (*WireMessageQuery) ProtoMessage()               {}
func (*WireMessageQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

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

func (m *WireMessageQuery) GetCustom() *Custom {
	if m != nil {
		return m.Custom
	}
	return nil
}

func (m *WireMessageQuery) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *WireMessageQuery) GetBeginTransmission() *BeginTransmission {
	if m != nil {
		return m.BeginTransmission
	}
	return nil
}

func (m *WireMessageQuery) GetQueryTransmissionStatus() *QueryTransmissionStatus {
	if m != nil {
		return m.QueryTransmissionStatus
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
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

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
	Custom             *Custom             `protobuf:"bytes,7,opt,name=custom" json:"custom,omitempty"`
	Data               *DataStatus         `protobuf:"bytes,8,opt,name=data" json:"data,omitempty"`
	TransmissionStatus *TransmissionStatus `protobuf:"bytes,9,opt,name=transmissionStatus" json:"transmissionStatus,omitempty"`
}

func (m *WireMessageReply) Reset()                    { *m = WireMessageReply{} }
func (m *WireMessageReply) String() string            { return proto.CompactTextString(m) }
func (*WireMessageReply) ProtoMessage()               {}
func (*WireMessageReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

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

func (m *WireMessageReply) GetCustom() *Custom {
	if m != nil {
		return m.Custom
	}
	return nil
}

func (m *WireMessageReply) GetData() *DataStatus {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *WireMessageReply) GetTransmissionStatus() *TransmissionStatus {
	if m != nil {
		return m.TransmissionStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryCapabilities)(nil), "fk_module.QueryCapabilities")
	proto.RegisterType((*QuerySensorCapabilities)(nil), "fk_module.QuerySensorCapabilities")
	proto.RegisterType((*Firmware)(nil), "fk_module.Firmware")
	proto.RegisterType((*Capabilities)(nil), "fk_module.Capabilities")
	proto.RegisterType((*SensorCapabilities)(nil), "fk_module.SensorCapabilities")
	proto.RegisterType((*BeginTakeReadings)(nil), "fk_module.BeginTakeReadings")
	proto.RegisterType((*QueryReadingStatus)(nil), "fk_module.QueryReadingStatus")
	proto.RegisterType((*ReadingStatus)(nil), "fk_module.ReadingStatus")
	proto.RegisterType((*SensorReadings)(nil), "fk_module.SensorReadings")
	proto.RegisterType((*SensorReading)(nil), "fk_module.SensorReading")
	proto.RegisterType((*Custom)(nil), "fk_module.Custom")
	proto.RegisterType((*Data)(nil), "fk_module.Data")
	proto.RegisterType((*DataStatus)(nil), "fk_module.DataStatus")
	proto.RegisterType((*BeginTransmission)(nil), "fk_module.BeginTransmission")
	proto.RegisterType((*QueryTransmissionStatus)(nil), "fk_module.QueryTransmissionStatus")
	proto.RegisterType((*TransmissionStatus)(nil), "fk_module.TransmissionStatus")
	proto.RegisterType((*WireMessageQuery)(nil), "fk_module.WireMessageQuery")
	proto.RegisterType((*Error)(nil), "fk_module.Error")
	proto.RegisterType((*WireMessageReply)(nil), "fk_module.WireMessageReply")
	proto.RegisterEnum("fk_module.QueryType", QueryType_name, QueryType_value)
	proto.RegisterEnum("fk_module.ReplyType", ReplyType_name, ReplyType_value)
	proto.RegisterEnum("fk_module.ModuleType", ModuleType_name, ModuleType_value)
	proto.RegisterEnum("fk_module.ReadingState", ReadingState_name, ReadingState_value)
	proto.RegisterEnum("fk_module.TransmissionState", TransmissionState_name, TransmissionState_value)
	proto.RegisterEnum("fk_module.ErrorType", ErrorType_name, ErrorType_value)
}

func init() { proto.RegisterFile("fk-module.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x57, 0x5d, 0x6f, 0xe3, 0x44,
	0x17, 0xde, 0xa4, 0xf9, 0x3c, 0xf9, 0xe8, 0x64, 0xde, 0xee, 0xd6, 0x2b, 0x6d, 0x5f, 0x55, 0x06,
	0x41, 0x1b, 0xb1, 0x45, 0x84, 0x0b, 0x90, 0x90, 0x90, 0xdc, 0xc4, 0x94, 0xb0, 0x8d, 0x93, 0x8e,
	0x1d, 0x41, 0x25, 0x44, 0xe4, 0x34, 0x93, 0xca, 0x6a, 0xbe, 0x6a, 0x3b, 0x0b, 0xe5, 0x77, 0xc0,
	0x35, 0x3f, 0x89, 0x1f, 0xc4, 0x0d, 0x9a, 0x0f, 0x27, 0xe3, 0xd8, 0xe9, 0x72, 0xd5, 0x39, 0x67,
	0xce, 0x3c, 0xe7, 0x9c, 0x67, 0x9e, 0x39, 0x75, 0xe0, 0x70, 0xfa, 0xf0, 0x76, 0xbe, 0x9c, 0xac,
	0x67, 0xf4, 0x62, 0xe5, 0x2f, 0xc3, 0x25, 0x2e, 0x4f, 0x1f, 0x46, 0xc2, 0xa1, 0xf7, 0xa0, 0x71,
	0xb3, 0xa6, 0xfe, 0x53, 0xdb, 0x5d, 0xb9, 0x63, 0x6f, 0xe6, 0x85, 0x1e, 0x0d, 0xb0, 0x06, 0xc5,
	0xf7, 0xd4, 0x0f, 0xbc, 0xe5, 0x42, 0xcb, 0x9c, 0x66, 0xce, 0xf2, 0x24, 0x32, 0xf1, 0xff, 0x01,
	0xee, 0xdc, 0xd9, 0x8c, 0xfa, 0x8e, 0x37, 0xa7, 0x5a, 0xf6, 0x34, 0x73, 0x56, 0x23, 0x8a, 0x47,
	0xff, 0x02, 0x8e, 0x39, 0x9c, 0x4d, 0x17, 0xc1, 0xd2, 0x8f, 0x81, 0xbe, 0x82, 0x42, 0xc0, 0xbd,
	0x12, 0x53, 0x5a, 0x7a, 0x0b, 0x4a, 0xdf, 0x79, 0xfe, 0xfc, 0x57, 0xd7, 0xa7, 0x18, 0xc1, 0xc1,
	0xbd, 0x17, 0xf2, 0x80, 0x32, 0x61, 0x4b, 0x7c, 0x04, 0xf9, 0xf1, 0xda, 0x9b, 0x4d, 0x78, 0xae,
	0x32, 0x11, 0x86, 0xfe, 0x47, 0x16, 0xaa, 0xff, 0xb1, 0xe2, 0x73, 0xc8, 0x85, 0x4f, 0x2b, 0x51,
	0x6b, 0xbd, 0xf5, 0xf2, 0x62, 0xd3, 0xfa, 0x45, 0x8f, 0xff, 0x71, 0x9e, 0x56, 0x94, 0xf0, 0x10,
	0x8c, 0x21, 0xb7, 0x70, 0xe7, 0x54, 0x3b, 0xe0, 0xa9, 0xf8, 0x9a, 0x55, 0x2d, 0xc2, 0xb5, 0x22,
	0xf7, 0x4a, 0x0b, 0x9f, 0xc1, 0xe1, 0x62, 0x3d, 0x1f, 0x53, 0xbf, 0x3f, 0x15, 0xbd, 0x06, 0x5a,
	0x8e, 0x27, 0xde, 0x75, 0xe3, 0xcf, 0xa1, 0x34, 0x95, 0xfd, 0x69, 0xf9, 0xd3, 0xcc, 0x59, 0xa5,
	0xf5, 0x3f, 0xa5, 0x88, 0xa8, 0x75, 0xb2, 0x09, 0xc2, 0x5f, 0xc3, 0xf1, 0xdc, 0x5b, 0x78, 0xf3,
	0xf5, 0xdc, 0x92, 0x50, 0x84, 0xba, 0x13, 0x6f, 0x71, 0x1f, 0x68, 0x05, 0x9e, 0x62, 0xdf, 0xb6,
	0xfe, 0x0b, 0xe0, 0x14, 0xe2, 0xeb, 0x90, 0xf5, 0x26, 0x92, 0x96, 0xac, 0x37, 0xd9, 0xb4, 0x99,
	0x55, 0xda, 0xfc, 0x18, 0x6a, 0xeb, 0x85, 0x17, 0xf6, 0xa7, 0x3d, 0xea, 0x06, 0x6b, 0x3f, 0xe2,
	0x20, 0xee, 0xd4, 0x5d, 0x68, 0x5c, 0xd2, 0x7b, 0x6f, 0xe1, 0xb8, 0x0f, 0x34, 0x4a, 0xca, 0x6e,
	0xc8, 0x5b, 0x4c, 0xe8, 0x6f, 0x32, 0x83, 0x30, 0x3e, 0x24, 0x14, 0xc6, 0xab, 0x20, 0x8a, 0x67,
	0xaa, 0x11, 0x69, 0xe9, 0x4d, 0xc0, 0x5c, 0x40, 0x12, 0xde, 0x0e, 0xdd, 0x70, 0xcd, 0x73, 0x04,
	0x33, 0x4a, 0x57, 0x3c, 0x47, 0x8d, 0x08, 0x43, 0xf7, 0xa1, 0x16, 0x0f, 0x7b, 0x0b, 0xf9, 0x20,
	0x74, 0x43, 0xca, 0xc3, 0xea, 0xad, 0x63, 0x85, 0x67, 0x25, 0x90, 0x12, 0x11, 0xc5, 0x44, 0x33,
	0x76, 0xef, 0x1e, 0x96, 0xd3, 0xa9, 0x2c, 0x30, 0x32, 0xd9, 0x0e, 0x9d, 0xb9, 0xab, 0x80, 0x4e,
	0x64, 0x79, 0x91, 0xa9, 0x23, 0xa8, 0x0b, 0x8a, 0x37, 0xa4, 0xdf, 0x40, 0x2d, 0xe6, 0xd9, 0x27,
	0x74, 0xc6, 0x7b, 0xb8, 0x25, 0x83, 0xaf, 0x59, 0x63, 0xef, 0xdd, 0xd9, 0x5a, 0xf0, 0x9d, 0x25,
	0xc2, 0xd0, 0x75, 0x28, 0xb4, 0xd7, 0x41, 0xb8, 0x9c, 0xb3, 0x42, 0xe6, 0x34, 0x08, 0xdc, 0x7b,
	0xd1, 0x53, 0x95, 0x44, 0xa6, 0x7e, 0x09, 0xb9, 0x8e, 0x1b, 0xba, 0x0c, 0x35, 0xf0, 0x7e, 0xa7,
	0x32, 0x17, 0x5f, 0x33, 0xdf, 0xc4, 0x0d, 0x5d, 0x9e, 0xa9, 0x4a, 0xf8, 0x9a, 0xf9, 0x1e, 0xbc,
	0x85, 0xe8, 0x27, 0x4f, 0xf8, 0x5a, 0x3f, 0x05, 0x60, 0x18, 0x92, 0xbd, 0x14, 0x24, 0xfd, 0xd3,
	0xe8, 0xc6, 0x7d, 0x77, 0x11, 0xcc, 0xbd, 0x80, 0x3f, 0xa9, 0x08, 0x2a, 0xa3, 0x40, 0xbd, 0x96,
	0x0f, 0x5f, 0x0d, 0x14, 0xb8, 0xfa, 0xf7, 0x80, 0x93, 0x5e, 0xdc, 0x8a, 0xdf, 0xd5, 0x1b, 0xe5,
	0xae, 0x76, 0xa3, 0xa3, 0x0b, 0xd3, 0xff, 0xc9, 0x01, 0xfa, 0xd1, 0xf3, 0x69, 0x4f, 0x70, 0xc0,
	0x13, 0xe2, 0x33, 0xf9, 0xc0, 0x05, 0xce, 0x91, 0x82, 0x23, 0x0a, 0xda, 0xbe, 0xef, 0x1f, 0xa0,
	0xf1, 0xb8, 0x3b, 0xeb, 0x38, 0x47, 0x95, 0x58, 0xfa, 0xc4, 0x3c, 0x24, 0xc9, 0x63, 0xf8, 0x67,
	0x38, 0x7e, 0x4c, 0x1f, 0x74, 0x9c, 0xe1, 0x4a, 0x4b, 0xdf, 0x45, 0x4c, 0x46, 0x92, 0x7d, 0x10,
	0xac, 0xd2, 0xf1, 0xee, 0x43, 0xe3, 0xf3, 0x25, 0x5e, 0x69, 0xe2, 0x31, 0x92, 0xe4, 0x31, 0xdc,
	0x03, 0xfc, 0x98, 0x78, 0x51, 0x72, 0x12, 0x9d, 0xec, 0x16, 0x19, 0x0b, 0x22, 0x29, 0x07, 0xf1,
	0x39, 0x14, 0xee, 0xb8, 0x36, 0xf9, 0x30, 0xaa, 0xb4, 0x1a, 0x0a, 0x84, 0x10, 0x2d, 0x91, 0x01,
	0xf8, 0x23, 0x29, 0xc3, 0x22, 0x0f, 0x3c, 0x54, 0x02, 0x99, 0xea, 0xa4, 0x2e, 0x37, 0xad, 0x2a,
	0x97, 0xae, 0x95, 0xf6, 0xb4, 0xaa, 0xc4, 0x90, 0xe4, 0xb1, 0xcd, 0xa5, 0x24, 0xe5, 0xa6, 0x95,
	0xd3, 0x2f, 0x25, 0x19, 0x49, 0xf6, 0x41, 0xe8, 0xef, 0x20, 0x6f, 0xfa, 0xfe, 0xd2, 0x7f, 0x46,
	0x71, 0x7c, 0x5f, 0x51, 0x9c, 0xf2, 0x7c, 0xc5, 0xb4, 0xdd, 0x3c, 0xdf, 0x3f, 0xe3, 0x52, 0x26,
	0x74, 0x35, 0x7b, 0x4e, 0xca, 0x7c, 0x5f, 0x01, 0xfe, 0x04, 0xf2, 0x94, 0xe5, 0x92, 0xf2, 0x45,
	0xbb, 0x35, 0x10, 0xb1, 0x8d, 0xbf, 0x81, 0xea, 0x5d, 0x52, 0x9b, 0xea, 0x60, 0x8c, 0x09, 0x32,
	0x16, 0xcc, 0x94, 0x13, 0x24, 0xe5, 0x9d, 0x4b, 0x28, 0x27, 0x45, 0xd9, 0x29, 0x07, 0xf1, 0xb7,
	0x50, 0xf3, 0x53, 0x34, 0xa8, 0xa5, 0x4f, 0xe9, 0x75, 0x40, 0xe2, 0xe1, 0xec, 0x7c, 0xa0, 0x0e,
	0x5a, 0x29, 0x40, 0x2d, 0x51, 0x89, 0xdc, 0x27, 0xf1, 0x70, 0x45, 0xb9, 0xc5, 0x0f, 0x29, 0xf7,
	0x5c, 0x2a, 0x57, 0xe8, 0xf0, 0xe5, 0x8e, 0x72, 0x65, 0x79, 0x42, 0xbf, 0x3d, 0xc0, 0xe1, 0x3e,
	0xb9, 0x9d, 0x3c, 0x33, 0xd4, 0xd8, 0xf3, 0x4a, 0x1e, 0x6c, 0xfe, 0x95, 0x85, 0xf2, 0x66, 0x6e,
	0xe1, 0x3a, 0xc0, 0xcd, 0xd0, 0x24, 0xb7, 0x23, 0xab, 0x6f, 0x99, 0xe8, 0x05, 0x7e, 0x05, 0x58,
	0xd8, 0x6d, 0x63, 0x60, 0x5c, 0x76, 0xaf, 0xbb, 0x4e, 0xd7, 0xb4, 0x51, 0x06, 0x9f, 0xc0, 0x6b,
	0xe1, 0xb7, 0x4d, 0xcb, 0xee, 0x93, 0xf8, 0x76, 0x76, 0xbb, 0x7d, 0x69, 0x5e, 0x75, 0xad, 0x91,
	0x63, 0xbc, 0x33, 0x47, 0xc4, 0x34, 0x3a, 0x5d, 0xeb, 0xca, 0x46, 0x07, 0x58, 0x83, 0x23, 0xb1,
	0x2d, 0x7d, 0x23, 0xdb, 0x31, 0x9c, 0xa1, 0x8d, 0x72, 0x18, 0x41, 0x55, 0xe6, 0x1b, 0xda, 0x4e,
	0xbf, 0x87, 0xf2, 0xf8, 0x25, 0x34, 0x84, 0xa7, 0x63, 0x38, 0xc6, 0xc8, 0x18, 0x0c, 0x4c, 0xab,
	0x83, 0x0a, 0xf8, 0x08, 0x90, 0xe2, 0x6e, 0x5f, 0x9b, 0x06, 0x41, 0x45, 0xfc, 0x06, 0xb4, 0x58,
	0x5e, 0x62, 0x58, 0x76, 0xaf, 0x6b, 0xdb, 0xdd, 0xbe, 0x85, 0x4a, 0xdb, 0xaa, 0x54, 0x7f, 0x94,
	0xbb, 0xbc, 0xed, 0x95, 0x43, 0x0e, 0x88, 0x39, 0x30, 0x88, 0x89, 0xa0, 0xf9, 0x77, 0x06, 0xca,
	0x9b, 0xe7, 0xc0, 0x18, 0x22, 0xe6, 0xe0, 0x7a, 0xc3, 0xd0, 0x21, 0x54, 0x84, 0x4d, 0x4c, 0x87,
	0xdc, 0xa2, 0xcc, 0xd6, 0x61, 0x12, 0xd2, 0x27, 0x28, 0xcb, 0x70, 0x85, 0x23, 0x46, 0xd2, 0x01,
	0x2b, 0x47, 0xf8, 0xd3, 0x38, 0xcc, 0x31, 0x92, 0x22, 0xe0, 0x18, 0x49, 0x79, 0x46, 0x92, 0x04,
	0x14, 0x24, 0x15, 0xb6, 0x45, 0xb1, 0xd2, 0x51, 0x71, 0x0b, 0x9d, 0xd6, 0x69, 0xa9, 0xf9, 0x19,
	0xc0, 0xf6, 0x5b, 0x14, 0x03, 0x14, 0x44, 0x05, 0xe8, 0x05, 0xc6, 0x50, 0x6f, 0xf7, 0x7b, 0xbd,
	0xa1, 0xd5, 0x6d, 0x1b, 0x4e, 0xb7, 0x6f, 0xd9, 0x28, 0xd3, 0xfc, 0x0a, 0xaa, 0xea, 0xc7, 0x0c,
	0x2e, 0x41, 0xae, 0xdb, 0xb9, 0x66, 0xbd, 0x97, 0x21, 0xcf, 0x89, 0x46, 0x19, 0xe6, 0xbc, 0x1c,
	0xda, 0xb7, 0x28, 0xcb, 0x56, 0x1d, 0x46, 0xcd, 0x41, 0x73, 0x05, 0x8d, 0xc4, 0x7f, 0x56, 0x76,
	0x9f, 0xb1, 0xa2, 0x24, 0xd4, 0x2b, 0xc0, 0x31, 0x77, 0x84, 0xbb, 0x1b, 0x2e, 0x93, 0xec, 0xba,
	0x65, 0xc6, 0x16, 0x94, 0x37, 0x13, 0x91, 0x15, 0x22, 0xef, 0xa8, 0x02, 0xc5, 0x2b, 0xd3, 0x32,
	0x89, 0x71, 0x8d, 0x32, 0x8c, 0xab, 0xa1, 0x65, 0xfe, 0x34, 0x30, 0xdb, 0x8e, 0xd9, 0x41, 0xd9,
	0x71, 0x81, 0xff, 0x44, 0xf9, 0xf2, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x41, 0xee, 0x68,
	0xb5, 0x0c, 0x00, 0x00,
}
