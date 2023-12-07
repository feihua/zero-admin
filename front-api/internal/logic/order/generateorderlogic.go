package order

import (
	"context"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateOrderLogic {
	return &GenerateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateOrderLogic) GenerateOrder(req *types.GenerateOrderReq) (resp *types.GenerateOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}
