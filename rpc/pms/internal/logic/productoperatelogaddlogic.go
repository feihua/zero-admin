package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/pmsmodel"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *ProductOperateLogAddLogic) ProductOperateLogAdd(in *pms.ProductOperateLogAddReq) (*pms.ProductOperateLogAddResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	_, err := l.svcCtx.PmsProductOperateLogModel.Insert(pmsmodel.PmsProductOperateLog{
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

	return &pms.ProductOperateLogAddResp{}, nil
}
