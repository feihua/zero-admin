package middleware

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zero-admin/rpc/sys/sys"
	"zero-admin/rpc/sys/sysclient"
)

type AddLogMiddleware struct {
	Sys sys.Sys
}

func NewAddLogMiddleware(Sys sys.Sys) *AddLogMiddleware {
	return &AddLogMiddleware{Sys: Sys}
}

func (m *AddLogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.RequestURI == "/api/sys/user/login" {
			next(w, r)
			return
		}

		userName := r.Context().Value("userName").(string)

		// 添加操作日志
		_, _ = m.Sys.SysLogAdd(r.Context(), &sysclient.SysLogAddReq{
			UserName:  userName,
			Operation: r.Method,
			Method:    r.RequestURI,
			Params:    "test",
			Time:      0,
			Ip:        httpx.GetRemoteAddr(r),
			CreateBy:  userName,
		})
		next(w, r)

	}
}
