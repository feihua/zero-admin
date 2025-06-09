package couponservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

// AddCouponLogic 添加优惠券
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type AddCouponLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCouponLogic {
	return &AddCouponLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddCoupon 添加优惠券
func (l *AddCouponLogic) AddCoupon(in *smsclient.AddCouponReq) (*smsclient.AddCouponResp, error) {
	q := query.SmsCoupon

	count, err := query.SmsCoupon.WithContext(l.ctx).Where(query.SmsCoupon.Name.Eq(in.Name)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "添加优惠券失败,参数：%+v, 异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("添加优惠券失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("优惠券：%s,已存在", in.Name))
	}

	startTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	endTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	item := &model.SmsCoupon{
		TypeID:      in.TypeId,             // 优惠券类型ID
		Name:        in.Name,               // 优惠券名称
		Code:        in.Code,               // 优惠券码
		Amount:      float64(in.Amount),    // 优惠金额/折扣率
		MinAmount:   float64(in.MinAmount), // 最低使用金额
		StartTime:   startTime,             // 生效时间
		EndTime:     endTime,               // 失效时间
		TotalCount:  in.TotalCount,         // 发放总量
		PerLimit:    in.PerLimit,           // 每人限领数量
		Status:      in.Status,             // 状态：0-未开始，1-进行中，2-已结束，3-已取消
		IsEnabled:   in.IsEnabled,          // 是否启用
		Description: in.Description,        // 使用说明
		CreateBy:    in.CreateBy,           // 创建人ID

	}

	err = q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加优惠券失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加优惠券失败")
	}

	return &smsclient.AddCouponResp{}, nil
}
