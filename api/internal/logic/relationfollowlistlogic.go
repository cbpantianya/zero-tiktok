package logic

import (
	"context"
	"strconv"
	error2 "zero-tiktok/internal/error"
	"zero-tiktok/service/interaction/pb/zero-tiktok/service/interaction"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"

	"zero-tiktok/api/internal/svc"
	"zero-tiktok/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowListLogic {
	return &RelationFollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFollowListLogic) RelationFollowList(req *types.RelationFollowList) (resp *types.RelationFollowListResp, err error) {
	// 显示粉丝列表
	// Token -> ID
	id, err := l.svcCtx.UserClient.GetIdByToken(l.ctx, &user.TokenToUserRequest{
		Token: req.Token,
	})

	if err != nil {
		// 报错
		err = error2.ErrAuth
		return
	}

	// 根据ID获取粉丝列表
	ids, err := l.svcCtx.Interaction.FollowList(l.ctx, &interaction.FollowListRequest{
		UserId: id.UserId,
	})

	// 根据ids获取用户信息
	users, err := l.svcCtx.UserClient.GetUsers(l.ctx, &user.GetUsersRequest{
		UserIds: ids.UserIdList,
	})

	list := make([]types.Author, 0)

	// 是否已关注
	isF, err := l.svcCtx.Interaction.HasFollowed(l.ctx, &interaction.HasFollowedRequest{
		UserId:   id.UserId,
		TargetId: ids.UserIdList,
	})
	if err != nil {
		return nil, error2.ErrInner
	}

	for k, v := range users.Users {
		list = append(list, types.Author{
			ID:             v.UserId,
			Name:           v.Name,
			FollowCount:    v.FollowCount,
			FollowerCount:  v.FollowerCount,
			IsFollow:       isF.Result[k],
			Avatar:         v.Avatar,
			Background:     v.Cover,
			Signature:      v.Signature,
			TotalFavorited: strconv.FormatInt(v.TotalFavorited, 10),
			WorkCount:      v.VideoCount,
			FavoriteCount:  v.FavoriteCount,
		})
	}

	// 返回信息
	resp = &types.RelationFollowListResp{
		Code: 0,
		Msg:  "success",
		List: list,
	}
	return

}
