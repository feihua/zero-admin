package seckillactivityservicelogic

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

// UpdateSeckillActivityLogic 更新秒杀活动
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:51
*/
type UpdateSeckillActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSeckillActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSeckillActivityLogic {
	return &UpdateSeckillActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateSeckillActivity 更新秒杀活动
func (l *UpdateSeckillActivityLogic) UpdateSeckillActivity(in *smsclient.UpdateSeckillActivityReq) (*smsclient.UpdateSeckillActivityResp, error) {
	q := query.SmsSeckillActivity.WithContext(l.ctx)

	// 1.根据秒杀活动id查询秒杀活动是否已存在
	detail, err := q.Where(query.SmsSeckillActivity.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "秒杀活动不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("秒杀活动不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询秒杀活动异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询秒杀活动异常")
	}

	count, err := q.Where(query.SmsSeckillActivity.Name.Eq(in.Name), query.SmsSeckillActivity.ID.Neq(in.Id)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "更新秒杀活动失败,参数：%+v, 异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("更新秒杀活动失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("活动名称：%s,已存在", in.Name))
	}

	startTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	endTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	now := time.Now()
	item := &model.SmsSeckillActivity{
		ID:          in.Id,             // 编号
		Name:        in.Name,           // 活动名称
		Description: in.Description,    // 活动描述
		StartTime:   startTime,         // 开始时间
		EndTime:     endTime,           // 结束时间
		Status:      in.Status,         // 状态:0-上线,1-下线
		IsEnabled:   in.IsEnabled,      // 是否启用
		CreateBy:    detail.CreateBy,   // 创建人ID
		CreateTime:  detail.CreateTime, // 创建时间
		UpdateBy:    &in.UpdateBy,      // 更新人ID
		UpdateTime:  &now,              // 更新时间
	}

	// 2.秒杀活动存在时,则直接更新秒杀活动
	err = l.svcCtx.DB.Model(&model.SmsSeckillActivity{}).WithContext(l.ctx).Where(query.SmsSeckillActivity.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新秒杀活动失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新秒杀活动失败")
	}

	return &smsclient.UpdateSeckillActivityResp{}, nil
}
