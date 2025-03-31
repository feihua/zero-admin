package userservicelogic

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

// QueryUserListLogic 查询用户列表
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

// QueryUserList 查询用户列表
func (l *QueryUserListLogic) QueryUserList(in *sysclient.QueryUserListReq) (*sysclient.QueryUserListResp, error) {

	user := query.SysUser
	q := user.WithContext(l.ctx)
	if in.DeptId != 0 {
		var deptIds []int64
		l.svcCtx.DB.Model(model.SysDept{}).WithContext(l.ctx).Select("id").Where("find_in_set(?, ancestors)", in.DeptId).Scan(&deptIds)
		deptIds = append(deptIds, in.DeptId)
		q = q.Where(user.DeptID.In(deptIds...))
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
	if len(in.UserName) > 0 {
		q = q.Where(user.UserName.Like("%" + in.UserName + "%"))
	}
	if in.Status != 2 {
		q = q.Where(user.Status.Eq(in.Status))
	}

	offset := (in.PageNum - 1) * in.PageSize
	result, count, err := q.FindByPage(int(offset), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询用户列表失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户列表失败")
	}

	var list = make([]*sysclient.UserListData, 0, len(result))
	for _, item := range result {
		list = append(list, &sysclient.UserListData{
			Id:            item.ID,                                    // 用户id
			Mobile:        item.Mobile,                                // 手机号码
			UserName:      item.UserName,                              // 用户账号
			NickName:      item.NickName,                              // 用户昵称
			UserType:      item.UserType,                              // 用户类型（00系统用户）
			Avatar:        item.Avatar,                                // 头像路径
			Email:         item.Email,                                 // 用户邮箱
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

	return &sysclient.QueryUserListResp{
		Total: count,
		List:  list,
	}, nil
}
