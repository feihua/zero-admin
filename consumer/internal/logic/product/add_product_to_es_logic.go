package product

import (
	"context"
	"encoding/json"

	"github.com/feihua/zero-admin/consumer/internal/svc"
	"github.com/feihua/zero-admin/consumer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductToEsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProductToEsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductToEsLogic {
	return &AddProductToEsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddProductToEs 同步商品到es
func (l *AddProductToEsLogic) AddProductToEs(req *types.ProductEsReq) (resp *types.Response, err error) {
	for _, id := range req.Ids {
		message := map[string]any{"id": id}
		body, _ := json.Marshal(message)
		err = l.svcCtx.RabbitMQ.SendMessage("syn.product.to.es.exchange", "syn.product.to.es.queue", "syn.product.to.es.queue", body)

	}

	return &types.Response{
		Message: "同步商品到es成功",
	}, nil
}
