package handler

import (
	"net/http"

	"github.com/feihua/zero-admin/api/internal/logic/sms/coupon"
	"github.com/feihua/zero-admin/api/internal/svc"
	"github.com/feihua/zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CouponDeleteHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteCouponReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCouponDeleteLogic(r.Context(), ctx)
		resp, err := l.CouponDelete(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
