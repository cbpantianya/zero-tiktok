package logic

import (
	"context"
	"strconv"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"

	"zero-tiktok/api/internal/svc"
	"zero-tiktok/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// Token转ID
	_, err = l.svcCtx.UserClient.GetIdByToken(l.ctx, &user.TokenToUserRequest{
		Token: req.Token,
	})
	if err != nil {
		return
	}
	// 获取用户信息
	userInfo, err := l.svcCtx.UserClient.GetUser(l.ctx, &user.GetUserRequest{
		UserId: req.UserID,
	})
	// 返回
	resp = &types.UserInfoResp{
		Code: 0,
		Msg:  "success",
		User: types.Author{
			ID:             userInfo.User.UserId,
			Name:           userInfo.User.Name,
			FollowCount:    userInfo.User.FollowCount,
			FollowerCount:  userInfo.User.FollowerCount,
			IsFollow:       false,
			Avatar:         userInfo.User.Avatar,
			Background:     userInfo.User.Cover,
			Signature:      userInfo.User.Signature,
			TotalFavorited: strconv.FormatInt(userInfo.User.TotalFavorited, 10),
			WorkCount:      userInfo.User.VideoCount,
			FavoriteCount:  userInfo.User.FavoriteCount,
		},
	}

	return
}
