package logic

import (
	"context"
	"strconv"
	"time"
	"zero-tiktok/api/internal/svc"
	"zero-tiktok/api/internal/types"
	"zero-tiktok/service/interaction/pb/zero-tiktok/service/interaction"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"
	"zero-tiktok/service/video/pb/zero-tiktok/service/video"

	e "zero-tiktok/internal/error"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedReq) (resp *types.FeedResp, err error) {

	var id *int64

	if req.Token != "" {
		idr, err := l.svcCtx.UserClient.GetIdByToken(l.ctx, &user.TokenToUserRequest{
			Token: req.Token,
		})
		if err != nil {
			// 忽略，以未登录用户处理
			l.Logger.Error(err)
		}

		id = &idr.UserId
	}

	latest := time.Now().Unix()
	if req.Latest != "" {
		t, err := strconv.ParseInt(req.Latest, 10, 64)

		if err != nil {
			l.Logger.Error(err)
			return nil, e.ErrInner
		}
		// 减小精度
		latest = t / 1000
	}

	// 向video rpc发起请求
	list, err := l.svcCtx.VideoClient.Feed(l.ctx, &video.FeedRequest{
		UserId: id,
		Latest: latest,
	})

	if err != nil {
		l.Logger.Error(err)
		return nil, e.ErrInner
	}

	// next_time
	// author list
	var next int64
	authorsID := map[int64]types.Author{}
	next = time.Now().Unix()
	for _, v := range list.List {
		if v.PublishAt < next {
			next = v.PublishAt
		}
		authorsID[v.AuthorId] = types.Author{}
	}

	// 使用批量接口获取数据
	userIDs := []int64{}
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

	// 查询是否关注
	if id != nil {
		isF, err := l.svcCtx.Interaction.HasFollowed(l.ctx, &interaction.HasFollowedRequest{
			UserId:   *id,
			TargetId: userIDs,
		})
		if err != nil {
			return nil, e.ErrInner
		}

		// range 列表
		for k, v := range userIDs {
			authorsID[v] = types.Author{
				ID:             authorsID[v].ID,
				Name:           authorsID[v].Name,
				FollowCount:    authorsID[v].FollowCount,
				FollowerCount:  authorsID[v].FollowerCount,
				IsFollow:       isF.Result[k],
				Avatar:         authorsID[v].Avatar,
				Background:     authorsID[v].Background,
				Signature:      authorsID[v].Signature,
				TotalFavorited: authorsID[v].TotalFavorited,
				WorkCount:      authorsID[v].WorkCount,
				FavoriteCount:  authorsID[v].FavoriteCount,
			}
		}

	}

	resp = &types.FeedResp{}

	resp.Code = 0
	resp.Msg = "success"
	resp.Next = next
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

	return resp, nil
}
