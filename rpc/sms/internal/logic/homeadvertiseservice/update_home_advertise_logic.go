package homeadvertiseservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

// UpdateHomeAdvertiseLogic 更新首页轮播广告
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type UpdateHomeAdvertiseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHomeAdvertiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomeAdvertiseLogic {
	return &UpdateHomeAdvertiseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateHomeAdvertise 更新首页轮播广告
func (l *UpdateHomeAdvertiseLogic) UpdateHomeAdvertise(in *smsclient.UpdateHomeAdvertiseReq) (*smsclient.UpdateHomeAdvertiseResp, error) {
	advertise := query.SmsHomeAdvertise
	q := advertise.WithContext(l.ctx)

	// 1.根据首页轮播广告id查询首页轮播广告是否已存在
	ad, err := q.Where(advertise.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "首页轮播广告不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("首页轮播广告不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询首页轮播广告异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询首页轮播广告异常")
	}

	count, err := advertise.WithContext(l.ctx).Where(advertise.Name.Eq(in.Name), advertise.ID.Neq(in.Id)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "更新首页轮播广告失败,参数：%+v, 异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("更新首页轮播广告失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("首页轮播广告：%s,已存在", in.Name))
	}

	startTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	endTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	now := time.Now()
	item := &model.SmsHomeAdvertise{
		ID:         in.Id,         // 编号
		Name:       in.Name,       // 名称
		Type:       in.Type,       // 轮播位置：0->PC首页轮播；1->app首页轮播
		Pic:        in.Pic,        // 图片地址
		StartTime:  startTime,     // 开始时间
		EndTime:    endTime,       // 结束时间
		Status:     in.Status,     // 上下线状态：0->下线；1->上线
		ClickCount: ad.ClickCount, // 点击数
		OrderCount: ad.OrderCount, // 下单数
		URL:        in.Url,        // 链接地址
		Remark:     in.Remark,     // 备注
		Sort:       in.Sort,       // 排序
		CreateTime: ad.CreateTime, // 创建时间
		UpdateTime: &now,
	}

	// 2.首页轮播广告存在时,则直接更新首页轮播广告
	err = l.svcCtx.DB.Model(&model.SmsHomeAdvertise{}).WithContext(l.ctx).Where(advertise.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新首页轮播广告失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新首页轮播广告失败")
	}

	return &smsclient.UpdateHomeAdvertiseResp{}, nil
}
