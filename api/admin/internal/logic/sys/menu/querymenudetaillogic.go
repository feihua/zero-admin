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
		Id:           detail.Id,           // 主键
		MenuName:     detail.MenuName,     // 菜单名称
		Ancestors:    detail.Ancestors,    // 祖级列表
		MenuType:     detail.MenuType,     // 菜单类型(1:目录,2:菜单,3:按钮)
		MenuUrl:      detail.MenuUrl,      // 路由路径
		MenuIcon:     detail.MenuIcon,     // 菜单图标
		MenuSort:     detail.MenuSort,     // 排序
		ParentId:     detail.ParentId,     // 父id
		ApiUrl:       detail.ApiUrl,       // 接口url
		Visible:      detail.Visible,      // 显示状态（0:隐藏,显示:1）
		Status:       detail.Status,       // 菜单状态(1:正常，0:禁用)
		Remark:       detail.Remark,       // 备注
		VuePath:      detail.VuePath,      // vue的path
		VueComponent: detail.VueComponent, // vue的页面
		VueIcon:      detail.VueIcon,      // vue的图标
		VueRedirect:  detail.VueRedirect,  // vue的路由重定向
		AngularIcon:  detail.AngularIcon,  // angular的图标
		ReactIcon:    detail.ReactIcon,    // antd react的图标
		CreateBy:     detail.CreateBy,     // 创建者
		CreateTime:   detail.CreateTime,   // 创建时间
		UpdateBy:     detail.UpdateBy,     // 更新者
		UpdateTime:   detail.UpdateTime,   // 更新时间
	}

	return &types.QueryMenuDetailResp{
		Code:    "000000",
		Message: "查询菜单详情成功",
		Data:    data,
	}, nil
}
