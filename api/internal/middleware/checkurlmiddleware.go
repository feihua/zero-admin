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
		if userId == "" {
			logx.Errorf("缺少必要参数x-user-id")
			httpx.Error(w, errorx.NewDefaultError("缺少必要参数x-user-id"))
			return
		}

		if r.RequestURI == "/api/sys/user/currentUser" || r.RequestURI == "/api/sys/user/selectAllData" || r.RequestURI == "/api/sys/role/queryMenuByRoleId" {
			logx.Infof("用户userId: %s,访问: %s路径", userId, r.RequestURI)
			next(w, r)
		} else {
			//获取用户能访问的url
			urls, err := m.Redis.Get(userId)
			if err != nil {
				logx.Errorf("用户：%s,获取redis连接异常", userId)
				httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("用户：%s,获取redis连接异常", userId)))
				return
			}

			if len(strings.TrimSpace(urls)) == 0 {
				logx.Errorf("用户userId: %s,还没有登录", userId)
				httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("用户userId: %s,还没有登录,请先登录", userId)))
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

			if b {
				logx.Infof("用户userId: %s,访问: %s路径", userId, r.RequestURI)
				next(w, r)
			} else {
				logx.Errorf("用户userId: %s,没有访问: %s路径的权限", userId, r.RequestURI)
				httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("用户userId: %s,没有访问: %s,路径的的权限,请联系管理员", userId, r.RequestURI)))
				return
			}

		}

	}
}
