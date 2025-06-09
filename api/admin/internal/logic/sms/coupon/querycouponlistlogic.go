package coupon

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

// QueryCouponListLogic 查询优惠券列表
/*
Author: LiuFeiHua
Date: 2025/06/12 10:40:14
*/
type QueryCouponListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCouponListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponListLogic {
	return &QueryCouponListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryCouponList 查询优惠券列表
func (l *QueryCouponListLogic) QueryCouponList(req *types.QueryCouponListReq) (resp *types.QueryCouponListResp, err error) {
	result, err := l.svcCtx.CouponService.QueryCouponList(l.ctx, &smsclient.QueryCouponListReq{
		PageNum:   req.Current,
		PageSize:  req.PageSize,
		TypeId:    req.TypeId,    // 优惠券类型ID
		Name:      req.Name,      // 优惠券名称
		Code:      req.Code,      // 优惠券码
		StartTime: req.StartTime, // 生效时间
		EndTime:   req.EndTime,   // 失效时间
		Status:    req.Status,    // 状态：0-未开始，1-进行中，2-已结束，3-已取消
		IsEnabled: req.IsEnabled, // 是否启用
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字优惠券列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryCouponListData

	for _, detail := range result.List {
		list = append(list, &types.QueryCouponListData{
			Id:            detail.Id,            // 优惠券ID
			TypeId:        detail.TypeId,        // 优惠券类型ID
			Name:          detail.Name,          // 优惠券名称
			Code:          detail.Code,          // 优惠券码
			Amount:        detail.Amount,        // 优惠金额/折扣率
			MinAmount:     detail.MinAmount,     // 最低使用金额
			StartTime:     detail.StartTime,     // 生效时间
			EndTime:       detail.EndTime,       // 失效时间
			TotalCount:    detail.TotalCount,    // 发放总量
			ReceivedCount: detail.ReceivedCount, // 已领取数量
			UsedCount:     detail.UsedCount,     // 已使用数量
			PerLimit:      detail.PerLimit,      // 每人限领数量
			Status:        detail.Status,        // 状态：0-未开始，1-进行中，2-已结束，3-已取消
			IsEnabled:     detail.IsEnabled,     // 是否启用
			Description:   detail.Description,   // 使用说明
			CreateBy:      detail.CreateBy,      // 创建人ID
			CreateTime:    detail.CreateTime,    // 创建时间
			UpdateBy:      detail.UpdateBy,      // 更新人ID
			UpdateTime:    detail.UpdateTime,    // 更新时间

		})
	}

	return &types.QueryCouponListResp{
		Code:     "000000",
		Message:  "查询优惠券列表成功",
		Data:     list,
		Current:  req.Current,
		PageSize: req.PageSize,
		Total:    result.Total,
		Success:  true,
	}, nil
}
