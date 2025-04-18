package member_integration_setting

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/ums/member_integration_setting"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryIntegrationConsumeSettingDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryIntegrationConsumeSettingDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := member_integration_setting.NewQueryIntegrationConsumeSettingDetailLogic(r.Context(), svcCtx)
		resp, err := l.QueryIntegrationConsumeSettingDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
