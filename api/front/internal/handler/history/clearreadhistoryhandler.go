package history

import (
	"net/http"

	"github.com/feihua/zero-admin/api/front/internal/logic/history"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ClearReadHistoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := history.NewClearReadHistoryLogic(r.Context(), svcCtx)
		resp, err := l.ClearReadHistory()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
