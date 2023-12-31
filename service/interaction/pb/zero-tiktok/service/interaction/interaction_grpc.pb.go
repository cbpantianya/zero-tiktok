// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: interaction.proto

package interaction

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
	InteractionService_Comment_FullMethodName      = "/interaction.InteractionService/Comment"
	InteractionService_CommentList_FullMethodName  = "/interaction.InteractionService/CommentList"
	InteractionService_Relation_FullMethodName     = "/interaction.InteractionService/Relation"
	InteractionService_FollowList_FullMethodName   = "/interaction.InteractionService/FollowList"
	InteractionService_FollowerList_FullMethodName = "/interaction.InteractionService/FollowerList"
	InteractionService_FriendList_FullMethodName   = "/interaction.InteractionService/FriendList"
	InteractionService_HasFollowed_FullMethodName  = "/interaction.InteractionService/HasFollowed"
)

// InteractionServiceClient is the client API for InteractionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InteractionServiceClient interface {
	Comment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error)
	CommentList(ctx context.Context, in *CommentListRequest, opts ...grpc.CallOption) (*CommentListResponse, error)
	Relation(ctx context.Context, in *RelationRequest, opts ...grpc.CallOption) (*RelationResponse, error)
	FollowList(ctx context.Context, in *FollowListRequest, opts ...grpc.CallOption) (*FollowListResponse, error)
	FollowerList(ctx context.Context, in *FollowerListRequest, opts ...grpc.CallOption) (*FollowerListResponse, error)
	FriendList(ctx context.Context, in *FriendListRequest, opts ...grpc.CallOption) (*FriendListResponse, error)
	HasFollowed(ctx context.Context, in *HasFollowedRequest, opts ...grpc.CallOption) (*HasFollowedResponse, error)
}

type interactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInteractionServiceClient(cc grpc.ClientConnInterface) InteractionServiceClient {
	return &interactionServiceClient{cc}
}

func (c *interactionServiceClient) Comment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, InteractionService_Comment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionServiceClient) CommentList(ctx context.Context, in *CommentListRequest, opts ...grpc.CallOption) (*CommentListResponse, error) {
	out := new(CommentListResponse)
	err := c.cc.Invoke(ctx, InteractionService_CommentList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionServiceClient) Relation(ctx context.Context, in *RelationRequest, opts ...grpc.CallOption) (*RelationResponse, error) {
	out := new(RelationResponse)
	err := c.cc.Invoke(ctx, InteractionService_Relation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionServiceClient) FollowList(ctx context.Context, in *FollowListRequest, opts ...grpc.CallOption) (*FollowListResponse, error) {
	out := new(FollowListResponse)
	err := c.cc.Invoke(ctx, InteractionService_FollowList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionServiceClient) FollowerList(ctx context.Context, in *FollowerListRequest, opts ...grpc.CallOption) (*FollowerListResponse, error) {
	out := new(FollowerListResponse)
	err := c.cc.Invoke(ctx, InteractionService_FollowerList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionServiceClient) FriendList(ctx context.Context, in *FriendListRequest, opts ...grpc.CallOption) (*FriendListResponse, error) {
	out := new(FriendListResponse)
	err := c.cc.Invoke(ctx, InteractionService_FriendList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionServiceClient) HasFollowed(ctx context.Context, in *HasFollowedRequest, opts ...grpc.CallOption) (*HasFollowedResponse, error) {
	out := new(HasFollowedResponse)
	err := c.cc.Invoke(ctx, InteractionService_HasFollowed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InteractionServiceServer is the server API for InteractionService service.
// All implementations must embed UnimplementedInteractionServiceServer
// for forward compatibility
type InteractionServiceServer interface {
	Comment(context.Context, *CommentRequest) (*CommentResponse, error)
	CommentList(context.Context, *CommentListRequest) (*CommentListResponse, error)
	Relation(context.Context, *RelationRequest) (*RelationResponse, error)
	FollowList(context.Context, *FollowListRequest) (*FollowListResponse, error)
	FollowerList(context.Context, *FollowerListRequest) (*FollowerListResponse, error)
	FriendList(context.Context, *FriendListRequest) (*FriendListResponse, error)
	HasFollowed(context.Context, *HasFollowedRequest) (*HasFollowedResponse, error)
	mustEmbedUnimplementedInteractionServiceServer()
}

// UnimplementedInteractionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInteractionServiceServer struct {
}

func (UnimplementedInteractionServiceServer) Comment(context.Context, *CommentRequest) (*CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Comment not implemented")
}
func (UnimplementedInteractionServiceServer) CommentList(context.Context, *CommentListRequest) (*CommentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentList not implemented")
}
func (UnimplementedInteractionServiceServer) Relation(context.Context, *RelationRequest) (*RelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Relation not implemented")
}
func (UnimplementedInteractionServiceServer) FollowList(context.Context, *FollowListRequest) (*FollowListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowList not implemented")
}
func (UnimplementedInteractionServiceServer) FollowerList(context.Context, *FollowerListRequest) (*FollowerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowerList not implemented")
}
func (UnimplementedInteractionServiceServer) FriendList(context.Context, *FriendListRequest) (*FriendListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendList not implemented")
}
func (UnimplementedInteractionServiceServer) HasFollowed(context.Context, *HasFollowedRequest) (*HasFollowedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasFollowed not implemented")
}
func (UnimplementedInteractionServiceServer) mustEmbedUnimplementedInteractionServiceServer() {}

// UnsafeInteractionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InteractionServiceServer will
// result in compilation errors.
type UnsafeInteractionServiceServer interface {
	mustEmbedUnimplementedInteractionServiceServer()
}

func RegisterInteractionServiceServer(s grpc.ServiceRegistrar, srv InteractionServiceServer) {
	s.RegisterService(&InteractionService_ServiceDesc, srv)
}

func _InteractionService_Comment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).Comment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractionService_Comment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).Comment(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionService_CommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).CommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractionService_CommentList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).CommentList(ctx, req.(*CommentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionService_Relation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).Relation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractionService_Relation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).Relation(ctx, req.(*RelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionService_FollowList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).FollowList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractionService_FollowList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).FollowList(ctx, req.(*FollowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionService_FollowerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowerListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).FollowerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractionService_FollowerList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).FollowerList(ctx, req.(*FollowerListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionService_FriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).FriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractionService_FriendList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).FriendList(ctx, req.(*FriendListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionService_HasFollowed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasFollowedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).HasFollowed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractionService_HasFollowed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).HasFollowed(ctx, req.(*HasFollowedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InteractionService_ServiceDesc is the grpc.ServiceDesc for InteractionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InteractionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "interaction.InteractionService",
	HandlerType: (*InteractionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Comment",
			Handler:    _InteractionService_Comment_Handler,
		},
		{
			MethodName: "CommentList",
			Handler:    _InteractionService_CommentList_Handler,
		},
		{
			MethodName: "Relation",
			Handler:    _InteractionService_Relation_Handler,
		},
		{
			MethodName: "FollowList",
			Handler:    _InteractionService_FollowList_Handler,
		},
		{
			MethodName: "FollowerList",
			Handler:    _InteractionService_FollowerList_Handler,
		},
		{
			MethodName: "FriendList",
			Handler:    _InteractionService_FriendList_Handler,
		},
		{
			MethodName: "HasFollowed",
			Handler:    _InteractionService_HasFollowed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interaction.proto",
}
