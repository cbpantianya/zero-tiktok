package handler

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"net/http"
	"zero-tiktok/service/user/pb/zero-tiktok/service/user"
	"zero-tiktok/service/video/pb/zero-tiktok/service/video"

	"github.com/h2non/filetype"
	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-tiktok/api/internal/logic"
	"zero-tiktok/api/internal/svc"
	"zero-tiktok/api/internal/types"
)

func UploadVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadVideo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 手动解析form中file字段的数据
		file, header, err := r.FormFile("data")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		defer func(file multipart.File) {
			_ = file.Close()
		}(file)

		name := uuid.New().String()

		err = svcCtx.OSSClient.PutObject("video/"+name, file)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 判断是否为视频
		tmp, err := header.Open()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, tmp); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		if !filetype.IsVideo(buf.Bytes()) {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		if err = tmp.Close(); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 抽取第一帧作为封面

		id, err := svcCtx.UserClient.GetIdByToken(r.Context(), &user.TokenToUserRequest{
			Token: req.Token,
		})
		fmt.Printf("%+v\n", id)
		if err != nil || id.UserId == 0 {
			httpx.OkJsonCtx(r.Context(), w, &types.UploadVideoResp{
				Msg:  "未登录",
				Code: -1,
			})
			return
		}

		url := fmt.Sprintf("%s/video/%s", svcCtx.Config.OSS.Domain, name)

		_, err = svcCtx.VideoClient.CreateVideo(r.Context(), &video.CreateVideoReq{
			UserId: id.UserId,
			Play:   url,
			Title:  req.Title,
		})
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUploadVideoLogic(r.Context(), svcCtx)
		resp, err := l.UploadVideo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
