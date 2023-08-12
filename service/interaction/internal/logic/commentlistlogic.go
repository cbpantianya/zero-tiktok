package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"sort"
	e "zero-tiktok/internal/error"
	"zero-tiktok/service/interaction/internal/model"
	"zero-tiktok/service/interaction/internal/svc"
	"zero-tiktok/service/interaction/pb/zero-tiktok/service/interaction"
)

type CommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentListLogic) CommentList(req *interaction.CommentListRequest) (*interaction.CommentListResponse, error) {
	commentlist := make([]*model.Comment, 0)
	if err := l.svcCtx.DB.Where("video_id=?", req.VideoId).Find(&commentlist).Error; err != nil {
		return nil, e.ErrDB
	}
	//获取评论信息
	comment_ids := make([]int64, 0)
	for _, comment := range commentlist {
		comment_ids = append(comment_ids, comment.CommentID)
	}

	res_list := make([]*interaction.Comment, 0)
	for _, comment := range commentlist {
		res_list = append(res_list, &interaction.Comment{
			CommentId: comment.CommentID,
			UserId:    comment.UserID,
			Content:   comment.CommentText,
			CreatedAt: comment.CreatedAt.Unix(),
		})
	}

	//按发布时间倒序
	sort.Slice(res_list, func(i, j int) bool {
		return res_list[i].CreatedAt > res_list[j].CreatedAt
	})

	return &interaction.CommentListResponse{CommentList: res_list}, nil
}
