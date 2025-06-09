package seckillproductservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QuerySeckillProductListLogic 查询秒杀商品列表
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:51
*/
type QuerySeckillProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySeckillProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySeckillProductListLogic {
	return &QuerySeckillProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QuerySeckillProductList 查询秒杀商品列表
func (l *QuerySeckillProductListLogic) QuerySeckillProductList(in *smsclient.QuerySeckillProductListReq) (*smsclient.QuerySeckillProductListResp, error) {
	seckillProduct := query.SmsSeckillProduct
	q := seckillProduct.WithContext(l.ctx)
	if in.ActivityId != 0 {
		q = q.Where(seckillProduct.ActivityID.Eq(in.ActivityId))
	}
	if in.SessionId != 0 {
		q = q.Where(seckillProduct.SessionID.Eq(in.SessionId))
	}

	if in.Status != 2 {
		q = q.Where(seckillProduct.Status.Eq(in.Status))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询秒杀商品列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询秒杀商品列表失败")
	}

	var list []*smsclient.SeckillProductListData

	for _, item := range result {
		list = append(list, &smsclient.SeckillProductListData{
			Id:           item.ID,                                          // ID
			ActivityId:   item.ActivityID,                                  // 活动ID
			SessionId:    item.SessionID,                                   // 秒杀场次ID
			SkuId:        item.SkuID,                                       // 商品SKU ID
			SkuName:      item.SkuName,                                     // 商品名称
			SeckillPrice: float32(item.SeckillPrice),                       // 秒杀价格
			SeckillStock: item.SeckillStock,                                // 秒杀库存
			StockLocked:  item.StockLocked,                                 // 锁定库存
			PerLimit:     item.PerLimit,                                    // 每人限购数量
			Sort:         item.Sort,                                        // 排序
			Status:       item.Status,                                      // 状态：0-未上架，1-已上架
			CreateBy:     item.CreateBy,                                    // 创建人ID
			CreateTime:   time_util.TimeToStr(item.CreateTime),             // 创建时间
			UpdateBy:     pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
			UpdateTime:   time_util.TimeToString(item.UpdateTime),          // 更新时间

		})
	}

	return &smsclient.QuerySeckillProductListResp{
		Total: count,
		List:  list,
	}, nil
}
