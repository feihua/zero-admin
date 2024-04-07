package category

import (
	"context"

	"github.com/feihua/zero-admin/api/web/internal/svc"
	"github.com/feihua/zero-admin/api/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryProductCateListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryProductCateListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductCateListLogic {
	return &QueryProductCateListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryProductCateListLogic) QueryProductCateList() (resp *types.QueryProductCateListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
