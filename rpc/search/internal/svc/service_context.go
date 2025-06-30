package svc

import (
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/feihua/zero-admin/pkg/es"
	"github.com/feihua/zero-admin/rpc/search/internal/config"
)

const IndexName = "products"

type ServiceContext struct {
	Config   config.Config
	ESClient *elasticsearch.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		ESClient: es.GetESClient(c.Es.Host),
	}
}
