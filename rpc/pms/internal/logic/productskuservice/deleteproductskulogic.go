package productskuservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductSkuLogic 删除商品SKU
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type DeleteProductSkuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductSkuLogic {
	return &DeleteProductSkuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductSku 删除商品SKU
func (l *DeleteProductSkuLogic) DeleteProductSku(in *pmsclient.DeleteProductSkuReq) (*pmsclient.DeleteProductSkuResp, error) {
	q := query.PmsProductSku

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除商品SKU失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除商品SKU失败")
	}

	return &pmsclient.DeleteProductSkuResp{}, nil
}
