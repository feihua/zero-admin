package prefrenceArea

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePreferredAreaStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePreferredAreaStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePreferredAreaStatusLogic {
	return &UpdatePreferredAreaStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePreferredAreaStatusLogic) UpdatePreferredAreaStatus(req *types.UpdatePreferredAreaStatusReq) (resp *types.UpdatePreferredAreaStatusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
