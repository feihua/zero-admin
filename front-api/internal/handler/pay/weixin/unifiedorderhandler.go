package handler

import (
	"net/http"

	"go-zero-admin/front-api/internal/logic/pay/weixin"
	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UnifiedOrderHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UnifiedOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUnifiedOrderLogic(r.Context(), ctx)
		resp, err := l.UnifiedOrder(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
