package logic

import (
	"context"
	myerror "zero-tiktok/internal/error"
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

		return nil, err
	}

	// 构建用户响应
	var userResponses []*user.User
	for _, usr := range users {
		//获得用户关注和粉丝
		var followerCount int64
		var followCount int64
		if err := tx.Model(&model.Relation{}).Where("user_id = ?", usr.UserId).Count(&followerCount).Error; err != nil {

			return nil, myerror.ErrDB
		}

		if err := tx.Model(&model.Relation{}).Where("follower_id = ?", usr.UserId).Count(&followCount).Error; err != nil {

			return nil, myerror.ErrDB
		}

		// 获得用户发布的视频总数
		var videoCount int64

		if err := tx.Model(&model.Video{}).Where("author_id = ?", usr.UserId).Count(&videoCount).Error; err != nil {
			return nil, myerror.ErrDB
		}

		//用户喜欢的总数
		var favoriteCount int64
		if err := tx.Model(&model.Favorite{}).Where("user_id = ?", usr.UserId).Count(&favoriteCount).Error; err != nil {
			return nil, myerror.ErrDB
		}

		//用户获赞总数
		var videos []model.Video
		var videoIds []int
		if err := tx.Model(&model.Video{}).Where("author_id = ?", usr.UserId).Find(&videos).Error; err != nil {
			return nil, myerror.ErrDB
		}

		for _, video := range videos {
			videoIds = append(videoIds, video.VideoId)
		}

		var count int64

		if err := tx.Model(&model.Favorite{}).Where("video_id IN ?", videoIds).Count(&count).Error; err != nil {
			return nil, myerror.ErrDB
		}
		userResponses = append(userResponses, &user.User{
			UserId:         int64(usr.UserId),
			Name:           usr.Name,
			Signature:      usr.Signature,
			Avatar:         usr.Avatar,
			Cover:          usr.Cover,
			FollowCount:    followCount,
			FollowerCount:  followerCount,
			FavoriteCount:  favoriteCount,
			TotalFavorited: count,
			VideoCount:     videoCount,
		})
	}

	return &user.GetUsersResponse{
		Users: userResponses,
	}, nil
}
