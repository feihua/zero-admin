package coupon_record

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

// QueryCouponRecordListLogic 查询优惠券领取记录列表
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:03
*/
type QueryCouponRecordListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCouponRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponRecordListLogic {
	return &QueryCouponRecordListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryCouponRecordList 查询优惠券领取记录列表
func (l *QueryCouponRecordListLogic) QueryCouponRecordList(req *types.QueryCouponRecordListReq) (resp *types.QueryCouponRecordListResp, err error) {
	result, err := l.svcCtx.CouponRecordService.QueryCouponRecordList(l.ctx, &smsclient.QueryCouponRecordListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
		CouponId: req.CouponId, // 优惠券ID
		MemberId: req.MemberId, // 用户ID
		GetType:  req.GetType,  // 获取类型：0->后台赠送；1->主动获取
		OrderId:  req.OrderId,  // 使用订单ID
		Status:   req.Status,   // 状态：0-未使用，1-已使用，2-已过期，3-已失效
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字优惠券领取记录列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryCouponRecordListData

	for _, detail := range result.List {
		list = append(list, &types.QueryCouponRecordListData{
			Id:             detail.Id,             // 主键ID
			CouponId:       detail.CouponId,       // 优惠券ID
			MemberId:       detail.MemberId,       // 用户ID
			GetTime:        detail.GetTime,        // 领取时间
			GetType:        detail.GetType,        // 获取类型：0->后台赠送；1->主动获取
			UseTime:        detail.UseTime,        // 使用时间
			OrderId:        detail.OrderId,        // 使用订单ID
			OrderAmount:    detail.OrderAmount,    // 订单金额
			DiscountAmount: detail.DiscountAmount, // 优惠金额
			Status:         detail.Status,         // 状态：0-未使用，1-已使用，2-已过期，3-已失效
			InvalidTime:    detail.InvalidTime,    // 失效时间
			InvalidReason:  detail.InvalidReason,  // 失效原因
			CreateTime:     detail.CreateTime,     // 创建时间

		})
	}

	return &types.QueryCouponRecordListResp{
		Code:     "000000",
		Message:  "查询优惠券领取记录列表成功",
		Data:     list,
		Current:  req.Current,
		PageSize: req.PageSize,
		Total:    result.Total,
		Success:  true,
	}, nil
}
