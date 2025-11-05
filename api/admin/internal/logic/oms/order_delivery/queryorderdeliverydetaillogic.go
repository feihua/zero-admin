// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package order_delivery

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryOrderDeliveryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryOrderDeliveryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderDeliveryDetailLogic {
	return &QueryOrderDeliveryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryOrderDeliveryDetailLogic) QueryOrderDeliveryDetail(req *types.QueryOrderDeliveryDetailReq) (resp *types.QueryOrderDeliveryDetailResp, err error) {
	detail, err := l.svcCtx.OrderDeliveryService.QueryOrderDeliveryDetail(l.ctx, &omsclient.QueryOrderDeliveryDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询订单收货地址详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryOrderDeliveryDetailData{
		Id:               detail.Id,               //
		OrderId:          detail.OrderId,          // 订单ID
		OrderNo:          detail.OrderNo,          // 订单编号
		ReceiverName:     detail.ReceiverName,     // 收货人姓名
		ReceiverPhone:    detail.ReceiverPhone,    // 收货人电话
		ReceiverProvince: detail.ReceiverProvince, // 省份
		ReceiverCity:     detail.ReceiverCity,     // 城市
		ReceiverDistrict: detail.ReceiverDistrict, // 区县
		ReceiverAddress:  detail.ReceiverAddress,  // 详细地址
		DeliveryCompany:  detail.DeliveryCompany,  // 物流公司
		DeliveryNo:       detail.DeliveryNo,       // 物流单号
		CreateTime:       detail.CreateTime,       // 创建时间
		UpdateTime:       detail.UpdateTime,       // 更新时间

	}
	return &types.QueryOrderDeliveryDetailResp{
		Code:    "000000",
		Message: "查询订单收货地址成功",
		Data:    data,
	}, nil
}
