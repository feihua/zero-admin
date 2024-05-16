package coupon

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCouponListLogic 查询会员优惠券成功
/*
Author: LiuFeiHua
Date: 2024/5/16 16:30
*/
type QueryCouponListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCouponListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponListLogic {
	return &QueryCouponListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryCouponList 获取用户优惠券列表
func (l *QueryCouponListLogic) QueryCouponList(req *types.ListCouponReq) (resp *types.ListCouponResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	couponList, err := l.svcCtx.CouponHistoryService.QueryMemberCouponList(l.ctx, &smsclient.QueryMemberCouponListReq{
		MemberId:  memberId,
		UseStatus: req.UseStatus,
	})
	if err != nil {
		return nil, err
	}
	var list []*types.ListCouponData

	for _, item := range couponList.List {
		list = append(list, &types.ListCouponData{
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

	return &types.ListCouponResp{
		Data:    list,
		Code:    0,
		Message: "查询会员优惠券成功",
	}, nil
}
