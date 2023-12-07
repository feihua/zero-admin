package order

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/oms/omsclient"
	"zero-admin/rpc/pms/pmsclient"
	"zero-admin/rpc/sms/smsclient"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

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
func (l *CancelUserOrderLogic) CancelUserOrder(req *types.CancelUserOrderReq) (resp *types.CancelUserOrderResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()

	//todo 暂时没有分布式事务
	//1.查询订单是否存在
	//2.修改订单状态
	result, err := l.svcCtx.OrderService.OrderCancel(l.ctx, &omsclient.OrderCancelReq{
		MemberId: memberId,
		OrderId:  req.OrderID,
	})

	var data []*pmsclient.ReleaseSkuStockLockData
	for _, item := range result.Data {
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

	//4.修改优惠券使用状态
	_, err = l.svcCtx.CouponHistoryService.UpdateCouponStatus(l.ctx, &smsclient.UpdateCouponStatusReq{
		MemberId:  memberId,
		UseStatus: 0,
		CouponId:  result.CouponId,
	})
	if err != nil {
		return nil, err
	}

	//5.返还使用积分
	member, _ := l.svcCtx.MemberService.QueryMemberById(l.ctx, &umsclient.MemberByIdReq{Id: memberId})
	i := member.Integration - result.Integration
	_, err = l.svcCtx.MemberService.UpdateMemberIntegration(l.ctx, &umsclient.UpdateMemberIntegrationReq{Id: memberId, Integration: i})
	if err != nil {
		return nil, err
	}

	return &types.CancelUserOrderResp{
		Code:    0,
		Message: "取消订单成功",
	}, nil
}
