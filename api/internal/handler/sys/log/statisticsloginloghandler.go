package handler

import (
	"net/http"
	logic "zero-admin/api/internal/logic/sys/log"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"
)

func StatisticsLoginLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StatisticsLoginLogReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewStatisticsLoginLogLogic(r.Context(), svcCtx)
		resp, err := l.StatisticsLoginLog(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
