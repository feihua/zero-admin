package user

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryUserRoleListLogic 查询用户与角色的关联
/*
Author: LiuFeiHua
Date: 2024/5/23 17:33
*/
type QueryUserRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserRoleListLogic {
	return &QueryUserRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryUserRoleList 查询用户与角色的关联
func (l *QueryUserRoleListLogic) QueryUserRoleList(req *types.QueryUserRoleListReq) (resp *types.QueryUserRoleListResp, err error) {
	//查询所有角色
	result, err := l.svcCtx.RoleService.RoleList(l.ctx, &sysclient.RoleListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Status:   1,
		Name:     req.RoleName,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询角色信息失败,参数:%+v,异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询角色信息失败")
	}

	var roleList []types.RoleListData
	var roleIds []int64

	for _, menu := range result.List {
		roleList = append(roleList, types.RoleListData{
			Id:         menu.Id,
			Name:       menu.Name,
			Remark:     menu.Remark,
			CreateBy:   menu.CreateBy,
			CreateTime: menu.CreateTime,
			UpdateBy:   menu.UpdateBy,
			UpdateTime: menu.UpdateTime,
			DelFlag:    menu.DelFlag,
			Status:     menu.Status,
		})
		//admin账号全部角色
		roleIds = append(roleIds, menu.Id)
	}

	//如果角色不是admin则根据roleId查询菜单
	//todo 待完善
	//if req.RoleId != 1 {
	//	QueryMenu, err1 := l.svcCtx.RoleService.QueryMenuByRoleId(l.ctx, &sysclient.QueryMenuByRoleIdReq{
	//		Id: req.RoleId,
	//	})
	//	if err1 != nil {
	//		logc.Errorf(l.ctx, "根据roleId查询菜单失败,参数:%+v,异常:%s", req, err1.Error())
	//		return nil, errorx.NewDefaultError("根据roleId查询菜单失败")
	//	}
	//	roleIds = QueryMenu.Ids
	//}

	return &types.QueryUserRoleListResp{
		Data: types.UserRoleListData{
			RoleList: roleList,
			RoleIds:  roleIds,
		},
		Code:    "000000",
		Message: "根据角色id查询菜单成功",
	}, nil
}
