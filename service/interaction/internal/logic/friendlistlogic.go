package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	e "zero-tiktok/internal/error"
	"zero-tiktok/service/interaction/internal/model"
	"zero-tiktok/service/interaction/internal/svc"
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
	if err := l.svcCtx.DB.Where("user_id=?", req.UserId).Find(&relations).Error; err != nil {
		return nil, e.ErrDB
	}
	friend_ids := make([]int64, 0)
	for _, relation := range relations {
		tmpRelation := &model.Relation{}
		if err := l.svcCtx.DB.Where("user_id=? and follower_id=?", relation.FollowerID, relation.UserID).Find(tmpRelation).Error; err != nil {
			continue
		} else {
			friend_ids = append(friend_ids, relation.FollowerID)
		}
	}
	return &interaction.FriendListResponse{UserIdList: friend_ids}, nil
}
