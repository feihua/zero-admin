package history

import (
	"net/http"

	"github.com/feihua/zero-admin/front-api/internal/logic/history"
	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ReadHistoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := history.NewReadHistoryListLogic(r.Context(), svcCtx)
		resp, err := l.ReadHistoryList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
