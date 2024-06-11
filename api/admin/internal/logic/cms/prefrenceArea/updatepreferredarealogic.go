package prefrenceArea

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePreferredAreaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePreferredAreaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePreferredAreaLogic {
	return &UpdatePreferredAreaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePreferredAreaLogic) UpdatePreferredArea(req *types.UpdatePreferredAreaReq) (resp *types.UpdatePreferredAreaResp, err error) {
	// todo: add your logic here and delete this line

	return
}
