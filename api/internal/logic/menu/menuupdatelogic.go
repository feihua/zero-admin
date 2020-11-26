package logic

import (
	"context"
	"go-zero-admin/rpc/sys/sysclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MenuUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MenuUpdateLogic {
	return MenuUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuUpdateLogic) MenuUpdate(req types.UpdateMenuReq) (*types.UpdateMenuResp, error) {
	_, _ = l.svcCtx.Sys.MenuUpdate(l.ctx, &sysclient.MenuUpdateReq{
		Id:       req.Id,
		Name:     req.Name,
		ParentId: req.ParentId,
		Url:      req.Url,
		Perms:    req.Perms,
		Type:     req.Type,
		Icon:     req.Icon,
		OrderNum: req.OrderNum,
		//todo 从token里面拿
		LastUpdateBy: "admin",
	})

	return &types.UpdateMenuResp{}, nil
}
