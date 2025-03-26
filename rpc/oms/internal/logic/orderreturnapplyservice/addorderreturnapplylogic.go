package orderreturnapplyservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddOrderReturnApplyLogic 添加订单退货申请
/*
Author: LiuFeiHua
Date: 2024/6/12 10:14
*/
type AddOrderReturnApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderReturnApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderReturnApplyLogic {
	return &AddOrderReturnApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddOrderReturnApply 添加订单退货申请
func (l *AddOrderReturnApplyLogic) AddOrderReturnApply(in *omsclient.AddOrderReturnApplyReq) (*omsclient.AddOrderReturnApplyResp, error) {
	err := query.OmsOrderReturnApply.WithContext(l.ctx).Create(&model.OmsOrderReturnApply{
		OrderID:          in.OrderId,          // 订单id
		ProductID:        in.ProductId,        // 退货商品id
		OrderSn:          in.OrderSn,          // 订单编号
		MemberUsername:   in.MemberUsername,   // 会员用户名
		ReturnName:       in.ReturnName,       // 退货人姓名
		ReturnPhone:      in.ReturnPhone,      // 退货人电话
		Status:           in.Status,           // 申请状态：0->待处理；1->退货中；2->已完成；3->已拒绝
		ProductPic:       in.ProductPic,       // 商品图片
		ProductName:      in.ProductName,      // 商品名称
		ProductBrand:     in.ProductBrand,     // 商品品牌
		ProductAttr:      in.ProductAttr,      // 商品销售属性：颜色：红色；尺码：xl;
		ProductCount:     in.ProductCount,     // 退货数量
		ProductPrice:     in.ProductPrice,     // 商品单价
		ProductRealPrice: in.ProductRealPrice, // 商品实际支付单价
		Reason:           in.Reason,           // 原因
		Description:      in.Description,      // 描述
		ProofPics:        in.ProofPics,        // 凭证图片，以逗号隔开
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加订单退货申请失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加订单退货申请失败")
	}

	return &omsclient.AddOrderReturnApplyResp{}, nil
}
