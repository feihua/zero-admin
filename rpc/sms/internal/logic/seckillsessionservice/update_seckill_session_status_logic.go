package seckillsessionservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateSeckillSessionStatusLogic 更新秒杀场次
/*
Author: LiuFeiHua
Date: 2025/06/11 10:29:58
*/
type UpdateSeckillSessionStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSeckillSessionStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSeckillSessionStatusLogic {
	return &UpdateSeckillSessionStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateSeckillSessionStatus 更新秒杀场次状态
func (l *UpdateSeckillSessionStatusLogic) UpdateSeckillSessionStatus(in *smsclient.UpdateSeckillSessionStatusReq) (*smsclient.UpdateSeckillSessionStatusResp, error) {
	q := query.SmsSeckillSession

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).UpdateSimple(q.Status.Value(in.Status), q.UpdateBy.Value(in.UpdateBy))

	if err != nil {
		logc.Errorf(l.ctx, "更新秒杀场次状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新秒杀场次状态失败")
	}

	return &smsclient.UpdateSeckillSessionStatusResp{}, nil
}
