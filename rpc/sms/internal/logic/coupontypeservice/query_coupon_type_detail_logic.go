package coupontypeservicelogic

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
	"gorm.io/gorm"
)

// QueryCouponTypeDetailLogic 查询优惠券类型详情
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type QueryCouponTypeDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCouponTypeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponTypeDetailLogic {
	return &QueryCouponTypeDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCouponTypeDetail 查询优惠券类型详情
func (l *QueryCouponTypeDetailLogic) QueryCouponTypeDetail(in *smsclient.QueryCouponTypeDetailReq) (*smsclient.QueryCouponTypeDetailResp, error) {
	item, err := query.SmsCouponType.WithContext(l.ctx).Where(query.SmsCouponType.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "优惠券类型不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("优惠券类型不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询优惠券类型异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询优惠券类型异常")
	}

	data := &smsclient.QueryCouponTypeDetailResp{
		Id:           item.ID,                                          // 主键ID
		Name:         item.Name,                                        // 类型名称
		Code:         item.Code,                                        // 类型编码
		Description:  item.Description,                                 // 描述
		DiscountType: item.DiscountType,                                // 优惠方式：1-固定金额，2-折扣率，3-第N件特惠，4-买赠，5-特价，6-套装优惠，7-搭配优惠，8-积分抵现，9-积分倍率，10-免运费，11-运费减免，12-限时特权，13-会员特权
		Status:       item.Status,                                      // 是否启用
		Sort:         item.Sort,                                        // 排序
		CreateBy:     item.CreateBy,                                    // 创建人ID
		CreateTime:   time_util.TimeToStr(item.CreateTime),             // 创建时间
		UpdateBy:     pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
		UpdateTime:   time_util.TimeToString(item.UpdateTime),          // 更新时间
	}

	return data, nil
}
