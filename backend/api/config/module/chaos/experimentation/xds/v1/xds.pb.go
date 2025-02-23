// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: config/module/chaos/experimentation/xds/v1/xds.proto

package xdsv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Interval in seconds between refreshes of xDS data from the backing store.
	CacheRefreshInterval *durationpb.Duration `protobuf:"bytes,1,opt,name=cache_refresh_interval,json=cacheRefreshInterval,proto3" json:"cache_refresh_interval,omitempty"`
	// Name of the RTDS layer in Envoy config i.e. envoy.yaml
	// https://www.envoyproxy.io/docs/envoy/latest/configuration/operations/runtime#config-runtime-rtds.
	RtdsLayerName string `protobuf:"bytes,2,opt,name=rtds_layer_name,json=rtdsLayerName,proto3" json:"rtds_layer_name,omitempty"`
	// The resource TTL to set for xDS resources.
	ResourceTtl *durationpb.Duration `protobuf:"bytes,3,opt,name=resource_ttl,json=resourceTtl,proto3" json:"resource_ttl,omitempty"`
	// The interval at which to send heartbeat responses for TTL'd resources.
	HeartbeatInterval *durationpb.Duration `protobuf:"bytes,4,opt,name=heartbeat_interval,json=heartbeatInterval,proto3" json:"heartbeat_interval,omitempty"`
	// Specifies which clusters ECDS is enabled for, causing faults to be communicated over ECDS instead of RTDS where
	// supported. Currently, allowing all clusters is not supported.
	EcdsAllowList *Config_ECDSAllowList `protobuf:"bytes,5,opt,name=ecds_allow_list,json=ecdsAllowList,proto3" json:"ecds_allow_list,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_module_chaos_experimentation_xds_v1_xds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_module_chaos_experimentation_xds_v1_xds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_config_module_chaos_experimentation_xds_v1_xds_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetCacheRefreshInterval() *durationpb.Duration {
	if x != nil {
		return x.CacheRefreshInterval
	}
	return nil
}

func (x *Config) GetRtdsLayerName() string {
	if x != nil {
		return x.RtdsLayerName
	}
	return ""
}

func (x *Config) GetResourceTtl() *durationpb.Duration {
	if x != nil {
		return x.ResourceTtl
	}
	return nil
}

func (x *Config) GetHeartbeatInterval() *durationpb.Duration {
	if x != nil {
		return x.HeartbeatInterval
	}
	return nil
}

func (x *Config) GetEcdsAllowList() *Config_ECDSAllowList {
	if x != nil {
		return x.EcdsAllowList
	}
	return nil
}

type Config_ECDSAllowList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnabledClusters []string `protobuf:"bytes,1,rep,name=enabled_clusters,json=enabledClusters,proto3" json:"enabled_clusters,omitempty"`
}

func (x *Config_ECDSAllowList) Reset() {
	*x = Config_ECDSAllowList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_module_chaos_experimentation_xds_v1_xds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config_ECDSAllowList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config_ECDSAllowList) ProtoMessage() {}

func (x *Config_ECDSAllowList) ProtoReflect() protoreflect.Message {
	mi := &file_config_module_chaos_experimentation_xds_v1_xds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config_ECDSAllowList.ProtoReflect.Descriptor instead.
func (*Config_ECDSAllowList) Descriptor() ([]byte, []int) {
	return file_config_module_chaos_experimentation_xds_v1_xds_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Config_ECDSAllowList) GetEnabledClusters() []string {
	if x != nil {
		return x.EnabledClusters
	}
	return nil
}

var File_config_module_chaos_experimentation_xds_v1_xds_proto protoreflect.FileDescriptor

var file_config_module_chaos_experimentation_xds_v1_xds_proto_rawDesc = []byte{
	0x0a, 0x34, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f,
	0x63, 0x68, 0x61, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x78, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x78, 0x64, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x31, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x63, 0x68, 0x61,
	0x6f, 0x73, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xed, 0x03, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x5b, 0x0a,
	0x16, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0xaa, 0x01, 0x04,
	0x08, 0x01, 0x2a, 0x00, 0x52, 0x14, 0x63, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x0f, 0x72, 0x74,
	0x64, 0x73, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x0d, 0x72, 0x74,
	0x64, 0x73, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0xfa, 0x42,
	0x07, 0xaa, 0x01, 0x04, 0x08, 0x01, 0x2a, 0x00, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x74, 0x6c, 0x12, 0x54, 0x0a, 0x12, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0xfa, 0x42,
	0x07, 0xaa, 0x01, 0x04, 0x08, 0x01, 0x2a, 0x00, 0x52, 0x11, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x6f, 0x0a, 0x0f, 0x65,
	0x63, 0x64, 0x73, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x6f,
	0x73, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x45, 0x43, 0x44, 0x53, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0d, 0x65,
	0x63, 0x64, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x44, 0x0a, 0x0d,
	0x45, 0x43, 0x44, 0x53, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x0a,
	0x10, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08,
	0x01, 0x52, 0x0f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x42, 0x55, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x78, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x78, 0x64, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_config_module_chaos_experimentation_xds_v1_xds_proto_rawDescOnce sync.Once
	file_config_module_chaos_experimentation_xds_v1_xds_proto_rawDescData = file_config_module_chaos_experimentation_xds_v1_xds_proto_rawDesc
)

func file_config_module_chaos_experimentation_xds_v1_xds_proto_rawDescGZIP() []byte {
	file_config_module_chaos_experimentation_xds_v1_xds_proto_rawDescOnce.Do(func() {
		file_config_module_chaos_experimentation_xds_v1_xds_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_module_chaos_experimentation_xds_v1_xds_proto_rawDescData)
	})
	return file_config_module_chaos_experimentation_xds_v1_xds_proto_rawDescData
}

var file_config_module_chaos_experimentation_xds_v1_xds_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_config_module_chaos_experimentation_xds_v1_xds_proto_goTypes = []interface{}{
	(*Config)(nil),               // 0: clutch.config.module.chaos.experimentation.xds.v1.Config
	(*Config_ECDSAllowList)(nil), // 1: clutch.config.module.chaos.experimentation.xds.v1.Config.ECDSAllowList
	(*durationpb.Duration)(nil),  // 2: google.protobuf.Duration
}
var file_config_module_chaos_experimentation_xds_v1_xds_proto_depIdxs = []int32{
	2, // 0: clutch.config.module.chaos.experimentation.xds.v1.Config.cache_refresh_interval:type_name -> google.protobuf.Duration
	2, // 1: clutch.config.module.chaos.experimentation.xds.v1.Config.resource_ttl:type_name -> google.protobuf.Duration
	2, // 2: clutch.config.module.chaos.experimentation.xds.v1.Config.heartbeat_interval:type_name -> google.protobuf.Duration
	1, // 3: clutch.config.module.chaos.experimentation.xds.v1.Config.ecds_allow_list:type_name -> clutch.config.module.chaos.experimentation.xds.v1.Config.ECDSAllowList
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_config_module_chaos_experimentation_xds_v1_xds_proto_init() }
func file_config_module_chaos_experimentation_xds_v1_xds_proto_init() {
	if File_config_module_chaos_experimentation_xds_v1_xds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_module_chaos_experimentation_xds_v1_xds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_config_module_chaos_experimentation_xds_v1_xds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config_ECDSAllowList); i {
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
			RawDescriptor: file_config_module_chaos_experimentation_xds_v1_xds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_module_chaos_experimentation_xds_v1_xds_proto_goTypes,
		DependencyIndexes: file_config_module_chaos_experimentation_xds_v1_xds_proto_depIdxs,
		MessageInfos:      file_config_module_chaos_experimentation_xds_v1_xds_proto_msgTypes,
	}.Build()
	File_config_module_chaos_experimentation_xds_v1_xds_proto = out.File
	file_config_module_chaos_experimentation_xds_v1_xds_proto_rawDesc = nil
	file_config_module_chaos_experimentation_xds_v1_xds_proto_goTypes = nil
	file_config_module_chaos_experimentation_xds_v1_xds_proto_depIdxs = nil
}
