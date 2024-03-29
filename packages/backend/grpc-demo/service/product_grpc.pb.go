// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: packages/backend/grpc-demo/proto/product.proto

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ProductService_GetProduct_FullMethodName          = "/service.ProductService/GetProduct"
	ProductService_UpdateProductClient_FullMethodName = "/service.ProductService/UpdateProductClient"
	ProductService_UpdateProductServer_FullMethodName = "/service.ProductService/UpdateProductServer"
	ProductService_UpdateProductBoth_FullMethodName   = "/service.ProductService/UpdateProductBoth"
)

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	GetProduct(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*Product, error)
	// 客户端流 只有传进来的参数是 带 stream 关键字
	UpdateProductClient(ctx context.Context, opts ...grpc.CallOption) (ProductService_UpdateProductClientClient, error)
	// 服务端流 只有返回的参数是 带 stream 关键字
	UpdateProductServer(ctx context.Context, in *Product, opts ...grpc.CallOption) (ProductService_UpdateProductServerClient, error)
	// 双向流 两边都是带关键子的
	UpdateProductBoth(ctx context.Context, opts ...grpc.CallOption) (ProductService_UpdateProductBothClient, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) GetProduct(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, ProductService_GetProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateProductClient(ctx context.Context, opts ...grpc.CallOption) (ProductService_UpdateProductClientClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProductService_ServiceDesc.Streams[0], ProductService_UpdateProductClient_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceUpdateProductClientClient{stream}
	return x, nil
}

type ProductService_UpdateProductClientClient interface {
	Send(*Product) error
	CloseAndRecv() (*Product, error)
	grpc.ClientStream
}

type productServiceUpdateProductClientClient struct {
	grpc.ClientStream
}

func (x *productServiceUpdateProductClientClient) Send(m *Product) error {
	return x.ClientStream.SendMsg(m)
}

func (x *productServiceUpdateProductClientClient) CloseAndRecv() (*Product, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Product)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productServiceClient) UpdateProductServer(ctx context.Context, in *Product, opts ...grpc.CallOption) (ProductService_UpdateProductServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProductService_ServiceDesc.Streams[1], ProductService_UpdateProductServer_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceUpdateProductServerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProductService_UpdateProductServerClient interface {
	Recv() (*Product, error)
	grpc.ClientStream
}

type productServiceUpdateProductServerClient struct {
	grpc.ClientStream
}

func (x *productServiceUpdateProductServerClient) Recv() (*Product, error) {
	m := new(Product)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productServiceClient) UpdateProductBoth(ctx context.Context, opts ...grpc.CallOption) (ProductService_UpdateProductBothClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProductService_ServiceDesc.Streams[2], ProductService_UpdateProductBoth_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceUpdateProductBothClient{stream}
	return x, nil
}

type ProductService_UpdateProductBothClient interface {
	Send(*Product) error
	Recv() (*Product, error)
	grpc.ClientStream
}

type productServiceUpdateProductBothClient struct {
	grpc.ClientStream
}

func (x *productServiceUpdateProductBothClient) Send(m *Product) error {
	return x.ClientStream.SendMsg(m)
}

func (x *productServiceUpdateProductBothClient) Recv() (*Product, error) {
	m := new(Product)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	GetProduct(context.Context, *UserInfo) (*Product, error)
	// 客户端流 只有传进来的参数是 带 stream 关键字
	UpdateProductClient(ProductService_UpdateProductClientServer) error
	// 服务端流 只有返回的参数是 带 stream 关键字
	UpdateProductServer(*Product, ProductService_UpdateProductServerServer) error
	// 双向流 两边都是带关键子的
	UpdateProductBoth(ProductService_UpdateProductBothServer) error
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) GetProduct(context.Context, *UserInfo) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductServiceServer) UpdateProductClient(ProductService_UpdateProductClientServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateProductClient not implemented")
}
func (UnimplementedProductServiceServer) UpdateProductServer(*Product, ProductService_UpdateProductServerServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateProductServer not implemented")
}
func (UnimplementedProductServiceServer) UpdateProductBoth(ProductService_UpdateProductBothServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateProductBoth not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProduct(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateProductClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProductServiceServer).UpdateProductClient(&productServiceUpdateProductClientServer{stream})
}

type ProductService_UpdateProductClientServer interface {
	SendAndClose(*Product) error
	Recv() (*Product, error)
	grpc.ServerStream
}

type productServiceUpdateProductClientServer struct {
	grpc.ServerStream
}

func (x *productServiceUpdateProductClientServer) SendAndClose(m *Product) error {
	return x.ServerStream.SendMsg(m)
}

func (x *productServiceUpdateProductClientServer) Recv() (*Product, error) {
	m := new(Product)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProductService_UpdateProductServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Product)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductServiceServer).UpdateProductServer(m, &productServiceUpdateProductServerServer{stream})
}

type ProductService_UpdateProductServerServer interface {
	Send(*Product) error
	grpc.ServerStream
}

type productServiceUpdateProductServerServer struct {
	grpc.ServerStream
}

func (x *productServiceUpdateProductServerServer) Send(m *Product) error {
	return x.ServerStream.SendMsg(m)
}

func _ProductService_UpdateProductBoth_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProductServiceServer).UpdateProductBoth(&productServiceUpdateProductBothServer{stream})
}

type ProductService_UpdateProductBothServer interface {
	Send(*Product) error
	Recv() (*Product, error)
	grpc.ServerStream
}

type productServiceUpdateProductBothServer struct {
	grpc.ServerStream
}

func (x *productServiceUpdateProductBothServer) Send(m *Product) error {
	return x.ServerStream.SendMsg(m)
}

func (x *productServiceUpdateProductBothServer) Recv() (*Product, error) {
	m := new(Product)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProduct",
			Handler:    _ProductService_GetProduct_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdateProductClient",
			Handler:       _ProductService_UpdateProductClient_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UpdateProductServer",
			Handler:       _ProductService_UpdateProductServer_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UpdateProductBoth",
			Handler:       _ProductService_UpdateProductBoth_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "packages/backend/grpc-demo/proto/product.proto",
}
