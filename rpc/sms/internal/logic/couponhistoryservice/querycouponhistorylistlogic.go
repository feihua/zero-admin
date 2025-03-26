package couponhistoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCouponHistoryListLogic 根据优惠券id，使用状态，订单编号分页获取领取记录
/*
Author: LiuFeiHua
Date: 2024/6/12 17:28
*/
type QueryCouponHistoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCouponHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponHistoryListLogic {
	return &QueryCouponHistoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCouponHistoryList 根据优惠券id，使用状态，订单编号分页获取领取记录
func (l *QueryCouponHistoryListLogic) QueryCouponHistoryList(in *smsclient.QueryCouponHistoryListReq) (*smsclient.QueryCouponHistoryListResp, error) {
	q := query.SmsCouponHistory.WithContext(l.ctx)

	if in.CouponId != 0 {
		q = q.Where(query.SmsCouponHistory.CouponID.Eq(in.CouponId))
	}
	if in.MemberId != 0 {
		q = q.Where(query.SmsCouponHistory.MemberID.Eq(in.MemberId))
	}
	if len(in.OrderSn) > 0 {
		q = q.Where(query.SmsCouponHistory.OrderSn.Eq(in.OrderSn))
	}
	if in.UseStatus != 3 {
		q = q.Where(query.SmsCouponHistory.UseStatus.Eq(in.UseStatus))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "根据优惠券id，使用状态，订单编号分页获取领取记录失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("根据优惠券id，使用状态，订单编号分页获取领取记录失败")
	}

	var list []*smsclient.CouponHistoryListData
	for _, item := range result {

		list = append(list, &smsclient.CouponHistoryListData{
			Id:             item.ID,                              //
			CouponId:       item.CouponID,                        // 优惠券id
			MemberId:       item.MemberID,                        // 会员id
			CouponCode:     item.CouponCode,                      // 优惠码
			MemberNickname: item.MemberNickname,                  // 领取人昵称
			GetType:        item.GetType,                         // 获取类型：0->后台赠送；1->主动获取
			CreateTime:     time_util.TimeToStr(item.CreateTime), // 领取时间
			UseStatus:      item.UseStatus,                       // 使用状态：0->未使用；1->已使用；2->已过期
			UseTime:        time_util.TimeToStr(item.UseTime),    // 使用时间
			OrderId:        item.OrderID,                         // 订单编号
			OrderSn:        item.OrderSn,                         // 订单号码
		})
	}

	return &smsclient.QueryCouponHistoryListResp{
		Total: count,
		List:  list,
	}, nil

}
