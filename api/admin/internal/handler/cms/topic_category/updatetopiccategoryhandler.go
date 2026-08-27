// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package topic_category

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/cms/topic_category"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateTopicCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateTopicCategoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := topic_category.NewUpdateTopicCategoryLogic(r.Context(), svcCtx)
		resp, err := l.UpdateTopicCategory(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
