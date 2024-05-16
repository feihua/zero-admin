package address

import (
	"net/http"

	"github.com/feihua/zero-admin/api/front/internal/logic/member/address"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryMemberAddressListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := address.NewQueryMemberAddressListLogic(r.Context(), svcCtx)
		resp, err := l.QueryMemberAddressList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
