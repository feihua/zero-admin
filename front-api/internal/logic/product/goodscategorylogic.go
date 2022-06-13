package product

import (
	"context"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) GoodsCategoryLogic {
	return GoodsCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsCategoryLogic) GoodsCategory(req types.GoodsCategoryReq) (resp *types.GoodsDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
