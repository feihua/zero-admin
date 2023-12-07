package order

import (
	"context"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateConfirmOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateConfirmOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateConfirmOrderLogic {
	return &GenerateConfirmOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateConfirmOrderLogic) GenerateConfirmOrder(req *types.GenerateConfirmOrderReq) (resp *types.GenerateConfirmOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}
