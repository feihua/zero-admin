// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4

package handler

import (
	"net/http"

	auth "github.com/feihua/zero-admin/consumer/internal/handler/auth"
	order "github.com/feihua/zero-admin/consumer/internal/handler/order"
	product "github.com/feihua/zero-admin/consumer/internal/handler/product"
	"github.com/feihua/zero-admin/consumer/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/consumer/auth",
				Handler: auth.ConsumerAuthHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/testOrder",
				Handler: order.TestOrderHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/consumer"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/addProductToEs",
				Handler: product.AddProductToEsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/deleteProductFromEs",
				Handler: product.DeleteProductFromEsHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/consumer"),
	)
}
