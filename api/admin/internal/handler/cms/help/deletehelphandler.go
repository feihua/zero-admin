// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package help

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/cms/help"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteHelpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteHelpReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := help.NewDeleteHelpLogic(r.Context(), svcCtx)
		resp, err := l.DeleteHelp(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
