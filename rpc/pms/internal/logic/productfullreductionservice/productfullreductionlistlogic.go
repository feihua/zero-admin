package productfullreductionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductFullReductionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductFullReductionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductFullReductionListLogic {
	return &ProductFullReductionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductFullReductionListLogic) ProductFullReductionList(in *pmsclient.ProductFullReductionListReq) (*pmsclient.ProductFullReductionListResp, error) {
	result, err := query.PmsProductFullReduction.WithContext(l.ctx).Where(query.PmsProductFullReduction.ProductID.Eq(in.ProductId)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询产品满减列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.ProductFullReductionListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductFullReductionListData{
			Id:          item.ID,
			ProductId:   item.ProductID,
			FullPrice:   item.FullPrice,
			ReducePrice: item.ReducePrice,
		})
	}

	logc.Infof(l.ctx, "查询产品满减列表信息,参数：%+v,响应：%+v", in, list)
	return &pmsclient.ProductFullReductionListResp{
		Total: 0,
		List:  list,
	}, nil
}
