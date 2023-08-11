package logic

import (
	"context"

	"zero-tiktok/api/internal/svc"
	"zero-tiktok/api/internal/types"
	"zero-tiktok/service/interaction/interactionservice"
	"zero-tiktok/service/interaction/pb/zero-tiktok/service/interaction"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"

	"github.com/spf13/cast"
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
	// // 1. 构造请求rpc结构体
	// // 1.1 Token to User ID
	// id, err := l.svcCtx.UserClient.GetIdByToken(l.ctx, &user.TokenToUserRequest{
	// 	Token: req.Token,
	// })
	// if err != nil {
	// 	l.Logger.Error(err)
	// 	return
	// }
	// 1. 获取评论列表
	list, err := l.svcCtx.Interaction.CommentList(l.ctx, &interaction.CommentListRequest{
		VideoId: req.VideoID,
	})

	if err != nil {
		l.Logger.Error(err)
		return
	}
	// 2. 构造响应结构体
	resp.Code = 0
	resp.Msg = "success"
	resp.List = []types.Comment{}
	for _, v := range list.CommentList {
		resp.List = append(resp.List, types.Comment{
			ID:   v.CommentId,
			User: v.UserId,
		})
	}

	return
}
