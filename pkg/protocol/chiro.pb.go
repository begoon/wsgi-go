// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: pkg/protocol/chiro.proto

package protocol

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Handedness int32

const (
	Handedness_LEFT  Handedness = 0
	Handedness_RIGHT Handedness = 1
)

// Enum value maps for Handedness.
var (
	Handedness_name = map[int32]string{
		0: "LEFT",
		1: "RIGHT",
	}
	Handedness_value = map[string]int32{
		"LEFT":  0,
		"RIGHT": 1,
	}
)

func (x Handedness) Enum() *Handedness {
	p := new(Handedness)
	*p = x
	return p
}

func (x Handedness) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Handedness) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_protocol_chiro_proto_enumTypes[0].Descriptor()
}

func (Handedness) Type() protoreflect.EnumType {
	return &file_pkg_protocol_chiro_proto_enumTypes[0]
}

func (x Handedness) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Handedness.Descriptor instead.
func (Handedness) EnumDescriptor() ([]byte, []int) {
	return file_pkg_protocol_chiro_proto_rawDescGZIP(), []int{0}
}

type CaptureType int32

const (
	CaptureType_FINGER CaptureType = 0
	CaptureType_THUMB  CaptureType = 1
)

// Enum value maps for CaptureType.
var (
	CaptureType_name = map[int32]string{
		0: "FINGER",
		1: "THUMB",
	}
	CaptureType_value = map[string]int32{
		"FINGER": 0,
		"THUMB":  1,
	}
)

func (x CaptureType) Enum() *CaptureType {
	p := new(CaptureType)
	*p = x
	return p
}

func (x CaptureType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CaptureType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_protocol_chiro_proto_enumTypes[1].Descriptor()
}

func (CaptureType) Type() protoreflect.EnumType {
	return &file_pkg_protocol_chiro_proto_enumTypes[1]
}

func (x CaptureType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CaptureType.Descriptor instead.
func (CaptureType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_protocol_chiro_proto_rawDescGZIP(), []int{1}
}

type EventType int32

const (
	EventType_ACK            EventType = 0 // // TODO DO we need ack in this early stage?
	EventType_CLIENT_START   EventType = 1 // ACK, we need to log Android ID
	EventType_PAD_FRAMES     EventType = 2
	EventType_PALM_SLAB      EventType = 3
	EventType_CAPTURE_RESULT EventType = 4
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "ACK",
		1: "CLIENT_START",
		2: "PAD_FRAMES",
		3: "PALM_SLAB",
		4: "CAPTURE_RESULT",
	}
	EventType_value = map[string]int32{
		"ACK":            0,
		"CLIENT_START":   1,
		"PAD_FRAMES":     2,
		"PALM_SLAB":      3,
		"CAPTURE_RESULT": 4,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_protocol_chiro_proto_enumTypes[2].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_pkg_protocol_chiro_proto_enumTypes[2]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_protocol_chiro_proto_rawDescGZIP(), []int{2}
}

type AckMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SequenceNumber uint32 `protobuf:"varint,1,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
}

func (x *AckMessage) Reset() {
	*x = AckMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protocol_chiro_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckMessage) ProtoMessage() {}

func (x *AckMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protocol_chiro_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckMessage.ProtoReflect.Descriptor instead.
func (*AckMessage) Descriptor() ([]byte, []int) {
	return file_pkg_protocol_chiro_proto_rawDescGZIP(), []int{0}
}

func (x *AckMessage) GetSequenceNumber() uint32 {
	if x != nil {
		return x.SequenceNumber
	}
	return 0
}

type CaptureResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO Rename something line isSuccessful
	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"` // Returns true if crops were successfully captured and stored
	// Handle result for different capture type and handedness
	PadResult     bool    `protobuf:"varint,2,opt,name=padResult,proto3" json:"padResult,omitempty"` // Should this be tied into the result field or be seperate??
	FailureReason *string `protobuf:"bytes,3,opt,name=failureReason,proto3,oneof" json:"failureReason,omitempty"`
}

func (x *CaptureResult) Reset() {
	*x = CaptureResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protocol_chiro_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaptureResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaptureResult) ProtoMessage() {}

func (x *CaptureResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protocol_chiro_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaptureResult.ProtoReflect.Descriptor instead.
func (*CaptureResult) Descriptor() ([]byte, []int) {
	return file_pkg_protocol_chiro_proto_rawDescGZIP(), []int{1}
}

func (x *CaptureResult) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *CaptureResult) GetPadResult() bool {
	if x != nil {
		return x.PadResult
	}
	return false
}

func (x *CaptureResult) GetFailureReason() string {
	if x != nil && x.FailureReason != nil {
		return *x.FailureReason
	}
	return ""
}

type ClientStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Make     string `protobuf:"bytes,2,opt,name=make,proto3" json:"make,omitempty"`
	Model    string `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	DeviceId string `protobuf:"bytes,4,opt,name=deviceId,proto3" json:"deviceId,omitempty"` // TODO Implement string userId = 4;
}

func (x *ClientStart) Reset() {
	*x = ClientStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protocol_chiro_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStart) ProtoMessage() {}

func (x *ClientStart) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protocol_chiro_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStart.ProtoReflect.Descriptor instead.
func (*ClientStart) Descriptor() ([]byte, []int) {
	return file_pkg_protocol_chiro_proto_rawDescGZIP(), []int{2}
}

func (x *ClientStart) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *ClientStart) GetMake() string {
	if x != nil {
		return x.Make
	}
	return ""
}

func (x *ClientStart) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *ClientStart) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

// Frames used for palm detection
type PadFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image       []byte      `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"` // As PNG
	ImageSize   *Size       `protobuf:"bytes,2,opt,name=imageSize,proto3" json:"imageSize,omitempty"`
	Landmarks   []*Landmark `protobuf:"bytes,3,rep,name=landmarks,proto3" json:"landmarks,omitempty"` // Expect 21 landmarks
	CaptureType CaptureType `protobuf:"varint,4,opt,name=captureType,proto3,enum=iproov.chiro.CaptureType" json:"captureType,omitempty"`
	Handedness  Handedness  `protobuf:"varint,5,opt,name=handedness,proto3,enum=iproov.chiro.Handedness" json:"handedness,omitempty"`
}

func (x *PadFrame) Reset() {
	*x = PadFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protocol_chiro_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PadFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PadFrame) ProtoMessage() {}

func (x *PadFrame) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protocol_chiro_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PadFrame.ProtoReflect.Descriptor instead.
func (*PadFrame) Descriptor() ([]byte, []int) {
	return file_pkg_protocol_chiro_proto_rawDescGZIP(), []int{3}
}

func (x *PadFrame) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *PadFrame) GetImageSize() *Size {
	if x != nil {
		return x.ImageSize
	}
	return nil
}

func (x *PadFrame) GetLandmarks() []*Landmark {
	if x != nil {
		return x.Landmarks
	}
	return nil
}

func (x *PadFrame) GetCaptureType() CaptureType {
	if x != nil {
		return x.CaptureType
	}
	return CaptureType_FINGER
}

func (x *PadFrame) GetHandedness() Handedness {
	if x != nil {
		return x.Handedness
	}
	return Handedness_LEFT
}

// Single high quality frame used for finger print extraction
type PalmSlab struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image         []byte      `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`                          // As PNG for now
	ImageSize     *Size       `protobuf:"bytes,2,opt,name=imageSize,proto3" json:"imageSize,omitempty"`                  // Image dimension
	QualityScores []float32   `protobuf:"fixed32,3,rep,packed,name=qualityScores,proto3" json:"qualityScores,omitempty"` // Should expect 10 an array of size 10; Score of -1 means failed to calculate score; Needs backend size checks
	NfiqScores    []float32   `protobuf:"fixed32,4,rep,packed,name=nfiqScores,proto3" json:"nfiqScores,omitempty"`       // array of size 10; scores can also be -1; Need backend size checks
	Landmarks     []*Landmark `protobuf:"bytes,5,rep,name=landmarks,proto3" json:"landmarks,omitempty"`                  // Expect array of size 21; Need to do backend checks
	CropBounds    []*Rect     `protobuf:"bytes,6,rep,name=cropBounds,proto3" json:"cropBounds,omitempty"`                // Expect array of size 4 or 1; Origin is top left and relative to image size
	Rotations     []int32     `protobuf:"varint,7,rep,packed,name=rotations,proto3" json:"rotations,omitempty"`          // Expect array of 4 or 1 (Depends on capturetype);angle degrees, rotate about center; Maybe radians are better?
	CaptureType   CaptureType `protobuf:"varint,8,opt,name=captureType,proto3,enum=iproov.chiro.CaptureType" json:"captureType,omitempty"`
	Handedness    Handedness  `protobuf:"varint,9,opt,name=handedness,proto3,enum=iproov.chiro.Handedness" json:"handedness,omitempty"`
}

func (x *PalmSlab) Reset() {
	*x = PalmSlab{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protocol_chiro_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PalmSlab) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PalmSlab) ProtoMessage() {}

func (x *PalmSlab) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protocol_chiro_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PalmSlab.ProtoReflect.Descriptor instead.
func (*PalmSlab) Descriptor() ([]byte, []int) {
	return file_pkg_protocol_chiro_proto_rawDescGZIP(), []int{4}
}

func (x *PalmSlab) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *PalmSlab) GetImageSize() *Size {
	if x != nil {
		return x.ImageSize
	}
	return nil
}

func (x *PalmSlab) GetQualityScores() []float32 {
	if x != nil {
		return x.QualityScores
	}
	return nil
}

func (x *PalmSlab) GetNfiqScores() []float32 {
	if x != nil {
		return x.NfiqScores
	}
	return nil
}

func (x *PalmSlab) GetLandmarks() []*Landmark {
	if x != nil {
		return x.Landmarks
	}
	return nil
}

func (x *PalmSlab) GetCropBounds() []*Rect {
	if x != nil {
		return x.CropBounds
	}
	return nil
}

func (x *PalmSlab) GetRotations() []int32 {
	if x != nil {
		return x.Rotations
	}
	return nil
}

func (x *PalmSlab) GetCaptureType() CaptureType {
	if x != nil {
		return x.CaptureType
	}
	return CaptureType_FINGER
}

func (x *PalmSlab) GetHandedness() Handedness {
	if x != nil {
		return x.Handedness
	}
	return Handedness_LEFT
}

type Landmark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Landmark) Reset() {
	*x = Landmark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protocol_chiro_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Landmark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Landmark) ProtoMessage() {}

func (x *Landmark) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protocol_chiro_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Landmark.ProtoReflect.Descriptor instead.
func (*Landmark) Descriptor() ([]byte, []int) {
	return file_pkg_protocol_chiro_proto_rawDescGZIP(), []int{5}
}

func (x *Landmark) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Landmark) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Rect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Left   int32 `protobuf:"varint,1,opt,name=left,proto3" json:"left,omitempty"`
	Top    int32 `protobuf:"varint,2,opt,name=top,proto3" json:"top,omitempty"`
	Right  int32 `protobuf:"varint,3,opt,name=right,proto3" json:"right,omitempty"`
	Bottom int32 `protobuf:"varint,4,opt,name=bottom,proto3" json:"bottom,omitempty"`
}

func (x *Rect) Reset() {
	*x = Rect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protocol_chiro_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rect) ProtoMessage() {}

func (x *Rect) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protocol_chiro_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rect.ProtoReflect.Descriptor instead.
func (*Rect) Descriptor() ([]byte, []int) {
	return file_pkg_protocol_chiro_proto_rawDescGZIP(), []int{6}
}

func (x *Rect) GetLeft() int32 {
	if x != nil {
		return x.Left
	}
	return 0
}

func (x *Rect) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *Rect) GetRight() int32 {
	if x != nil {
		return x.Right
	}
	return 0
}

func (x *Rect) GetBottom() int32 {
	if x != nil {
		return x.Bottom
	}
	return 0
}

type Size struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width  int32 `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height int32 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *Size) Reset() {
	*x = Size{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protocol_chiro_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Size) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Size) ProtoMessage() {}

func (x *Size) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protocol_chiro_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Size.ProtoReflect.Descriptor instead.
func (*Size) Descriptor() ([]byte, []int) {
	return file_pkg_protocol_chiro_proto_rawDescGZIP(), []int{7}
}

func (x *Size) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Size) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version        uint32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	SequenceNumber uint32    `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"` // TODO required?? https://developers.google.com/protocol-buffers/docs/proto#specifying-rules
	EventType      EventType `protobuf:"varint,3,opt,name=event_type,json=eventType,proto3,enum=iproov.chiro.EventType" json:"event_type,omitempty"`
	// Types that are assignable to Payload:
	//
	//	*Envelope_ClientStart
	//	*Envelope_PalmSlab
	//	*Envelope_PadFrame
	//	*Envelope_AckMessage
	//	*Envelope_CaptureResult
	Payload isEnvelope_Payload `protobuf_oneof:"payload"`
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protocol_chiro_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protocol_chiro_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_pkg_protocol_chiro_proto_rawDescGZIP(), []int{8}
}

func (x *Envelope) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Envelope) GetSequenceNumber() uint32 {
	if x != nil {
		return x.SequenceNumber
	}
	return 0
}

func (x *Envelope) GetEventType() EventType {
	if x != nil {
		return x.EventType
	}
	return EventType_ACK
}

func (m *Envelope) GetPayload() isEnvelope_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Envelope) GetClientStart() *ClientStart {
	if x, ok := x.GetPayload().(*Envelope_ClientStart); ok {
		return x.ClientStart
	}
	return nil
}

func (x *Envelope) GetPalmSlab() *PalmSlab {
	if x, ok := x.GetPayload().(*Envelope_PalmSlab); ok {
		return x.PalmSlab
	}
	return nil
}

func (x *Envelope) GetPadFrame() *PadFrame {
	if x, ok := x.GetPayload().(*Envelope_PadFrame); ok {
		return x.PadFrame
	}
	return nil
}

func (x *Envelope) GetAckMessage() *AckMessage {
	if x, ok := x.GetPayload().(*Envelope_AckMessage); ok {
		return x.AckMessage
	}
	return nil
}

func (x *Envelope) GetCaptureResult() *CaptureResult {
	if x, ok := x.GetPayload().(*Envelope_CaptureResult); ok {
		return x.CaptureResult
	}
	return nil
}

type isEnvelope_Payload interface {
	isEnvelope_Payload()
}

type Envelope_ClientStart struct {
	ClientStart *ClientStart `protobuf:"bytes,4,opt,name=clientStart,proto3,oneof"`
}

type Envelope_PalmSlab struct {
	PalmSlab *PalmSlab `protobuf:"bytes,5,opt,name=palmSlab,proto3,oneof"`
}

type Envelope_PadFrame struct {
	PadFrame *PadFrame `protobuf:"bytes,6,opt,name=padFrame,proto3,oneof"`
}

type Envelope_AckMessage struct {
	AckMessage *AckMessage `protobuf:"bytes,7,opt,name=ackMessage,proto3,oneof"`
}

type Envelope_CaptureResult struct {
	CaptureResult *CaptureResult `protobuf:"bytes,9,opt,name=captureResult,proto3,oneof"`
}

func (*Envelope_ClientStart) isEnvelope_Payload() {}

func (*Envelope_PalmSlab) isEnvelope_Payload() {}

func (*Envelope_PadFrame) isEnvelope_Payload() {}

func (*Envelope_AckMessage) isEnvelope_Payload() {}

func (*Envelope_CaptureResult) isEnvelope_Payload() {}

var File_pkg_protocol_chiro_proto protoreflect.FileDescriptor

var file_pkg_protocol_chiro_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63,
	0x68, 0x69, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x69, 0x70, 0x72, 0x6f,
	0x6f, 0x76, 0x2e, 0x63, 0x68, 0x69, 0x72, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x0a, 0x41, 0x63, 0x6b, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x73,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x82, 0x01,
	0x0a, 0x0d, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x70, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x29, 0x0a, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x22, 0x6b, 0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x6b,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x6b, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22,
	0xff, 0x01, 0x0a, 0x08, 0x50, 0x61, 0x64, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x70, 0x72, 0x6f, 0x6f, 0x76, 0x2e, 0x63,
	0x68, 0x69, 0x72, 0x6f, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x64, 0x6d, 0x61, 0x72, 0x6b,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x70, 0x72, 0x6f, 0x6f, 0x76,
	0x2e, 0x63, 0x68, 0x69, 0x72, 0x6f, 0x2e, 0x4c, 0x61, 0x6e, 0x64, 0x6d, 0x61, 0x72, 0x6b, 0x52,
	0x09, 0x6c, 0x61, 0x6e, 0x64, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x61,
	0x70, 0x74, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x69, 0x70, 0x72, 0x6f, 0x6f, 0x76, 0x2e, 0x63, 0x68, 0x69, 0x72, 0x6f, 0x2e, 0x43,
	0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x61, 0x70, 0x74,
	0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x68, 0x61, 0x6e, 0x64, 0x65,
	0x64, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x69, 0x70,
	0x72, 0x6f, 0x6f, 0x76, 0x2e, 0x63, 0x68, 0x69, 0x72, 0x6f, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x65,
	0x64, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x0a, 0x68, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x6e, 0x65, 0x73,
	0x73, 0x22, 0x97, 0x03, 0x0a, 0x08, 0x50, 0x61, 0x6c, 0x6d, 0x53, 0x6c, 0x61, 0x62, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x70, 0x72, 0x6f, 0x6f, 0x76,
	0x2e, 0x63, 0x68, 0x69, 0x72, 0x6f, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x09, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x02, 0x52, 0x0d, 0x71,
	0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x6e, 0x66, 0x69, 0x71, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x02,
	0x52, 0x0a, 0x6e, 0x66, 0x69, 0x71, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x09,
	0x6c, 0x61, 0x6e, 0x64, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x69, 0x70, 0x72, 0x6f, 0x6f, 0x76, 0x2e, 0x63, 0x68, 0x69, 0x72, 0x6f, 0x2e, 0x4c,
	0x61, 0x6e, 0x64, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x09, 0x6c, 0x61, 0x6e, 0x64, 0x6d, 0x61, 0x72,
	0x6b, 0x73, 0x12, 0x32, 0x0a, 0x0a, 0x63, 0x72, 0x6f, 0x70, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x70, 0x72, 0x6f, 0x6f, 0x76, 0x2e,
	0x63, 0x68, 0x69, 0x72, 0x6f, 0x2e, 0x52, 0x65, 0x63, 0x74, 0x52, 0x0a, 0x63, 0x72, 0x6f, 0x70,
	0x42, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x72, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x69, 0x70, 0x72, 0x6f,
	0x6f, 0x76, 0x2e, 0x63, 0x68, 0x69, 0x72, 0x6f, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x38, 0x0a, 0x0a, 0x68, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x6e, 0x65, 0x73, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x69, 0x70, 0x72, 0x6f, 0x6f, 0x76, 0x2e, 0x63,
	0x68, 0x69, 0x72, 0x6f, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x6e, 0x65, 0x73, 0x73, 0x52,
	0x0a, 0x68, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x6e, 0x65, 0x73, 0x73, 0x22, 0x26, 0x0a, 0x08, 0x4c,
	0x61, 0x6e, 0x64, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x01, 0x79, 0x22, 0x5a, 0x0a, 0x04, 0x52, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x65, 0x66, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x74, 0x74, 0x6f,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x22,
	0x34, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0xbc, 0x03, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x69, 0x70, 0x72, 0x6f,
	0x6f, 0x76, 0x2e, 0x63, 0x68, 0x69, 0x72, 0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a,
	0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x70, 0x72, 0x6f, 0x6f, 0x76, 0x2e, 0x63, 0x68, 0x69, 0x72,
	0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52,
	0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x34, 0x0a, 0x08,
	0x70, 0x61, 0x6c, 0x6d, 0x53, 0x6c, 0x61, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x69, 0x70, 0x72, 0x6f, 0x6f, 0x76, 0x2e, 0x63, 0x68, 0x69, 0x72, 0x6f, 0x2e, 0x50, 0x61,
	0x6c, 0x6d, 0x53, 0x6c, 0x61, 0x62, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x6c, 0x6d, 0x53, 0x6c,
	0x61, 0x62, 0x12, 0x34, 0x0a, 0x08, 0x70, 0x61, 0x64, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x70, 0x72, 0x6f, 0x6f, 0x76, 0x2e, 0x63, 0x68,
	0x69, 0x72, 0x6f, 0x2e, 0x50, 0x61, 0x64, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x08,
	0x70, 0x61, 0x64, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x61, 0x63, 0x6b, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69,
	0x70, 0x72, 0x6f, 0x6f, 0x76, 0x2e, 0x63, 0x68, 0x69, 0x72, 0x6f, 0x2e, 0x41, 0x63, 0x6b, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x63, 0x6b, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x70,
	0x72, 0x6f, 0x6f, 0x76, 0x2e, 0x63, 0x68, 0x69, 0x72, 0x6f, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x75,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x61, 0x70, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x2a, 0x21, 0x0a, 0x0a, 0x48, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x6e, 0x65,
	0x73, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x45, 0x46, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x01, 0x2a, 0x24, 0x0a, 0x0b, 0x43, 0x61, 0x70, 0x74, 0x75,
	0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x49, 0x4e, 0x47, 0x45, 0x52,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x48, 0x55, 0x4d, 0x42, 0x10, 0x01, 0x2a, 0x59, 0x0a,
	0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x43,
	0x4b, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54,
	0x41, 0x52, 0x54, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x41, 0x44, 0x5f, 0x46, 0x52, 0x41,
	0x4d, 0x45, 0x53, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x41, 0x4c, 0x4d, 0x5f, 0x53, 0x4c,
	0x41, 0x42, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x50, 0x54, 0x55, 0x52, 0x45, 0x5f,
	0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x10, 0x04, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_protocol_chiro_proto_rawDescOnce sync.Once
	file_pkg_protocol_chiro_proto_rawDescData = file_pkg_protocol_chiro_proto_rawDesc
)

func file_pkg_protocol_chiro_proto_rawDescGZIP() []byte {
	file_pkg_protocol_chiro_proto_rawDescOnce.Do(func() {
		file_pkg_protocol_chiro_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_protocol_chiro_proto_rawDescData)
	})
	return file_pkg_protocol_chiro_proto_rawDescData
}

var file_pkg_protocol_chiro_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_pkg_protocol_chiro_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_protocol_chiro_proto_goTypes = []interface{}{
	(Handedness)(0),       // 0: iproov.chiro.Handedness
	(CaptureType)(0),      // 1: iproov.chiro.CaptureType
	(EventType)(0),        // 2: iproov.chiro.EventType
	(*AckMessage)(nil),    // 3: iproov.chiro.AckMessage
	(*CaptureResult)(nil), // 4: iproov.chiro.CaptureResult
	(*ClientStart)(nil),   // 5: iproov.chiro.ClientStart
	(*PadFrame)(nil),      // 6: iproov.chiro.PadFrame
	(*PalmSlab)(nil),      // 7: iproov.chiro.PalmSlab
	(*Landmark)(nil),      // 8: iproov.chiro.Landmark
	(*Rect)(nil),          // 9: iproov.chiro.Rect
	(*Size)(nil),          // 10: iproov.chiro.Size
	(*Envelope)(nil),      // 11: iproov.chiro.Envelope
}
var file_pkg_protocol_chiro_proto_depIdxs = []int32{
	10, // 0: iproov.chiro.PadFrame.imageSize:type_name -> iproov.chiro.Size
	8,  // 1: iproov.chiro.PadFrame.landmarks:type_name -> iproov.chiro.Landmark
	1,  // 2: iproov.chiro.PadFrame.captureType:type_name -> iproov.chiro.CaptureType
	0,  // 3: iproov.chiro.PadFrame.handedness:type_name -> iproov.chiro.Handedness
	10, // 4: iproov.chiro.PalmSlab.imageSize:type_name -> iproov.chiro.Size
	8,  // 5: iproov.chiro.PalmSlab.landmarks:type_name -> iproov.chiro.Landmark
	9,  // 6: iproov.chiro.PalmSlab.cropBounds:type_name -> iproov.chiro.Rect
	1,  // 7: iproov.chiro.PalmSlab.captureType:type_name -> iproov.chiro.CaptureType
	0,  // 8: iproov.chiro.PalmSlab.handedness:type_name -> iproov.chiro.Handedness
	2,  // 9: iproov.chiro.Envelope.event_type:type_name -> iproov.chiro.EventType
	5,  // 10: iproov.chiro.Envelope.clientStart:type_name -> iproov.chiro.ClientStart
	7,  // 11: iproov.chiro.Envelope.palmSlab:type_name -> iproov.chiro.PalmSlab
	6,  // 12: iproov.chiro.Envelope.padFrame:type_name -> iproov.chiro.PadFrame
	3,  // 13: iproov.chiro.Envelope.ackMessage:type_name -> iproov.chiro.AckMessage
	4,  // 14: iproov.chiro.Envelope.captureResult:type_name -> iproov.chiro.CaptureResult
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_pkg_protocol_chiro_proto_init() }
func file_pkg_protocol_chiro_proto_init() {
	if File_pkg_protocol_chiro_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_protocol_chiro_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_protocol_chiro_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaptureResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_protocol_chiro_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientStart); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_protocol_chiro_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PadFrame); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_protocol_chiro_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PalmSlab); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_protocol_chiro_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Landmark); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_protocol_chiro_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rect); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_protocol_chiro_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Size); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_protocol_chiro_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Envelope); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_pkg_protocol_chiro_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_pkg_protocol_chiro_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*Envelope_ClientStart)(nil),
		(*Envelope_PalmSlab)(nil),
		(*Envelope_PadFrame)(nil),
		(*Envelope_AckMessage)(nil),
		(*Envelope_CaptureResult)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_protocol_chiro_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_protocol_chiro_proto_goTypes,
		DependencyIndexes: file_pkg_protocol_chiro_proto_depIdxs,
		EnumInfos:         file_pkg_protocol_chiro_proto_enumTypes,
		MessageInfos:      file_pkg_protocol_chiro_proto_msgTypes,
	}.Build()
	File_pkg_protocol_chiro_proto = out.File
	file_pkg_protocol_chiro_proto_rawDesc = nil
	file_pkg_protocol_chiro_proto_goTypes = nil
	file_pkg_protocol_chiro_proto_depIdxs = nil
}