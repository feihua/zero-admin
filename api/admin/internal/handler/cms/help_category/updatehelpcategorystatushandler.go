// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package help_category

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/cms/help_category"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateHelpCategoryStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateHelpCategoryStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := help_category.NewUpdateHelpCategoryStatusLogic(r.Context(), svcCtx)
		resp, err := l.UpdateHelpCategoryStatus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
