package seckillproductservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateSeckillProductStatusLogic 更新秒杀商品
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:51
*/
type UpdateSeckillProductStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSeckillProductStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSeckillProductStatusLogic {
	return &UpdateSeckillProductStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateSeckillProductStatus 更新秒杀商品状态
func (l *UpdateSeckillProductStatusLogic) UpdateSeckillProductStatus(in *smsclient.UpdateSeckillProductStatusReq) (*smsclient.UpdateSeckillProductStatusResp, error) {
	q := query.SmsSeckillProduct

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "更新秒杀商品状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新秒杀商品状态失败")
	}

	return &smsclient.UpdateSeckillProductStatusResp{}, nil
}
