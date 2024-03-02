// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: packages/backend/grpc-demo/proto/product.proto

package service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// so , should use package name + message name
	User *User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	// define a Any type, should import google/protobuf/any.proto
	Dto *anypb.Any `protobuf:"bytes,3,opt,name=dto,proto3" json:"dto,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_backend_grpc_demo_proto_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_packages_backend_grpc_demo_proto_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_packages_backend_grpc_demo_proto_product_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Product) GetDto() *anypb.Any {
	if x != nil {
		return x.Dto
	}
	return nil
}

var File_packages_backend_grpc_demo_proto_product_proto protoreflect.FileDescriptor

var file_packages_backend_grpc_demo_proto_product_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x2b, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x73, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x68, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x03, 0x64, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x03, 0x64, 0x74, 0x6f, 0x32, 0x82, 0x02, 0x0a, 0x0e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x10, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00,
	0x28, 0x01, 0x12, 0x3d, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x3d, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x42, 0x6f, 0x74, 0x68, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_packages_backend_grpc_demo_proto_product_proto_rawDescOnce sync.Once
	file_packages_backend_grpc_demo_proto_product_proto_rawDescData = file_packages_backend_grpc_demo_proto_product_proto_rawDesc
)

func file_packages_backend_grpc_demo_proto_product_proto_rawDescGZIP() []byte {
	file_packages_backend_grpc_demo_proto_product_proto_rawDescOnce.Do(func() {
		file_packages_backend_grpc_demo_proto_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_packages_backend_grpc_demo_proto_product_proto_rawDescData)
	})
	return file_packages_backend_grpc_demo_proto_product_proto_rawDescData
}

var file_packages_backend_grpc_demo_proto_product_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_packages_backend_grpc_demo_proto_product_proto_goTypes = []interface{}{
	(*Product)(nil),   // 0: service.Product
	(*User)(nil),      // 1: service.User
	(*anypb.Any)(nil), // 2: google.protobuf.Any
	(*UserInfo)(nil),  // 3: service.UserInfo
}
var file_packages_backend_grpc_demo_proto_product_proto_depIdxs = []int32{
	1, // 0: service.Product.user:type_name -> service.User
	2, // 1: service.Product.dto:type_name -> google.protobuf.Any
	3, // 2: service.ProductService.GetProduct:input_type -> service.UserInfo
	0, // 3: service.ProductService.UpdateProductClient:input_type -> service.Product
	0, // 4: service.ProductService.UpdateProductServer:input_type -> service.Product
	0, // 5: service.ProductService.UpdateProductBoth:input_type -> service.Product
	0, // 6: service.ProductService.GetProduct:output_type -> service.Product
	0, // 7: service.ProductService.UpdateProductClient:output_type -> service.Product
	0, // 8: service.ProductService.UpdateProductServer:output_type -> service.Product
	0, // 9: service.ProductService.UpdateProductBoth:output_type -> service.Product
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_packages_backend_grpc_demo_proto_product_proto_init() }
func file_packages_backend_grpc_demo_proto_product_proto_init() {
	if File_packages_backend_grpc_demo_proto_product_proto != nil {
		return
	}
	file_packages_backend_grpc_demo_proto_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_packages_backend_grpc_demo_proto_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
			RawDescriptor: file_packages_backend_grpc_demo_proto_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_packages_backend_grpc_demo_proto_product_proto_goTypes,
		DependencyIndexes: file_packages_backend_grpc_demo_proto_product_proto_depIdxs,
		MessageInfos:      file_packages_backend_grpc_demo_proto_product_proto_msgTypes,
	}.Build()
	File_packages_backend_grpc_demo_proto_product_proto = out.File
	file_packages_backend_grpc_demo_proto_product_proto_rawDesc = nil
	file_packages_backend_grpc_demo_proto_product_proto_goTypes = nil
	file_packages_backend_grpc_demo_proto_product_proto_depIdxs = nil
}
