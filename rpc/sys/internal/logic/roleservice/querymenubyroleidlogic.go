package roleservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMenuByRoleIdLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:52
*/
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

// QueryMenuByRoleId 查询角色权限
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
