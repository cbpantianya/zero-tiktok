package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	e "zero-tiktok/internal/error"
	"zero-tiktok/service/interaction/internal/model"
	"zero-tiktok/service/interaction/internal/svc"
	"zero-tiktok/service/interaction/pb/zero-tiktok/service/interaction"
)

type CommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentLogic {
	return &CommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentLogic) Comment(req *interaction.CommentRequest) (*interaction.CommentResponse, error) {
	//发布评论
	if req.ActionType == 1 {
		comment := &model.Comment{
			VideoID:     req.VideoId,
			UserID:      req.UserId,
			CommentText: *req.CommentText,
			// CommentID:   *req.CommentId,
			//CreatedAt:
		}
		if err := l.svcCtx.DB.Create(comment).Error; err != nil {
			return nil, e.ErrDB
		}
	} else if req.ActionType == 2 { //删除评论
		if err := l.svcCtx.DB.Where("comment_id=?", req.CommentId).Delete(&model.Comment{}).Error; err != nil {
			return nil, e.ErrDB
		}
	}
	return &interaction.CommentResponse{}, nil
}
