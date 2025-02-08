package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strings"
)

type CheckUrlMiddleware struct {
	Redis *redis.Redis
}

func NewCheckUrlMiddleware(Redis *redis.Redis) *CheckUrlMiddleware {
	return &CheckUrlMiddleware{Redis: Redis}
}

func (m *CheckUrlMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		uri := strings.Split(r.RequestURI, "?")[0]
		//判断请求header中是否携带了x-user-id
		userId := r.Context().Value("userId").(json.Number).String()
		userName := r.Context().Value("userName").(string)
		if userId == "" || userName == "" {
			logc.Errorf(r.Context(), "缺少必要参数x-user-id")
			httpx.Error(w, errorx.NewDefaultError("缺少必要参数x-user-id"))
			return
		}

		if uri == "/api/sys/user/info" || uri == "/api/sys/user/queryAllRelations" || uri == "/api/sys/role/queryMenuByRoleId" {
			next(w, r)
			return
		}

		//获取用户能访问的url
		urls, err := m.Redis.Get("zero:mall:token:" + userId)
		if err != nil {
			logc.Errorf(r.Context(), "用户：%s,获取redis连接异常", userName)
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("用户：%s,获取redis连接异常", userName)))
			return
		}

		if len(strings.TrimSpace(urls)) == 0 {
			logc.Errorf(r.Context(), "用户: %s,还没有登录", userName)
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("用户: %s,还没有登录,请先登录", userName)))
			return
		}

		backUrls := strings.Split(urls, ",")

		b := false
		for _, url := range backUrls {
			if url == uri {
				b = true
				break
			}
		}

		if !b {
			logc.Errorf(r.Context(), "用户: %s,没有访问: %s路径的权限", userName, uri)
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("用户: %s,没有访问: %s,路径的的权限,请联系管理员", userName, uri)))
			return
		}

		next(w, r)
	}
}
