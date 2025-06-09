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

// QueryCouponScopeDetailLogic 查询优惠券使用范围详情
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:03
*/
type QueryCouponScopeDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCouponScopeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponScopeDetailLogic {
	return &QueryCouponScopeDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryCouponScopeDetail 查询优惠券使用范围详情
func (l *QueryCouponScopeDetailLogic) QueryCouponScopeDetail(req *types.QueryCouponScopeDetailReq) (resp *types.QueryCouponScopeDetailResp, err error) {

	detail, err := l.svcCtx.CouponScopeService.QueryCouponScopeDetail(l.ctx, &smsclient.QueryCouponScopeDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询优惠券使用范围详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryCouponScopeDetailData{
		Id:         detail.Id,         // 主键ID
		CouponId:   detail.CouponId,   // 优惠券ID
		ScopeType:  detail.ScopeType,  // 范围类型：0-全场，1-分类，2-商品
		ScopeId:    detail.ScopeId,    // 范围ID（分类ID或商品ID）
		CreateTime: detail.CreateTime, // 创建时间

	}
	return &types.QueryCouponScopeDetailResp{
		Code:    "000000",
		Message: "查询优惠券使用范围成功",
		Data:    data,
	}, nil
}
