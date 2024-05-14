package homerecommendproduct

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/sms/homerecommendproduct"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HomeRecommendProductListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListHomeRecommendProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := homerecommendproduct.NewHomeRecommendProductListLogic(r.Context(), ctx)
		resp, err := l.HomeRecommendProductList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
