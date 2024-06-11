package prefrenceArea

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPreferredAreaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPreferredAreaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPreferredAreaLogic {
	return &AddPreferredAreaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPreferredAreaLogic) AddPreferredArea(req *types.AddPreferredAreaReq) (resp *types.AddPreferredAreaResp, err error) {
	// todo: add your logic here and delete this line

	return
}
