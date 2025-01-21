package roleservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryRoleUserListLogic 查询角色与用户的关联
/*
Author: LiuFeiHua
Date: 2024/6/02 17:42
*/
type QueryRoleUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryRoleUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleUserListLogic {
	return &QueryRoleUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryRoleUserList 查询角色与用户的关联
// 1.超级管理员不需要分配用户
// 2.如果角色不是admin则根据roleId查询用户
func (l *QueryRoleUserListLogic) QueryRoleUserList(in *sysclient.QueryRoleUserListReq) (*sysclient.QueryRoleUserListResp, error) {
	var result []model.SysUser

	userRole := query.SysUserRole
	sysUser := query.SysUser
	q := userRole.WithContext(l.ctx).LeftJoin(sysUser, sysUser.ID.EqCol(userRole.UserID)).Select(sysUser.ALL)
	if len(in.Mobile) > 0 {
		q = q.Where(sysUser.Mobile.Like("%" + in.Mobile + "%"))
	}
	if len(in.UserName) > 0 {
		q = q.Where(sysUser.UserName.Like("%" + in.UserName + "%"))
	}

	// IsExist 1:表示拥有的用户
	if in.IsExist == 1 {
		q = q.Where(userRole.RoleID.Eq(in.RoleId))
	} else {
		// IsExist 0:表示没拥有的用户
		q = q.Where(userRole.RoleID.Neq(in.RoleId))
	}
	offset := (in.PageNum - 1) * in.PageSize
	count, err := q.ScanByPage(&result, int(offset), int(in.PageSize))
	if err != nil {
		logc.Errorf(l.ctx, "查询用户列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户列表信息失败")
	}

	var list []*sysclient.UserData
	for _, item := range result {
		loginTime := item.LoginTime.Format("2006-01-02 15:04:05")
		createTime := item.CreateTime.Format("2006-01-02 15:04:05")
		list = append(list, &sysclient.UserData{
			Id:           item.ID,                                 // 编号
			UserName:     item.UserName,                           // 用户名
			NickName:     item.NickName,                           // 昵称
			Avatar:       item.Avatar,                             // 头像
			Email:        item.Email,                              // 邮箱
			Mobile:       item.Mobile,                             // 手机号
			UserStatus:   item.UserStatus,                         // 帐号状态（0正常 1停用）
			DeptId:       item.DeptID,                             // 部门id
			Remark:       item.Remark,                             // 备注信息
			IsDeleted:    item.IsDeleted,                          // 是否删除  0：否  1：是
			LoginTime:    loginTime,                               // 登录时间
			LoginIp:      item.LoginIP,                            // 登录ip
			LoginOs:      item.LoginOs,                            // 登录os
			LoginBrowser: item.LoginBrowser,                       // 登录浏览器
			CreateBy:     item.CreateBy,                           // 创建者
			CreateTime:   createTime,                              // 创建时间
			UpdateBy:     item.UpdateBy,                           // 更新者
			UpdateTime:   time_util.TimeToString(item.UpdateTime), // 更新时间
		})
	}

	return &sysclient.QueryRoleUserListResp{
		List:  list,
		Total: count,
	}, nil
}
