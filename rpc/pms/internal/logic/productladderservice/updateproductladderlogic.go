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

// UpdateProductLadderLogic 更新产品阶梯价格(只针对同商品)
/*
Author: LiuFeiHua
Date: 2024/6/12 17:08
*/
type UpdateProductLadderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductLadderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLadderLogic {
	return &UpdateProductLadderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductLadder 更新产品阶梯价格(只针对同商品)
func (l *UpdateProductLadderLogic) UpdateProductLadder(in *pmsclient.UpdateProductLadderReq) (*pmsclient.UpdateProductLadderResp, error) {
	q := query.PmsProductLadder
	_, err := q.WithContext(l.ctx).Updates(&model.PmsProductLadder{
		ID:        in.Id,
		ProductID: in.ProductId, // 商品id
		Count:     in.Count,     // 满足的商品数量
		Discount:  in.Discount,  // 折扣
		Price:     in.Price,     // 折后价格
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新产品阶梯价格失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新产品阶梯价格失败")
	}

	return &pmsclient.UpdateProductLadderResp{}, nil
}
