package logic

import (
	"context"
	"strconv"
	"time"

	"zero-tiktok/api/internal/svc"
	"zero-tiktok/api/internal/types"
	"zero-tiktok/service/interaction/pb/zero-tiktok/service/interaction"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentActionLogic) CommentAction(req *types.CommentActionReq) (resp *types.CommentActionResp, err error) {
	// 0 初始化rpc请求结构体
	_req := interaction.CommentRequest{}
	// 1. 鉴权
	// 1.1 未登录用户不能评论
	if req.Token == "" {
		return &types.CommentActionResp{
			Code: -1,
			Msg:  "未登录用户无法评论",
		}, nil
	} else {
		// 1.2 已登录用户查询id赋值
		_resp, err := l.svcCtx.UserClient.GetIdByToken(l.ctx, &user.TokenToUserRequest{Token: req.Token})
		if err != nil {
			l.Logger.Error(err)
			return nil, err
		}
		_req.UserId = _resp.UserId
	}
	// 2. 调用rpc服务
	_req.VideoId = req.VideoID
	_req.ActionType = int32(req.ActionType)
	_req.CommentText = &req.CommentText
	_req.CommentId = &req.CommentID
	// 2.1 评论操作
	_resp_comment, err := l.svcCtx.Interaction.Comment(l.ctx, &_req)
	if err != nil {
		l.Logger.Error(err)
		return
	}
	_comment := _resp_comment.Comment
	// 2.2 获取评论用户信息
	_user_resp, err := l.svcCtx.UserClient.GetUser(l.ctx, &user.GetUserRequest{
		UserId: _req.UserId,
	})
	_user := types.Author{
		ID:             _user_resp.User.UserId,
		Name:           _user_resp.User.Name,
		FollowCount:    _user_resp.User.FollowCount,
		FollowerCount:  _user_resp.User.FollowerCount,
		IsFollow:       false,
		Avatar:         _user_resp.User.Avatar,
		Background:     _user_resp.User.Cover,
		Signature:      _user_resp.User.Signature,
		TotalFavorited: strconv.FormatInt(_user_resp.User.TotalFavorited, 10),
		WorkCount:      -1,
		FavoriteCount:  _user_resp.User.FavoriteCount,
	}
	// 3 构造结构体返回
	resp.Code = 0
	resp.Msg = "success"
	resp.Comment = types.Comment{
		ID:         _comment.CommentId,
		User:       _user,
		Content:    _comment.Content,
		CreateDate: time.Unix(_comment.CreatedAt, 0).Format("01-02"),
	}
	return
}
