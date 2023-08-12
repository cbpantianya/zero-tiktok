package logic

import (
	"context"

	"zero-tiktok/api/internal/svc"
	"zero-tiktok/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVideoLogic {
	return &UploadVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadVideoLogic) UploadVideo(req *types.UploadVideo) (resp *types.UploadVideoResp, err error) {
	// 此处不应该填写任何逻辑，请前往handle
	resp = &types.UploadVideoResp{
		Msg:  "ok",
		Code: 0,
	}
	return
}
