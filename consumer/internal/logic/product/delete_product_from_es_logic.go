package product

import (
	"context"
	"encoding/json"

	"github.com/feihua/zero-admin/consumer/internal/svc"
	"github.com/feihua/zero-admin/consumer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductFromEsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductFromEsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductFromEsLogic {
	return &DeleteProductFromEsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteProductFromEs 测试删除商品
func (l *DeleteProductFromEsLogic) DeleteProductFromEs(req *types.ProductEsReq) (resp *types.Response, err error) {
	// message := map[string]any{"ids": nil}// 清空所有商品索引
	message := map[string]any{"ids": req.Ids} // 根据id删除商品索引
	body, _ := json.Marshal(message)
	err = l.svcCtx.RabbitMQ.SendMessage("syn.product.to.es.exchange", "delete.product.from.es.queue", "delete.product.from.es.queue", body)

	return &types.Response{
		Message: "从es删除商品索引成功",
	}, nil
}
