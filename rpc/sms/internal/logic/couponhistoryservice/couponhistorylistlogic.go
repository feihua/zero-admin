package couponhistoryservicelogic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logc"
	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/smsclient"

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
	all, err := l.svcCtx.SmsCouponHistoryModel.FindAll(l.ctx, in)
	count, _ := l.svcCtx.SmsCouponHistoryModel.Count(l.ctx, in)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logc.Errorf(l.ctx, "查询优惠券使用历史列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*smsclient.CouponHistoryListData
	for _, item := range *all {

		list = append(list, &smsclient.CouponHistoryListData{
			Id:             item.Id,
			CouponId:       item.CouponId,
			MemberId:       item.MemberId,
			CouponCode:     item.CouponCode,
			MemberNickname: item.MemberNickname,
			GetType:        item.GetType,
			CreateTime:     item.CreateTime.Format("2006-01-02 15:04:05"),
			UseStatus:      item.UseStatus,
			UseTime:        item.UseTime.Format("2006-01-02 15:04:05"),
			OrderId:        item.OrderId,
			OrderSn:        item.OrderSn,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logc.Infof(l.ctx, "查询优惠券使用历史列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &smsclient.CouponHistoryListResp{
		Total: count,
		List:  list,
	}, nil
}
