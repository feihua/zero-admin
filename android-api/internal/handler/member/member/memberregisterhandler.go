package handler

import (
	"net/http"

	"go-zero-admin/android-api/internal/logic/member/member"
	"go-zero-admin/android-api/internal/svc"
	"go-zero-admin/android-api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func MemberRegisterHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MemberRegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMemberRegisterLogic(r.Context(), ctx)
		resp, err := l.MemberRegister(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
