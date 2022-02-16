package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/pmsmodel"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *ProductOperateLogUpdateLogic) ProductOperateLogUpdate(in *pms.ProductOperateLogUpdateReq) (*pms.ProductOperateLogUpdateResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	err := l.svcCtx.PmsProductOperateLogModel.Update(pmsmodel.PmsProductOperateLog{
		Id:               in.Id,
		ProductId:        in.ProductId,
		PriceOld:         float64(in.PriceOld),
		PriceNew:         float64(in.PriceNew),
		SalePriceOld:     float64(in.SalePriceOld),
		SalePriceNew:     float64(in.SalePriceNew),
		GiftPointOld:     in.GiftPointOld,
		GiftPointNew:     in.GiftPointNew,
		UsePointLimitOld: in.UsePointLimitOld,
		UsePointLimitNew: in.UsePointLimitNew,
		OperateMan:       in.OperateMan,
		CreateTime:       CreateTime,
	})
	if err != nil {
		return nil, err
	}

	return &pms.ProductOperateLogUpdateResp{}, nil
}
