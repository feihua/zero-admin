package homeadvertiseservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeAdvertiseDeleteLogic 首页广告
/*
Author: LiuFeiHua
Date: 2024/5/8 10:17
*/
type HomeAdvertiseDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeAdvertiseDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseDeleteLogic {
	return &HomeAdvertiseDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// HomeAdvertiseDelete 删除首页广告
func (l *HomeAdvertiseDeleteLogic) HomeAdvertiseDelete(in *smsclient.HomeAdvertiseDeleteReq) (*smsclient.HomeAdvertiseDeleteResp, error) {
	q := query.SmsHomeAdvertise
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &smsclient.HomeAdvertiseDeleteResp{}, nil
}
