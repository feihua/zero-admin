package seckill_activity

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/sms/seckill_activity"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateSeckillActivityStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSeckillActivityStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := seckill_activity.NewUpdateSeckillActivityStatusLogic(r.Context(), svcCtx)
		resp, err := l.UpdateSeckillActivityStatus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
