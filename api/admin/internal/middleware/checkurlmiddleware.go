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
	Redis       *redis.Redis
	ExcludeUrls string
}

func NewCheckUrlMiddleware(Redis *redis.Redis, excludeUrls string) *CheckUrlMiddleware {
	return &CheckUrlMiddleware{Redis: Redis, ExcludeUrls: excludeUrls}
}

func (m *CheckUrlMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		uri := strings.Split(r.RequestURI, "?")[0]
		// 判断请求header中是否携带了x-user-id
		userId := r.Context().Value("userId").(json.Number).String()
		userName := r.Context().Value("userName").(string)
		if userId == "" || userName == "" {
			logc.Errorf(r.Context(), "缺少必要参数x-user-id")
			httpx.Error(w, errorx.NewDefaultError("缺少必要参数x-user-id"))
			return
		}

		// 获取用户能访问的url
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

		backUrls := strings.Split(fmt.Sprintf("%s,%s", urls, m.ExcludeUrls), ",")

		if !isValueExist(backUrls, uri) {
			logc.Errorf(r.Context(), "用户: %s,没有访问: %s路径的权限", userName, uri)
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("用户: %s,没有访问: %s,路径的的权限,请联系管理员", userName, uri)))
			return
		}

		next(w, r)
	}
}

// isValueExist 检查一个字符串切片中是否存在指定的字符串值。
// 这个函数通过遍历切片中的每个元素来搜索匹配的字符串。
// 如果找到匹配项，它会立即返回true，表示值存在。
// 如果遍历完所有元素都没有找到匹配项，则返回false，表示值不存在。
//
// 参数:
//
//	arr []string: 要搜索的字符串切片。
//	value string: 要查找的字符串值。
//
// 返回值:
//
//	bool: 如果切片中存在指定的值，则返回true，否则返回false。
func isValueExist(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}
