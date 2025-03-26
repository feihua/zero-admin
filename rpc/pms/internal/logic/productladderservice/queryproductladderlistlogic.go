package productladderservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductLadderListLogic 查询产品阶梯价格(只针对同商品)列表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:08
*/
type QueryProductLadderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductLadderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductLadderListLogic {
	return &QueryProductLadderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductLadderList 查询产品阶梯价格(只针对同商品)列表
func (l *QueryProductLadderListLogic) QueryProductLadderList(in *pmsclient.QueryProductLadderListReq) (*pmsclient.QueryProductLadderListResp, error) {
	result, err := query.PmsProductLadder.WithContext(l.ctx).Where(query.PmsProductLadder.ProductID.Eq(in.ProductId)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询产品阶梯价格失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询产品阶梯价格失败")
	}

	var list []*pmsclient.ProductLadderListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductLadderListData{
			Id:        item.ID,        //
			ProductId: item.ProductID, // 商品id
			Count:     item.Count,     // 满足的商品数量
			Discount:  item.Discount,  // 折扣
			Price:     item.Price,     // 折后价格
		})
	}

	return &pmsclient.QueryProductLadderListResp{
		Total: 0,
		List:  list,
	}, nil

}
