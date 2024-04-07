package order

import (
	"context"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayCallbackLogic {
	return &PayCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayCallbackLogic) PayCallback(req *types.PayCallbackReq) (resp *types.PayCallbackResp, err error) {
	// todo: add your logic here and delete this line

	return
}
