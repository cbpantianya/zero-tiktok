package logic

import (
	"context"
	"sort"
	"time"
	"zero-tiktok/service/video/internal/model"

	"zero-tiktok/service/video/internal/svc"
	"zero-tiktok/service/video/pb/zero-tiktok/service/video"

	"github.com/zeromicro/go-zero/core/logx"
	e "zero-tiktok/internal/error"
)

type FeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeedLogic) Feed(in *video.FeedRequest) (*video.VideoListResp, error) {

	// 1. 根据时间查询视频列表
	var videos []model.Video
	var latest time.Time
	latest = time.Unix(in.Latest, 0)
	// 时间戳比较
	err := l.svcCtx.DB.Model(model.Video{}).Where("publish_at < ?", latest).Order("publish_at desc").Limit(10).Find(&videos).Error
	if err != nil {
		return nil, e.ErrDB
	}

	var ids []int64
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

	// 根据视频id查询喜欢数
	var favorites []model.Favorite
	err = l.svcCtx.DB.Model(model.Favorite{}).Where("video_id in (?)", ids).Find(&favorites).Error
	if err != nil {
		return nil, e.ErrDB
	}

	for _, v := range favorites {
		resp[v.VideoID].FavoriteCount++
	}

	// 根据视频id查询评论数
	var comments []model.Comment
	err = l.svcCtx.DB.Model(model.Comment{}).Where("video_id in (?)", ids).Find(&comments).Error
	if err != nil {
		return nil, e.ErrDB
	}

	for _, v := range comments {
		resp[v.VideoID].CommentCount++
	}

	// 是否有用户ID
	if in.UserId != nil {
		// favourite中是否有用户ID
		for _, v := range favorites {
			if v.UserID == *in.UserId {
				resp[v.VideoID].IsFavorite = true
			}
		}
	}
	// map 转List
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
