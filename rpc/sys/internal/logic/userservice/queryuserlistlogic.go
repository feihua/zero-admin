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

// QueryUserListLogic 查询用户列表信息
/*
Author: LiuFeiHua
Date: 2023/12/18 14:35
*/
type QueryUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserListLogic {
	return &QueryUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryUserList 查询用户列表信息
func (l *QueryUserListLogic) QueryUserList(in *sysclient.QueryUserListReq) (*sysclient.QueryUserListResp, error) {

	user := query.SysUser
	q := user.WithContext(l.ctx)
	if in.DeptId != 0 {
		q = q.Where(user.DeptID.Eq(in.DeptId))
	}
	if len(in.Email) > 0 {
		q = q.Where(user.Email.Like("%" + in.Email + "%"))
	}
	if len(in.Mobile) > 0 {
		q = q.Where(user.Mobile.Like("%" + in.Mobile + "%"))
	}
	if len(in.NickName) > 0 {
		q = q.Where(user.NickName.Like("%" + in.NickName + "%"))
	}
	if in.UserStatus != 2 {
		q = q.Where(user.UserStatus.Eq(in.UserStatus))
	}

	offset := (in.PageNum - 1) * in.PageSize
	result, count, err := q.FindByPage(int(offset), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询用户列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户列表信息失败")
	}

	var list []*sysclient.UserListData
	for _, item := range result {
		loginTime := item.LoginTime.Format("2006-01-02 15:04:05")
		createTime := item.CreateTime.Format("2006-01-02 15:04:05")
		list = append(list, &sysclient.UserListData{
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

	logc.Infof(l.ctx, "查询用户列表信息,参数：%+v,响应：%+v", in, list)
	return &sysclient.QueryUserListResp{
		Total: count,
		List:  list,
	}, nil
}
