package coupon

import (
	"context"
	"zero-admin/rpc/sms/smsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *CouponHistoryListLogic) CouponHistoryList(req *types.ListCouponHistoryReq) (resp *types.ListCouponHistoryResp, err error) {
	historyList, err := l.svcCtx.CouponHistoryService.CouponHistoryList(l.ctx, &smsclient.CouponHistoryListReq{
		Current:   req.Current,
		PageSize:  req.PageSize,
		CouponId:  0,
		MemberId:  l.ctx.Value("memberId").(int64),
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
		Message: "查询优惠券使用记录成功",
	}, nil
}
