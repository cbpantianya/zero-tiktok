package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	e "zero-tiktok/internal/error"
	"zero-tiktok/service/interaction/internal/model"
	"zero-tiktok/service/interaction/internal/svc"
	"zero-tiktok/service/interaction/pb/zero-tiktok/service/interaction"
)

type FollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerListLogic {
	return &FollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowerListLogic) FollowerList(req *interaction.FollowerListRequest) (*interaction.FollowerListResponse, error) {
	relations := make([]*model.Relation, 0)
	if err := l.svcCtx.DB.Where("user_id=?", req.UserId).Find(&relations).Error; err != nil {
		return nil, e.ErrDB
	}

	follower_ids := make([]int64, 0)
	for _, relation := range relations {
		follower_ids = append(follower_ids, relation.FollowerID)
	}

	return &interaction.FollowerListResponse{UserIdList: follower_ids}, nil
}
