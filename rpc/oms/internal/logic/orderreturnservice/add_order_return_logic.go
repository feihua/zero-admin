package orderreturnservicelogic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// AddOrderReturnLogic 添加退货/售后
/*
Author: LiuFeiHua
Date: 2025/06/30 15:49:28
*/
type AddOrderReturnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderReturnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderReturnLogic {
	return &AddOrderReturnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddOrderReturn 添加退货/售后
func (l *AddOrderReturnLogic) AddOrderReturn(in *omsclient.OrderReturnReq) (*omsclient.OrderReturnResp, error) {
	order, err := query.OmsOrderMain.WithContext(l.ctx).Where(query.OmsOrderMain.ID.Eq(in.OrderId), query.OmsOrderMain.IsDeleted.Eq(0)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "订单不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("订单不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询订单异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询订单异常")
	}

	if order.OrderStatus != 4 {
		logc.Errorf(l.ctx, "订单状态异常, 订单ID：%d, 订单状态：%d", in.OrderId, order.OrderStatus)
		return nil, errors.New("订单状态异常,申请退货/售后明细失败")
	}

	q := query.OmsOrderReturn

	item := &model.OmsOrderReturn{
		OrderID:        in.OrderId,               // 关联订单ID
		ReturnNo:       in.ReturnNo,              // 退货单号
		MemberID:       in.MemberId,              // 会员ID
		Status:         in.Status,                // 退货状态（0待审核 1审核通过 2已收货 3已退款 4已拒绝 5已关闭）
		Type:           in.Type,                  // 售后类型（0退货退款 1仅退款 2换货）
		Reason:         in.Reason,                // 退货原因
		Description:    in.Description,           // 问题描述
		ProofPic:       in.ProofPic,              // 凭证图片，逗号分隔
		RefundAmount:   float64(in.RefundAmount), // 退款金额
		ReturnName:     in.ReturnName,            // 退货人姓名
		ReturnPhone:    in.ReturnPhone,           // 退货人电话
		CompanyAddress: in.CompanyAddress,        // 退货收货地址
		Remark:         in.Remark,                // 备注
	}

	err = q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加退货/售后失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加退货/售后失败")
	}

	orderReturnItem := query.OmsOrderReturnItem
	_, _ = orderReturnItem.WithContext(l.ctx).Where(orderReturnItem.ReturnID.Eq(item.ID)).Delete()
	for _, data := range in.OrderReturnItem {

		returnItem := &model.OmsOrderReturnItem{
			ReturnID:     item.ID,                    // 退货单ID（关联oms_order_return.id）
			OrderID:      data.OrderId,               // 订单ID
			OrderItemID:  data.OrderItemId,           // 订单明细ID
			SkuID:        data.SkuId,                 // 商品SKU ID
			SkuName:      data.SkuName,               // 商品名称
			SkuPic:       data.SkuPic,                // 商品图片
			SkuAttrs:     data.SkuAttrs,              // 商品销售属性
			Quantity:     data.Quantity,              // 退货数量
			ProductPrice: float64(data.ProductPrice), // 商品单价
			RealAmount:   float64(data.RealAmount),   // 实际退款金额
			Reason:       data.Reason,                // 退货原因
			Remark:       data.Remark,                // 备注
		}

		err = orderReturnItem.WithContext(l.ctx).Create(returnItem)
		if err != nil {
			logc.Errorf(l.ctx, "添加退货/售后明细失败,参数:%+v,异常:%s", item, err.Error())
			return nil, errors.New("添加退货/售后明细失败")
		}
	}

	message := map[string]any{"id": in.OrderId}
	body, _ := json.Marshal(message)
	err = l.svcCtx.RabbitMQ.SendMessage("order.return.exchange", "order.return.queue", "order.return.key", body)

	return &omsclient.OrderReturnResp{}, nil
}
