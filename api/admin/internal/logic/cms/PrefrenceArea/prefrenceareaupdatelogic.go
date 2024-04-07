package PrefrenceArea

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrefrenceAreaUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrefrenceAreaUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrefrenceAreaUpdateLogic {
	return &PrefrenceAreaUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PrefrenceAreaUpdateLogic) PrefrenceAreaUpdate(req *types.UpdatePrefrenceAreaReq) (resp *types.UpdatePrefrenceAreaResp, err error) {
	// todo: add your logic here and delete this line

	return
}
