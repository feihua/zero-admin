package member

import (
	"net/http"

	"github.com/feihua/zero-admin/api/web/internal/logic/member/member"
	"github.com/feihua/zero-admin/api/web/internal/svc"
	"github.com/feihua/zero-admin/api/web/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateMemberHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateMemberReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := member.NewUpdateMemberLogic(r.Context(), svcCtx)
		resp, err := l.UpdateMember(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
