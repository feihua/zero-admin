package home

import (
	"net/http"

	"github.com/feihua/zero-admin/front-api/internal/logic/home"
	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HomeIndexHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := home.NewHomeIndexLogic(r.Context(), svcCtx)
		resp, err := l.HomeIndex()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
