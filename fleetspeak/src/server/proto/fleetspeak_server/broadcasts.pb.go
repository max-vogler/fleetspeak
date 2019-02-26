// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fleetspeak/src/server/proto/fleetspeak_server/broadcasts.proto

package fleetspeak_server

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/any"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import fleetspeak "github.com/google/fleetspeak/fleetspeak/src/common/proto/fleetspeak"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A Broadcast is a template to build messages to send to a number of machines.
type Broadcast struct {
	BroadcastId []byte `protobuf:"bytes,1,opt,name=broadcast_id,json=broadcastId,proto3" json:"broadcast_id,omitempty"`
	// The source of the broadcast, it should only be a server side service. The
	// destinations for the resulting broadcast messages will be different clients
	// with the same service name.
	Source      *fleetspeak.Address `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	MessageType string              `protobuf:"bytes,3,opt,name=message_type,json=messageType" json:"message_type,omitempty"`
	// A client will only be sent this broadcast if it has been marked with all of
	// the required labels.
	RequiredLabels []*fleetspeak.Label `protobuf:"bytes,4,rep,name=required_labels,json=requiredLabels" json:"required_labels,omitempty"`
	// A broadcast will stop being sent at this time.
	ExpirationTime *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=expiration_time,json=expirationTime" json:"expiration_time,omitempty"`
	Data           *google_protobuf1.Any      `protobuf:"bytes,6,opt,name=data" json:"data,omitempty"`
}

func (m *Broadcast) Reset()                    { *m = Broadcast{} }
func (m *Broadcast) String() string            { return proto.CompactTextString(m) }
func (*Broadcast) ProtoMessage()               {}
func (*Broadcast) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Broadcast) GetBroadcastId() []byte {
	if m != nil {
		return m.BroadcastId
	}
	return nil
}

func (m *Broadcast) GetSource() *fleetspeak.Address {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Broadcast) GetMessageType() string {
	if m != nil {
		return m.MessageType
	}
	return ""
}

func (m *Broadcast) GetRequiredLabels() []*fleetspeak.Label {
	if m != nil {
		return m.RequiredLabels
	}
	return nil
}

func (m *Broadcast) GetExpirationTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

func (m *Broadcast) GetData() *google_protobuf1.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Broadcast)(nil), "fleetspeak.server.Broadcast")
}

func init() {
	proto.RegisterFile("fleetspeak/src/server/proto/fleetspeak_server/broadcasts.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x6b, 0xc2, 0x30,
	0x18, 0xc5, 0xa9, 0x3a, 0xc1, 0x54, 0x14, 0xbb, 0x1d, 0x32, 0x2f, 0xeb, 0x76, 0x2a, 0x0c, 0x52,
	0xd0, 0xdb, 0x0e, 0x03, 0xb7, 0xd3, 0x60, 0xa7, 0xe0, 0xbd, 0xa4, 0xcd, 0xa7, 0x94, 0xb5, 0x4d,
	0x96, 0x2f, 0x8e, 0xf5, 0x4f, 0xda, 0x7f, 0x39, 0x9a, 0xb6, 0x2a, 0x7a, 0xcc, 0xfb, 0x7e, 0xef,
	0xf1, 0xf2, 0xc8, 0xeb, 0xae, 0x00, 0xb0, 0xa8, 0x41, 0x7c, 0xc5, 0x68, 0xb2, 0x18, 0xc1, 0xfc,
	0x80, 0x89, 0xb5, 0x51, 0x56, 0xc5, 0xa7, 0x5b, 0xd2, 0xe9, 0xa9, 0x51, 0x42, 0x66, 0x02, 0x2d,
	0x32, 0x87, 0x04, 0x8b, 0x13, 0xc3, 0x5a, 0x66, 0x79, 0xbf, 0x57, 0x6a, 0x5f, 0x40, 0x9b, 0x91,
	0x1e, 0x76, 0xb1, 0xa8, 0xea, 0x96, 0x5e, 0x3e, 0x5c, 0x9e, 0x6c, 0x5e, 0x02, 0x5a, 0x51, 0xea,
	0x0e, 0x58, 0x5f, 0xd4, 0xc9, 0x54, 0x59, 0xaa, 0xea, 0xaa, 0x4e, 0xa7, 0xb7, 0xa6, 0xa7, 0xbf,
	0x01, 0x99, 0xbc, 0xf5, 0xc5, 0x82, 0x47, 0x32, 0x3d, 0xb6, 0x4c, 0x72, 0x49, 0xbd, 0xd0, 0x8b,
	0xa6, 0xdc, 0x3f, 0x6a, 0x1f, 0x32, 0x78, 0x26, 0x63, 0x54, 0x07, 0x93, 0x01, 0x1d, 0x84, 0x5e,
	0xe4, 0xaf, 0x6e, 0xd9, 0xd9, 0x2f, 0x36, 0x52, 0x1a, 0x40, 0xe4, 0x1d, 0xd2, 0xe4, 0x95, 0x80,
	0x28, 0xf6, 0x90, 0xd8, 0x5a, 0x03, 0x1d, 0x86, 0x5e, 0x34, 0xe1, 0x7e, 0xa7, 0x6d, 0x6b, 0x0d,
	0xc1, 0x0b, 0x99, 0x1b, 0xf8, 0x3e, 0xe4, 0x06, 0x64, 0x52, 0x88, 0x14, 0x0a, 0xa4, 0xa3, 0x70,
	0x18, 0xf9, 0xab, 0xc5, 0x79, 0xf0, 0x67, 0x73, 0xe1, 0xb3, 0x9e, 0x74, 0x4f, 0x0c, 0xde, 0xc9,
	0x1c, 0x7e, 0x75, 0x6e, 0x84, 0xcd, 0x55, 0x95, 0x34, 0x7b, 0xd0, 0x1b, 0x57, 0x6a, 0xc9, 0xda,
	0xb1, 0x58, 0x3f, 0x16, 0xdb, 0xf6, 0x63, 0xf1, 0xd9, 0xc9, 0xd2, 0x88, 0x41, 0x44, 0x46, 0x52,
	0x58, 0x41, 0xc7, 0xce, 0x79, 0x77, 0xe5, 0xdc, 0x54, 0x35, 0x77, 0x44, 0x3a, 0x76, 0xda, 0xfa,
	0x3f, 0x00, 0x00, 0xff, 0xff, 0x74, 0xa5, 0x54, 0x74, 0xf8, 0x01, 0x00, 0x00,
}
