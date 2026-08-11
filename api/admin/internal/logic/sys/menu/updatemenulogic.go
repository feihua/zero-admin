package menu

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMenuLogic 更新菜单
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
func (l *UpdateMenuLogic) UpdateMenu(req *types.UpdateMenuReq) (*types.BaseResp, error) {
	updateBy := l.ctx.Value("userName").(string)
	menuReq := sysclient.UpdateMenuReq{
		Id:           req.Id,           // 主键
		MenuName:     req.MenuName,     // 菜单名称
		MenuType:     req.MenuType,     // 菜单类型(1:目录,2:菜单,3:按钮)
		MenuUrl:      req.MenuUrl,      // 路由路径
		MenuIcon:     req.MenuIcon,     // 菜单图标
		MenuSort:     req.MenuSort,     // 排序
		ParentId:     req.ParentId,     // 父id
		ApiUrl:       req.ApiUrl,       // 接口url
		Visible:      req.Visible,      // 显示状态（0:隐藏,显示:1）
		Status:       req.Status,       // 菜单状态(1:正常，0:禁用)
		Remark:       req.Remark,       // 备注
		VuePath:      req.VuePath,      // vue的path
		VueComponent: req.VueComponent, // vue的页面
		VueIcon:      req.VueIcon,      // vue的图标
		VueRedirect:  req.VueRedirect,  // vue的路由重定向
		AngularIcon:  req.AngularIcon,  // angular的图标
		ReactIcon:    req.ReactIcon,    // antd react的图标
		UpdateBy:     updateBy,         // 更新者
	}
	if _, err := l.svcCtx.MenuService.UpdateMenu(l.ctx, &menuReq); err != nil {
		logc.Errorf(l.ctx, "更新菜单信息失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
