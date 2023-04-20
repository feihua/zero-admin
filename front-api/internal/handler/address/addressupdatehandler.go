package address

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin/front-api/internal/logic/address"
	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"
)

func AddressUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddressUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := address.NewAddressUpdateLogic(r.Context(), svcCtx)
		resp, err := l.AddressUpdate(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
