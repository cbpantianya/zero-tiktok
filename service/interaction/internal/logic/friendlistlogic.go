package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	e "zero-tiktok/internal/error"
	"zero-tiktok/service/interaction/internal/model"
	"zero-tiktok/service/interaction/internal/svc"
	"zero-tiktok/service/interaction/internal/utils"
	"zero-tiktok/service/interaction/pb/zero-tiktok/service/interaction"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *FriendListLogic) FriendList(req *interaction.FriendListRequest) (*interaction.FriendListResponse, error) {
	//这里认为互关的两个人为好友

	relations := make([]*model.Relation, 0)
	if err := l.svcCtx.DB.Where("user_id=? or follower_id=?", req.UserId, req.UserId).Find(&relations).Error; err != nil {
		return nil, e.ErrDB
	}
	friend_ids := make([]int64, 0)

	for _, relation := range relations {
		for _, subrelation := range relations {
			if relation.FollowerID == subrelation.UserID && relation.UserID == subrelation.FollowerID {
				friend_ids = append(friend_ids, relation.FollowerID)
			} else {
				continue
			}
		}
	}
	friend_ids = utils.DeleteSlice(friend_ids, req.UserId)
	return &interaction.FriendListResponse{UserIdList: friend_ids}, nil
}
