package logic

import (
	"context"
	"zero-tiktok/service/user/internal/utils"

	"zero-tiktok/service/user/internal/svc"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIdByTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIdByTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIdByTokenLogic {
	return &GetIdByTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 识别用户（token转id）
func (l *GetIdByTokenLogic) GetIdByToken(in *user.TokenToUserRequest) (*user.TokenToUserResponse, error) {
	// todo: add your logic here and delete this line
	token := in.Token
	uid := utils.GetUserIdByToken(token)
	return &user.TokenToUserResponse{
		UserId: int64(uid),
	}, nil
}
