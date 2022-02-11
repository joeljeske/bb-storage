// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: pkg/proto/configuration/tls/tls.proto

package tls

import (
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

type ClientConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerCertificateAuthorities string   `protobuf:"bytes,1,opt,name=server_certificate_authorities,json=serverCertificateAuthorities,proto3" json:"server_certificate_authorities,omitempty"`
	ClientCertificate            string   `protobuf:"bytes,2,opt,name=client_certificate,json=clientCertificate,proto3" json:"client_certificate,omitempty"`
	ClientPrivateKey             string   `protobuf:"bytes,3,opt,name=client_private_key,json=clientPrivateKey,proto3" json:"client_private_key,omitempty"`
	CipherSuites                 []string `protobuf:"bytes,4,rep,name=cipher_suites,json=cipherSuites,proto3" json:"cipher_suites,omitempty"`
	ServerName                   string   `protobuf:"bytes,5,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
}

func (x *ClientConfiguration) Reset() {
	*x = ClientConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_tls_tls_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConfiguration) ProtoMessage() {}

func (x *ClientConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_tls_tls_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConfiguration.ProtoReflect.Descriptor instead.
func (*ClientConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_tls_tls_proto_rawDescGZIP(), []int{0}
}

func (x *ClientConfiguration) GetServerCertificateAuthorities() string {
	if x != nil {
		return x.ServerCertificateAuthorities
	}
	return ""
}

func (x *ClientConfiguration) GetClientCertificate() string {
	if x != nil {
		return x.ClientCertificate
	}
	return ""
}

func (x *ClientConfiguration) GetClientPrivateKey() string {
	if x != nil {
		return x.ClientPrivateKey
	}
	return ""
}

func (x *ClientConfiguration) GetCipherSuites() []string {
	if x != nil {
		return x.CipherSuites
	}
	return nil
}

func (x *ClientConfiguration) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

type ServerConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerCertificate string   `protobuf:"bytes,1,opt,name=server_certificate,json=serverCertificate,proto3" json:"server_certificate,omitempty"`
	ServerPrivateKey  string   `protobuf:"bytes,2,opt,name=server_private_key,json=serverPrivateKey,proto3" json:"server_private_key,omitempty"`
	CipherSuites      []string `protobuf:"bytes,3,rep,name=cipher_suites,json=cipherSuites,proto3" json:"cipher_suites,omitempty"`
}

func (x *ServerConfiguration) Reset() {
	*x = ServerConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_tls_tls_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfiguration) ProtoMessage() {}

func (x *ServerConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_tls_tls_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfiguration.ProtoReflect.Descriptor instead.
func (*ServerConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_tls_tls_proto_rawDescGZIP(), []int{1}
}

func (x *ServerConfiguration) GetServerCertificate() string {
	if x != nil {
		return x.ServerCertificate
	}
	return ""
}

func (x *ServerConfiguration) GetServerPrivateKey() string {
	if x != nil {
		return x.ServerPrivateKey
	}
	return ""
}

func (x *ServerConfiguration) GetCipherSuites() []string {
	if x != nil {
		return x.CipherSuites
	}
	return nil
}

var File_pkg_proto_configuration_tls_tls_proto protoreflect.FileDescriptor

var file_pkg_proto_configuration_tls_tls_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x6c, 0x73, 0x2f, 0x74, 0x6c,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61,
	0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x74, 0x6c, 0x73, 0x22, 0xfe, 0x01, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x1e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x1c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x53, 0x75,
	0x69, 0x74, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a,
	0x12, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x12,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x69,
	0x70, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x53, 0x75, 0x69, 0x74, 0x65, 0x73, 0x42,
	0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2f, 0x62, 0x62, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x6c, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_configuration_tls_tls_proto_rawDescOnce sync.Once
	file_pkg_proto_configuration_tls_tls_proto_rawDescData = file_pkg_proto_configuration_tls_tls_proto_rawDesc
)

func file_pkg_proto_configuration_tls_tls_proto_rawDescGZIP() []byte {
	file_pkg_proto_configuration_tls_tls_proto_rawDescOnce.Do(func() {
		file_pkg_proto_configuration_tls_tls_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_configuration_tls_tls_proto_rawDescData)
	})
	return file_pkg_proto_configuration_tls_tls_proto_rawDescData
}

var file_pkg_proto_configuration_tls_tls_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_proto_configuration_tls_tls_proto_goTypes = []interface{}{
	(*ClientConfiguration)(nil), // 0: buildbarn.configuration.tls.ClientConfiguration
	(*ServerConfiguration)(nil), // 1: buildbarn.configuration.tls.ServerConfiguration
}
var file_pkg_proto_configuration_tls_tls_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_configuration_tls_tls_proto_init() }
func file_pkg_proto_configuration_tls_tls_proto_init() {
	if File_pkg_proto_configuration_tls_tls_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_configuration_tls_tls_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConfiguration); i {
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
		file_pkg_proto_configuration_tls_tls_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConfiguration); i {
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
			RawDescriptor: file_pkg_proto_configuration_tls_tls_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_configuration_tls_tls_proto_goTypes,
		DependencyIndexes: file_pkg_proto_configuration_tls_tls_proto_depIdxs,
		MessageInfos:      file_pkg_proto_configuration_tls_tls_proto_msgTypes,
	}.Build()
	File_pkg_proto_configuration_tls_tls_proto = out.File
	file_pkg_proto_configuration_tls_tls_proto_rawDesc = nil
	file_pkg_proto_configuration_tls_tls_proto_goTypes = nil
	file_pkg_proto_configuration_tls_tls_proto_depIdxs = nil
}
