package main

import (
	"flag"
	"fmt"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/config"
	"github.com/feihua/zero-admin/api/admin/internal/handler"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

var configFile = flag.String("f", "api/admin/etc/admin-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	server.Use(ctx.AddLog)

	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	// httpx.SetErrorHandler 仅在调用了 httpx.Error 处理响应时才有效
	//httpx.SetErrorHandler(func(err error) (int, any) {
	//	switch e := err.(type) {
	//	case *errors.CodeMsg:
	//		return http.StatusOK, xhttp.BaseResponse[types.Nil]{
	//			Code: e.Code,
	//			Msg:  e.Msg,
	//		}
	//	default:
	//		return http.StatusInternalServerError, nil
	//	}
	//})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
