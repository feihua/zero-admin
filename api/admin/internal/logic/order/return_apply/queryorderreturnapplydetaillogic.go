package return_apply

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

// QueryOrderReturnApplyDetailLogic 查询订单退货申请详情
/*
Author: 刘飞华
Date: 2025/02/05 11:40:41
*/
type QueryOrderReturnApplyDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryOrderReturnApplyDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderReturnApplyDetailLogic {
	return &QueryOrderReturnApplyDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryOrderReturnApplyDetail 查询订单退货申请详情
func (l *QueryOrderReturnApplyDetailLogic) QueryOrderReturnApplyDetail(req *types.QueryOrderReturnApplyDetailReq) (resp *types.QueryOrderReturnApplyDetailResp, err error) {

	detail, err := l.svcCtx.OrderReturnApplyService.QueryOrderReturnApplyDetail(l.ctx, &omsclient.QueryOrderReturnApplyDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询订单退货申请详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryOrderReturnApplyDetailData{
		Id:               detail.Id,               //
		OrderId:          detail.OrderId,          // 订单id
		CompanyAddressId: detail.CompanyAddressId, // 收货地址表id
		ProductId:        detail.ProductId,        // 退货商品id
		OrderSn:          detail.OrderSn,          // 订单编号
		CreateTime:       detail.CreateTime,       // 申请时间
		MemberUsername:   detail.MemberUsername,   // 会员用户名
		ReturnAmount:     detail.ReturnAmount,     // 退款金额
		ReturnName:       detail.ReturnName,       // 退货人姓名
		ReturnPhone:      detail.ReturnPhone,      // 退货人电话
		Status:           detail.Status,           // 申请状态：0->待处理；1->退货中；2->已完成；3->已拒绝
		HandleTime:       detail.HandleTime,       // 处理时间
		ProductPic:       detail.ProductPic,       // 商品图片
		ProductName:      detail.ProductName,      // 商品名称
		ProductBrand:     detail.ProductBrand,     // 商品品牌
		ProductAttr:      detail.ProductAttr,      // 商品销售属性：颜色：红色；尺码：xl;
		ProductCount:     detail.ProductCount,     // 退货数量
		ProductPrice:     detail.ProductPrice,     // 商品单价
		ProductRealPrice: detail.ProductRealPrice, // 商品实际支付单价
		Reason:           detail.Reason,           // 原因
		Description:      detail.Description,      // 描述
		ProofPics:        detail.ProofPics,        // 凭证图片，以逗号隔开
		HandleNote:       detail.HandleNote,       // 处理备注
		HandleMan:        detail.HandleMan,        // 处理人员
		ReceiveMan:       detail.ReceiveMan,       // 收货人
		ReceiveTime:      detail.ReceiveTime,      // 收货时间
		ReceiveNote:      detail.ReceiveNote,      // 收货备注
	}
	return &types.QueryOrderReturnApplyDetailResp{
		Code:    "000000",
		Message: "查询订单退货申请成功",
		Data:    data,
	}, nil
}
