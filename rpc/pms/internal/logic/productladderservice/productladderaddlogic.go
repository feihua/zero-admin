package productladderservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductLadderAddLogic 产品阶梯价格
/*
Author: LiuFeiHua
Date: 2024/5/8 11:02
*/
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

// ProductLadderAdd 添加产品阶梯价格
func (l *ProductLadderAddLogic) ProductLadderAdd(in *pmsclient.ProductLadderAddReq) (*pmsclient.ProductLadderAddResp, error) {
	err := query.PmsProductLadder.WithContext(l.ctx).Create(&model.PmsProductLadder{
		ProductID: in.ProductId,
		Count:     in.Count,
		Discount:  in.Discount,
		Price:     in.Price,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductLadderAddResp{}, nil
}
