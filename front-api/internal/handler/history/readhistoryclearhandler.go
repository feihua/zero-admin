package history

import (
	"net/http"

	"github.com/feihua/zero-admin/front-api/internal/logic/history"
	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ReadHistoryClearHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := history.NewReadHistoryClearLogic(r.Context(), svcCtx)
		resp, err := l.ReadHistoryClear()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
