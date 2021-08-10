// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package userfriend

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

// UserFriendRPCClient is the client API for UserFriendRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserFriendRPCClient interface {
	GetListFriends(ctx context.Context, in *ConditionRequest, opts ...grpc.CallOption) (*ListFriendIdsResponse, error)
}

type userFriendRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewUserFriendRPCClient(cc grpc.ClientConnInterface) UserFriendRPCClient {
	return &userFriendRPCClient{cc}
}

func (c *userFriendRPCClient) GetListFriends(ctx context.Context, in *ConditionRequest, opts ...grpc.CallOption) (*ListFriendIdsResponse, error) {
	out := new(ListFriendIdsResponse)
	err := c.cc.Invoke(ctx, "/userfriend.UserFriendRPC/GetListFriends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserFriendRPCServer is the server API for UserFriendRPC service.
// All implementations must embed UnimplementedUserFriendRPCServer
// for forward compatibility
type UserFriendRPCServer interface {
	GetListFriends(context.Context, *ConditionRequest) (*ListFriendIdsResponse, error)
	mustEmbedUnimplementedUserFriendRPCServer()
}

// UnimplementedUserFriendRPCServer must be embedded to have forward compatible implementations.
type UnimplementedUserFriendRPCServer struct {
}

func (UnimplementedUserFriendRPCServer) GetListFriends(context.Context, *ConditionRequest) (*ListFriendIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListFriends not implemented")
}
func (UnimplementedUserFriendRPCServer) mustEmbedUnimplementedUserFriendRPCServer() {}

// UnsafeUserFriendRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserFriendRPCServer will
// result in compilation errors.
type UnsafeUserFriendRPCServer interface {
	mustEmbedUnimplementedUserFriendRPCServer()
}

func RegisterUserFriendRPCServer(s grpc.ServiceRegistrar, srv UserFriendRPCServer) {
	s.RegisterService(&UserFriendRPC_ServiceDesc, srv)
}

func _UserFriendRPC_GetListFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConditionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserFriendRPCServer).GetListFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userfriend.UserFriendRPC/GetListFriends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserFriendRPCServer).GetListFriends(ctx, req.(*ConditionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserFriendRPC_ServiceDesc is the grpc.ServiceDesc for UserFriendRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserFriendRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userfriend.UserFriendRPC",
	HandlerType: (*UserFriendRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetListFriends",
			Handler:    _UserFriendRPC_GetListFriends_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/userfriend/userfriend.proto",
}
