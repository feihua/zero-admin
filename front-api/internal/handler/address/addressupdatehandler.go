package address

import (
	"net/http"
	"zero-admin/common/result"
	"zero-admin/front-api/internal/logic/address"
	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddressUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddressUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := address.NewAddressUpdateLogic(r.Context(), svcCtx)
		resp, err := l.AddressUpdate(&req)
		result.HttpResult(r, w, resp, err)
	}
}
