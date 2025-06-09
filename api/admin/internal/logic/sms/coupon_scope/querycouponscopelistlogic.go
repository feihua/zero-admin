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

// QueryCouponScopeListLogic 查询优惠券使用范围列表
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:03
*/
type QueryCouponScopeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCouponScopeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponScopeListLogic {
	return &QueryCouponScopeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryCouponScopeList 查询优惠券使用范围列表
func (l *QueryCouponScopeListLogic) QueryCouponScopeList(req *types.QueryCouponScopeListReq) (resp *types.QueryCouponScopeListResp, err error) {
	result, err := l.svcCtx.CouponScopeService.QueryCouponScopeList(l.ctx, &smsclient.QueryCouponScopeListReq{
		PageNum:   req.Current,
		PageSize:  req.PageSize,
		CouponId:  req.CouponId,  // 优惠券ID
		ScopeType: req.ScopeType, // 范围类型：0-全场，1-分类，2-商品
		ScopeId:   req.ScopeId,   // 范围ID（分类ID或商品ID）
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字优惠券使用范围列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryCouponScopeListData

	for _, detail := range result.List {
		list = append(list, &types.QueryCouponScopeListData{
			Id:         detail.Id,         // 主键ID
			CouponId:   detail.CouponId,   // 优惠券ID
			ScopeType:  detail.ScopeType,  // 范围类型：0-全场，1-分类，2-商品
			ScopeId:    detail.ScopeId,    // 范围ID（分类ID或商品ID）
			CreateTime: detail.CreateTime, // 创建时间

		})
	}

	return &types.QueryCouponScopeListResp{
		Code:     "000000",
		Message:  "查询优惠券使用范围列表成功",
		Data:     list,
		Current:  req.Current,
		PageSize: req.PageSize,
		Total:    result.Total,
		Success:  true,
	}, nil
}
