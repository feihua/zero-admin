package orderreturnservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

// UpdateOrderReturnLogic 更新退货/售后
/*
Author: LiuFeiHua
Date: 2025/06/30 15:49:28
*/
type UpdateOrderReturnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderReturnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderReturnLogic {
	return &UpdateOrderReturnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateOrderReturn 更新退货/售后
func (l *UpdateOrderReturnLogic) UpdateOrderReturn(in *omsclient.OrderReturnReq) (*omsclient.OrderReturnResp, error) {
	q := query.OmsOrderReturn.WithContext(l.ctx)

	// 1.根据退货/售后id查询退货/售后是否已存在
	_, err := q.Where(query.OmsOrderReturn.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "退货/售后不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("退货/售后不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询退货/售后异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询退货/售后异常")
	}

	item := &model.OmsOrderReturn{
		ID:             in.Id,                    // 主键ID
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

	if len(in.HandleTime) > 0 {
		t, _ := time.Parse("2006-01-02 15:04:05", in.HandleTime)
		item.HandleTime = &t
	}
	if len(in.ReceiveTime) > 0 {
		t, _ := time.Parse("2006-01-02 15:04:05", in.ReceiveTime)
		item.ReceiveTime = &t
	}
	if len(in.RefundTime) > 0 {
		t, _ := time.Parse("2006-01-02 15:04:05", in.RefundTime)
		item.RefundTime = &t
	}
	if len(in.CloseTime) > 0 {
		t, _ := time.Parse("2006-01-02 15:04:05", in.CloseTime)
		item.CloseTime = &t
	}
	// 2.退货/售后存在时,则直接更新退货/售后
	_, err = q.Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新退货/售后失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新退货/售后失败")
	}

	orderReturnItem := query.OmsOrderReturnItem
	_, _ = orderReturnItem.WithContext(l.ctx).Where(orderReturnItem.ReturnID.Eq(in.Id)).Delete()
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
	return &omsclient.OrderReturnResp{}, nil
}
