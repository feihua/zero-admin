package productoperatelogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

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

func (l *ProductOperateLogListLogic) ProductOperateLogList(in *pmsclient.ProductOperateLogListReq) (*pmsclient.ProductOperateLogListResp, error) {
	q := query.PmsProductOperateLog.WithContext(l.ctx)

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询商品操作历史列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.ProductOperateLogListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductOperateLogListData{
			Id:               item.ID,
			ProductId:        item.ProductID,
			PriceOld:         float32(item.PriceOld),
			PriceNew:         float32(item.PriceNew),
			SalePriceOld:     float32(item.SalePriceOld),
			SalePriceNew:     float32(item.SalePriceNew),
			GiftPointOld:     item.GiftPointOld,
			GiftPointNew:     item.GiftPointNew,
			UsePointLimitOld: item.UsePointLimitOld,
			UsePointLimitNew: item.UsePointLimitNew,
			OperateMan:       item.OperateMan,
			CreateTime:       item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	logc.Infof(l.ctx, "查询商品操作历史列表信息,参数：%+v,响应：%+v", in, list)
	return &pmsclient.ProductOperateLogListResp{
		Total: count,
		List:  list,
	}, nil
}
