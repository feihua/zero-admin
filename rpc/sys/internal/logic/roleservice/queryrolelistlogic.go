package roleservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
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

	if in.Status != 2 {
		q = q.Where(query.SysRole.Status.Eq(in.Status))
	}
	if in.DataScope != 0 {
		q = q.Where(query.SysRole.DataScope.Eq(in.DataScope))
	}

	offset := (in.PageNum - 1) * in.PageSize
	result, count, err := q.FindByPage(int(offset), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询角色列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询角色列表失败")
	}

	var list = make([]*sysclient.RoleListData, 0, len(result))
	for _, item := range result {
		list = append(list, &sysclient.RoleListData{
			Id:         item.ID,                                 // 角色id
			RoleName:   item.RoleName,                           // 名称
			RoleKey:    item.RoleKey,                            // 角色权限字符串
			DataScope:  item.DataScope,                          // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
			Status:     item.Status,                             // 状态(1:正常，0:禁用)
			Remark:     item.Remark,                             // 备注
			DelFlag:    item.DelFlag,                            // 删除标志（0代表删除 1代表存在）
			CreateBy:   item.CreateBy,                           // 创建者
			CreateTime: time_util.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:   item.UpdateBy,                           // 更新者
			UpdateTime: time_util.TimeToString(item.UpdateTime), // 更新时间
		})
	}

	return &sysclient.QueryRoleListResp{
		Total: count,
		List:  list,
	}, nil

}
