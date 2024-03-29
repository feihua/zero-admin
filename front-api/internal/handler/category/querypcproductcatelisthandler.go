package category

import (
	"net/http"

	"github.com/feihua/zero-admin/front-api/internal/logic/category"
	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryPcProductCateListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := category.NewQueryPcProductCateListLogic(r.Context(), svcCtx)
		resp, err := l.QueryPcProductCateList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
