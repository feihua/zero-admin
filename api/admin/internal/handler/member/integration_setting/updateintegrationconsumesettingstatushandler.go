package integration_setting

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/member/integration_setting"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateIntegrationConsumeSettingStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateIntegrationConsumeSettingStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := integration_setting.NewUpdateIntegrationConsumeSettingStatusLogic(r.Context(), svcCtx)
		resp, err := l.UpdateIntegrationConsumeSettingStatus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
