package handler

import (
	"net/http"

	"github.com/feihua/zero-admin/job/internal/logic"
	"github.com/feihua/zero-admin/job/internal/svc"
	"github.com/feihua/zero-admin/job/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func JobHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewJobLogic(r.Context(), svcCtx)
		resp, err := l.Job(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
