package order

import (
	"net/http"

	"github.com/feihua/zero-admin/api/front/internal/logic/order/order"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ConfirmReceiveOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ConfirmReceiveOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewConfirmReceiveOrderLogic(r.Context(), svcCtx)
		resp, err := l.ConfirmReceiveOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
