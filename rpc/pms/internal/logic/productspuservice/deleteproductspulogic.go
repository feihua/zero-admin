package productspuservicelogic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductSpuLogic 删除商品SPU
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:38
*/
type DeleteProductSpuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductSpuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductSpuLogic {
	return &DeleteProductSpuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductSpu 删除商品SPU
func (l *DeleteProductSpuLogic) DeleteProductSpu(in *pmsclient.DeleteProductSpuReq) (*pmsclient.DeleteProductSpuResp, error) {
	q := query.PmsProductSpu

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除商品SPU失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除商品SPU失败")
	}

	message := map[string]any{"ids": in.Ids}
	body, _ := json.Marshal(message)
	err = l.svcCtx.RabbitMQ.SendMessage("product.event.exchange", "delete.product.from.es.queue", "delete.product.key", body)

	return &pmsclient.DeleteProductSpuResp{}, nil
}
