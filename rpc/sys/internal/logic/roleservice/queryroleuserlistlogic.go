package roleservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/logic/common"
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
	//role := query.SysRole
	//total, _ := role.WithContext(l.ctx).Where(role.ID.Eq(in.RoleId), role.IsAdmin.Eq(1)).Count()
	//
	//// 1.超级管理员不需要分配用户
	//if total != 0 {
	//	return nil, errors.New("超级管理员不需要分配用户")
	//}

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

	//IsExist 1:表示拥有的用户
	if in.IsExist == 1 {
		q = q.Where(userRole.RoleID.Eq(in.RoleId))
	} else {
		//IsExist 0:表示没拥有的用户
		q = q.Where(userRole.RoleID.Neq(in.RoleId))
	}
	offset := (in.PageNum - 1) * in.PageSize
	count, err := q.ScanByPage(&result, int(offset), int(in.PageSize))
	if err != nil {
		logc.Errorf(l.ctx, "查询用户列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户列表信息失败")
	}

	var list []*sysclient.UserListData
	for _, item := range result {
		list = append(list, &sysclient.UserListData{
			Avatar:     item.Avatar,
			CreateBy:   item.CreateBy,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
			DeptId:     item.DeptID,
			Email:      item.Email,
			Id:         item.ID,
			UserStatus: item.UserStatus,
			LoginIp:    item.LoginIP,
			LoginTime:  common.TimeToString(item.LoginTime),
			Mobile:     item.Mobile,
			NickName:   item.NickName,
			Remark:     item.Remark,
			UpdateBy:   item.UpdateBy,
			UpdateTime: common.TimeToString(item.UpdateTime),
			UserName:   item.UserName,
		})
	}

	return &sysclient.QueryRoleUserListResp{
		List:  list,
		Total: count,
	}, nil
}
