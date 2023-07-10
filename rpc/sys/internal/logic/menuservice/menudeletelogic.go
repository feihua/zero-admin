package menuservicelogic

import (
	"context"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuDeleteLogic {
	return &MenuDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuDeleteLogic) MenuDelete(in *sysclient.MenuDeleteReq) (*sysclient.MenuDeleteResp, error) {
	err := l.svcCtx.MenuModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &sysclient.MenuDeleteResp{}, nil
}
