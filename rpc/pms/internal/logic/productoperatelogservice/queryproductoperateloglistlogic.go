package productoperatelogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductOperateLogListLogic 查询商品操作日志列表
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

// QueryProductOperateLogList 查询商品操作日志列表
func (l *QueryProductOperateLogListLogic) QueryProductOperateLogList(in *pmsclient.QueryProductOperateLogListReq) (*pmsclient.QueryProductOperateLogListResp, error) {
	q := query.PmsProductOperateLog.WithContext(l.ctx)

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询商品操作日志列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询商品操作日志列表失败")
	}

	var list []*pmsclient.ProductOperateLogListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductOperateLogListData{
			Id:               item.ID,                              //
			ProductId:        item.ProductID,                       // 产品id
			PriceOld:         item.PriceOld,                        // 原价
			PriceNew:         item.PriceNew,                        // 新价格
			SalePriceOld:     item.SalePriceOld,                    // 原售价
			SalePriceNew:     item.SalePriceNew,                    // 新售价
			GiftPointOld:     item.GiftPointOld,                    // 赠送的积分
			GiftPointNew:     item.GiftPointNew,                    // 新的积分
			UsePointLimitOld: item.UsePointLimitOld,                // 使用积分限制
			UsePointLimitNew: item.UsePointLimitNew,                // 新的使用积分限制
			OperateMan:       item.OperateMan,                      // 操作人
			CreateTime:       time_util.TimeToStr(item.CreateTime), // 创建时间
		})
	}

	return &pmsclient.QueryProductOperateLogListResp{
		Total: count,
		List:  list,
	}, nil

}
