package product

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNewStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateNewStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNewStatusLogic {
	return &UpdateNewStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateNewStatusLogic) UpdateNewStatus(req *types.UpdateProductStatusReq) (resp *types.UpdateProductStatusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
