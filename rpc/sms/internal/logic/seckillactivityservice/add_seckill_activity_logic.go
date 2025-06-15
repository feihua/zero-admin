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
	"time"
)

// AddSeckillActivityLogic 添加秒杀活动
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:51
*/
type AddSeckillActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSeckillActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSeckillActivityLogic {
	return &AddSeckillActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddSeckillActivity 添加秒杀活动
func (l *AddSeckillActivityLogic) AddSeckillActivity(in *smsclient.AddSeckillActivityReq) (*smsclient.AddSeckillActivityResp, error) {
	q := query.SmsSeckillActivity

	count, err := q.WithContext(l.ctx).Where(q.Name.Eq(in.Name)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "添加秒杀活动失败,参数：%+v, 异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("添加秒杀活动失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("活动名称：%s,已存在", in.Name))
	}

	startTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	endTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	item := &model.SmsSeckillActivity{
		Name:        in.Name,        // 活动名称
		Description: in.Description, // 活动描述
		StartTime:   startTime,      // 开始时间
		EndTime:     endTime,        // 结束时间
		Status:      in.Status,      // 状态:0-上线,1-下线
		IsEnabled:   in.IsEnabled,   // 是否启用
		CreateBy:    in.CreateBy,    // 创建人ID
	}

	err = q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加秒杀活动失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加秒杀活动失败")
	}

	return &smsclient.AddSeckillActivityResp{}, nil
}
