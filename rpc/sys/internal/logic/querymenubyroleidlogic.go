package logic

import (
	"context"
	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

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

func (l *QueryMenuByRoleIdLogic) QueryMenuByRoleId(in *sys.QueryMenuByRoleIdReq) (*sys.QueryMenuByRoleIdResp, error) {
	RoleMenus, _ := l.svcCtx.RoleMenuModel.FindByRoleId(in.Id)

	var list []int64
	for _, user := range *RoleMenus {

		list = append(list, user.MenuId)
	}

	return &sys.QueryMenuByRoleIdResp{
		Ids: list,
	}, nil
}
