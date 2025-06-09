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

// DeleteSeckillSessionLogic 删除秒杀场次
/*
Author: LiuFeiHua
Date: 2025/06/11 10:29:58
*/
type DeleteSeckillSessionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSeckillSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSeckillSessionLogic {
	return &DeleteSeckillSessionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteSeckillSession 删除秒杀场次
func (l *DeleteSeckillSessionLogic) DeleteSeckillSession(in *smsclient.DeleteSeckillSessionReq) (*smsclient.DeleteSeckillSessionResp, error) {
	q := query.SmsSeckillSession

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除秒杀场次失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除秒杀场次失败")
	}

	return &smsclient.DeleteSeckillSessionResp{}, nil
}
