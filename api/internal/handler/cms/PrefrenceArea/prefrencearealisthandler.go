package PrefrenceArea

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-admin/api/internal/logic/cms/PrefrenceArea"
	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"
)

func PrefrenceAreaListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListPrefrenceAreaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := PrefrenceArea.NewPrefrenceAreaListLogic(r.Context(), svcCtx)
		resp, err := l.PrefrenceAreaList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
