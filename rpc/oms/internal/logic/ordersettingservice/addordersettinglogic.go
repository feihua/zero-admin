package ordersettingservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddOrderSettingLogic 添加订单设置
/*
Author: LiuFeiHua
Date: 2024/6/12 9:35
*/
type AddOrderSettingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderSettingLogic {
	return &AddOrderSettingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddOrderSetting 添加订单设置
func (l *AddOrderSettingLogic) AddOrderSetting(in *omsclient.AddOrderSettingReq) (*omsclient.AddOrderSettingResp, error) {
	q := query.OmsOrderSetting
	if in.IsDefault == 1 {
		if _, err := q.WithContext(l.ctx).Where(q.IsDefault.Eq(1)).Update(q.IsDefault, 0); err != nil {
			return nil, errors.New("添加订单设置失败")
		}
	}

	err := q.WithContext(l.ctx).Create(&model.OmsOrderSetting{
		FlashOrderOvertime:  in.FlashOrderOvertime,  // 秒杀订单超时关闭时间(分)
		NormalOrderOvertime: in.NormalOrderOvertime, // 正常订单超时时间(分)
		ConfirmOvertime:     in.ConfirmOvertime,     // 发货后自动确认收货时间（天）
		FinishOvertime:      in.FinishOvertime,      // 自动完成交易时间，不能申请售后（天）
		Status:              in.Status,              // 状态：0->禁用；1->启用
		IsDefault:           in.IsDefault,           // 是否默认：0->否；1->是
		CommentOvertime:     in.CommentOvertime,     // 订单完成后自动好评时间（天）
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加订单设置失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加订单设置失败")
	}

	return &omsclient.AddOrderSettingResp{}, nil
}
