package productladderservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductLadderDetailLogic 查询产品阶梯价格(只针对同商品)详情
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:06
*/
type QueryProductLadderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductLadderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductLadderDetailLogic {
	return &QueryProductLadderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductLadderDetail 查询产品阶梯价格(只针对同商品)详情
func (l *QueryProductLadderDetailLogic) QueryProductLadderDetail(in *pmsclient.QueryProductLadderDetailReq) (*pmsclient.QueryProductLadderDetailResp, error) {
	item, err := query.PmsProductLadder.WithContext(l.ctx).Where(query.PmsProductLadder.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询产品阶梯价格(只针对同商品)详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询产品阶梯价格(只针对同商品)详情失败")
	}

	data := &pmsclient.QueryProductLadderDetailResp{
		Id:        item.ID,        //
		ProductId: item.ProductID, // 商品id
		Count:     item.Count,     // 满足的商品数量
		Discount:  item.Discount,  // 折扣
		Price:     item.Price,     // 折后价格
	}

	return data, nil
}
