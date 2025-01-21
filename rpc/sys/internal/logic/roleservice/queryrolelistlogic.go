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
	if in.IsAdmin != 2 {
		q = q.Where(query.SysRole.IsAdmin.Eq(in.IsAdmin))
	}
	if in.RoleStatus != 2 {
		q = q.Where(query.SysRole.RoleStatus.Eq(in.RoleStatus))
	}

	offset := (in.PageNum - 1) * in.PageSize
	result, count, err := q.FindByPage(int(offset), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询角色列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询角色列表失败")
	}

	var list []*sysclient.RoleListData
	for _, item := range result {
		createTime := item.CreateTime.Format("2006-01-02 15:04:05")
		list = append(list, &sysclient.RoleListData{
			Id:         item.ID,                                 // 编号
			RoleName:   item.RoleName,                           // 角色名称
			RoleKey:    item.RoleKey,                            // 权限字符
			RoleStatus: item.RoleStatus,                         // 角色状态
			RoleSort:   item.RoleSort,                           // 角色排序
			DataScope:  item.DataScope,                          // 数据权限
			IsDeleted:  item.IsDeleted,                          // 是否删除  0：否  1：是
			IsAdmin:    item.IsAdmin,                            // 是否超级管理员:  0：否  1：是
			Remark:     item.Remark,                             // 备注
			CreateBy:   item.CreateBy,                           // 创建者
			CreateTime: createTime,                              // 创建时间
			UpdateBy:   item.UpdateBy,                           // 更新者
			UpdateTime: time_util.TimeToString(item.UpdateTime), // 更新时间
		})
	}

	logc.Infof(l.ctx, "查询角色列表,参数：%+v,响应：%+v", in, list)
	return &sysclient.QueryRoleListResp{
		Total: count,
		List:  list,
	}, nil

}
