package attention

import (
	"net/http"

	"github.com/feihua/zero-admin/front-api/internal/logic/member/attention"
	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AttentionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := attention.NewAttentionListLogic(r.Context(), svcCtx)
		resp, err := l.AttentionList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
