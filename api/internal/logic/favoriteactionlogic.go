package logic

import (
	"context"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"
	"zero-tiktok/service/video/pb/zero-tiktok/service/video"

	"zero-tiktok/api/internal/svc"
	"zero-tiktok/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteActionLogic) FavoriteAction(req *types.FavoriteActionReq) (resp *types.FavoriteActionResp, err error) {
	id, err := l.svcCtx.UserClient.GetIdByToken(l.ctx, &user.TokenToUserRequest{Token: req.Token})
	if err != nil {
		l.Logger.Error(err)
		return
	}

	_, err = l.svcCtx.VideoClient.Favorite(l.ctx, &video.FavoriteAction{
		UserId:  id.UserId,
		VideoId: req.VideoID,
		Action:  req.ActionType == 1,
	})

	if err != nil {
		l.Logger.Error(err)
		resp = &types.FavoriteActionResp{
			Code: -1,
			Msg:  "您已经点过赞了",
		}
		return resp, nil
	}

	resp = &types.FavoriteActionResp{
		Code: 0,
		Msg:  "success",
	}

	return
}
