package couponhistoryservicelogic

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

// QueryCouponHistoryDetailLogic 查询优惠券使用、领取历史详情
/*
Author: LiuFeiHua
Date: 2025/01/23 16:18:56
*/
type QueryCouponHistoryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCouponHistoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponHistoryDetailLogic {
	return &QueryCouponHistoryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCouponHistoryDetail 查询优惠券使用、领取历史详情
func (l *QueryCouponHistoryDetailLogic) QueryCouponHistoryDetail(in *smsclient.QueryCouponHistoryDetailReq) (*smsclient.QueryCouponHistoryDetailResp, error) {
	item, err := query.SmsCouponHistory.WithContext(l.ctx).Where(query.SmsCouponHistory.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询优惠券使用、领取历史详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询优惠券使用、领取历史详情失败")
	}

	data := &smsclient.QueryCouponHistoryDetailResp{
		Id:             item.ID,                                          //
		CouponId:       item.CouponID,                                    // 优惠券id
		MemberId:       item.MemberID,                                    // 会员id
		CouponCode:     item.CouponCode,                                  // 优惠码
		MemberNickname: item.MemberNickname,                              // 领取人昵称
		GetType:        item.GetType,                                     // 获取类型：0->后台赠送；1->主动获取
		CreateTime:     time_util.TimeToStr(item.CreateTime),             // 领取时间
		UseStatus:      item.UseStatus,                                   // 使用状态：0->未使用；1->已使用；2->已过期
		UseTime:        time_util.TimeToString(item.UseTime),             // 使用时间
		OrderId:        pointerprocess.DefaltData(item.OrderID).(int64),  // 订单编号
		OrderSn:        pointerprocess.DefaltData(item.OrderSn).(string), // 订单号码
	}

	logc.Infof(l.ctx, "查询优惠券使用、领取历史详情,参数：%+v,响应：%+v", in, data)
	return data, nil
}
