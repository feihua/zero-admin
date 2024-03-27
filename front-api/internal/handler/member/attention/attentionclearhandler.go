package attention

import (
	"net/http"

	"github.com/feihua/zero-admin/front-api/internal/logic/member/attention"
	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AttentionClearHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := attention.NewAttentionClearLogic(r.Context(), svcCtx)
		resp, err := l.AttentionClear()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
