package roleservicelogic

import (
	"context"
	"errors"
	"fmt"
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
// 1.查询角色是否已存在
// 2.超级管理员不需要分配用户
// 3.如果角色不是admin则根据roleId查询用户
func (l *QueryRoleUserListLogic) QueryRoleUserList(in *sysclient.QueryRoleUserListReq) (*sysclient.QueryRoleUserListResp, error) {

	if in.RoleId == 1 {
		return nil, errors.New(fmt.Sprintf("不允许操作超级管理员角色"))
	}

	// 1.查询角色角色是否已存在
	count, err := query.SysRole.WithContext(l.ctx).Where(query.SysRole.ID.Eq(in.RoleId)).Count()

	if err != nil {
		return nil, errors.New("查询角色失败")
	}

	if count == 0 {
		return nil, errors.New("角色不存在")
	}

	var result []*model.SysUser
	var count1 int64
	// IsExist 1:表示拥有的用户,0:表示没拥有的用户
	if in.IsExist == 1 {
		result, count1, err = QueryAllocatedList(l.ctx, in)
	} else if in.IsExist == 0 {
		result, count1, err = QueryUnallocatedList(l.ctx, in)
	} else {
		return nil, errors.New("参数错误,IsExist只能是0或者1")
	}

	if err != nil {
		logc.Errorf(l.ctx, "查询用户列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户列表信息失败")
	}

	var list = make([]*sysclient.UserData, 0, len(result))

	for _, item := range result {
		list = append(list, &sysclient.UserData{
			Id:            item.ID,                                    // 用户id
			Mobile:        item.Mobile,                                // 手机号码
			UserName:      item.UserName,                              // 用户账号
			NickName:      item.NickName,                              // 用户昵称
			UserType:      item.UserType,                              // 用户类型（00系统用户）
			Avatar:        item.Avatar,                                // 头像路径
			Email:         item.Email,                                 // 用户邮箱
			Password:      item.Password,                              // 密码
			Status:        item.Status,                                // 状态(1:正常，0:禁用)
			DeptId:        item.DeptID,                                // 部门ID
			LoginIp:       item.LoginIP,                               // 最后登录IP
			LoginDate:     time_util.TimeToString(item.LoginDate),     // 最后登录时间
			LoginBrowser:  item.LoginBrowser,                          // 浏览器类型
			LoginOs:       item.LoginOs,                               // 操作系统
			PwdUpdateDate: time_util.TimeToString(item.PwdUpdateDate), // 密码最后更新时间
			Remark:        item.Remark,                                // 备注
			DelFlag:       item.DelFlag,                               // 删除标志（0代表删除 1代表存在）
			CreateBy:      item.CreateBy,                              // 创建者
			CreateTime:    time_util.TimeToStr(item.CreateTime),       // 创建时间
			UpdateBy:      item.UpdateBy,                              // 更新者
			UpdateTime:    time_util.TimeToString(item.UpdateTime),    // 更新时间
		})
	}

	return &sysclient.QueryRoleUserListResp{
		List:  list,
		Total: count1,
	}, nil
}

func QueryAllocatedList(ctx context.Context, in *sysclient.QueryRoleUserListReq) ([]*model.SysUser, int64, error) {
	ur := query.SysUserRole
	u := query.SysUser

	q := u.WithContext(ctx).LeftJoin(ur, ur.UserID.EqCol(u.ID)).Select(u.ALL)
	q = q.Where(u.DelFlag.Eq(1))
	if len(in.Mobile) > 0 {
		q = q.Where(u.Mobile.Like("%" + in.Mobile + "%"))
	}
	if len(in.UserName) > 0 {
		q = q.Where(u.UserName.Like("%" + in.UserName + "%"))
	}
	q = q.Where(ur.RoleID.Eq(in.RoleId))

	offset := (in.PageNum - 1) * in.PageSize
	result, count, err := q.FindByPage(int(offset), int(in.PageSize))
	if err != nil {
		logc.Errorf(ctx, "查询用户列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, 0, errors.New("查询用户列表信息失败")
	}

	return result, count, nil
}

func QueryUnallocatedList(ctx context.Context, in *sysclient.QueryRoleUserListReq) ([]*model.SysUser, int64, error) {
	ur := query.SysUserRole
	u := query.SysUser
	q := u.WithContext(ctx).LeftJoin(ur, ur.UserID.EqCol(u.ID))
	q = q.Where(u.DelFlag.Eq(1))

	q = q.Select(u.ID).Where(ur.RoleID.Eq(in.RoleId))

	var ids []int64
	err := q.Scan(&ids)
	if err != nil {
		logc.Errorf(ctx, "查询用户列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, 0, errors.New("查询用户列表信息失败")
	}

	offset := (in.PageNum - 1) * in.PageSize

	q1 := u.WithContext(ctx).Where(u.ID.NotIn(ids...))
	if len(in.Mobile) > 0 {
		q1 = q1.Where(u.Mobile.Like("%" + in.Mobile + "%"))
	}
	if len(in.UserName) > 0 {
		q1 = q1.Where(u.UserName.Like("%" + in.UserName + "%"))
	}
	users, count, err := q1.FindByPage(int(offset), int(in.PageSize))
	if err != nil {
		logc.Errorf(ctx, "查询用户列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, 0, errors.New("查询用户列表信息失败")
	}

	return users, count, nil

}
