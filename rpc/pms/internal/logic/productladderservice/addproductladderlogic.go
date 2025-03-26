package productladderservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddProductLadderLogic 添加产品阶梯价格(只针对同商品)
/*
Author: LiuFeiHua
Date: 2024/6/12 17:07
*/
type AddProductLadderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductLadderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductLadderLogic {
	return &AddProductLadderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddProductLadder 添加产品阶梯价格(只针对同商品)
func (l *AddProductLadderLogic) AddProductLadder(in *pmsclient.AddProductLadderReq) (*pmsclient.AddProductLadderResp, error) {
	err := query.PmsProductLadder.WithContext(l.ctx).Create(&model.PmsProductLadder{
		ProductID: in.ProductId, // 商品id
		Count:     in.Count,     // 满足的商品数量
		Discount:  in.Discount,  // 折扣
		Price:     in.Price,     // 折后价格
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加产品阶梯价格失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加产品阶梯价格失败")
	}

	return &pmsclient.AddProductLadderResp{}, nil
}
