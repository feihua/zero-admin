package flash_promotion_session

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/sms/flash_promotion_session"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateFlashPromotionSessionStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateFlashPromotionSessionStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := flash_promotion_session.NewUpdateFlashPromotionSessionStatusLogic(r.Context(), svcCtx)
		resp, err := l.UpdateFlashPromotionSessionStatus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
