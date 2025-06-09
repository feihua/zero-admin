package seckillsessionservicelogic

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
)

// AddSeckillSessionLogic 添加秒杀场次
/*
Author: LiuFeiHua
Date: 2025/06/11 10:29:58
*/
type AddSeckillSessionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSeckillSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSeckillSessionLogic {
	return &AddSeckillSessionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddSeckillSession 添加秒杀场次
func (l *AddSeckillSessionLogic) AddSeckillSession(in *smsclient.AddSeckillSessionReq) (*smsclient.AddSeckillSessionResp, error) {
	q := query.SmsSeckillSession

	count, err := q.WithContext(l.ctx).Where(q.Name.Eq(in.Name)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "添加秒杀场次失败,参数：%+v, 异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("添加秒杀场次失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("秒杀场次：%s,已存在", in.Name))
	}

	item := &model.SmsSeckillSession{
		Name:      in.Name,      // 场次名称
		StartTime: in.StartTime, // 开始时间
		EndTime:   in.EndTime,   // 结束时间
		Status:    in.Status,    // 状态：0-禁用，1-启用
		Sort:      in.Sort,      // 排序
		CreateBy:  in.CreateBy,  // 创建人ID
	}

	err = q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加秒杀场次失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加秒杀场次失败")
	}

	return &smsclient.AddSeckillSessionResp{}, nil
}
