package seckillproductservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySeckillProductBySkuIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySeckillProductBySkuIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySeckillProductBySkuIdLogic {
	return &QuerySeckillProductBySkuIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QuerySeckillProductBySkuId 查询秒杀商品详情(app)
func (l *QuerySeckillProductBySkuIdLogic) QuerySeckillProductBySkuId(in *smsclient.QuerySeckillProductBySkuIdReq) (*smsclient.QuerySeckillProductDetailResp, error) {
	item, err := query.SmsSeckillProduct.WithContext(l.ctx).Where(query.SmsSeckillProduct.SkuID.Eq(in.SkuId)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "秒杀商品不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("秒杀商品不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询秒杀商品异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询秒杀商品异常")
	}

	data := &smsclient.QuerySeckillProductDetailResp{
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
	}

	return data, nil
}
