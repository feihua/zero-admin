package auth

import (
	"net/http"
	"zero-admin/common/result"
	"zero-admin/front-api/internal/logic/auth"
	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := auth.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		result.HttpResult(r, w, resp, err)
	}
}
