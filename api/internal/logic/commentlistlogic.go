package logic

import (
	"context"
	"strconv"
	"time"
	"zero-tiktok/api/internal/svc"
	"zero-tiktok/api/internal/types"
	e "zero-tiktok/internal/error"
	"zero-tiktok/service/interaction/pb/zero-tiktok/service/interaction"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentListLogic) CommentList(req *types.CommentListReq) (resp *types.CommentListResp, err error) {
	// 1. 获取评论列表
	list, err := l.svcCtx.Interaction.CommentList(l.ctx, &interaction.CommentListRequest{
		VideoId: req.VideoID,
	})
	if err != nil {
		l.Logger.Error(err)
		return
	}
	// 2. 构造响应结构体
	// 2.1 准备 _users map 方便后续构造
	_users := make(map[int64]types.Author, len(list.CommentList))
	for _, v := range list.CommentList {
		_resp, err := l.svcCtx.UserClient.GetUser(l.ctx, &user.GetUserRequest{
			UserId: v.UserId,
		})
		if err != nil {
			return nil, e.ErrInner
		}
		_user := _resp.User
		_users[v.UserId] = types.Author{
			ID:             _user.UserId,
			Name:           _user.Name,
			FollowCount:    _user.FollowCount,
			FollowerCount:  _user.FollowerCount,
			IsFollow:       false,
			Avatar:         _user.Avatar,
			Background:     _user.Cover,
			Signature:      _user.Signature,
			TotalFavorited: strconv.FormatInt(_user.TotalFavorited, 10),
			WorkCount:      -1,
			FavoriteCount:  _user.FavoriteCount,
		}
	}
	// 2.2 实际构造
	_resp := &types.CommentListResp{
		Code: 0,
		Msg:  "success",
	}
	for _, v := range list.CommentList {
		_resp.List = append(_resp.List, types.Comment{
			ID:         v.CommentId,
			User:       _users[v.UserId],
			Content:    v.Content,
			CreateDate: time.Unix(v.CreatedAt, 0).Format("01-02"),
		})
	}
	resp = _resp
	return
}
