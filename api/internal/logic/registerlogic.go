package logic

import (
	"context"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"

	"zero-tiktok/api/internal/svc"
	"zero-tiktok/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.UserRegisterReq) (resp *types.UserTokenResp, err error) {
	// 参数筛选
	if req.Username == "" || req.Password == "" {
		return &types.UserTokenResp{
			Code: -40000,
			Msg:  "用户名或密码不能为空",
		}, nil
	}

	if len(req.Username) > 32 || len(req.Password) > 32 {
		return &types.UserTokenResp{
			Code: -40001,
			Msg:  "用户名长度应小于32位",
		}, nil
	}

	// 调用服务
	register, err := l.svcCtx.UserClient.Register(l.ctx, &user.LoginOrRegisterRequest{
		Name: req.Username,
		Pass: req.Password,
	})

	if err != nil {
		return &types.UserTokenResp{
			Code: -40002,
			Msg:  "用户名重复",
		}, nil
	}

	id, err := l.svcCtx.UserClient.GetIdByToken(l.ctx, &user.TokenToUserRequest{
		Token: register.Token,
	})

	// 返回
	resp = &types.UserTokenResp{
		Token:  register.Token,
		UserID: id.UserId,
		Code:   0,
		Msg:    "success",
	}

	return
}
