package product

import (
	"context"

	"github.com/feihua/zero-admin/api/web/internal/svc"
	"github.com/feihua/zero-admin/api/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductLogic {
	return &QueryProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryProductLogic) QueryProduct(req *types.QueryProductReq) (resp *types.QueryProductResp, err error) {
	// todo: add your logic here and delete this line

	return
}
