package couponhistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// CouponHistoryListLogic
/*
Author: LiuFeiHua
Date: 2023/11/29 11:33
*/
type CouponHistoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponHistoryListLogic {
	return &CouponHistoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CouponHistoryList 查询会员的优惠券
func (l *CouponHistoryListLogic) CouponHistoryList(in *smsclient.CouponHistoryListReq) (*smsclient.CouponHistoryListResp, error) {
	q := query.SmsCouponHistory.WithContext(l.ctx)

	if in.CouponId != 0 {
		q = q.Where(query.SmsCouponHistory.CouponID.Eq(in.CouponId))
	}
	if in.MemberId != 0 {
		q = q.Where(query.SmsCouponHistory.MemberID.Eq(in.MemberId))
	}
	if in.UseStatus != 3 {
		q = q.Where(query.SmsCouponHistory.UseStatus.Eq(in.UseStatus))
	}

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询优惠券使用历史列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*smsclient.CouponHistoryListData
	for _, item := range result {

		list = append(list, &smsclient.CouponHistoryListData{
			Id:             item.ID,
			CouponId:       item.CouponID,
			MemberId:       item.MemberID,
			CouponCode:     item.CouponCode,
			MemberNickname: item.MemberNickname,
			GetType:        item.GetType,
			CreateTime:     item.CreateTime.Format("2006-01-02 15:04:05"),
			UseStatus:      item.UseStatus,
			UseTime:        item.UseTime.Format("2006-01-02 15:04:05"),
			OrderId:        item.OrderID,
			OrderSn:        item.OrderSn,
		})
	}

	logc.Infof(l.ctx, "查询优惠券使用历史列表信息,参数：%+v,响应：%+v", in, list)
	return &smsclient.CouponHistoryListResp{
		Total: count,
		List:  list,
	}, nil
}
