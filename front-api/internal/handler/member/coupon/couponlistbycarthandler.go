package coupon

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin/front-api/internal/logic/member/coupon"
	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"
)

func CouponListByCartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CouponListByCartReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := coupon.NewCouponListByCartLogic(r.Context(), svcCtx)
		resp, err := l.CouponListByCart(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
