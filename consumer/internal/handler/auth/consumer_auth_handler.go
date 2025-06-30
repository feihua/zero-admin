package auth

import (
	"net/http"

	"github.com/feihua/zero-admin/consumer/internal/logic/auth"
	"github.com/feihua/zero-admin/consumer/internal/svc"
	"github.com/feihua/zero-admin/consumer/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ConsumerAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewConsumerAuthLogic(r.Context(), svcCtx)
		resp, err := l.ConsumerAuth(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
