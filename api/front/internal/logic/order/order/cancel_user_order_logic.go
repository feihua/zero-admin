package order

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CancelUserOrderLogic 取消订单
/*
Author: LiuFeiHua
Date: 2025/6/20 9:41
*/
type CancelUserOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelUserOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelUserOrderLogic {
	return &CancelUserOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CancelUserOrder 取消订单
func (l *CancelUserOrderLogic) CancelUserOrder(req *types.CancelUserOrderReq) (resp1 *types.CancelUserOrderResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}

	// todo 暂时没有分布式事务
	// 1.查询订单是否存在
	// 2.修改订单状态
	resp, err := l.svcCtx.OrderService.CancelOrder(l.ctx, &omsclient.CancelOrderReq{
		MemberId: memberId,
		OrderId:  req.OrderId,
	})
	if err != nil {
		return nil, err
	}

	couponIds := resp.CouponIds
	integration := resp.Integration
	stockLockData := resp.Data

	var data []*pmsclient.UpdateSkuStockData
	for _, item := range stockLockData {
		data = append(data, &pmsclient.UpdateSkuStockData{
			Id:              item.ProductSkuId,
			ProductQuantity: item.ProductQuantity,
		})
	}
	// 3.释放库存
	_, err = l.svcCtx.ProductSkuService.ReleaseSkuStockLock(l.ctx, &pmsclient.UpdateSkuStockReq{
		Data: data,
	})
	if err != nil {
		return nil, err
	}

	// 4.如果使用优惠券,更新优惠券使用状态
	// 4.1修改sms_coupon_history表的use_status字段
	// 4.2记得修改sms_coupon的use_count字段,下单的时候要加1,取消订单的时候,要减1
	if len(couponIds) > 0 {
		_, err = l.svcCtx.CouponRecordService.UpdateCouponRecord(l.ctx, &smsclient.UpdateCouponRecordReq{
			MemberId:  memberId,
			Status:    0,
			CouponIds: couponIds,
		})
		if err != nil {
			return nil, err
		}
	}

	// 5.返还使用积分
	member, _ := l.svcCtx.MemberService.QueryMemberInfoDetail(l.ctx, &umsclient.QueryMemberInfoDetailReq{MemberId: memberId})
	i := member.Points + integration
	_, err = l.svcCtx.MemberService.UpdateMemberPoints(l.ctx, &umsclient.UpdateMemberPointsReq{MemberId: memberId, Points: i})
	if err != nil {
		return nil, err
	}

	return &types.CancelUserOrderResp{
		Code:    0,
		Message: "取消订单成功",
	}, nil
}
