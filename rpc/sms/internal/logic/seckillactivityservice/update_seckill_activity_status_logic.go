package seckillactivityservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateSeckillActivityStatusLogic 更新秒杀活动
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:51
*/
type UpdateSeckillActivityStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSeckillActivityStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSeckillActivityStatusLogic {
	return &UpdateSeckillActivityStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateSeckillActivityStatus 更新秒杀活动状态
func (l *UpdateSeckillActivityStatusLogic) UpdateSeckillActivityStatus(in *smsclient.UpdateSeckillActivityStatusReq) (*smsclient.UpdateSeckillActivityStatusResp, error) {
	q := query.SmsSeckillActivity

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "更新秒杀活动状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新秒杀活动状态失败")
	}

	return &smsclient.UpdateSeckillActivityStatusResp{}, nil
}
