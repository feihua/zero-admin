package roleservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryRoleDetailLogic 查询角色详情
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

// QueryRoleDetail 查询角色详情
// 1.判断角色是否存在
func (l *QueryRoleDetailLogic) QueryRoleDetail(in *sysclient.QueryRoleDetailReq) (*sysclient.QueryRoleDetailResp, error) {
	item, err := query.SysRole.WithContext(l.ctx).Where(query.SysRole.ID.Eq(in.Id)).First()

	// 1.判断角色是否存在
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errors.New("角色不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询角色异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询角色异常")
	}

	data := &sysclient.QueryRoleDetailResp{
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
	}

	return data, nil
}
