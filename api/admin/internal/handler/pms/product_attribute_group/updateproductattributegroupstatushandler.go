package product_attribute_group

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/pms/product_attribute_group"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateProductAttributeGroupStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateProductAttributeGroupStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := product_attribute_group.NewUpdateProductAttributeGroupStatusLogic(r.Context(), svcCtx)
		resp, err := l.UpdateProductAttributeGroupStatus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
