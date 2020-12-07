package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductOperateLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductOperateLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductOperateLogListLogic {
	return &ProductOperateLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductOperateLogListLogic) ProductOperateLogList(in *pms.ProductOperateLogListReq) (*pms.ProductOperateLogListResp, error) {
	all, _ := l.svcCtx.PmsProductOperateLogModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	var list []*pms.ProductOperateLogListData
	for _, item := range *all {

		list = append(list, &pms.ProductOperateLogListData{
			Id:               item.Id,
			ProductId:        item.ProductId,
			PriceOld:         int64(item.PriceOld),
			PriceNew:         int64(item.PriceNew),
			SalePriceOld:     int64(item.SalePriceOld),
			SalePriceNew:     int64(item.SalePriceNew),
			GiftPointOld:     item.GiftPointOld,
			GiftPointNew:     item.GiftPointNew,
			UsePointLimitOld: item.UsePointLimitOld,
			UsePointLimitNew: item.UsePointLimitNew,
			OperateMan:       item.OperateMan,
			CreateTime:       item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	fmt.Println(list)
	return &pms.ProductOperateLogListResp{
		Total: 10,
		List:  list,
	}, nil
}
