// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: orders/order.proto

package orders

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

// OrdersClient is the client API for Orders service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrdersClient interface {
	AddOrder(ctx context.Context, in *PayloadWithSingleOrder, opts ...grpc.CallOption) (*Empty, error)
	GetOrder(ctx context.Context, in *PayloadWithOrderID, opts ...grpc.CallOption) (*PayloadWithSingleOrder, error)
	UpdateOrder(ctx context.Context, in *PayloadWithSingleOrder, opts ...grpc.CallOption) (*Empty, error)
	RemoveOrder(ctx context.Context, in *PayloadWithOrderID, opts ...grpc.CallOption) (*Empty, error)
}

type ordersClient struct {
	cc grpc.ClientConnInterface
}

func NewOrdersClient(cc grpc.ClientConnInterface) OrdersClient {
	return &ordersClient{cc}
}

func (c *ordersClient) AddOrder(ctx context.Context, in *PayloadWithSingleOrder, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Orders/AddOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersClient) GetOrder(ctx context.Context, in *PayloadWithOrderID, opts ...grpc.CallOption) (*PayloadWithSingleOrder, error) {
	out := new(PayloadWithSingleOrder)
	err := c.cc.Invoke(ctx, "/Orders/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersClient) UpdateOrder(ctx context.Context, in *PayloadWithSingleOrder, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Orders/UpdateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersClient) RemoveOrder(ctx context.Context, in *PayloadWithOrderID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Orders/RemoveOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrdersServer is the server API for Orders service.
// All implementations must embed UnimplementedOrdersServer
// for forward compatibility
type OrdersServer interface {
	AddOrder(context.Context, *PayloadWithSingleOrder) (*Empty, error)
	GetOrder(context.Context, *PayloadWithOrderID) (*PayloadWithSingleOrder, error)
	UpdateOrder(context.Context, *PayloadWithSingleOrder) (*Empty, error)
	RemoveOrder(context.Context, *PayloadWithOrderID) (*Empty, error)
	mustEmbedUnimplementedOrdersServer()
}

// UnimplementedOrdersServer must be embedded to have forward compatible implementations.
type UnimplementedOrdersServer struct {
}

func (UnimplementedOrdersServer) AddOrder(context.Context, *PayloadWithSingleOrder) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrder not implemented")
}
func (UnimplementedOrdersServer) GetOrder(context.Context, *PayloadWithOrderID) (*PayloadWithSingleOrder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrdersServer) UpdateOrder(context.Context, *PayloadWithSingleOrder) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedOrdersServer) RemoveOrder(context.Context, *PayloadWithOrderID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveOrder not implemented")
}
func (UnimplementedOrdersServer) mustEmbedUnimplementedOrdersServer() {}

// UnsafeOrdersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrdersServer will
// result in compilation errors.
type UnsafeOrdersServer interface {
	mustEmbedUnimplementedOrdersServer()
}

func RegisterOrdersServer(s grpc.ServiceRegistrar, srv OrdersServer) {
	s.RegisterService(&Orders_ServiceDesc, srv)
}

func _Orders_AddOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayloadWithSingleOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServer).AddOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Orders/AddOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServer).AddOrder(ctx, req.(*PayloadWithSingleOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orders_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayloadWithOrderID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Orders/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServer).GetOrder(ctx, req.(*PayloadWithOrderID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orders_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayloadWithSingleOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Orders/UpdateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServer).UpdateOrder(ctx, req.(*PayloadWithSingleOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orders_RemoveOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayloadWithOrderID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServer).RemoveOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Orders/RemoveOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServer).RemoveOrder(ctx, req.(*PayloadWithOrderID))
	}
	return interceptor(ctx, in, info, handler)
}

// Orders_ServiceDesc is the grpc.ServiceDesc for Orders service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Orders_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Orders",
	HandlerType: (*OrdersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddOrder",
			Handler:    _Orders_AddOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _Orders_GetOrder_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _Orders_UpdateOrder_Handler,
		},
		{
			MethodName: "RemoveOrder",
			Handler:    _Orders_RemoveOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orders/order.proto",
}
