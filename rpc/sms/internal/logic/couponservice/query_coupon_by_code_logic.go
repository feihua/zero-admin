package couponservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCouponByCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCouponByCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponByCodeLogic {
	return &QueryCouponByCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCouponByCode 根据优惠券类型的code查询优惠券
func (l *QueryCouponByCodeLogic) QueryCouponByCode(in *smsclient.QueryCouponByCodeReq) (*smsclient.QueryCouponByCodeResp, error) {
	sql := `select t1.*
	from sms_coupon t1,
		 sms_coupon_type t2
	where t1.type_id = t2.id
	  and t1.status=1
	  and t2.status=1
	  and t2.code = ?`

	var result []model.SmsCoupon
	db := l.svcCtx.DB
	err := db.WithContext(l.ctx).Raw(sql, in.Code).Scan(&result).Error

	if err != nil {
		logc.Errorf(l.ctx, "根据优惠券类型的code查询优惠券失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("根据优惠券类型的code查询优惠券失败")
	}

	var list []*smsclient.CouponListData

	for _, item := range result {
		list = append(list, &smsclient.CouponListData{
			Id:            item.ID,                                          // 优惠券ID
			TypeId:        item.TypeID,                                      // 优惠券类型ID
			Name:          item.Name,                                        // 优惠券名称
			Code:          item.Code,                                        // 优惠券码
			Amount:        float32(item.Amount),                             // 优惠金额/折扣率
			MinAmount:     float32(item.MinAmount),                          // 最低使用金额
			StartTime:     time_util.TimeToStr(item.StartTime),              // 开始时间
			EndTime:       time_util.TimeToStr(item.EndTime),                // 结束时间
			TotalCount:    item.TotalCount,                                  // 发放总量
			ReceivedCount: item.ReceivedCount,                               // 已领取数量
			UsedCount:     item.UsedCount,                                   // 已使用数量
			PerLimit:      item.PerLimit,                                    // 每人限领数量
			Status:        item.Status,                                      // 状态：0-未开始，1-进行中，2-已结束，3-已取消
			IsEnabled:     item.IsEnabled,                                   // 是否启用
			Description:   item.Description,                                 // 使用说明
			CreateBy:      item.CreateBy,                                    // 创建人ID
			CreateTime:    time_util.TimeToStr(item.CreateTime),             // 创建时间
			UpdateBy:      pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
			UpdateTime:    time_util.TimeToString(item.UpdateTime),          // 更新时间

		})
	}

	return &smsclient.QueryCouponByCodeResp{
		List: list,
	}, nil

}
