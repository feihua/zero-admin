package homerecommendsubject

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/sms/homerecommendsubject"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateRecommendSubjectSortHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateRecommendSubjectSortReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := homerecommendsubject.NewUpdateRecommendSubjectSortLogic(r.Context(), svcCtx)
		resp, err := l.UpdateRecommendSubjectSort(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
