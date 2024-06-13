package productoperatelogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddProductOperateLogLogic 添加
/*
Author: LiuFeiHua
Date: 2024/6/12 17:09
*/
type AddProductOperateLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductOperateLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductOperateLogLogic {
	return &AddProductOperateLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddProductOperateLog 添加
func (l *AddProductOperateLogLogic) AddProductOperateLog(in *pmsclient.AddProductOperateLogReq) (*pmsclient.AddProductOperateLogResp, error) {
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

	return &pmsclient.AddProductOperateLogResp{}, nil
}
