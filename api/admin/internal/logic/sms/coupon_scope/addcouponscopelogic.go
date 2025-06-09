package coupon_scope

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddCouponScopeLogic 添加优惠券使用范围
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:03
*/
type AddCouponScopeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCouponScopeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCouponScopeLogic {
	return &AddCouponScopeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddCouponScope 添加优惠券使用范围
func (l *AddCouponScopeLogic) AddCouponScope(req *types.AddCouponScopeReq) (resp *types.BaseResp, err error) {

	var list []*smsclient.AddCouponScopeData
	for _, item := range req.Data {
		list = append(list, &smsclient.AddCouponScopeData{
			CouponId:  item.CouponId,  // 优惠券ID
			ScopeType: item.ScopeType, // 范围类型：0-全场，1-分类，2-商品
			ScopeId:   item.ScopeId,   // 范围ID（分类ID或商品ID）
		})
	}
	_, err = l.svcCtx.CouponScopeService.AddCouponScope(l.ctx, &smsclient.AddCouponScopeReq{
		Data: list,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加优惠券使用范围失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "添加优惠券使用范围成功",
	}, nil
}
