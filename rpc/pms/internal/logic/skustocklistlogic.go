package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuStockListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSkuStockListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuStockListLogic {
	return &SkuStockListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SkuStockListLogic) SkuStockList(in *pms.SkuStockListReq) (*pms.SkuStockListResp, error) {
	all, err := l.svcCtx.PmsSkuStockModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.PmsSkuStockModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询库存列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*pms.SkuStockListData
	for _, item := range *all {

		list = append(list, &pms.SkuStockListData{
			Id:             item.Id,
			ProductId:      item.ProductId,
			SkuCode:        item.SkuCode,
			Price:          int64(item.Price),
			Stock:          item.Stock,
			LowStock:       item.LowStock,
			Pic:            item.Pic,
			Sale:           item.Sale,
			PromotionPrice: int64(item.PromotionPrice),
			LockStock:      item.LockStock,
			SpData:         item.SpData,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询库存列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &pms.SkuStockListResp{
		Total: count,
		List:  list,
	}, nil
}
