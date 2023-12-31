package logic

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	myerror "zero-tiktok/internal/error"
	"zero-tiktok/service/user/internal/model"
	"zero-tiktok/service/user/internal/svc"
	"zero-tiktok/service/user/internal/utils"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 登录
func (l *LoginLogic) Login(in *user.LoginOrRegisterRequest) (*user.LoginOrRegisterResponse, error) {
	// todo: add your logic here and delete this line
	tx := l.svcCtx.DBList.Mysql
	username := in.Name
	pass := in.Pass
	//todo 根据username去数据库查找用户，找不到返回用户不存在
	// 检查是否已经存在
	var count int64

	if err := tx.Model(&model.User{}).Where("name = ?", in.Name).Count(&count).Error; err != nil {

		return nil, myerror.ErrDB
	}

	if count == 0 {

		return nil, myerror.ErrDB
	}
	var usr model.User
	if err := tx.Model(&model.User{}).Where("name=?", username).First(&usr).Error; err != nil {

		return nil, myerror.ErrDB
	}
	fmt.Println("123")
	err := bcrypt.CompareHashAndPassword([]byte(usr.Pass), []byte(pass)) //验证（对比）
	if err != nil {

		return nil, myerror.ErrParam
	}
	fmt.Println("123")
	//todo 生成token
	token, err := utils.SetToken(usr.UserId)
	if err != nil {
		return nil, myerror.ErrDB
	}

	return &user.LoginOrRegisterResponse{
		Token: token,
	}, nil

}
