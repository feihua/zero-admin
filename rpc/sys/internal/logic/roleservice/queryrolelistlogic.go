package roleservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/logic/common"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryRoleListLogic 查询角色列表
/*
Author: LiuFeiHua
Date: 2023/12/18 16:06
*/
type QueryRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleListLogic {
	return &QueryRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryRoleList 查询角色列表
func (l *QueryRoleListLogic) QueryRoleList(in *sysclient.QueryRoleListReq) (*sysclient.QueryRoleListResp, error) {
	q := query.SysRole.WithContext(l.ctx)
	if len(in.RoleName) > 0 {
		q = q.Where(query.SysRole.RoleName.Like("%" + in.RoleName + "%"))
	}
	if len(in.RoleKey) > 0 {
		q = q.Where(query.SysRole.RoleKey.Like("%" + in.RoleKey + "%"))
	}
	if in.IsAdmin != 2 {
		q = q.Where(query.SysRole.IsAdmin.Eq(in.IsAdmin))
	}
	if in.RoleStatus != 2 {
		q = q.Where(query.SysRole.RoleStatus.Eq(in.RoleStatus))
	}

	offset := (in.PageNum - 1) * in.PageSize
	result, count, err := q.FindByPage(int(offset), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询角色列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询角色列表信息失败")
	}

	var list []*sysclient.RoleListData
	for _, role := range result {
		list = append(list, &sysclient.RoleListData{
			CreateBy:   role.CreateBy,
			CreateTime: role.CreateTime.Format("2006-01-02 15:04:05"),
			DataScope:  role.DataScope,
			Id:         role.ID,
			IsAdmin:    role.IsAdmin,
			UpdateTime: common.TimeToString(role.UpdateTime),
			Remark:     role.Remark,
			RoleKey:    role.RoleKey,
			RoleName:   role.RoleName,
			RoleSort:   role.RoleSort,
			RoleStatus: role.RoleStatus,
			UpdateBy:   role.UpdateBy,
		})
	}

	logc.Infof(l.ctx, "查询角色列表信息,参数：%+v,响应：%+v", in, list)
	return &sysclient.QueryRoleListResp{
		Total: count,
		List:  list,
	}, nil

}
