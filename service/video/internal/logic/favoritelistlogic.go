package logic

import (
	"context"

	"zero-tiktok/service/video/internal/svc"
	"zero-tiktok/service/video/pb/zero-tiktok/service/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteListLogic) FavoriteList(in *video.FavoriteRequest) (*video.VideoListResp, error) {
	// todo: add your logic here and delete this line

	return &video.VideoListResp{}, nil
}
