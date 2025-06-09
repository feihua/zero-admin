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

// DeleteSeckillProductLogic 删除秒杀商品
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:51
*/
type DeleteSeckillProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSeckillProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSeckillProductLogic {
	return &DeleteSeckillProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteSeckillProduct 删除秒杀商品
func (l *DeleteSeckillProductLogic) DeleteSeckillProduct(in *smsclient.DeleteSeckillProductReq) (*smsclient.DeleteSeckillProductResp, error) {
	q := query.SmsSeckillProduct

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除秒杀商品失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除秒杀商品失败")
	}

	return &smsclient.DeleteSeckillProductResp{}, nil
}
