// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: pb/rolesvc/v1/rolesvc.proto

package v1

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

// ModelSvcClient is the client API for ModelSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelSvcClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	List(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type modelSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewModelSvcClient(cc grpc.ClientConnInterface) ModelSvcClient {
	return &modelSvcClient{cc}
}

func (c *modelSvcClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/pb.rolesvc.v1.ModelSvc/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelSvcClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/pb.rolesvc.v1.ModelSvc/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelSvcClient) List(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/pb.rolesvc.v1.ModelSvc/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelSvcClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/pb.rolesvc.v1.ModelSvc/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelSvcClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/pb.rolesvc.v1.ModelSvc/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelSvcServer is the server API for ModelSvc service.
// All implementations must embed UnimplementedModelSvcServer
// for forward compatibility
type ModelSvcServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	List(context.Context, *EmptyRequest) (*ListResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*EmptyResponse, error)
	mustEmbedUnimplementedModelSvcServer()
}

// UnimplementedModelSvcServer must be embedded to have forward compatible implementations.
type UnimplementedModelSvcServer struct {
}

func (UnimplementedModelSvcServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedModelSvcServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedModelSvcServer) List(context.Context, *EmptyRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedModelSvcServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedModelSvcServer) Delete(context.Context, *DeleteRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedModelSvcServer) mustEmbedUnimplementedModelSvcServer() {}

// UnsafeModelSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelSvcServer will
// result in compilation errors.
type UnsafeModelSvcServer interface {
	mustEmbedUnimplementedModelSvcServer()
}

func RegisterModelSvcServer(s grpc.ServiceRegistrar, srv ModelSvcServer) {
	s.RegisterService(&ModelSvc_ServiceDesc, srv)
}

func _ModelSvc_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelSvcServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.rolesvc.v1.ModelSvc/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelSvcServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelSvc_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelSvcServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.rolesvc.v1.ModelSvc/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelSvcServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelSvc_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelSvcServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.rolesvc.v1.ModelSvc/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelSvcServer).List(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelSvc_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelSvcServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.rolesvc.v1.ModelSvc/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelSvcServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelSvc_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelSvcServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.rolesvc.v1.ModelSvc/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelSvcServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ModelSvc_ServiceDesc is the grpc.ServiceDesc for ModelSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.rolesvc.v1.ModelSvc",
	HandlerType: (*ModelSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ModelSvc_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ModelSvc_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ModelSvc_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ModelSvc_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ModelSvc_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/rolesvc/v1/rolesvc.proto",
}
