// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package menu

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMenuResourceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMenuResourceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuResourceListLogic {
	return &QueryMenuResourceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryMenuResourceListLogic) QueryMenuResourceList(req *types.QueryMenuListReq) (resp *types.QueryMenuListResp, err error) {
	result, err := l.svcCtx.MenuService.QueryMenuResourceList(l.ctx, &sysclient.QueryMenuListReq{
		PageNum:    req.Current,
		PageSize:   req.PageSize,
		MenuName:   req.MenuName,
		ParentId:   req.ParentId,
		MenuStatus: req.MenuStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询菜单列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryMenuListData

	for _, item := range result.List {
		menuItem := &types.QueryMenuListData{
			Id:           item.Id,           // 主键
			MenuName:     item.MenuName,     // 菜单名称
			Ancestors:    item.Ancestors,    // 祖级列表
			MenuType:     item.MenuType,     // 菜单类型(1:目录,2:菜单,3:按钮)
			MenuUrl:      item.MenuUrl,      // 路由路径
			MenuIcon:     item.MenuIcon,     // 菜单图标
			MenuSort:     item.MenuSort,     // 排序
			ParentId:     item.ParentId,     // 父id
			ApiUrl:       item.ApiUrl,       // 接口url
			Visible:      item.Visible,      // 显示状态（0:隐藏,显示:1）
			Status:       item.Status,       // 菜单状态(1:正常，0:禁用)
			Remark:       item.Remark,       // 备注
			VuePath:      item.VuePath,      // vue的path
			VueComponent: item.VueComponent, // vue的页面
			VueIcon:      item.VueIcon,      // vue的图标
			VueRedirect:  item.VueRedirect,  // vue的路由重定向
			AngularIcon:  item.AngularIcon,  // angular的图标
			ReactIcon:    item.ReactIcon,    // antd react的图标
			CreateBy:     item.CreateBy,     // 创建者
			CreateTime:   item.CreateTime,   // 创建时间
			UpdateBy:     item.UpdateBy,     // 更新者
			UpdateTime:   item.UpdateTime,   // 更新时间
		}

		list = append(list, menuItem)
	}

	return &types.QueryMenuListResp{
		Code:    "000000",
		Message: "查询菜单成功",
		Data:    list,
		Success: true,
		Total:   result.Total,
	}, nil
}
