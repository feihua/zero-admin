package setting

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

// QueryOrderSettingDetailLogic 查询订单设置表详情
/*
Author: 刘飞华
Date: 2025/02/05 11:40:41
*/
type QueryOrderSettingDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryOrderSettingDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderSettingDetailLogic {
	return &QueryOrderSettingDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryOrderSettingDetail 查询订单设置表详情
func (l *QueryOrderSettingDetailLogic) QueryOrderSettingDetail(req *types.QueryOrderSettingDetailReq) (resp *types.QueryOrderSettingDetailResp, err error) {

	detail, err := l.svcCtx.OrderSettingService.QueryOrderSettingDetail(l.ctx, &omsclient.QueryOrderSettingDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询订单设置表详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryOrderSettingDetailData{
		Id:                  detail.Id,                  //
		FlashOrderOvertime:  detail.FlashOrderOvertime,  // 秒杀订单超时关闭时间(分)
		NormalOrderOvertime: detail.NormalOrderOvertime, // 正常订单超时时间(分)
		ConfirmOvertime:     detail.ConfirmOvertime,     // 发货后自动确认收货时间（天）
		FinishOvertime:      detail.FinishOvertime,      // 自动完成交易时间，不能申请售后（天）
		Status:              detail.Status,              // 状态：0->禁用；1->启用
		IsDefault:           detail.IsDefault,           // 是否默认：0->否；1->是
		CommentOvertime:     detail.CommentOvertime,     // 订单完成后自动好评时间（天）
	}
	return &types.QueryOrderSettingDetailResp{
		Code:    "000000",
		Message: "查询订单设置表成功",
		Data:    data,
	}, nil
}
