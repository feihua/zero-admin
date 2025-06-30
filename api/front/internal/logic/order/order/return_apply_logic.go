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

// ReturnApplyLogic 申请退货
/*
Author: LiuFeiHua
Date: 2025/6/20 10:59
*/
type ReturnApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnApplyLogic {
	return &ReturnApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ReturnApply 申请退货
func (l *ReturnApplyLogic) ReturnApply(req *types.ReturnApplyReq) (resp *types.ReturnApplyResp, err error) {
	userId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}

	var items []*omsclient.OrderReturnItemData
	for _, item := range req.ReturnItemData {
		items = append(items, &omsclient.OrderReturnItemData{
			ReturnId:     item.ReturnId,     // 退货单ID（关联oms_order_return.id）
			OrderId:      item.OrderId,      // 订单ID
			OrderItemId:  item.OrderItemId,  // 订单明细ID
			SkuId:        item.SkuId,        // 商品SKU ID
			SkuName:      item.SkuName,      // 商品名称
			SkuPic:       item.SkuPic,       // 商品图片
			SkuAttrs:     item.SkuAttrs,     // 商品销售属性
			Quantity:     item.Quantity,     // 退货数量
			ProductPrice: item.ProductPrice, // 商品单价
			RealAmount:   item.RealAmount,   // 实际退款金额
			Reason:       item.Reason,       // 退货原因
			Remark:       item.Remark,       // 备注
		})
	}
	_, err = l.svcCtx.OrderReturnService.AddOrderReturn(l.ctx, &omsclient.OrderReturnReq{
		OrderId:         req.OrderId,        // 关联订单ID
		MemberId:        userId,             // 会员ID
		Status:          req.Status,         // 退货状态（0待审核 1审核通过 2已收货 3已退款 4已拒绝 5已关闭）
		Type:            req.Type,           // 售后类型（0退货退款 1仅退款 2换货）
		Reason:          req.Reason,         // 退货原因
		Description:     req.Description,    // 问题描述
		ProofPic:        req.ProofPic,       // 凭证图片，逗号分隔
		RefundAmount:    req.RefundAmount,   // 退款金额
		ReturnName:      req.ReturnName,     // 退货人姓名
		ReturnPhone:     req.ReturnPhone,    // 退货人电话
		CompanyAddress:  req.CompanyAddress, // 退货收货地址
		Remark:          req.Remark,         // 备注
		OrderReturnItem: items,              // 退货商品
	})
	if err != nil {
		logc.Errorf(l.ctx, "申请退货失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.ReturnApplyResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
