package productoperatelogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductOperateLogListLogic 查询列表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:10
*/
type QueryProductOperateLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductOperateLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductOperateLogListLogic {
	return &QueryProductOperateLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductOperateLogList 查询列表
func (l *QueryProductOperateLogListLogic) QueryProductOperateLogList(in *pmsclient.QueryProductOperateLogListReq) (*pmsclient.QueryProductOperateLogListResp, error) {
	q := query.PmsProductOperateLog.WithContext(l.ctx)

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询商品操作历史列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.ProductOperateLogListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductOperateLogListData{
			Id:               item.ID,
			ProductId:        item.ProductID,
			PriceOld:         item.PriceOld,
			PriceNew:         item.PriceNew,
			SalePriceOld:     item.SalePriceOld,
			SalePriceNew:     item.SalePriceNew,
			GiftPointOld:     item.GiftPointOld,
			GiftPointNew:     item.GiftPointNew,
			UsePointLimitOld: item.UsePointLimitOld,
			UsePointLimitNew: item.UsePointLimitNew,
			OperateMan:       item.OperateMan,
			CreateTime:       item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &pmsclient.QueryProductOperateLogListResp{
		Total: count,
		List:  list,
	}, nil

}
