package middleware

import (
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"net/http"
)

type CheckUrlMiddleware struct {
	Redis *redis.Redis
}

func NewCheckUrlMiddleware(Redis *redis.Redis) *CheckUrlMiddleware {
	return &CheckUrlMiddleware{Redis: Redis}
}

func (m *CheckUrlMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		logx.Infof("example middle: %s", r.RequestURI)
		get, _ := m.Redis.Get("1")
		logx.Infof("example middle: %s", get)
		next(w, r)
	}
}
