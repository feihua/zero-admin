// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package sys_notice

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/sys/sys_notice"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateNoticeStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateNoticeStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sys_notice.NewUpdateNoticeStatusLogic(r.Context(), svcCtx)
		resp, err := l.UpdateNoticeStatus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
