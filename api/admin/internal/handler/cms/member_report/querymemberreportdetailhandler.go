// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package member_report

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/cms/member_report"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryMemberReportDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryMemberReportDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := member_report.NewQueryMemberReportDetailLogic(r.Context(), svcCtx)
		resp, err := l.QueryMemberReportDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
