package collection

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin/front-api/internal/logic/collection"
	"zero-admin/front-api/internal/svc"
)

func ProductCollectionClearHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := collection.NewProductCollectionClearLogic(r.Context(), svcCtx)
		resp, err := l.ProductCollectionClear()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
