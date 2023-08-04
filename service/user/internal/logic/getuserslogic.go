package logic

import (
	"context"
	"zero-tiktok/service/user/internal/model"

	"zero-tiktok/service/user/internal/svc"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersLogic {
	return &GetUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取多个用户信息
func (l *GetUsersLogic) GetUsers(in *user.GetUsersRequest) (*user.GetUsersResponse, error) {
	// todo: add your logic here and delete this line
	tx := l.svcCtx.DBList.Mysql.Begin()
	var users []*model.User

	// 构建一个空的切片用于存储查询条件
	ids := []int64{}
	ids = append(ids, in.UserIds...)

	// 查询符合条件的用户
	if err := tx.Model(&model.User{}).Where("user_id in (?)", ids).Find(&users).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	// 构建用户响应
	var userResponses []*user.User
	for _, usr := range users {
		userResponses = append(userResponses, &user.User{
			UserId:    int64(usr.ID),
			Name:      usr.Name,
			Signature: usr.Signature,
			Avatar:    usr.Avatar,
			Cover:     usr.Cover,
		})
	}

	return &user.GetUsersResponse{
		Users: userResponses,
	}, nil
}
