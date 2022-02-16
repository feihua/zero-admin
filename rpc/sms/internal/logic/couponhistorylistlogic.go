package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *CouponHistoryListLogic) CouponHistoryList(in *sms.CouponHistoryListReq) (*sms.CouponHistoryListResp, error) {
	all, err := l.svcCtx.SmsCouponHistoryModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.SmsCouponHistoryModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询优惠券使用历史列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*sms.CouponHistoryListData
	for _, item := range *all {

		list = append(list, &sms.CouponHistoryListData{
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
	logx.WithContext(l.ctx).Infof("查询优惠券使用历史列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &sms.CouponHistoryListResp{
		Total: count,
		List:  list,
	}, nil
}
