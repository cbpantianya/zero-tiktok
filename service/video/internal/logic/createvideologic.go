package logic

import (
	"context"
	"time"
	"zero-tiktok/service/video/internal/model"

	"zero-tiktok/service/video/internal/svc"
	"zero-tiktok/service/video/pb/zero-tiktok/service/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateVideoLogic {
	return &CreateVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateVideoLogic) CreateVideo(in *video.CreateVideoReq) (*video.CreateVideoResp, error) {
	vi := model.Video{
		Title:     in.Title,
		PublishAt: time.Now(),
		AuthorID:  in.UserId,
		Play:      in.Play,
		Cover:     "https://zero-tiktok.oss-cn-hongkong.aliyuncs.com/cover/video-cover-01.jpg",
	}

	err := l.svcCtx.DB.Create(&vi).Error
	if err != nil {
		return nil, err
	}

	return &video.CreateVideoResp{
		VideoId: vi.VideoID,
	}, nil
}
