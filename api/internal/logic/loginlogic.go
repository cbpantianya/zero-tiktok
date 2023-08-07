package logic

import (
	"context"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"

	"zero-tiktok/api/internal/svc"
	"zero-tiktok/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.UserLoginReq) (resp *types.UserTokenResp, err error) {
	login, err := l.svcCtx.UserClient.Login(l.ctx, &user.LoginOrRegisterRequest{
		Name: req.Username,
		Pass: req.Password,
	})
	if err != nil {
		return &types.UserTokenResp{
			Code: -40300,
			Msg:  "用户名或密码错误",
		}, nil
	}

	id, err := l.svcCtx.UserClient.GetIdByToken(l.ctx, &user.TokenToUserRequest{
		Token: login.Token,
	})
	if err != nil {
		return &types.UserTokenResp{
			Code: -50000,
			Msg:  "内部服务错误",
		}, nil
	}

	resp = &types.UserTokenResp{
		Token:  login.Token,
		UserID: id.UserId,
		Code:   0,
		Msg:    "success",
	}

	return resp, nil
}
