package productoperatelogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductOperateLogUpdateLogic 商品操作历史
/*
Author: LiuFeiHua
Date: 2024/5/8 11:07
*/
type ProductOperateLogUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductOperateLogUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductOperateLogUpdateLogic {
	return &ProductOperateLogUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductOperateLogUpdate 更新商品操作历史
func (l *ProductOperateLogUpdateLogic) ProductOperateLogUpdate(in *pmsclient.ProductOperateLogUpdateReq) (*pmsclient.ProductOperateLogUpdateResp, error) {
	q := query.PmsProductOperateLog
	_, err := q.WithContext(l.ctx).Updates(&model.PmsProductOperateLog{
		ID:               in.Id,
		ProductID:        in.ProductId,
		PriceOld:         in.PriceOld,
		PriceNew:         in.PriceNew,
		SalePriceOld:     in.SalePriceOld,
		SalePriceNew:     in.SalePriceNew,
		GiftPointOld:     in.GiftPointOld,
		GiftPointNew:     in.GiftPointNew,
		UsePointLimitOld: in.UsePointLimitOld,
		UsePointLimitNew: in.UsePointLimitNew,
		OperateMan:       in.OperateMan,
		CreateTime:       time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductOperateLogUpdateResp{}, nil
}
