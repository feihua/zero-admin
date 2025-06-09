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

// QueryCouponTypeDetailLogic 查询优惠券类型详情
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:03
*/
type QueryCouponTypeDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCouponTypeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponTypeDetailLogic {
	return &QueryCouponTypeDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryCouponTypeDetail 查询优惠券类型详情
func (l *QueryCouponTypeDetailLogic) QueryCouponTypeDetail(req *types.QueryCouponTypeDetailReq) (resp *types.QueryCouponTypeDetailResp, err error) {

	detail, err := l.svcCtx.CouponTypeService.QueryCouponTypeDetail(l.ctx, &smsclient.QueryCouponTypeDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询优惠券类型详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryCouponTypeDetailData{
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

	}
	return &types.QueryCouponTypeDetailResp{
		Code:    "000000",
		Message: "查询优惠券类型成功",
		Data:    data,
	}, nil
}
