package product

import (
	"context"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) GoodsDetailLogic {
	return GoodsDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsDetailLogic) GoodsDetail(req types.GoodsDetailReq) (resp *types.GoodsDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
