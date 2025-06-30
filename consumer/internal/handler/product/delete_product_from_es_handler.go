package product

import (
	"net/http"

	"github.com/feihua/zero-admin/consumer/internal/logic/product"
	"github.com/feihua/zero-admin/consumer/internal/svc"
	"github.com/feihua/zero-admin/consumer/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteProductFromEsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductEsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := product.NewDeleteProductFromEsLogic(r.Context(), svcCtx)
		resp, err := l.DeleteProductFromEs(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
