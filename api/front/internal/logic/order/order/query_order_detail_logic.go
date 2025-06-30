package order

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOrderDetailLogic 订单详情
/*
Author: LiuFeiHua
Date: 2025/6/20 11:35
*/
type QueryOrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderDetailLogic {
	return &QueryOrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryOrderDetail 订单详情
func (l *QueryOrderDetailLogic) QueryOrderDetail(req *types.OrderDetailReq) (resp1 *types.OrderDetailResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	res, err := l.svcCtx.OrderService.QueryOrderDetail(l.ctx, &omsclient.QueryOrderDetailReq{
		Id:     req.OrderId,
		UserId: memberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询订单详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	detail := res.Data
	data := types.QueryOrderData{
		Id:                 detail.Id,                 //
		OrderNo:            detail.OrderNo,            // 订单编号
		UserId:             detail.UserId,             // 用户ID
		OrderStatus:        detail.OrderStatus,        // 订单状态：1-待支付,2-已支付,3-已发货,4-已完成,5-已取消,6-已退款,7-售后中
		TotalAmount:        detail.TotalAmount,        // 订单总金额
		PromotionAmount:    detail.PromotionAmount,    // 促销金额
		CouponAmount:       detail.CouponAmount,       // 优惠券金额
		PointsAmount:       detail.PointsAmount,       // 积分金额
		DiscountAmount:     detail.DiscountAmount,     // 优惠金额
		FreightAmount:      detail.FreightAmount,      // 运费金额
		PayAmount:          detail.PayAmount,          // 实付金额
		PayType:            detail.PayType,            // 支付方式：1-支付宝,2-微信,3-银联
		PayTime:            detail.PayTime,            // 支付时间
		DeliveryTime:       detail.DeliveryTime,       // 发货时间
		ReceiveTime:        detail.ReceiveTime,        // 收货时间
		CommentTime:        detail.CommentTime,        // 评价时间
		SourceType:         detail.SourceType,         // 订单来源：1-APP,2-PC,3-小程序
		ExpressOrderNumber: detail.ExpressOrderNumber, // 快递单号
		UsePoints:          detail.UsePoints,          // 下单时使用的积分
		ReceiveStatus:      detail.ReceiveStatus,      // 是否确认收货：0->否,1->是
		Remark:             detail.Remark,             // 订单备注
		CreateTime:         detail.CreateTime,         // 提交时间
		UpdateTime:         detail.UpdateTime,         //

	}

	itemData := detail.OrderItemData

	var orderItemData []*types.OrderItemData

	for _, d := range itemData {
		orderItemData = append(orderItemData, &types.OrderItemData{
			Id:              d.Id,              //
			OrderId:         d.OrderId,         // 订单ID
			OrderNo:         d.OrderNo,         // 订单编号
			OrderItemStatus: d.OrderItemStatus, // 订单商品状态：1-正常,2-退货申请中,3-已退货,4-已拒绝
			SkuId:           d.SkuId,           // 商品SKU ID
			SkuName:         d.SkuName,         // 商品名称
			SkuPic:          d.SkuPic,          // 商品图片
			SkuPrice:        d.SkuPrice,        // 商品单价
			SkuQuantity:     d.SkuQuantity,     // 商品数量
			SpecData:        d.SpecData,        // 规格数据
			SkuTotalAmount:  d.SkuTotalAmount,  // 商品总金额
			PromotionAmount: d.PromotionAmount, // 促销分摊金额
			CouponAmount:    d.CouponAmount,    // 优惠券分摊金额
			PointsAmount:    d.PointsAmount,    // 积分分摊金额
			DiscountAmount:  d.DiscountAmount,  // 优惠分摊金额
			RealAmount:      d.RealAmount,      // 实付金额

		})
	}
	data.OrderItemData = orderItemData

	return &types.OrderDetailResp{
		Code:    0,
		Message: "操作成功",
		Data:    data,
	}, nil

}
