package homeadvertiseservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteHomeAdvertiseLogic 删除首页轮播广告表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:51
*/
type DeleteHomeAdvertiseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteHomeAdvertiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHomeAdvertiseLogic {
	return &DeleteHomeAdvertiseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteHomeAdvertise 删除首页轮播广告表
func (l *DeleteHomeAdvertiseLogic) DeleteHomeAdvertise(in *smsclient.DeleteHomeAdvertiseReq) (*smsclient.DeleteHomeAdvertiseResp, error) {
	q := query.SmsHomeAdvertise
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &smsclient.DeleteHomeAdvertiseResp{}, nil
}
