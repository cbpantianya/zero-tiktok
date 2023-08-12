package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
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
	var commentId int64
	if req.ActionType == 1 {
		comment := model.Comment{
			VideoID:     req.VideoId,
			UserID:      req.UserId,
			CommentText: *req.CommentText,
		}
		if err := l.svcCtx.DB.Create(&comment).Error; err != nil {
			return nil, e.ErrDB
		}
		fmt.Println(comment)
		commentId = comment.CommentID
	} else if req.ActionType == 2 { //删除评论
		commentId = *req.CommentId
		if err := l.svcCtx.DB.Where("comment_id=?", req.CommentId).Delete(&model.Comment{}).Error; err != nil {
			return nil, e.ErrDB
		}
	}
	return &interaction.CommentResponse{
		Comment: &interaction.Comment{
			CommentId: commentId,
			UserId:    req.UserId,
			Content:   *req.CommentText,
			CreatedAt: time.Now().Unix(),
		},
	}, nil
}
