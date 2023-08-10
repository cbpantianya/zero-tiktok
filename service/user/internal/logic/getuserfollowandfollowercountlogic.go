package logic

import (
	"context"
	myerror "zero-tiktok/internal/error"
	"zero-tiktok/service/user/internal/model"

	"zero-tiktok/service/user/internal/svc"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFollowAndFollowerCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFollowAndFollowerCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFollowAndFollowerCountLogic {
	return &GetUserFollowAndFollowerCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获得用户关注和粉丝数量
func (l *GetUserFollowAndFollowerCountLogic) GetUserFollowAndFollowerCount(in *user.GetUserFollowAndFollowerCountRequest) (*user.GetUserFollowAndFollowerCountResponse, error) {
	// todo: add your logic here and delete this line
	tx := l.svcCtx.DBList.Mysql
	var followerCount int64
	var followCount int64
	if err := tx.Model(&model.Relation{}).Where("user_id = ?", in.UserId).Count(&followerCount).Error; err != nil {

		return nil, myerror.ErrDB
	}

	if err := tx.Model(&model.Relation{}).Where("follower_id = ?", in.UserId).Count(&followCount).Error; err != nil {

		return nil, myerror.ErrDB
	}
	return &user.GetUserFollowAndFollowerCountResponse{
		FollowCount:   followCount,
		FollowerCount: followerCount,
	}, nil
}
