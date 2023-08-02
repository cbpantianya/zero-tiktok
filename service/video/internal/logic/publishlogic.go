package logic

import (
	"context"
	"sort"
	"zero-tiktok/service/video/internal/model"

	"zero-tiktok/service/video/internal/svc"
	"zero-tiktok/service/video/pb/zero-tiktok/service/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishLogic) Publish(in *video.PublishRequest) (*video.VideoListResp, error) {
	var videos []*model.Video
	err := l.svcCtx.DB.Where("author_id = ?", in.TargetId).Find(&videos).Error
	if err != nil {
		return nil, err
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
		return nil, err
	}

	for _, v := range favorites {
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

	// favourite中是否有用户ID
	for _, v := range favorites {
		if v.UserID == in.UserId {
			resp[v.VideoID].IsFavorite = true
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
