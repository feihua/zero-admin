package ordersettingservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

// UpdateOrderSettingLogic 更新订单设置
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:44
*/
type UpdateOrderSettingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderSettingLogic {
	return &UpdateOrderSettingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateOrderSetting 更新订单设置
func (l *UpdateOrderSettingLogic) UpdateOrderSetting(in *omsclient.UpdateOrderSettingReq) (*omsclient.UpdateOrderSettingResp, error) {
	setting := query.OmsOrderSetting
	q := setting.WithContext(l.ctx)

	// 1.根据订单设置id查询订单设置是否已存在
	detail, err := q.Where(setting.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "订单设置不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("订单设置不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询订单设置异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询订单设置异常")
	}
	if in.IsDefault == 1 {
		if _, err = q.WithContext(l.ctx).Where(setting.IsDefault.Eq(1)).Update(setting.IsDefault, 0); err != nil {
			return nil, err
		}
	}

	now := time.Now()
	item := &model.OmsOrderSetting{
		ID:                  in.Id,                  // 主键ID
		FlashOrderOvertime:  in.FlashOrderOvertime,  // 秒杀订单超时关闭时间(分)
		NormalOrderOvertime: in.NormalOrderOvertime, // 正常订单超时时间(分)
		ConfirmOvertime:     in.ConfirmOvertime,     // 发货后自动确认收货时间（天）
		FinishOvertime:      in.FinishOvertime,      // 自动完成交易时间，不能申请售后（天）
		Status:              in.Status,              // 状态：0->禁用；1->启用
		IsDefault:           in.IsDefault,           // 是否默认：0->否；1->是
		CommentOvertime:     in.CommentOvertime,     // 订单完成后自动好评时间（天）
		CreateBy:            detail.CreateBy,        // 创建者
		CreateTime:          detail.CreateTime,      // 创建时间
		UpdateBy:            &in.UpdateBy,           // 更新者
		UpdateTime:          &now,                   // 更新时间
	}

	// 2.订单设置存在时,则直接更新订单设置
	err = l.svcCtx.DB.Model(&model.OmsOrderSetting{}).WithContext(l.ctx).Where(setting.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新订单设置失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新订单设置失败")
	}

	return &omsclient.UpdateOrderSettingResp{}, nil
}
