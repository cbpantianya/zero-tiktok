package logic

import (
	"context"
	"strconv"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"
	"zero-tiktok/service/video/pb/zero-tiktok/service/video"

	"zero-tiktok/api/internal/svc"
	"zero-tiktok/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishListLogic) PublishList(req *types.UserInfoReq) (resp *types.PublishListResp, err error) {
	// Token to id
	_, err = l.svcCtx.UserClient.GetIdByToken(l.ctx, &user.TokenToUserRequest{
		Token: req.Token,
	})

	if err != nil {
		l.Logger.Error(err)
		return
	}

	list, err := l.svcCtx.VideoClient.Publish(l.ctx, &video.PublishRequest{
		UserId:   req.UserID,
		TargetId: req.UserID,
	})

	if err != nil {
		l.Logger.Error(err)
		return
	}

	authorInfo, err := l.svcCtx.UserClient.GetUser(l.ctx, &user.GetUserRequest{
		UserId: req.UserID,
	})

	if err != nil {
		l.Logger.Error(err)
		return
	}

	author := types.Author{
		ID:             authorInfo.User.UserId,
		Name:           authorInfo.User.Name,
		FollowCount:    authorInfo.User.FollowCount,
		FollowerCount:  authorInfo.User.FollowerCount,
		IsFollow:       false,
		Avatar:         authorInfo.User.Avatar,
		Background:     authorInfo.User.Cover,
		Signature:      authorInfo.User.Signature,
		TotalFavorited: strconv.FormatInt(authorInfo.User.TotalFavorited, 10),
		WorkCount:      -1,
		FavoriteCount:  authorInfo.User.FavoriteCount,
	}

	resp = &types.PublishListResp{
		Code: 0,
		Msg:  "success",
		List: make([]types.Video, 0),
	}

	for _, v := range list.List {
		resp.List = append(resp.List, types.Video{
			ID:             v.VideoId,
			Author:         author,
			PlayUrl:        v.Play,
			CoverUrl:       v.Cover,
			FavouriteCount: v.FavoriteCount,
			CommentCount:   v.CommentCount,
			IsFavourite:    false,
			Title:          v.Title,
		})
	}

	return
}
