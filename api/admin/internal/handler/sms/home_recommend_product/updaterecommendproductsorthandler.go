package home_recommend_product

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/sms/home_recommend_product"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateRecommendProductSortHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateRecommendProductSortReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := home_recommend_product.NewUpdateRecommendProductSortLogic(r.Context(), svcCtx)
		resp, err := l.UpdateRecommendProductSort(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
