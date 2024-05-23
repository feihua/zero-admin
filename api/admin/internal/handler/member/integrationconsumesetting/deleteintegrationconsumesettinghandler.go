package integrationconsumesetting

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/member/integrationconsumesetting"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteIntegrationConsumeSettingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteIntegrationConsumeSettingReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := integrationconsumesetting.NewDeleteIntegrationConsumeSettingLogic(r.Context(), svcCtx)
		resp, err := l.DeleteIntegrationConsumeSetting(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
