package attribute

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/product/attribute"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryByproductCategoryIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryByproductCategoryIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := attribute.NewQueryByproductCategoryIdLogic(r.Context(), svcCtx)
		resp, err := l.QueryByproductCategoryId(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
