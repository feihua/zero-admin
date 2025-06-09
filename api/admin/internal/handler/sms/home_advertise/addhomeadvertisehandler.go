package home_advertise

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/sms/home_advertise"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddHomeAdvertiseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddHomeAdvertiseReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := home_advertise.NewAddHomeAdvertiseLogic(r.Context(), svcCtx)
		resp, err := l.AddHomeAdvertise(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
