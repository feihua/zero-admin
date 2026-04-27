package middleware

import (
	"context"
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
		userIdValue := r.Context().Value("userId")
		userNameValue := r.Context().Value("userName")

		userId, ok := userIdValue.(json.Number)
		if !ok || userId == "" {
			logc.Errorf(r.Context(), "缺少必要参数userId")
			httpx.Error(w, errorx.NewDefaultError("缺少必要参数userId"))
			return
		}

		userName, ok := userNameValue.(string)
		if !ok || userName == "" {
			logc.Errorf(r.Context(), "缺少必要参数userName")
			httpx.Error(w, errorx.NewDefaultError("缺少必要参数userName"))
			return
		}

		key := fmt.Sprintf("zero:user:%s", userId)
		
		//// 检查用户登录状态
		loginStatus, err := m.checkLoginStatus(r.Context(), key, userName, uri)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		// 如果是管理员，直接放行
		if loginStatus.isAdmin {
			next(w, r)
			return
		}

		// 验证用户权限
		if !loginStatus.hasPermission {
			logc.Errorf(r.Context(), "用户: %s,没有访问: %s 路径的权限", userName, uri)
			httpx.Error(w, errorx.NewDefaultError(
				fmt.Sprintf("用户: %s,没有访问: %s 路径的权限,请联系管理员", userName, uri),
			))
			return
		}

		// 权限验证通过，继续执行下一个处理器
		next(w, r)
	}
}

type LoginStatus struct {
	isAdmin       bool
	hasPermission bool
}

// checkLoginStatus 检查用户登录状态
func (m *CheckUrlMiddleware) checkLoginStatus(ctx context.Context, key, userName, uri string) (*LoginStatus, error) {
	// 检查 token 是否存在
	exists, err := m.Redis.HexistsCtx(ctx, key, "token")
	if err != nil {
		logc.Errorf(ctx, "用户：%s,获取redis连接异常", userName)
		return nil, errorx.NewDefaultError(fmt.Sprintf("用户：%s,获取redis连接异常", userName))
	}

	if !exists {
		logc.Errorf(ctx, "用户：%s,登录已超时", userName)
		return nil, errorx.NewDefaultError(fmt.Sprintf("用户：%s,登录已超时，请求重新登录", userName))
	}

	// 获取 token
	token, err := m.Redis.HgetCtx(ctx, key, "token")
	if err != nil {
		logc.Errorf(ctx, "用户：%s,获取redis连接异常", userName)
		return nil, errorx.NewDefaultError(fmt.Sprintf("用户：%s,获取redis连接异常", userName))
	}

	if len(strings.TrimSpace(token)) == 0 {
		logc.Errorf(ctx, "用户: %s,还没有登录", userName)
		return nil, errorx.NewDefaultError(fmt.Sprintf("用户: %s,还没有登录,请先登录", userName))
	}

	// 检查是否为管理员
	isAdmin, err := m.Redis.HgetCtx(ctx, key, "isAdmin")
	if err != nil {
		logc.Errorf(ctx, "用户：%s,获取redis连接异常", userName)
		return nil, errorx.NewDefaultError(fmt.Sprintf("用户：%s,获取redis连接异常", userName))
	}

	if isAdmin == "1" {
		return &LoginStatus{isAdmin: true}, nil
	}

	// 获取用户权限
	permissions, err := m.Redis.HgetCtx(ctx, key, "permissions")
	if err != nil {
		logc.Errorf(ctx, "用户：%s,获取redis连接异常", userName)
		return nil, errorx.NewDefaultError(fmt.Sprintf("用户：%s,获取redis连接异常", userName))
	}

	if len(strings.TrimSpace(permissions)) == 0 {
		logc.Errorf(ctx, "用户: %s,还没有登录", userName)
		return nil, errorx.NewDefaultError(fmt.Sprintf("用户: %s,还没有登录,请先登录", userName))
	}

	allUrls := strings.Split(fmt.Sprintf("%s,%s", permissions, m.ExcludeUrls), ",")
	return &LoginStatus{isAdmin: false, hasPermission: isValueExist(allUrls, uri)}, nil
}

// bool: 如果切片中存在指定的值，则返回true，否则返回false。
func isValueExist(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}
