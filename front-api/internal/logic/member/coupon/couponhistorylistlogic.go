package coupon

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/feihua/zero-admin/front-api/internal/types"

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
	historyList, err := l.svcCtx.CouponHistoryService.QueryMemberCouponList(l.ctx, &smsclient.QueryMemberCouponListReq{
		MemberId:  memberId,
		UseStatus: req.UseStatus,
	})
	if err != nil {
		return nil, err
	}
	var list []*types.ListCouponHistoryData

	for _, item := range historyList.List {
		list = append(list, &types.ListCouponHistoryData{
			Id:           item.Id,
			Type:         item.Type,
			Name:         item.Name,
			Platform:     item.Platform,
			Count:        item.Count,
			Amount:       item.Amount,
			PerLimit:     item.PerLimit,
			MinPoint:     item.MinPoint,
			StartTime:    item.StartTime,
			EndTime:      item.EndTime,
			UseType:      item.UseType,
			Note:         item.Note,
			PublishCount: item.PublishCount,
			UseCount:     item.UseCount,
			ReceiveCount: item.ReceiveCount,
			EnableTime:   item.EnableTime,
			Code:         item.Code,
			MemberLevel:  item.MemberLevel,
		})
	}

	return &types.ListCouponHistoryResp{
		Data:    list,
		Code:    0,
		Message: "查询会员优惠券成功",
	}, nil
}
