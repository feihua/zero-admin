package payservicelogic

import (
	"context"

	"zero-admin/rpc/pay/internal/svc"
	"zero-admin/rpc/pay/payclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayLogic {
	return &PayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PayLogic) Pay(in *payclient.PayReq) (*payclient.PayResp, error) {
	// todo: add your logic here and delete this line

	return &payclient.PayResp{}, nil
}
