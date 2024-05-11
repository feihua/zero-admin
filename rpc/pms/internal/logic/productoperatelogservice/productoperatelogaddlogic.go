package productoperatelogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductOperateLogAddLogic 商品操作历史
/*
Author: LiuFeiHua
Date: 2024/5/8 11:04
*/
type ProductOperateLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductOperateLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductOperateLogAddLogic {
	return &ProductOperateLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductOperateLogAdd 添加商品操作历史
func (l *ProductOperateLogAddLogic) ProductOperateLogAdd(in *pmsclient.ProductOperateLogAddReq) (*pmsclient.ProductOperateLogAddResp, error) {
	err := query.PmsProductOperateLog.WithContext(l.ctx).Create(&model.PmsProductOperateLog{
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

	return &pmsclient.ProductOperateLogAddResp{}, nil
}
