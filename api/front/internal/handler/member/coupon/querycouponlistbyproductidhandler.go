package coupon

import (
	"net/http"

	"github.com/feihua/zero-admin/api/front/internal/logic/member/coupon"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryCouponListByProductIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryCouponListByProductIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := coupon.NewQueryCouponListByProductIdLogic(r.Context(), svcCtx)
		resp, err := l.QueryCouponListByProductId(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
