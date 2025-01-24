package productoperatelogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
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
		ProductID:        in.ProductId,        // 产品id
		PriceOld:         in.PriceOld,         // 原价
		PriceNew:         in.PriceNew,         // 新价格
		SalePriceOld:     in.SalePriceOld,     // 原售价
		SalePriceNew:     in.SalePriceNew,     // 新售价
		GiftPointOld:     in.GiftPointOld,     // 赠送的积分
		GiftPointNew:     in.GiftPointNew,     // 新的积分
		UsePointLimitOld: in.UsePointLimitOld, // 使用积分限制
		UsePointLimitNew: in.UsePointLimitNew, // 新的使用积分限制
		OperateMan:       in.OperateMan,       // 操作人
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.AddProductOperateLogResp{}, nil
}
