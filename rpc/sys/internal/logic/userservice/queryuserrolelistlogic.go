package userservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryUserRoleListLogic 查询用户与角色的关联
/*
Author: LiuFeiHua
Date: 2024/5/24 18:05
*/
type QueryUserRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUserRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserRoleListLogic {
	return &QueryUserRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryUserRoleList 查询用户与角色的关联
func (l *QueryUserRoleListLogic) QueryUserRoleList(in *sysclient.QueryUserRoleListReq) (*sysclient.QueryUserRoleListResp, error) {
	//1.查询所有角色
	//roles, err := query.SysRole.WithContext(l.ctx).Find()
	//
	//if err != nil {
	//	logc.Errorf(l.ctx, "查询角色信息失败,参数:%+v,异常:%s", in, err.Error())
	//	return nil, errors.New("查询角色信息失败")
	//}
	//
	//var roleList []types.RoleListData
	//var roleIds []int64
	//
	//for _, menu := range result.List {
	//	roleList = append(roleList, types.RoleListData{
	//		Id:         menu.Id,
	//		Name:       menu.Name,
	//		Remark:     menu.Remark,
	//		CreateBy:   menu.CreateBy,
	//		CreateTime: menu.CreateTime,
	//		UpdateBy:   menu.UpdateBy,
	//		UpdateTime: menu.UpdateTime,
	//		DelFlag:    menu.DelFlag,
	//		Status:     menu.Status,
	//	})
	//	//admin账号全部角色
	//	roleIds = append(roleIds, menu.Id)
	//}

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

	return &sysclient.QueryUserRoleListResp{}, nil
}
