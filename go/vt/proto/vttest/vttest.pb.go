// Code generated by protoc-gen-go.
// source: vttest.proto
// DO NOT EDIT!

/*
Package vttest is a generated protocol buffer package.

It is generated from these files:
	vttest.proto

It has these top-level messages:
	Shard
	Keyspace
	VTTestTopology
*/
package vttest

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

// Shard describes a single shard in a keyspace.
type Shard struct {
	// name has to be unique in a keyspace. For unsharded keyspaces, it
	// should be '0'. For sharded keyspace, it should be derived from
	// the keyrange, like '-80' or '40-80'.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// db_name_override is the mysql db name for this shard. Has to be
	// globally unique. If not specified, we will by default use
	// 'vt_<keyspace>_<shard>'.
	DbNameOverride string `protobuf:"bytes,2,opt,name=db_name_override,json=dbNameOverride" json:"db_name_override,omitempty"`
}

func (m *Shard) Reset()                    { *m = Shard{} }
func (m *Shard) String() string            { return proto.CompactTextString(m) }
func (*Shard) ProtoMessage()               {}
func (*Shard) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Keyspace describes a single keyspace.
type Keyspace struct {
	// name has to be unique in a VTTestTopology.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// shards inside this keyspace. Ignored if redirect is set.
	Shards []*Shard `protobuf:"bytes,2,rep,name=shards" json:"shards,omitempty"`
	// sharding_column_name for this keyspace. Used for v2 calls, but not for v3.
	ShardingColumnName string `protobuf:"bytes,3,opt,name=sharding_column_name,json=shardingColumnName" json:"sharding_column_name,omitempty"`
	// sharding_column_type for this keyspace. Used for v2 calls, but not for v3.
	ShardingColumnType string `protobuf:"bytes,4,opt,name=sharding_column_type,json=shardingColumnType" json:"sharding_column_type,omitempty"`
	// redirects all traffic to another keyspace. If set, shards is ignored.
	ServedFrom string `protobuf:"bytes,5,opt,name=served_from,json=servedFrom" json:"served_from,omitempty"`
}

func (m *Keyspace) Reset()                    { *m = Keyspace{} }
func (m *Keyspace) String() string            { return proto.CompactTextString(m) }
func (*Keyspace) ProtoMessage()               {}
func (*Keyspace) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Keyspace) GetShards() []*Shard {
	if m != nil {
		return m.Shards
	}
	return nil
}

// VTTestTopology describes the keyspaces in the topology.
type VTTestTopology struct {
	// all keyspaces in the topology.
	Keyspaces []*Keyspace `protobuf:"bytes,1,rep,name=keyspaces" json:"keyspaces,omitempty"`
}

func (m *VTTestTopology) Reset()                    { *m = VTTestTopology{} }
func (m *VTTestTopology) String() string            { return proto.CompactTextString(m) }
func (*VTTestTopology) ProtoMessage()               {}
func (*VTTestTopology) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *VTTestTopology) GetKeyspaces() []*Keyspace {
	if m != nil {
		return m.Keyspaces
	}
	return nil
}

func init() {
	proto.RegisterType((*Shard)(nil), "vttest.Shard")
	proto.RegisterType((*Keyspace)(nil), "vttest.Keyspace")
	proto.RegisterType((*VTTestTopology)(nil), "vttest.VTTestTopology")
}

func init() { proto.RegisterFile("vttest.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x50, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x26, 0xb6, 0x0d, 0x76, 0xaa, 0xa5, 0x0c, 0x1e, 0xf6, 0xa6, 0x04, 0x84, 0x9c, 0x82, 0xe8,
	0x0b, 0x08, 0xa2, 0x17, 0x41, 0x21, 0x06, 0xaf, 0x21, 0xe9, 0x8e, 0xb5, 0xd8, 0x64, 0xc2, 0xee,
	0x1a, 0xc8, 0x2b, 0xfa, 0x54, 0x6e, 0xa6, 0x09, 0xbd, 0xe4, 0x36, 0xfb, 0x7d, 0x33, 0xdf, 0xcf,
	0xc2, 0x45, 0xeb, 0x1c, 0x59, 0x97, 0x34, 0x86, 0x1d, 0x63, 0x78, 0x7c, 0x45, 0xcf, 0xb0, 0xf8,
	0xf8, 0x2e, 0x8c, 0x46, 0x84, 0x79, 0x5d, 0x54, 0xa4, 0x82, 0x9b, 0x20, 0x5e, 0xa6, 0x32, 0x63,
	0x0c, 0x1b, 0x5d, 0xe6, 0xfd, 0x98, 0x73, 0x4b, 0xc6, 0xec, 0x35, 0xa9, 0x33, 0xe1, 0xd7, 0xba,
	0x7c, 0xf3, 0xf0, 0xfb, 0x80, 0x46, 0x7f, 0x01, 0x9c, 0xbf, 0x52, 0x67, 0x9b, 0x62, 0x4b, 0x93,
	0x52, 0xb7, 0x10, 0xda, 0xde, 0xc7, 0x7a, 0x81, 0x59, 0xbc, 0xba, 0xbf, 0x4c, 0x86, 0x38, 0xe2,
	0x9e, 0x0e, 0x24, 0xde, 0xc1, 0x95, 0x4c, 0xfb, 0x7a, 0x97, 0x6f, 0xf9, 0xf0, 0x5b, 0xd5, 0x62,
	0xaf, 0x66, 0x22, 0x85, 0x23, 0xf7, 0x24, 0x54, 0x9f, 0x60, 0xea, 0xc2, 0x75, 0x0d, 0xa9, 0xf9,
	0xd4, 0x45, 0xe6, 0x19, 0xbc, 0x86, 0x95, 0x25, 0xd3, 0x92, 0xce, 0xbf, 0x0c, 0x57, 0x6a, 0x21,
	0x8b, 0x70, 0x84, 0x5e, 0x3c, 0x12, 0x3d, 0xc2, 0xfa, 0x33, 0xcb, 0x7c, 0xb8, 0x8c, 0x1b, 0x3e,
	0xf0, 0xae, 0xc3, 0x04, 0x96, 0x3f, 0x43, 0x3b, 0xeb, 0x6b, 0xf5, 0x05, 0x36, 0x63, 0x81, 0xb1,
	0x76, 0x7a, 0x5a, 0x29, 0x43, 0xf9, 0xe4, 0x87, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x49,
	0x37, 0xec, 0x74, 0x01, 0x00, 0x00,
}
