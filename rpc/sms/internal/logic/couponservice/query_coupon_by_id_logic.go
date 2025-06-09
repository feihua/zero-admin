package couponservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCouponByIdLogic 根据商品Id和分类id查询可用的优惠券(app)
/*
Author: LiuFeiHua
Date: 2025/6/11 16:15
*/
type QueryCouponByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCouponByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponByIdLogic {
	return &QueryCouponByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCouponById 根据商品Id和分类id查询可用的优惠券(app)
func (l *QueryCouponByIdLogic) QueryCouponById(in *smsclient.QueryCouponByIdReq) (*smsclient.QueryCouponByIdResp, error) {
	sql := `select t1.*
	from sms_coupon t1,
		 sms_coupon_scope t2
	where t1.id = t2.coupon_id
	  and t2.scope_id = ?`

	var result []model.SmsCoupon
	db := l.svcCtx.DB
	err := db.WithContext(l.ctx).Raw(sql, in.Id).Scan(&result).Error

	if err != nil {
		logc.Errorf(l.ctx, "根据商品Id和分类id查询可用的优惠券(app)失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("根据商品Id和分类id查询可用的优惠券(app)失败")
	}

	var list []*smsclient.CouponListData

	for _, item := range result {
		list = append(list, &smsclient.CouponListData{
			Id:            item.ID,                                 // 优惠券ID
			TypeId:        item.TypeID,                             // 优惠券类型ID
			Name:          item.Name,                               // 优惠券名称
			Code:          item.Code,                               // 优惠券码
			Amount:        float32(item.Amount),                    // 优惠金额/折扣率
			MinAmount:     float32(item.MinAmount),                 // 最低使用金额
			StartTime:     time_util.TimeToStr(item.StartTime),     // 开始时间
			EndTime:       time_util.TimeToStr(item.EndTime),       // 结束时间
			TotalCount:    item.TotalCount,                         // 发放总量
			ReceivedCount: item.ReceivedCount,                      // 已领取数量
			UsedCount:     item.UsedCount,                          // 已使用数量
			PerLimit:      item.PerLimit,                           // 每人限领数量
			Status:        item.Status,                             // 状态：0-未开始，1-进行中，2-已结束，3-已取消
			IsEnabled:     item.IsEnabled,                          // 是否启用
			Description:   item.Description,                        // 使用说明
			CreateBy:      item.CreateBy,                           // 创建人ID
			CreateTime:    time_util.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:      *item.UpdateBy,                          // 更新人ID
			UpdateTime:    time_util.TimeToString(item.UpdateTime), // 更新时间

		})
	}

	return &smsclient.QueryCouponByIdResp{
		List: list,
	}, nil
}
