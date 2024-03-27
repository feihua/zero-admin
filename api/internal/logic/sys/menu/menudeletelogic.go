package menu

import (
	"context"
	"github.com/feihua/zero-admin/api/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/internal/svc"
	"github.com/feihua/zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// MenuDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:25
*/
type MenuDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MenuDeleteLogic {
	return MenuDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// MenuDelete 删除菜单
func (l *MenuDeleteLogic) MenuDelete(req types.DeleteMenuReq) (*types.DeleteMenuResp, error) {
	if _, err := l.svcCtx.MenuService.MenuDelete(l.ctx, &sysclient.MenuDeleteReq{
		Ids: req.Ids,
	}); err != nil {
		logc.Errorf(l.ctx, "根据menuId: %+v,删除菜单异常:%s", req.Ids, err.Error())
		return nil, errorx.NewDefaultError("删除菜单失败")
	}

	return &types.DeleteMenuResp{
		Code:    "000000",
		Message: "删除菜单成功",
	}, nil
}
