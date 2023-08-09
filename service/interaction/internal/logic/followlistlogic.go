package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	e "zero-tiktok/internal/error"
	"zero-tiktok/service/interaction/internal/model"
	"zero-tiktok/service/interaction/internal/svc"
	"zero-tiktok/service/interaction/pb/zero-tiktok/service/interaction"
)

type FollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowListLogic {
	return &FollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowListLogic) FollowList(req *interaction.FollowListRequest) (*interaction.FollowListResponse, error) {
	relations := make([]*model.Relation, 0)
	if err := l.svcCtx.DB.Where("follower_id=?", req.UserId).Find(&relations).Error; err != nil {
		return nil, e.ErrDB
	}

	follow_ids := make([]int64, 0)
	for _, relation := range relations {
		follow_ids = append(follow_ids, relation.UserID)
	}

	return &interaction.FollowListResponse{UserIdList: follow_ids}, nil
}
