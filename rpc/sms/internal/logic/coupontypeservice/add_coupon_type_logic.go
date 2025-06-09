package coupontypeservicelogic

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
)

// AddCouponTypeLogic 添加优惠券类型
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type AddCouponTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCouponTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCouponTypeLogic {
	return &AddCouponTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddCouponType 添加优惠券类型
func (l *AddCouponTypeLogic) AddCouponType(in *smsclient.AddCouponTypeReq) (*smsclient.AddCouponTypeResp, error) {
	q := query.SmsCouponType
	count, err := q.WithContext(l.ctx).Where(q.Name.Eq(in.Name)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "添加优惠券类型失败,参数：%+v, 异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("添加优惠券类型失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("优惠券类型：%s,已存在", in.Name))
	}
	item := &model.SmsCouponType{
		Name:         in.Name,         // 类型名称
		Code:         in.Code,         // 类型编码
		Description:  in.Description,  // 描述
		DiscountType: in.DiscountType, // 优惠方式：1-固定金额，2-折扣率，3-第N件特惠，4-买赠，5-特价，6-套装优惠，7-搭配优惠，8-积分抵现，9-积分倍率，10-免运费，11-运费减免，12-限时特权，13-会员特权
		Status:       in.Status,       // 是否启用
		Sort:         in.Sort,         // 排序
		CreateBy:     in.CreateBy,     // 创建人ID
	}

	err = q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加优惠券类型失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加优惠券类型失败")
	}

	return &smsclient.AddCouponTypeResp{}, nil
}
