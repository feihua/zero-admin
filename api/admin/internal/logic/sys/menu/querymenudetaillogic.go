package menu

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMenuDetailLogic 查询菜单详情
/*
Author: LiuFeiHua
Date: 2024/5/29 17:54
*/
type QueryMenuDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMenuDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuDetailLogic {
	return &QueryMenuDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMenuDetail 查询菜单详情
func (l *QueryMenuDetailLogic) QueryMenuDetail(req *types.QueryMenuDetailReq) (resp *types.QueryMenuDetailResp, err error) {
	detail, err := l.svcCtx.MenuService.QueryMenuDetail(l.ctx, &sysclient.QueryMenuDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询菜单详情,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryMenuDetailData{
		Id:            detail.Id,            // 编号
		MenuName:      detail.MenuName,      // 菜单名称
		ParentId:      detail.ParentId,      // 父菜单ID，一级菜单为0
		MenuPath:      detail.MenuPath,      // 前端路由
		MenuPerms:     detail.MenuPerms,     // 权限标识
		MenuType:      detail.MenuType,      // 类型 0：目录,1：菜单,2：按钮,3：外链
		MenuIcon:      detail.MenuIcon,      // 菜单图标
		MenuSort:      detail.MenuSort,      // 菜单排序
		CreateBy:      detail.CreateBy,      // 创建者
		CreateTime:    detail.CreateTime,    // 创建时间
		UpdateBy:      detail.UpdateBy,      // 更新者
		UpdateTime:    detail.UpdateTime,    // 更新时间
		MenuStatus:    detail.MenuStatus,    // 菜单状态
		IsDeleted:     detail.IsDeleted,     // 是否删除  0：否  1：是
		IsVisible:     detail.IsVisible,     // 是否可见  0：否  1：是
		Remark:        detail.Remark,        // 备注信息
		VuePath:       detail.VuePath,       // vue系统的path
		VueComponent:  detail.VueComponent,  // vue的页面
		VueIcon:       detail.VueIcon,       // vue的图标
		VueRedirect:   detail.VueRedirect,   // vue的路由重定向
		BackgroundUrl: detail.BackgroundUrl, // 接口地址
	}

	return &types.QueryMenuDetailResp{
		Code:    "000000",
		Message: "查询菜单详情成功",
		Data:    data,
	}, nil
}
