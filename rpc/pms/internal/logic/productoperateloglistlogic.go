package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
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
	all, err := l.svcCtx.PmsProductOperateLogModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.PmsProductOperateLogModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询商品操作历史列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

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

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询商品操作历史列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &pms.ProductOperateLogListResp{
		Total: count,
		List:  list,
	}, nil
}
