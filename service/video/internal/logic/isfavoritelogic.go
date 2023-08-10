package logic

import (
	"context"
	"zero-tiktok/service/video/internal/model"

	"zero-tiktok/service/video/internal/svc"
	"zero-tiktok/service/video/pb/zero-tiktok/service/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFavoriteLogic {
	return &IsFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsFavoriteLogic) IsFavorite(in *video.IsFavoriteReq) (*video.IsFavoriteResp, error) {
	var f []model.Favorite
	f = make([]model.Favorite, 0)
	err := l.svcCtx.DB.Where("user_id = ?", in.UserId).Find(&f).Error
	if err != nil {
		return nil, err
	}
	resp := &video.IsFavoriteResp{
		IsFavorite: make([]bool, 0),
	}
	for _, v := range in.VideoId {
		flag := false
		for _, fv := range f {
			if fv.VideoID == v {
				resp.IsFavorite = append(resp.IsFavorite, true)
				flag = true
				break
			}
		}
		if !flag {
			resp.IsFavorite = append(resp.IsFavorite, false)
		}
	}

	return resp, nil
}
