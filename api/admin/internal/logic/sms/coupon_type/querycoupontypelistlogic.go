package coupon_type

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

// QueryCouponTypeListLogic 查询优惠券类型列表
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:03
*/
type QueryCouponTypeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCouponTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponTypeListLogic {
	return &QueryCouponTypeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryCouponTypeList 查询优惠券类型列表
func (l *QueryCouponTypeListLogic) QueryCouponTypeList(req *types.QueryCouponTypeListReq) (resp *types.QueryCouponTypeListResp, err error) {
	result, err := l.svcCtx.CouponTypeService.QueryCouponTypeList(l.ctx, &smsclient.QueryCouponTypeListReq{
		PageNum:      req.Current,
		PageSize:     req.PageSize,
		Name:         req.Name,         // 类型名称
		Code:         req.Code,         // 类型编码
		DiscountType: req.DiscountType, // 优惠方式：1-固定金额，2-折扣率，3-第N件特惠，4-买赠，5-特价，6-套装优惠，7-搭配优惠，8-积分抵现，9-积分倍率，10-免运费，11-运费减免，12-限时特权，13-会员特权
		Status:       req.Status,       // 是否启用
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字优惠券类型列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryCouponTypeListData

	for _, detail := range result.List {
		list = append(list, &types.QueryCouponTypeListData{
			Id:           detail.Id,           // 主键ID
			Name:         detail.Name,         // 类型名称
			Code:         detail.Code,         // 类型编码
			Description:  detail.Description,  // 描述
			DiscountType: detail.DiscountType, // 优惠方式：1-固定金额，2-折扣率，3-第N件特惠，4-买赠，5-特价，6-套装优惠，7-搭配优惠，8-积分抵现，9-积分倍率，10-免运费，11-运费减免，12-限时特权，13-会员特权
			Status:       detail.Status,       // 是否启用
			Sort:         detail.Sort,         // 排序
			CreateBy:     detail.CreateBy,     // 创建人ID
			CreateTime:   detail.CreateTime,   // 创建时间
			UpdateBy:     detail.UpdateBy,     // 更新人ID
			UpdateTime:   detail.UpdateTime,   // 更新时间

		})
	}

	return &types.QueryCouponTypeListResp{
		Code:     "000000",
		Message:  "查询优惠券类型列表成功",
		Data:     list,
		Current:  req.Current,
		PageSize: req.PageSize,
		Total:    result.Total,
		Success:  true,
	}, nil
}
