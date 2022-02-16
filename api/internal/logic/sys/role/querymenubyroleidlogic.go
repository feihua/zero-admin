package logic

import (
	"context"
	"strconv"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMenuByRoleIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMenuByRoleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryMenuByRoleIdLogic {
	return QueryMenuByRoleIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryMenuByRoleIdLogic) QueryMenuByRoleId(req types.RoleMenuReq) (*types.RoleMenuResp, error) {
	resp, _ := l.svcCtx.Sys.MenuList(l.ctx, &sysclient.MenuListReq{
		Name: "",
		Url:  "",
	})

	var list []*types.ListMenuData

	for _, menu := range resp.List {
		list = append(list, &types.ListMenuData{
			Key:      strconv.FormatInt(menu.Id, 10),
			Title:    menu.Name,
			ParentId: menu.ParentId,
			Id:       menu.Id,
			Label:    menu.Name,
		})
	}

	QueryMenu, _ := l.svcCtx.Sys.QueryMenuByRoleId(l.ctx, &sysclient.QueryMenuByRoleIdReq{
		Id: req.Id,
	})

	return &types.RoleMenuResp{
		AllData:  list,
		RoleData: QueryMenu.Ids,
		Code:     "000000",
		Message:  "",
	}, nil
}
