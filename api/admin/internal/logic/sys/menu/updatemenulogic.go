package menu

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMenuLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:28
*/
type UpdateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateMenuLogic {
	return UpdateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateMenu 更新菜单
func (l *UpdateMenuLogic) UpdateMenu(req *types.UpdateMenuReq) (*types.UpdateMenuResp, error) {
	menuUpdateReq := sysclient.MenuUpdateReq{
		Id:            req.Id,
		Name:          req.Name,
		ParentId:      req.ParentId,
		Url:           req.Url,
		Perms:         req.Perms,
		Type:          req.Type,
		Icon:          req.Icon,
		OrderNum:      req.OrderNum,
		UpdateBy:      l.ctx.Value("userName").(string),
		VuePath:       req.VuePath,
		VueComponent:  req.VueComponent,
		VueIcon:       req.VueIcon,
		VueRedirect:   req.VueRedirect,
		DelFlag:       req.DelFlag,
		BackgroundUrl: req.BackgroundUrl,
	}
	if _, err := l.svcCtx.MenuService.MenuUpdate(l.ctx, &menuUpdateReq); err != nil {
		logc.Errorf(l.ctx, "更新菜单信息失败,参数:%+v,异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新菜单信息失败")
	}

	return &types.UpdateMenuResp{
		Code:    "000000",
		Message: "更新菜单信息成功!",
	}, nil
}
