package userservicelogic

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
// 1.查询所有角色
// 2.如果角色不是admin则根据userId查询角色
func (l *QueryUserRoleListLogic) QueryUserRoleList(in *sysclient.QueryUserRoleListReq) (*sysclient.QueryUserRoleListResp, error) {
	// 1.查询所有角色
	offset := (in.Current - 1) * in.PageSize
	result, total, err := query.SysRole.WithContext(l.ctx).FindByPage(int(offset), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询角色失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询角色失败")
	}

	var roleList = make([]*sysclient.RoleData, 0, len(result))
	var roleIds = make([]int64, 0, len(result))

	for _, item := range result {
		roleList = append(roleList, &sysclient.RoleData{
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
			CreateTime: time_util.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:   item.UpdateBy,                           // 更新者
			UpdateTime: time_util.TimeToString(item.UpdateTime), // 更新时间
		})

		// admin账号全部角色
		roleIds = append(roleIds, item.ID)
	}

	// 2.如果角色不是admin则根据userId查询角色
	q := query.SysUserRole
	if in.UserId != 1 {
		roleIds = roleIds[:0]
		_ = q.WithContext(l.ctx).Select(q.RoleID).Where(q.UserID.Eq(in.UserId)).Scan(&roleIds)
	}

	return &sysclient.QueryUserRoleListResp{
		List:    roleList,
		RoleIds: roleIds,
		Total:   total,
	}, nil
}
