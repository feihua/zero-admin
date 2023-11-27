package payservicelogic

import (
	"context"

	"zero-admin/rpc/pay/internal/svc"
	"zero-admin/rpc/pay/payclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPayQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayQueryLogic {
	return &PayQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PayQueryLogic) PayQuery(in *payclient.PayQueryResp) (*payclient.PayQueryReq, error) {
	// todo: add your logic here and delete this line

	return &payclient.PayQueryReq{}, nil
}
