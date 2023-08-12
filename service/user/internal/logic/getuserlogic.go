package logic

import (
	"context"
	myerror "zero-tiktok/internal/error"
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

	//获得用户关注和粉丝
	var followerCount int64
	var followCount int64
	if err := tx.Model(&model.Relation{}).Where("user_id = ?", in.UserId).Count(&followerCount).Error; err != nil {
		tx.Rollback()
		return nil, myerror.ErrDB
	}

	if err := tx.Model(&model.Relation{}).Where("follower_id = ?", in.UserId).Count(&followCount).Error; err != nil {
		tx.Rollback()
		return nil, myerror.ErrDB
	}

	// 获得用户发布的视频总数
	var videoCount int64

	if err := tx.Model(&model.Video{}).Where("author_id = ?", in.UserId).Count(&videoCount).Error; err != nil {
		return nil, myerror.ErrDB
	}

	//用户喜欢的总数
	var favoriteCount int64
	if err := tx.Model(&model.Favorite{}).Where("user_id = ?", in.UserId).Count(&favoriteCount).Error; err != nil {
		return nil, myerror.ErrDB
	}

	//用户获赞总数
	var videos []model.Video
	var videoIds []int
	if err := tx.Model(&model.Video{}).Where("author_id = ?", in.UserId).Find(&videos).Error; err != nil {
		return nil, myerror.ErrDB
	}

	for _, video := range videos {
		videoIds = append(videoIds, video.VideoId)
	}

	var count int64

	if err := tx.Model(&model.Favorite{}).Where("video_id IN ?", videoIds).Count(&count).Error; err != nil {
		return nil, myerror.ErrDB
	}
	tx.Commit()
	return &user.GetUserResponse{
		User: &user.User{
			UserId:         uid,
			Name:           usr.Name,
			Signature:      usr.Signature,
			Avatar:         usr.Avatar,
			Cover:          usr.Cover,
			FollowCount:    followCount,
			FollowerCount:  followerCount,
			FavoriteCount:  favoriteCount,
			TotalFavorited: count,
			VideoCount:     videoCount,
		},
	}, nil
}
