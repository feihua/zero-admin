package skustockservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

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

func (l *SkuStockListLogic) SkuStockList(in *pmsclient.SkuStockListReq) (*pmsclient.SkuStockListResp, error) {
	result, err := query.PmsSkuStock.WithContext(l.ctx).Where(query.PmsSkuStock.ProductID.Eq(in.ProductId)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询库存列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.SkuStockListData
	for _, item := range result {

		list = append(list, &pmsclient.SkuStockListData{
			Id:             item.ID,
			ProductId:      item.ProductID,
			SkuCode:        item.SkuCode,
			Price:          item.Price,
			Stock:          item.Stock,
			LowStock:       item.LowStock,
			Pic:            item.Pic,
			Sale:           item.Sale,
			PromotionPrice: item.PromotionPrice,
			LockStock:      item.LockStock,
			SpData:         item.SpData,
		})
	}

	logc.Infof(l.ctx, "查询库存列表信息,参数：%+v,响应：%+v", in, list)
	return &pmsclient.SkuStockListResp{
		Total: 0,
		List:  list,
	}, nil
}
