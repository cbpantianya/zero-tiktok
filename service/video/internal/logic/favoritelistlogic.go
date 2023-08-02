package logic

import (
	"context"
	"sort"
	"zero-tiktok/service/video/internal/model"

	"zero-tiktok/service/video/internal/svc"
	"zero-tiktok/service/video/pb/zero-tiktok/service/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteListLogic) FavoriteList(in *video.FavoriteRequest) (*video.VideoListResp, error) {
	// 喜欢列表
	var favoriteList []*model.Favorite
	err := l.svcCtx.DB.Where("user_id = ?", in.TargetId).Find(&favoriteList).Error
	if err != nil {
		return nil, err
	}

	var ids []int64
	for _, favorite := range favoriteList {
		ids = append(ids, favorite.VideoID)
	}

	// 获取视频信息
	var videos []*model.Video
	err = l.svcCtx.DB.Where("video_id in (?)", ids).Find(&videos).Error
	if err != nil {
		return nil, err
	}

	var resp map[int64]*video.Video
	resp = make(map[int64]*video.Video)
	for _, v := range videos {
		ids = append(ids, v.VideoID)
		resp[v.VideoID] = &video.Video{
			VideoId:       v.VideoID,
			AuthorId:      v.AuthorID,
			Play:          v.Play,
			Cover:         v.Cover,
			Title:         v.Title,
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
			PublishAt:     v.PublishAt.Unix(),
		}
	}

	for _, v := range favoriteList {
		resp[v.VideoID].FavoriteCount++
	}

	// 根据视频id查询评论数
	var comments []model.Comment
	err = l.svcCtx.DB.Model(model.Comment{}).Where("video_id in (?)", ids).Find(&comments).Error
	if err != nil {
		return nil, err
	}

	for _, v := range comments {
		resp[v.VideoID].CommentCount++
	}

	var list []*video.Video
	for _, v := range resp {
		list = append(list, v)
	}

	// 按照ID排序
	sort.Slice(list, func(i, j int) bool {
		return list[i].VideoId > list[j].VideoId
	})

	return &video.VideoListResp{
		List: list,
	}, nil
}
