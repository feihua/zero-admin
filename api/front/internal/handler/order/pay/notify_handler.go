package pay

import (
	"net/http"

	"github.com/feihua/zero-admin/api/front/internal/logic/order/pay"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func NotifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := pay.NewNotifyLogic(w, r, r.Context(), svcCtx)
		err := l.Notify()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
