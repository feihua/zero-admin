package productfullreductionservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

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
	all, err := l.svcCtx.PmsProductFullReductionModel.FindAll(l.ctx, in.ProductId)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询产品满减列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*pmsclient.ProductFullReductionListData
	for _, item := range *all {

		list = append(list, &pmsclient.ProductFullReductionListData{
			Id:          item.Id,
			ProductId:   item.ProductId,
			FullPrice:   float32(item.FullPrice),
			ReducePrice: float32(item.ReducePrice),
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询产品满减列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &pmsclient.ProductFullReductionListResp{
		Total: 0,
		List:  list,
	}, nil
}
