package user

import (
	"github.com/ua-parser/uap-go/uaparser"
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/sys/user"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserLoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		//获取客户端真实的ip
		//nginx的配置
		//location /api {
		//		proxy_set_header Host $host;
		//		proxy_set_header X-Real-IP $remote_addr;
		//		proxy_set_header REMOTE-HOST $remote_addr;
		//		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		//		proxy_pass http://127.0.0.1:8888;
		//}
		userAgent := r.Header.Get("User-Agent")
		parser := uaparser.NewFromSaved()
		ua := parser.Parse(userAgent)

		browser := ua.UserAgent.Family + " " + ua.UserAgent.Major
		os := ua.Os.Family + " " + ua.Os.Major

		l := user.NewUserLoginLogic(r.Context(), ctx)
		resp, err := l.UserLogin(&req, httpx.GetRemoteAddr(r), browser, os)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
