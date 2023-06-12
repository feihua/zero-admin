package middleware

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"time"
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

		// 解析 JSON 参数
		//var requestBody map[string]interface{}
		//err := httpx.ParseJsonBody(r, &requestBody)
		//req, _ := json.Marshal(requestBody)
		//logx.WithContext(r.Context()).Infov(string(req))
		//if err != nil {
		//	logx.Errorf("Failed to parse JSON body: %v", err)
		//	httpx.Error(w, errorx.NewDefaultError("Invalid JSON body"))
		//	return
		//}

		//all, _ := io.ReadAll(r.Body)
		startTime := time.Now()
		next(w, r)
		endTime := time.Now()
		duration := endTime.Sub(startTime)
		// 添加操作日志
		_, _ = m.Sys.SysLogAdd(r.Context(), &sysclient.SysLogAddReq{
			UserName:  userName,
			Operation: r.Method,
			Method:    r.RequestURI,
			Params:    "string(all)",
			Time:      duration.Microseconds(),
			Ip:        httpx.GetRemoteAddr(r),
			CreateBy:  userName,
		})
	}
}
