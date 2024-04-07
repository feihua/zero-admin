package brand

import (
	"context"

	"github.com/feihua/zero-admin/api/web/internal/svc"
	"github.com/feihua/zero-admin/api/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandListLogic {
	return &BrandListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BrandListLogic) BrandList() (resp *types.BrandListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
