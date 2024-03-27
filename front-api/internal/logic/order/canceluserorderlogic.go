package order

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/feihua/zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CancelUserOrderLogic
/*
Author: LiuFeiHua
Date: 2023/12/7 9:35
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
func (l *CancelUserOrderLogic) CancelUserOrder(req *types.CancelUserOrderReq) (*types.CancelUserOrderResp, error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()

	//todo 暂时没有分布式事务
	//1.查询订单是否存在
	//2.修改订单状态
	resp, err := l.svcCtx.OrderService.OrderCancel(l.ctx, &omsclient.OrderCancelReq{
		MemberId: memberId,
		OrderId:  req.OrderId,
	})

	couponId := resp.CouponId
	integration := resp.Integration
	stockLockData := resp.Data

	var data []*pmsclient.ReleaseSkuStockLockData
	for _, item := range stockLockData {
		data = append(data, &pmsclient.ReleaseSkuStockLockData{
			ProductSkuId:    item.ProductSkuId,
			ProductQuantity: item.ProductQuantity,
		})
	}
	//3.释放库存
	_, err = l.svcCtx.SkuStockService.ReleaseSkuStockLock(l.ctx, &pmsclient.ReleaseSkuStockLockReq{
		Data: data,
	})
	if err != nil {
		return nil, err
	}

	//4.如果使用优惠券,更新优惠券使用状态
	//4.1修改sms_coupon_history表的use_status字段
	//4.2记得修改sms_coupon的use_count字段,下单的时候要加1,取消订单的时候,要减1
	if couponId > 0 {
		_, err = l.svcCtx.CouponHistoryService.UpdateCouponStatus(l.ctx, &smsclient.UpdateCouponStatusReq{
			MemberId:  memberId,
			UseStatus: 0,
			CouponId:  couponId,
		})
		if err != nil {
			return nil, err
		}
	}

	//5.返还使用积分
	member, _ := l.svcCtx.MemberService.QueryMemberById(l.ctx, &umsclient.MemberByIdReq{Id: memberId})
	i := member.Integration + integration
	_, err = l.svcCtx.MemberService.UpdateMemberIntegration(l.ctx, &umsclient.UpdateMemberIntegrationReq{Id: memberId, Integration: i})
	if err != nil {
		return nil, err
	}

	return &types.CancelUserOrderResp{
		Code:    0,
		Message: "取消订单成功",
	}, nil
}
