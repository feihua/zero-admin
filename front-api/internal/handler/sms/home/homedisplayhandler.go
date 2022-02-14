package handler

import (
	"net/http"

	"go-zero-admin/front-api/internal/logic/sms/home"
	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HomeDisplayHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListHomeDisplayReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHomeDisplayLogic(r.Context(), ctx)
		resp, err := l.HomeDisplay(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
