package collection

import (
	"net/http"

	"github.com/feihua/zero-admin/api/front/internal/logic/collection"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryProductCollectionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := collection.NewQueryProductCollectionListLogic(r.Context(), svcCtx)
		resp, err := l.QueryProductCollectionList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
