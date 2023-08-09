package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	e "zero-tiktok/internal/error"
	"zero-tiktok/service/interaction/internal/model"
	"zero-tiktok/service/interaction/internal/svc"
	"zero-tiktok/service/interaction/pb/zero-tiktok/service/interaction"
)

type RelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationLogic {
	return &RelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *RelationLogic) Relation(req *interaction.RelationRequest) (*interaction.RelationResponse, error) {
	var count int64
	if req.ActionType == 1 { //关注
		if err := l.svcCtx.DB.Where("user_id=? AND follower_id=?", req.TargetId, req.UserId).Find(&model.Relation{}).Count(&count).Error; err != nil {
			return nil, e.ErrDB
		} else {
			if count == 0 {
				if err := l.svcCtx.DB.Create(&model.Relation{
					UserID:     req.TargetId,
					FollowerID: req.UserId,
				}).Error; err != nil {
					return nil, e.ErrDB
				}
			} else {
				return nil, e.ErrDB
			}

		}
	} else if req.ActionType == 2 { //取关
		if err := l.svcCtx.DB.Where("user_id=? and follower_id=?", req.TargetId, req.UserId).Delete(&model.Relation{}).Error; err != nil {
			return nil, e.ErrDB
		}
	}

	return &interaction.RelationResponse{}, nil
}
