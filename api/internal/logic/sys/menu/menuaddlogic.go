package logic

import (
	"context"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/sys/sysclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MenuAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MenuAddLogic {
	return MenuAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuAddLogic) MenuAdd(req types.AddMenuReq) (*types.AddMenuResp, error) {
	_, err := l.svcCtx.Sys.MenuAdd(l.ctx, &sysclient.MenuAddReq{
		Name:     req.Name,
		ParentId: req.ParentId,
		Url:      req.Url,
		Perms:    req.Perms,
		Type:     req.Type,
		Icon:     req.Icon,
		OrderNum: req.OrderNum,
		//todo 从token里面拿
		CreateBy:     "admin",
		VuePath:      req.VuePath,
		VueComponent: req.VueComponent,
		VueIcon:      req.VueIcon,
		VueRedirect:  req.VueRedirect,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("添加菜单失败")
	}

	return &types.AddMenuResp{
		Code:    "000000",
		Message: "添加菜单成功",
	}, nil
}
