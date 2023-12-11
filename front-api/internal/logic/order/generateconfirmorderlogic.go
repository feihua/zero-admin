package order

import (
	"context"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// GenerateConfirmOrderLogic
/*
Author: LiuFeiHua
Date: 2023/12/8 14:04
*/
type GenerateConfirmOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateConfirmOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateConfirmOrderLogic {
	return &GenerateConfirmOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GenerateConfirmOrder 根据用户购物车信息生成确认单信息
//1.获取购物车信息
//2.获取用户收货地址列表
//3。获取用户可用优惠券列表
//4。获取用户积分
//5.获取积分使用规则
//6.计算总金额、活动优惠、应付金额
func (l *GenerateConfirmOrderLogic) GenerateConfirmOrder(req *types.GenerateConfirmOrderReq) (resp *types.GenerateConfirmOrderResp, err error) {
	//1.获取购物车信息
	//2.获取用户收货地址列表
	//3。获取用户可用优惠券列表
	//4。获取用户积分
	//5.获取积分使用规则
	//6.计算总金额、活动优惠、应付金额
	return
}
