package seckill_product

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/sms/seckill_product"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteSeckillProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteSeckillProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := seckill_product.NewDeleteSeckillProductLogic(r.Context(), svcCtx)
		resp, err := l.DeleteSeckillProduct(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
