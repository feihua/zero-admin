package collection

import (
	"net/http"

	"github.com/feihua/zero-admin/api/front/internal/logic/collection"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteProductCollectionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductCollectionDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := collection.NewDeleteProductCollectionLogic(r.Context(), svcCtx)
		resp, err := l.DeleteProductCollection(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
