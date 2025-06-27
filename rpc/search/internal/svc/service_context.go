package svc

import "github.com/feihua/zero-admin/rpc/search/internal/config"

const IndexName = "products"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
