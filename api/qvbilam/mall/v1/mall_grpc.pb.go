// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: mall.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MallClient is the client API for Mall service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MallClient interface {
	GetCategory(ctx context.Context, in *CategoryListRequest, opts ...grpc.CallOption) (*CategoryListResponse, error)
	GetGoodsList(ctx context.Context, in *GoodsListRequest, opts ...grpc.CallOption) (*GoodsDetailResponse, error)
	GetGoodsDetail(ctx context.Context, in *GoodsDetailRequest, opts ...grpc.CallOption) (*GoodsListResponse, error)
	Sell(ctx context.Context, in *SellRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Rollback(ctx context.Context, in *SellRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type mallClient struct {
	cc grpc.ClientConnInterface
}

func NewMallClient(cc grpc.ClientConnInterface) MallClient {
	return &mallClient{cc}
}

func (c *mallClient) GetCategory(ctx context.Context, in *CategoryListRequest, opts ...grpc.CallOption) (*CategoryListResponse, error) {
	out := new(CategoryListResponse)
	err := c.cc.Invoke(ctx, "/userPb.v1.Mall/GetCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallClient) GetGoodsList(ctx context.Context, in *GoodsListRequest, opts ...grpc.CallOption) (*GoodsDetailResponse, error) {
	out := new(GoodsDetailResponse)
	err := c.cc.Invoke(ctx, "/userPb.v1.Mall/GetGoodsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallClient) GetGoodsDetail(ctx context.Context, in *GoodsDetailRequest, opts ...grpc.CallOption) (*GoodsListResponse, error) {
	out := new(GoodsListResponse)
	err := c.cc.Invoke(ctx, "/userPb.v1.Mall/GetGoodsDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallClient) Sell(ctx context.Context, in *SellRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/userPb.v1.Mall/Sell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallClient) Rollback(ctx context.Context, in *SellRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/userPb.v1.Mall/Rollback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MallServer is the server API for Mall service.
// All implementations must embed UnimplementedMallServer
// for forward compatibility
type MallServer interface {
	GetCategory(context.Context, *CategoryListRequest) (*CategoryListResponse, error)
	GetGoodsList(context.Context, *GoodsListRequest) (*GoodsDetailResponse, error)
	GetGoodsDetail(context.Context, *GoodsDetailRequest) (*GoodsListResponse, error)
	Sell(context.Context, *SellRequest) (*emptypb.Empty, error)
	Rollback(context.Context, *SellRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMallServer()
}

// UnimplementedMallServer must be embedded to have forward compatible implementations.
type UnimplementedMallServer struct {
}

func (UnimplementedMallServer) GetCategory(context.Context, *CategoryListRequest) (*CategoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (UnimplementedMallServer) GetGoodsList(context.Context, *GoodsListRequest) (*GoodsDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsList not implemented")
}
func (UnimplementedMallServer) GetGoodsDetail(context.Context, *GoodsDetailRequest) (*GoodsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsDetail not implemented")
}
func (UnimplementedMallServer) Sell(context.Context, *SellRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sell not implemented")
}
func (UnimplementedMallServer) Rollback(context.Context, *SellRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rollback not implemented")
}
func (UnimplementedMallServer) mustEmbedUnimplementedMallServer() {}

// UnsafeMallServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MallServer will
// result in compilation errors.
type UnsafeMallServer interface {
	mustEmbedUnimplementedMallServer()
}

func RegisterMallServer(s grpc.ServiceRegistrar, srv MallServer) {
	s.RegisterService(&Mall_ServiceDesc, srv)
}

func _Mall_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userPb.v1.Mall/GetCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallServer).GetCategory(ctx, req.(*CategoryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mall_GetGoodsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallServer).GetGoodsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userPb.v1.Mall/GetGoodsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallServer).GetGoodsList(ctx, req.(*GoodsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mall_GetGoodsDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallServer).GetGoodsDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userPb.v1.Mall/GetGoodsDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallServer).GetGoodsDetail(ctx, req.(*GoodsDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mall_Sell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallServer).Sell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userPb.v1.Mall/Sell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallServer).Sell(ctx, req.(*SellRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mall_Rollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallServer).Rollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userPb.v1.Mall/Rollback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallServer).Rollback(ctx, req.(*SellRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mall_ServiceDesc is the grpc.ServiceDesc for Mall service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mall_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userPb.v1.Mall",
	HandlerType: (*MallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCategory",
			Handler:    _Mall_GetCategory_Handler,
		},
		{
			MethodName: "GetGoodsList",
			Handler:    _Mall_GetGoodsList_Handler,
		},
		{
			MethodName: "GetGoodsDetail",
			Handler:    _Mall_GetGoodsDetail_Handler,
		},
		{
			MethodName: "Sell",
			Handler:    _Mall_Sell_Handler,
		},
		{
			MethodName: "Rollback",
			Handler:    _Mall_Rollback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mall.proto",
}
