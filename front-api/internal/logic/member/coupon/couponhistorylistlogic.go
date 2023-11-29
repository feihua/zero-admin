package coupon

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/sms/smsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CouponHistoryListLogic
/*
Author: LiuFeiHua
Date: 2023/11/29 11:24
*/
type CouponHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponHistoryListLogic {
	return &CouponHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CouponHistoryList 获取用户优惠券列表
func (l *CouponHistoryListLogic) CouponHistoryList(req *types.ListCouponHistoryReq) (resp *types.ListCouponHistoryResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	historyList, err := l.svcCtx.CouponHistoryService.CouponHistoryList(l.ctx, &smsclient.CouponHistoryListReq{
		Current:   1,
		PageSize:  100,
		CouponId:  0,
		MemberId:  memberId,
		UseStatus: req.UseStatus,
	})
	if err != nil {
		return nil, err
	}
	var list []*types.ListCouponHistoryData

	for _, item := range historyList.List {
		list = append(list, &types.ListCouponHistoryData{
			Id:             item.Id,
			CouponId:       item.CouponId,
			MemberId:       item.MemberId,
			CouponCode:     item.CouponCode,
			MemberNickname: item.MemberNickname,
			GetType:        item.GetType,
			CreateTime:     item.CreateTime,
			UseStatus:      item.UseStatus,
			UseTime:        item.UseTime,
			OrderId:        item.OrderId,
			OrderSn:        item.OrderSn,
		})
	}

	return &types.ListCouponHistoryResp{
		Data:    list,
		Code:    0,
		Message: "查询会员优惠券成功",
	}, nil
}
