package logic

import (
	"context"

	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeBrandDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeBrandDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeBrandDeleteLogic {
	return &HomeBrandDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeBrandDeleteLogic) HomeBrandDelete(in *sms.HomeBrandDeleteReq) (*sms.HomeBrandDeleteResp, error) {
	err := l.svcCtx.SmsHomeBrandModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sms.HomeBrandDeleteResp{}, nil
}
