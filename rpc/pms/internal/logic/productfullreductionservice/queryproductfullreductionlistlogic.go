package productfullreductionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductFullReductionListLogic 查询产品满减表(只针对同商品)列表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:05
*/
type QueryProductFullReductionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductFullReductionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductFullReductionListLogic {
	return &QueryProductFullReductionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductFullReductionList 查询产品满减表(只针对同商品)列表
func (l *QueryProductFullReductionListLogic) QueryProductFullReductionList(in *pmsclient.QueryProductFullReductionListReq) (*pmsclient.QueryProductFullReductionListResp, error) {
	result, err := query.PmsProductFullReduction.WithContext(l.ctx).Where(query.PmsProductFullReduction.ProductID.Eq(in.ProductId)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询产品满减列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.ProductFullReductionListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductFullReductionListData{
			Id:          item.ID,          //
			ProductId:   item.ProductID,   // 商品id
			FullPrice:   item.FullPrice,   // 商品满多少
			ReducePrice: item.ReducePrice, // 商品减多少
		})
	}

	return &pmsclient.QueryProductFullReductionListResp{
		Total: 0,
		List:  list,
	}, nil

}
