// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userservice

import (
	"context"

	"zero-tiktok/service/user/pb/zero-tiktok/service/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetUserFollowAndFollowerCountRequest  = user.GetUserFollowAndFollowerCountRequest
	GetUserFollowAndFollowerCountResponse = user.GetUserFollowAndFollowerCountResponse
	GetUserRequest                        = user.GetUserRequest
	GetUserResponse                       = user.GetUserResponse
	GetUserTotalFavoritedRequest          = user.GetUserTotalFavoritedRequest
	GetUserTotalFavoritedResponse         = user.GetUserTotalFavoritedResponse
	GetUserVideoCountRequest              = user.GetUserVideoCountRequest
	GetUserVideoCountResponse             = user.GetUserVideoCountResponse
	GetUsersRequest                       = user.GetUsersRequest
	GetUsersResponse                      = user.GetUsersResponse
	GetVideoFavoritedRequest              = user.GetVideoFavoritedRequest
	GetVideolFavoritedResponse            = user.GetVideolFavoritedResponse
	LoginOrRegisterRequest                = user.LoginOrRegisterRequest
	LoginOrRegisterResponse               = user.LoginOrRegisterResponse
	TokenToUserRequest                    = user.TokenToUserRequest
	TokenToUserResponse                   = user.TokenToUserResponse
	User                                  = user.User

	UserService interface {
		// 登录
		Login(ctx context.Context, in *LoginOrRegisterRequest, opts ...grpc.CallOption) (*LoginOrRegisterResponse, error)
		// 注册
		Register(ctx context.Context, in *LoginOrRegisterRequest, opts ...grpc.CallOption) (*LoginOrRegisterResponse, error)
		// 获取单个用户信息
		GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
		// 获取多个用户信息
		GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
		// 识别用户（token转id）
		GetIdByToken(ctx context.Context, in *TokenToUserRequest, opts ...grpc.CallOption) (*TokenToUserResponse, error)
		// 获得用户关注和粉丝数量
		GetUserFollowAndFollowerCount(ctx context.Context, in *GetUserFollowAndFollowerCountRequest, opts ...grpc.CallOption) (*GetUserFollowAndFollowerCountResponse, error)
		// 获得用户发布的视频总数
		GetUserVideoCount(ctx context.Context, in *GetUserVideoCountRequest, opts ...grpc.CallOption) (*GetUserVideoCountResponse, error)
		// 用户的获赞总数
		GetUserTotalFavorited(ctx context.Context, in *GetUserTotalFavoritedRequest, opts ...grpc.CallOption) (*GetUserTotalFavoritedResponse, error)
		// 单个视频的点赞量
		GetVideoFavorited(ctx context.Context, in *GetVideoFavoritedRequest, opts ...grpc.CallOption) (*GetVideolFavoritedResponse, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

// 登录
func (m *defaultUserService) Login(ctx context.Context, in *LoginOrRegisterRequest, opts ...grpc.CallOption) (*LoginOrRegisterResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

// 注册
func (m *defaultUserService) Register(ctx context.Context, in *LoginOrRegisterRequest, opts ...grpc.CallOption) (*LoginOrRegisterResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

// 获取单个用户信息
func (m *defaultUserService) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

// 获取多个用户信息
func (m *defaultUserService) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.GetUsers(ctx, in, opts...)
}

// 识别用户（token转id）
func (m *defaultUserService) GetIdByToken(ctx context.Context, in *TokenToUserRequest, opts ...grpc.CallOption) (*TokenToUserResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.GetIdByToken(ctx, in, opts...)
}

// 获得用户关注和粉丝数量
func (m *defaultUserService) GetUserFollowAndFollowerCount(ctx context.Context, in *GetUserFollowAndFollowerCountRequest, opts ...grpc.CallOption) (*GetUserFollowAndFollowerCountResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.GetUserFollowAndFollowerCount(ctx, in, opts...)
}

// 获得用户发布的视频总数
func (m *defaultUserService) GetUserVideoCount(ctx context.Context, in *GetUserVideoCountRequest, opts ...grpc.CallOption) (*GetUserVideoCountResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.GetUserVideoCount(ctx, in, opts...)
}

// 用户的获赞总数
func (m *defaultUserService) GetUserTotalFavorited(ctx context.Context, in *GetUserTotalFavoritedRequest, opts ...grpc.CallOption) (*GetUserTotalFavoritedResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.GetUserTotalFavorited(ctx, in, opts...)
}

// 单个视频的点赞量
func (m *defaultUserService) GetVideoFavorited(ctx context.Context, in *GetVideoFavoritedRequest, opts ...grpc.CallOption) (*GetVideolFavoritedResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.GetVideoFavorited(ctx, in, opts...)
}
