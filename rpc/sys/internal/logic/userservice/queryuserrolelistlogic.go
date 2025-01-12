package userservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/logic/common"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

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
	offset := (in.Current - 1) * in.PageSize
	result, total, err := query.SysRole.WithContext(l.ctx).FindByPage(int(offset), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询角色信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询角色信息失败")
	}

	var roleList []*sysclient.RoleData
	var roleIds []int64

	for _, role := range result {
		roleList = append(roleList, &sysclient.RoleData{
			CreateBy:   role.CreateBy,
			CreateTime: role.CreateTime.Format("2006-01-02 15:04:05"),
			DataScope:  role.DataScope,
			Id:         role.ID,
			IsAdmin:    role.IsAdmin,
			UpdateTime: time_util.TimeToString(role.UpdateTime),
			Remark:     role.Remark,
			RoleKey:    role.RoleKey,
			RoleName:   role.RoleName,
			RoleSort:   role.RoleSort,
			RoleStatus: role.RoleStatus,
			UpdateBy:   role.UpdateBy,
		})
		//admin账号全部角色
		roleIds = append(roleIds, role.ID)
	}

	//2.如果角色不是admin则根据userId查询角色
	q := query.SysUserRole
	if !common.IsAdmin(l.ctx, in.UserId, l.svcCtx.DB) {
		var ids []int64
		_ = q.WithContext(l.ctx).Select(q.RoleID).Where(q.UserID.Eq(in.UserId)).Scan(&ids)
		roleIds = ids
	}

	return &sysclient.QueryUserRoleListResp{
		List:    roleList,
		RoleIds: roleIds,
		Total:   total,
	}, nil
}
