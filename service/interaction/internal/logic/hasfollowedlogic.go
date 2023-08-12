package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	e "zero-tiktok/internal/error"
	"zero-tiktok/service/interaction/internal/model"
	"zero-tiktok/service/interaction/internal/svc"
	"zero-tiktok/service/interaction/pb/zero-tiktok/service/interaction"
)

type HasFollowedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHasFollowedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HasFollowedLogic {
	return &HasFollowedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *HasFollowedLogic) HasFollowed(req *interaction.HasFollowedRequest) (*interaction.HasFollowedResponse, error) {

	relations := make([]*model.Relation, 0)

	if err := l.svcCtx.DB.Where("follower_id=?", req.UserId).Find(&relations).Error; err != nil {
		return nil, e.ErrDB
	}
	res := make([]bool, 0)
	flag := false
	for _, target_id := range req.TargetId {
		for _, relation := range relations {
			if relation.UserID == target_id {
				res = append(res, true)
				flag = true
				break
			}
		}
		if flag == false {
			res = append(res, false)
		} else {
			flag = !flag
		}
	}
	return &interaction.HasFollowedResponse{Result: res}, nil
}
