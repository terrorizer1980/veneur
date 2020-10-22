// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: tdigest/tdigest.proto

package tdigest

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// MergingDigestData contains all fields necessary to generate a MergingDigest.
// This type should generally just be used when serializing MergingDigest's,
// and doesn't have much of a purpose on its own.
type MergingDigestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// With the standard Go protobuf generation, we have to use pointers
	// for the Centroid array, even though values would avoid a ton of code changes
	// and probably a lot of allocations as well.
	MainCentroids []*Centroid `protobuf:"bytes,1,rep,name=main_centroids,json=mainCentroids,proto3" json:"main_centroids,omitempty"`
	Compression   float64     `protobuf:"fixed64,2,opt,name=compression,proto3" json:"compression,omitempty"`
	Min           float64     `protobuf:"fixed64,3,opt,name=min,proto3" json:"min,omitempty"`
	Max           float64     `protobuf:"fixed64,4,opt,name=max,proto3" json:"max,omitempty"`
	ReciprocalSum float64     `protobuf:"fixed64,5,opt,name=reciprocalSum,proto3" json:"reciprocalSum,omitempty"`
}

func (x *MergingDigestData) Reset() {
	*x = MergingDigestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdigest_tdigest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MergingDigestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MergingDigestData) ProtoMessage() {}

func (x *MergingDigestData) ProtoReflect() protoreflect.Message {
	mi := &file_tdigest_tdigest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MergingDigestData.ProtoReflect.Descriptor instead.
func (*MergingDigestData) Descriptor() ([]byte, []int) {
	return file_tdigest_tdigest_proto_rawDescGZIP(), []int{0}
}

func (x *MergingDigestData) GetMainCentroids() []*Centroid {
	if x != nil {
		return x.MainCentroids
	}
	return nil
}

func (x *MergingDigestData) GetCompression() float64 {
	if x != nil {
		return x.Compression
	}
	return 0
}

func (x *MergingDigestData) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *MergingDigestData) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *MergingDigestData) GetReciprocalSum() float64 {
	if x != nil {
		return x.ReciprocalSum
	}
	return 0
}

type Centroid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mean    float64   `protobuf:"fixed64,1,opt,name=mean,proto3" json:"mean,omitempty"`
	Weight  float64   `protobuf:"fixed64,2,opt,name=weight,proto3" json:"weight,omitempty"`
	Samples []float64 `protobuf:"fixed64,3,rep,packed,name=samples,proto3" json:"samples,omitempty"`
}

func (x *Centroid) Reset() {
	*x = Centroid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdigest_tdigest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Centroid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Centroid) ProtoMessage() {}

func (x *Centroid) ProtoReflect() protoreflect.Message {
	mi := &file_tdigest_tdigest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Centroid.ProtoReflect.Descriptor instead.
func (*Centroid) Descriptor() ([]byte, []int) {
	return file_tdigest_tdigest_proto_rawDescGZIP(), []int{1}
}

func (x *Centroid) GetMean() float64 {
	if x != nil {
		return x.Mean
	}
	return 0
}

func (x *Centroid) GetWeight() float64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Centroid) GetSamples() []float64 {
	if x != nil {
		return x.Samples
	}
	return nil
}

var File_tdigest_tdigest_proto protoreflect.FileDescriptor

var file_tdigest_tdigest_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x64, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x22, 0xb9, 0x01, 0x0a, 0x11, 0x4d, 0x65, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x0e, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x72, 0x6f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x74, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x69,
	0x64, 0x52, 0x0d, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x69, 0x64, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x69, 0x70, 0x72,
	0x6f, 0x63, 0x61, 0x6c, 0x53, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x72, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x75, 0x6d, 0x22, 0x50, 0x0a, 0x08,
	0x43, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x61, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6d, 0x65, 0x61, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x01, 0x52, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x42, 0x22,
	0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72,
	0x69, 0x70, 0x65, 0x2f, 0x76, 0x65, 0x6e, 0x65, 0x75, 0x72, 0x2f, 0x74, 0x64, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tdigest_tdigest_proto_rawDescOnce sync.Once
	file_tdigest_tdigest_proto_rawDescData = file_tdigest_tdigest_proto_rawDesc
)

func file_tdigest_tdigest_proto_rawDescGZIP() []byte {
	file_tdigest_tdigest_proto_rawDescOnce.Do(func() {
		file_tdigest_tdigest_proto_rawDescData = protoimpl.X.CompressGZIP(file_tdigest_tdigest_proto_rawDescData)
	})
	return file_tdigest_tdigest_proto_rawDescData
}

var file_tdigest_tdigest_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tdigest_tdigest_proto_goTypes = []interface{}{
	(*MergingDigestData)(nil), // 0: tdigest.MergingDigestData
	(*Centroid)(nil),          // 1: tdigest.Centroid
}
var file_tdigest_tdigest_proto_depIdxs = []int32{
	1, // 0: tdigest.MergingDigestData.main_centroids:type_name -> tdigest.Centroid
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tdigest_tdigest_proto_init() }
func file_tdigest_tdigest_proto_init() {
	if File_tdigest_tdigest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tdigest_tdigest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MergingDigestData); i {
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
		file_tdigest_tdigest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Centroid); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tdigest_tdigest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tdigest_tdigest_proto_goTypes,
		DependencyIndexes: file_tdigest_tdigest_proto_depIdxs,
		MessageInfos:      file_tdigest_tdigest_proto_msgTypes,
	}.Build()
	File_tdigest_tdigest_proto = out.File
	file_tdigest_tdigest_proto_rawDesc = nil
	file_tdigest_tdigest_proto_goTypes = nil
	file_tdigest_tdigest_proto_depIdxs = nil
}
