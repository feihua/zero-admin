// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package brand

import (
	"net/http"

	"github.com/feihua/zero-admin/api/front/internal/logic/product/brand"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryBrandDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryBrandDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := brand.NewQueryBrandDetailLogic(r.Context(), svcCtx)
		resp, err := l.QueryBrandDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
