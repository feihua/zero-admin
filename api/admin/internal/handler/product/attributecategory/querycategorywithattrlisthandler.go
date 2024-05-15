package attributecategory

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/product/attributecategory"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryCategoryWithAttrListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := attributecategory.NewQueryCategoryWithAttrListLogic(r.Context(), svcCtx)
		resp, err := l.QueryCategoryWithAttrList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
