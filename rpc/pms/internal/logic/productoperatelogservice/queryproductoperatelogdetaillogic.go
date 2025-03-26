package productoperatelogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductOperateLogDetailLogic 查询商品操作日志详情
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:06
*/
type QueryProductOperateLogDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductOperateLogDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductOperateLogDetailLogic {
	return &QueryProductOperateLogDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductOperateLogDetail 查询商品操作日志详情
func (l *QueryProductOperateLogDetailLogic) QueryProductOperateLogDetail(in *pmsclient.QueryProductOperateLogDetailReq) (*pmsclient.QueryProductOperateLogDetailResp, error) {
	item, err := query.PmsProductOperateLog.WithContext(l.ctx).Where(query.PmsProductOperateLog.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询商品操作日志详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询商品操作日志详情失败")
	}

	data := &pmsclient.QueryProductOperateLogDetailResp{
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
	}

	return data, nil
}
