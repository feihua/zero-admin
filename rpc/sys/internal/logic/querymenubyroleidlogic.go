package logic

import (
	"context"
	"strconv"

	"go-zero-admin/rpc/sys/internal/svc"
	"go-zero-admin/rpc/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
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

	var list []string
	for _, user := range *RoleMenus {

		list = append(list, strconv.FormatInt(user.MenuId, 10))
	}

	return &sys.QueryMenuByRoleIdResp{
		Ids: nil,
	}, nil
}
