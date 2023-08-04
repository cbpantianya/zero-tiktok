package logic

import (
	"context"
	"zero-tiktok/service/user/internal/model"

	"zero-tiktok/service/user/internal/svc"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取单个用户信息
func (l *GetUserLogic) GetUser(in *user.GetUserRequest) (*user.GetUserResponse, error) {
	// todo: add your logic here and delete this line
	tx := l.svcCtx.DBList.Mysql.Begin()
	uid := in.UserId
	var usr model.User
	if err := tx.Model(&model.User{}).Where("user_id=?", uid).Find(&usr).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &user.GetUserResponse{
		User: &user.User{
			UserId:    uid,
			Name:      usr.Name,
			Signature: usr.Signature,
			Avatar:    usr.Avatar,
			Cover:     usr.Cover,
		},
	}, nil
}
