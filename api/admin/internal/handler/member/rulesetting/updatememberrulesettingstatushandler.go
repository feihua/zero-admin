package rulesetting

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/member/rulesetting"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateMemberRuleSettingStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateMemberRuleSettingStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := rulesetting.NewUpdateMemberRuleSettingStatusLogic(r.Context(), svcCtx)
		resp, err := l.UpdateMemberRuleSettingStatus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
