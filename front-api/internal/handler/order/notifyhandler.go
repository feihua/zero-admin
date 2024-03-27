package order

import (
	"net/http"

	"github.com/feihua/zero-admin/front-api/internal/logic/order"
	"github.com/feihua/zero-admin/front-api/internal/svc"
)

func NotifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := order.NewNotifyLogic(w, r, r.Context(), svcCtx)
		l.Notify()
		//注释下面,因为支付回调响应用w写出,所以不要下面的
		//err := l.Notify()
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.Ok(w)
		//}
	}
}
