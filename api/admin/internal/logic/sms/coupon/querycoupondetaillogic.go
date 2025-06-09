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

// QueryCouponDetailLogic 查询优惠券详情
/*
Author: LiuFeiHua
Date: 2025/06/12 10:40:14
*/
type QueryCouponDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCouponDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponDetailLogic {
	return &QueryCouponDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryCouponDetail 查询优惠券详情
func (l *QueryCouponDetailLogic) QueryCouponDetail(req *types.QueryCouponDetailReq) (resp *types.QueryCouponDetailResp, err error) {

	detail, err := l.svcCtx.CouponService.QueryCouponDetail(l.ctx, &smsclient.QueryCouponDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询优惠券详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryCouponDetailData{
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

	}
	return &types.QueryCouponDetailResp{
		Code:    "000000",
		Message: "查询优惠券成功",
		Data:    data,
	}, nil
}
