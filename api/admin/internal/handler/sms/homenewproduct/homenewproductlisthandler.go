package homenewproduct

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/sms/homenewproduct"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HomeNewProductListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListHomeNewProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := homenewproduct.NewHomeNewProductListLogic(r.Context(), ctx)
		resp, err := l.HomeNewProductList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
