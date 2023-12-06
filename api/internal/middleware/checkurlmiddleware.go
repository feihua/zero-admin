package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strings"
	"zero-admin/api/internal/common/errorx"
)

type CheckUrlMiddleware struct {
	Redis *redis.Redis
}

func NewCheckUrlMiddleware(Redis *redis.Redis) *CheckUrlMiddleware {
	return &CheckUrlMiddleware{Redis: Redis}
}

func (m *CheckUrlMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//判断请求header中是否携带了x-user-id
		userId := r.Context().Value("userId").(json.Number).String()
		userName := r.Context().Value("userName").(string)
		if userId == "" || userName == "" {
			logx.WithContext(r.Context()).Errorf("缺少必要参数x-user-id")
			httpx.Error(w, errorx.NewDefaultError("缺少必要参数x-user-id"))
			return
		}

		if r.RequestURI == "/api/sys/user/currentUser" || r.RequestURI == "/api/sys/user/selectAllData" || r.RequestURI == "/api/sys/role/queryMenuByRoleId" {
			next(w, r)
			return
		}

		//获取用户能访问的url
		urls, err := m.Redis.Get(userId)
		if err != nil {
			logx.WithContext(r.Context()).Errorf("用户：%s,获取redis连接异常", userName)
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("用户：%s,获取redis连接异常", userName)))
			return
		}

		if len(strings.TrimSpace(urls)) == 0 {
			logx.WithContext(r.Context()).Errorf("用户: %s,还没有登录", userName)
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("用户: %s,还没有登录,请先登录", userName)))
			return
		}

		backUrls := strings.Split(urls, ",")

		b := false
		for _, url := range backUrls {
			if url == r.RequestURI {
				b = true
				break
			}
		}

		if !b {
			logx.WithContext(r.Context()).Errorf("用户: %s,没有访问: %s路径的权限", userName, r.RequestURI)
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("用户: %s,没有访问: %s,路径的的权限,请联系管理员", userName, r.RequestURI)))
			return
		}

		next(w, r)
	}
}
