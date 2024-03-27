package attributecategory

import (
	"net/http"

	"github.com/feihua/zero-admin/api/internal/logic/product/attributecategory"
	"github.com/feihua/zero-admin/api/internal/svc"
	"github.com/feihua/zero-admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ProductAttributecategoryAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddProductAttributecategoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := attributecategory.NewProductAttributecategoryAddLogic(r.Context(), svcCtx)
		resp, err := l.ProductAttributecategoryAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
