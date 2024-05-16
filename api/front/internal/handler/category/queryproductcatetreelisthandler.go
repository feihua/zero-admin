package category

import (
	"net/http"

	"github.com/feihua/zero-admin/api/front/internal/logic/category"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryProductCateTreeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := category.NewQueryProductCateTreeListLogic(r.Context(), svcCtx)
		resp, err := l.QueryProductCateTreeList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
