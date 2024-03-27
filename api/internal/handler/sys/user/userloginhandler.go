package user

import (
	"net/http"

	"github.com/feihua/zero-admin/api/internal/logic/sys/user"
	"github.com/feihua/zero-admin/api/internal/svc"
	"github.com/feihua/zero-admin/api/internal/types"

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
		l := user.NewUserLoginLogic(r.Context(), ctx)
		resp, err := l.UserLogin(req, httpx.GetRemoteAddr(r))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
