package logic

import (
	"context"

	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeNewProductDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeNewProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeNewProductDeleteLogic {
	return &HomeNewProductDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeNewProductDeleteLogic) HomeNewProductDelete(in *sms.HomeNewProductDeleteReq) (*sms.HomeNewProductDeleteResp, error) {
	err := l.svcCtx.SmsHomeNewProductModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sms.HomeNewProductDeleteResp{}, nil
}
