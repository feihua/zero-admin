package productladderservicelogic

import (
	"context"
	"zero-admin/rpc/model/pmsmodel"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/rpc/pms/internal/svc"

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

func (l *ProductLadderAddLogic) ProductLadderAdd(in *pmsclient.ProductLadderAddReq) (*pmsclient.ProductLadderAddResp, error) {
	_, err := l.svcCtx.PmsProductLadderModel.Insert(l.ctx, &pmsmodel.PmsProductLadder{
		ProductId: in.ProductId,
		Count:     in.Count,
		Discount:  float64(in.Discount),
		Price:     float64(in.Price),
	})
	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductLadderAddResp{}, nil
}
