package userservicelogic

import (
	"context"
	"errors"

	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

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
	_, err := query.SysUser.WithContext(l.ctx).Where(query.SysUser.ID.Eq(in.UserId)).First()

	// 1.判断用户是否存在
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "用户不存在, 参数：%+v", in)
		return nil, errors.New("用户不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询用户信息, 参数：%+v, 异常: %s", in, err.Error())
		return nil, errors.New("查询用户信息异常")
	}

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
