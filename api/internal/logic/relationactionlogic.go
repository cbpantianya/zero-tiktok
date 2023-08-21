package logic

import (
	"context"
	error2 "zero-tiktok/internal/error"
	"zero-tiktok/service/interaction/pb/zero-tiktok/service/interaction"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"

	"zero-tiktok/api/internal/svc"
	"zero-tiktok/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationActionLogic {
	return &RelationActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationActionLogic) RelationAction(req *types.RelationAction) (resp *types.RelationActionResp, err error) {
	// Token -> ID
	id, err := l.svcCtx.UserClient.GetIdByToken(l.ctx, &user.TokenToUserRequest{
		Token: req.Token,
	})
	if err != nil {
		// 报错
		err = error2.ErrAuth
		return
	}

	// action
	_, err = l.svcCtx.Interaction.Relation(l.ctx, &interaction.RelationRequest{
		UserId:     id.UserId,
		TargetId:   req.UserID,
		ActionType: int32(req.ActionType),
	})

	if err != nil {
		l.Logger.Error(err)
		return
	}

	resp = &types.RelationActionResp{
		Code: 0,
		Msg:  "success",
	}

	return
}
