package middleware

import (
	"github.com/tal-tech/go-zero/core/logx"
	"net/http"
)

type CheckUrlMiddleware struct {
}

func NewCheckUrlMiddleware() *CheckUrlMiddleware {
	return &CheckUrlMiddleware{}
}

func (m *CheckUrlMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		logx.Info("example middle")
		next(w, r)
	}
}
