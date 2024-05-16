package product

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRecommendStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRecommendStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRecommendStatusLogic {
	return &UpdateRecommendStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRecommendStatusLogic) UpdateRecommendStatus(req *types.UpdateProductStatusReq) (resp *types.UpdateProductStatusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
