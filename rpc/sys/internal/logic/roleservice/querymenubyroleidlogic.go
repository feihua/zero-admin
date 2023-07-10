package roleservicelogic

import (
	"context"
	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMenuByRoleIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMenuByRoleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuByRoleIdLogic {
	return &QueryMenuByRoleIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryMenuByRoleIdLogic) QueryMenuByRoleId(in *sysclient.QueryMenuByRoleIdReq) (*sysclient.QueryMenuByRoleIdResp, error) {
	RoleMenus, _ := l.svcCtx.RoleMenuModel.FindByRoleId(l.ctx, in.Id)

	var list []int64
	for _, user := range *RoleMenus {

		list = append(list, user.MenuId)
	}

	return &sysclient.QueryMenuByRoleIdResp{
		Ids: list,
	}, nil
}
