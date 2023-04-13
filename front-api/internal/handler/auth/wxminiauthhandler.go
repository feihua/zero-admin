package auth

import (
	"net/http"

	"zero-admin/common/result"
	"zero-admin/front-api/internal/logic/auth"
	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func WxMiniAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WXMiniAuthReq
		if err := httpx.Parse(r, &req); err != nil {
			// httpx.Error(w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		l := auth.NewWxMiniAuthLogic(r.Context(), svcCtx)
		resp, err := l.WxMiniAuth(&req)
		result.HttpResult(r, w, resp, err)
		// if err != nil {
		// 	httpx.Error(w, err)
		// } else {
		// 	httpx.OkJson(w, resp)
		// }
	}
}
