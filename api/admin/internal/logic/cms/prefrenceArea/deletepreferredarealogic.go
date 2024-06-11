package prefrenceArea

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePreferredAreaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePreferredAreaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePreferredAreaLogic {
	return &DeletePreferredAreaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePreferredAreaLogic) DeletePreferredArea(req *types.DeletePreferredAreaReq) (resp *types.DeletePreferredAreaResp, err error) {
	// todo: add your logic here and delete this line

	return
}
