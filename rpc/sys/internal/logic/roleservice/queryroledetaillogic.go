package roleservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/logic/common"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryRoleDetailLogic 查询角色信息表详情
/*
Author: LiuFeiHua
Date: 2024/5/30 14:18
*/
type QueryRoleDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryRoleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleDetailLogic {
	return &QueryRoleDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryRoleDetail 查询角色信息表详情
func (l *QueryRoleDetailLogic) QueryRoleDetail(in *sysclient.QueryRoleDetailReq) (*sysclient.QueryRoleDetailResp, error) {
	role, err := query.SysRole.WithContext(l.ctx).Where(query.SysRole.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询角色信息表详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询角色信息表详情失败")
	}

	data := &sysclient.QueryRoleDetailResp{
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
	}

	logc.Infof(l.ctx, "查询角色信息表详情,参数：%+v,响应：%+v", in, data)
	return data, nil
}
