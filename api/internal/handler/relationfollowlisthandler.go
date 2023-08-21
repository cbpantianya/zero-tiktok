package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-tiktok/api/internal/logic"
	"zero-tiktok/api/internal/svc"
	"zero-tiktok/api/internal/types"
)

func RelationFollowListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RelationFollowList
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRelationFollowListLogic(r.Context(), svcCtx)
		resp, err := l.RelationFollowList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
