package homeadvertiseservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteHomeAdvertiseLogic 删除首页轮播广告
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
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

// DeleteHomeAdvertise 删除首页轮播广告
func (l *DeleteHomeAdvertiseLogic) DeleteHomeAdvertise(in *smsclient.DeleteHomeAdvertiseReq) (*smsclient.DeleteHomeAdvertiseResp, error) {
	q := query.SmsHomeAdvertise

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除首页轮播广告失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除首页轮播广告失败")
	}

	return &smsclient.DeleteHomeAdvertiseResp{}, nil
}
