package logic

import (
	"context"
	"fmt"
	myerror "zero-tiktok/internal/error"
	"zero-tiktok/service/user/internal/model"
	"zero-tiktok/service/user/internal/svc"
	"zero-tiktok/service/user/internal/utils"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注册
func (l *RegisterLogic) Register(in *user.LoginOrRegisterRequest) (*user.LoginOrRegisterResponse, error) {
	// todo: add your logic here and delete this line
	tx := l.svcCtx.DBList.Mysql
	username := in.Name
	// 检查是否已经存在
	var count int64
	if err := tx.Model(&model.User{}).Where("name = ?", in.Name).Count(&count).Error; err != nil {

		return nil, myerror.ErrDB
	}
	if count > 0 {

		return nil, myerror.ErrDB
	}
	pass := in.Pass
	salt := bcrypt.MinCost
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), salt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	encodePWD := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	fmt.Println(encodePWD)
	usr := model.User{
		Name:      username,
		Pass:      encodePWD,
		Signature: "这个人很懒，什么都没有留下",
		Avatar:    "https://pic1.zhimg.com/50/v2-6afa72220d29f045c15217aa6b275808_hd.jpg?source=1940ef5c",
		Cover:     "https://tse3-mm.cn.bing.net/th/id/OIP-C.K03bmv8yI_CYhvW_-_unRgHaEN?w=329&h=187&c=7&r=0&o=5&dpr=1.3&pid=1.7",
	}

	//TODO 这里写插入数据库，并根据jwt生成token
	if err = tx.Model(&model.User{}).Create(&usr).Error; err != nil {
		return nil, err
	}

	token, err := utils.SetToken(usr.UserId)
	if err != nil {
		return nil, err
	}
	return &user.LoginOrRegisterResponse{
		Token: token,
	}, nil
}
