package logic

import (
	"context"
	"go-zero-admin/rpc/model/pmsmodel"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductLadderAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLadderAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLadderAddLogic {
	return &ProductLadderAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductLadderAddLogic) ProductLadderAdd(in *pms.ProductLadderAddReq) (*pms.ProductLadderAddResp, error) {
	_, err := l.svcCtx.PmsProductLadderModel.Insert(pmsmodel.PmsProductLadder{
		ProductId: in.ProductId,
		Count:     in.Count,
		Discount:  float64(in.Discount),
		Price:     float64(in.Price),
	})
	if err != nil {
		return nil, err
	}

	return &pms.ProductLadderAddResp{}, nil
}
