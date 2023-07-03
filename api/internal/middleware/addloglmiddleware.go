package middleware

import (
	"bytes"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io/ioutil"
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

// Handle 参考chatgpt实现
func (m *AddLogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		uri := r.RequestURI
		if uri == "/api/sys/user/login" || uri == "/api/sys/upload" {
			logx.Infof("Request: %s %s", r.Method, uri)
			next(w, r)
			return
		}

		userName := r.Context().Value("userName").(string)

		startTime := time.Now()

		// 读取请求主体
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logx.Errorf("Failed to read request body: %v", err)
		}

		// 创建一个新的请求主体用于后续读取
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// 打印请求参数日志
		logx.Infof("Request: %s %s %s", r.Method, uri, body)

		// 继续处理请求
		next(w, r)

		// 记录响应日志
		//response := httpx.Response{ResponseWriter: w}
		//responseData, err := response.DumpBody()
		//if err != nil {
		//	logx.Errorf("Failed to read response body: %v", err)
		//}

		//// 打印响应日志
		//logx.Infof("Response: %s %s %s", r.Method, r.RequestURI, responseData)

		// 打印请求和响应耗时
		duration := time.Since(startTime)
		//// 添加操作日志
		_, _ = m.Sys.SysLogAdd(r.Context(), &sysclient.SysLogAddReq{
			UserName:  userName,
			Operation: r.Method,
			Method:    uri,
			Params:    string(body),
			Time:      duration.Milliseconds(),
			Ip:        httpx.GetRemoteAddr(r),
			CreateBy:  userName,
		})
	}
}
