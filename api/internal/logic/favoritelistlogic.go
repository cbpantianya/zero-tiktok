package logic

import (
	"context"
	"strconv"
	e "zero-tiktok/internal/error"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"
	"zero-tiktok/service/video/pb/zero-tiktok/service/video"

	"zero-tiktok/api/internal/svc"
	"zero-tiktok/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListLogic) FavoriteList(req *types.UserInfoReq) (resp *types.PublishListResp, err error) {
	// Token to id
	id, err := l.svcCtx.UserClient.GetIdByToken(l.ctx, &user.TokenToUserRequest{
		Token: req.Token,
	})

	if err != nil {
		l.Logger.Error(err)
		return
	}

	list, err := l.svcCtx.VideoClient.FavoriteList(l.ctx, &video.FavoriteRequest{
		UserId:   id.UserId,
		TargetId: req.UserID,
	})

	if err != nil {
		l.Logger.Error(err)
		return
	}

	authorsID := map[int64]types.Author{}
	for _, v := range list.List {
		authorsID[v.AuthorId] = types.Author{}
	}

	// 使用批量接口获取数据
	var userIDs []int64
	for k := range authorsID {
		userIDs = append(userIDs, k)
	}

	users, err := l.svcCtx.UserClient.GetUsers(l.ctx, &user.GetUsersRequest{
		UserIds: userIDs,
	})

	if err != nil {
		return nil, e.ErrInner
	}

	for _, v := range users.Users {
		authorsID[v.UserId] = types.Author{
			ID:             v.UserId,
			Name:           v.Name,
			FollowCount:    v.FollowCount,
			FollowerCount:  v.FollowerCount,
			IsFollow:       false,
			Avatar:         v.Avatar,
			Background:     v.Cover,
			Signature:      v.Signature,
			TotalFavorited: strconv.FormatInt(v.TotalFavorited, 10),
			WorkCount:      -1, // TODO: Users
			FavoriteCount:  v.FavoriteCount,
		}
	}

	resp = &types.PublishListResp{}

	resp.Code = 0
	resp.Msg = "success"
	resp.List = []types.Video{}
	for _, v := range list.List {
		resp.List = append(resp.List, types.Video{
			ID:             v.VideoId,
			Author:         authorsID[v.AuthorId],
			PlayUrl:        v.Play,
			CoverUrl:       v.Cover,
			FavouriteCount: v.FavoriteCount,
			CommentCount:   v.CommentCount,
			IsFavourite:    v.IsFavorite,
			Title:          v.Title,
		})
	}

	return
}
