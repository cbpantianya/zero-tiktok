package logic

import (
	"context"
	"zero-tiktok/service/video/internal/model"

	"zero-tiktok/service/video/internal/svc"
	"zero-tiktok/service/video/pb/zero-tiktok/service/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteLogic {
	return &FavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteLogic) Favorite(in *video.FavoriteAction) (*video.FavoriteActionResp, error) {
	// 喜欢操作
	if in.Action == true {
		// 点赞
		var favorite = &model.Favorite{
			UserID:  in.UserId,
			VideoID: in.VideoId,
		}

		err := l.svcCtx.DB.Create(favorite).Error
		if err != nil {
			return nil, err
		}
	} else {
		// 删除点赞
		err := l.svcCtx.DB.Where("user_id = ? and video_id = ?", in.UserId, in.VideoId).Delete(&model.Favorite{}).Error
		if err != nil {
			return nil, err
		}
	}
	return &video.FavoriteActionResp{}, nil
}
