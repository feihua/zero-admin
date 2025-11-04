package order_return

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOrderReturnDetailLogic 查询退货/售后主详情
/*
Author: LiuFeiHua
Date: 2025/07/01 10:06:44
*/
type QueryOrderReturnDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryOrderReturnDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderReturnDetailLogic {
	return &QueryOrderReturnDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryOrderReturnDetail 查询退货/售后主详情
func (l *QueryOrderReturnDetailLogic) QueryOrderReturnDetail(req *types.QueryOrderReturnDetailReq) (resp *types.QueryOrderReturnDetailResp, err error) {

	detail, err := l.svcCtx.OrderReturnService.QueryOrderReturnDetail(l.ctx, &omsclient.QueryOrderReturnDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询退货/售后主详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryOrderReturnDetailData{
		Id:             detail.Id,                    // 主键ID
		OrderId:        detail.OrderId,               // 关联订单ID
		ReturnNo:       detail.ReturnNo,              // 退货单号
		MemberId:       detail.MemberId,              // 会员ID
		Status:         detail.Status,                // 退货状态（0待审核 1审核通过 2已收货 3已退款 4已拒绝 5已关闭）
		Type:           detail.Type,                  // 售后类型（0退货退款 1仅退款 2换货）
		Reason:         detail.Reason,                // 退货原因
		Description:    detail.Description,           // 问题描述
		ProofPic:       detail.ProofPic,              // 凭证图片，逗号分隔
		RefundAmount:   float64(detail.RefundAmount), // 退款金额
		ReturnName:     detail.ReturnName,            // 退货人姓名
		ReturnPhone:    detail.ReturnPhone,           // 退货人电话
		CompanyAddress: detail.CompanyAddress,        // 退货收货地址
		CreateTime:     detail.CreateTime,            // 申请时间
		HandleTime:     detail.HandleTime,            // 处理时间
		HandleNote:     detail.HandleNote,            // 处理备注
		HandleMan:      detail.HandleMan,             // 处理人员
		ReceiveTime:    detail.ReceiveTime,           // 收货时间
		ReceiveNote:    detail.ReceiveNote,           // 收货备注
		ReceiveMan:     detail.ReceiveMan,            // 收货人
		RefundTime:     detail.RefundTime,            // 退款时间
		CloseTime:      detail.CloseTime,             // 关闭时间
		Remark:         detail.Remark,                // 备注

	}
	var list []*types.ReturnItemListData

	for _, item := range detail.OrderReturnItem {
		list = append(list, &types.ReturnItemListData{
			Id:           item.Id,           // 主键ID
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

	data.ReturnItemListData = list
	return &types.QueryOrderReturnDetailResp{
		Code:    "000000",
		Message: "查询退货/售后主成功",
		Data:    data,
	}, nil
}
