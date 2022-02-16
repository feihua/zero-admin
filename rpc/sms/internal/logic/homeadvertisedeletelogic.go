package logic

import (
	"context"

	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *HomeAdvertiseDeleteLogic) HomeAdvertiseDelete(in *sms.HomeAdvertiseDeleteReq) (*sms.HomeAdvertiseDeleteResp, error) {
	err := l.svcCtx.SmsHomeAdvertiseModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sms.HomeAdvertiseDeleteResp{}, nil
}
