// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package order_delivery

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/oms/order_delivery"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateOrderDeliveryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateOrderDeliveryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order_delivery.NewUpdateOrderDeliveryLogic(r.Context(), svcCtx)
		resp, err := l.UpdateOrderDelivery(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
