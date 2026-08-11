package menu

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMenuLogic 新增菜单
/*
Author: LiuFeiHua
Date: 2023/12/18 15:24
*/
type AddMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddMenuLogic {
	return AddMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddMenu 新增菜单
func (l *AddMenuLogic) AddMenu(req *types.AddMenuReq) (*types.BaseResp, error) {
	createBy := l.ctx.Value("userName").(string)
	menuAddReq := sysclient.AddMenuReq{
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
		CreateBy:     createBy,         // 创建者
	}
	if _, err := l.svcCtx.MenuService.AddMenu(l.ctx, &menuAddReq); err != nil {
		logc.Errorf(l.ctx, "添加菜单信息失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
