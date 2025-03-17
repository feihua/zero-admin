package coupon_history

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CouponHistoryListLogic 优惠券领取记录管理
/*
Author: LiuFeiHua
Date: 2024/5/14 10:30
*/
type CouponHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) CouponHistoryListLogic {
	return CouponHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CouponHistoryList 根据优惠券id，使用状态，订单编号分页获取领取记录
func (l *CouponHistoryListLogic) CouponHistoryList(req *types.ListCouponHistoryReq) (*types.ListCouponHistoryResp, error) {
	resp, err := l.svcCtx.CouponHistoryService.QueryCouponHistoryList(l.ctx, &smsclient.QueryCouponHistoryListReq{
		PageNum:   req.Current,
		PageSize:  req.PageSize,
		UseStatus: req.UseStatus,
		CouponId:  req.CouponId,
		OrderSn:   req.OrderSn,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询优惠券使用记录列表异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.ListCouponHistoryData

	for _, item := range resp.List {
		list = append(list, &types.ListCouponHistoryData{
			Id:             item.Id,             //
			CouponId:       item.CouponId,       // 优惠券id
			MemberId:       item.MemberId,       // 会员id
			CouponCode:     item.CouponCode,     // 优惠码
			MemberNickname: item.MemberNickname, // 领取人昵称
			GetType:        item.GetType,        // 获取类型：0->后台赠送；1->主动获取
			CreateTime:     item.CreateTime,     // 领取时间
			UseStatus:      item.UseStatus,      // 使用状态：0->未使用；1->已使用；2->已过期
			UseTime:        item.UseTime,        // 使用时间
			OrderId:        item.OrderId,        // 订单编号
			OrderSn:        item.OrderSn,        // 订单号码
		})
	}

	return &types.ListCouponHistoryResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询优惠券使用记录成功",
	}, nil
}
