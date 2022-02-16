package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

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

func (l *ProductFullReductionListLogic) ProductFullReductionList(in *pms.ProductFullReductionListReq) (*pms.ProductFullReductionListResp, error) {
	all, err := l.svcCtx.PmsProductFullReductionModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.PmsProductFullReductionModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询产品满减列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*pms.ProductFullReductionListData
	for _, item := range *all {

		list = append(list, &pms.ProductFullReductionListData{
			Id:          item.Id,
			ProductId:   item.ProductId,
			FullPrice:   int64(item.FullPrice),
			ReducePrice: int64(item.ReducePrice),
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询产品满减列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &pms.ProductFullReductionListResp{
		Total: count,
		List:  list,
	}, nil
}
