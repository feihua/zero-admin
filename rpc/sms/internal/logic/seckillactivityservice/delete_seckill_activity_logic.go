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

// DeleteSeckillActivityLogic 删除秒杀活动
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:51
*/
type DeleteSeckillActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSeckillActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSeckillActivityLogic {
	return &DeleteSeckillActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteSeckillActivity 删除秒杀活动
func (l *DeleteSeckillActivityLogic) DeleteSeckillActivity(in *smsclient.DeleteSeckillActivityReq) (*smsclient.DeleteSeckillActivityResp, error) {
	q := query.SmsSeckillActivity

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除秒杀活动失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除秒杀活动失败")
	}

	return &smsclient.DeleteSeckillActivityResp{}, nil
}
