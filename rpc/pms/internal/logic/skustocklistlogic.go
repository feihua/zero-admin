package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
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
	all, _ := l.svcCtx.PmsSkuStockModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

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

	fmt.Println(list)
	return &pms.SkuStockListResp{
		Total: 10,
		List:  list,
	}, nil
}
