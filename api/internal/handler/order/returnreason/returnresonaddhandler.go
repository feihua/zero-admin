package handler

import (
	"net/http"

	"github.com/feihua/zero-admin/api/internal/logic/order/returnreason"
	"github.com/feihua/zero-admin/api/internal/svc"
	"github.com/feihua/zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ReturnResonAddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddReturnResonReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewReturnResonAddLogic(r.Context(), ctx)
		resp, err := l.ReturnResonAdd(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
