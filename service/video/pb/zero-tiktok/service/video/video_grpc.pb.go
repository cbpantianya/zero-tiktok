// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: video.proto

package video

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

// VideoServiceClient is the client API for VideoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoServiceClient interface {
	Feed(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*VideoListResp, error)
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*VideoListResp, error)
	Favorite(ctx context.Context, in *FavoriteAction, opts ...grpc.CallOption) (*FavoriteActionResp, error)
	FavoriteList(ctx context.Context, in *FavoriteRequest, opts ...grpc.CallOption) (*VideoListResp, error)
}

type videoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoServiceClient(cc grpc.ClientConnInterface) VideoServiceClient {
	return &videoServiceClient{cc}
}

func (c *videoServiceClient) Feed(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*VideoListResp, error) {
	out := new(VideoListResp)
	err := c.cc.Invoke(ctx, "/user.VideoService/Feed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*VideoListResp, error) {
	out := new(VideoListResp)
	err := c.cc.Invoke(ctx, "/user.VideoService/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) Favorite(ctx context.Context, in *FavoriteAction, opts ...grpc.CallOption) (*FavoriteActionResp, error) {
	out := new(FavoriteActionResp)
	err := c.cc.Invoke(ctx, "/user.VideoService/Favorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) FavoriteList(ctx context.Context, in *FavoriteRequest, opts ...grpc.CallOption) (*VideoListResp, error) {
	out := new(VideoListResp)
	err := c.cc.Invoke(ctx, "/user.VideoService/FavoriteList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServiceServer is the server API for VideoService service.
// All implementations must embed UnimplementedVideoServiceServer
// for forward compatibility
type VideoServiceServer interface {
	Feed(context.Context, *FeedRequest) (*VideoListResp, error)
	Publish(context.Context, *PublishRequest) (*VideoListResp, error)
	Favorite(context.Context, *FavoriteAction) (*FavoriteActionResp, error)
	FavoriteList(context.Context, *FavoriteRequest) (*VideoListResp, error)
	mustEmbedUnimplementedVideoServiceServer()
}

// UnimplementedVideoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVideoServiceServer struct {
}

func (UnimplementedVideoServiceServer) Feed(context.Context, *FeedRequest) (*VideoListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Feed not implemented")
}
func (UnimplementedVideoServiceServer) Publish(context.Context, *PublishRequest) (*VideoListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedVideoServiceServer) Favorite(context.Context, *FavoriteAction) (*FavoriteActionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Favorite not implemented")
}
func (UnimplementedVideoServiceServer) FavoriteList(context.Context, *FavoriteRequest) (*VideoListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteList not implemented")
}
func (UnimplementedVideoServiceServer) mustEmbedUnimplementedVideoServiceServer() {}

// UnsafeVideoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServiceServer will
// result in compilation errors.
type UnsafeVideoServiceServer interface {
	mustEmbedUnimplementedVideoServiceServer()
}

func RegisterVideoServiceServer(s grpc.ServiceRegistrar, srv VideoServiceServer) {
	s.RegisterService(&VideoService_ServiceDesc, srv)
}

func _VideoService_Feed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).Feed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.VideoService/Feed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).Feed(ctx, req.(*FeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.VideoService/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_Favorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteAction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).Favorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.VideoService/Favorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).Favorite(ctx, req.(*FavoriteAction))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_FavoriteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).FavoriteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.VideoService/FavoriteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).FavoriteList(ctx, req.(*FavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoService_ServiceDesc is the grpc.ServiceDesc for VideoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.VideoService",
	HandlerType: (*VideoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Feed",
			Handler:    _VideoService_Feed_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _VideoService_Publish_Handler,
		},
		{
			MethodName: "Favorite",
			Handler:    _VideoService_Favorite_Handler,
		},
		{
			MethodName: "FavoriteList",
			Handler:    _VideoService_FavoriteList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "video.proto",
}
