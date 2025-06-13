package productoperatelogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddProductOperateLogLogic 添加商品操作日志
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

// AddProductOperateLog 添加商品操作日志
func (l *AddProductOperateLogLogic) AddProductOperateLog(in *pmsclient.AddProductOperateLogReq) (*pmsclient.AddProductOperateLogResp, error) {
	err := l.svcCtx.ProductOperateLogModel.Insert(l.ctx, &model.ProductOperateLog{
		ProductId:        in.ProductId,        // 产品id
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
		logc.Errorf(l.ctx, "添加商品操作日志失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加商品操作日志失败")
	}

	return &pmsclient.AddProductOperateLogResp{}, nil
}
