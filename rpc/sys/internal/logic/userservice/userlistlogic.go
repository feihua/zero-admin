package userservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UserListLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 14:35
*/
type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserList 查询用户列表信息
func (l *UserListLogic) UserList(in *sysclient.UserListReq) (*sysclient.UserListResp, error) {

	q := query.SysUser.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(query.SysUser.Name.Like("%" + in.Name + "%"))
	}
	if len(in.Mobile) > 0 {
		q = q.Where(query.SysUser.Name.Like("%" + in.Mobile + "%"))
	}

	if in.Status != 2 {
		q = q.Where(query.SysUser.Status.Eq(in.Status))
	}

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询用户列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*sysclient.UserListData
	for _, user := range result {
		list = append(list, &sysclient.UserListData{
			Id:         user.ID,
			Name:       user.Name,
			NickName:   *user.NickName,
			Avatar:     *user.Avatar,
			Email:      *user.Email,
			Mobile:     *user.Mobile,
			DeptId:     user.DeptID,
			Status:     user.Status,
			CreateBy:   user.CreateBy,
			CreateTime: user.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateBy:   *user.UpdateBy,
			UpdateTime: user.UpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:    user.DelFlag,
			JobId:      user.JobID,
			//RoleId:         user.RoleId,
			//RoleName:       user.RoleName,
			//JobName:        user.JobName,
			//DeptName:       user.DeptName,
		})
	}

	logc.Infof(l.ctx, "查询用户列表信息,参数：%+v,响应：%+v", in, list)
	return &sysclient.UserListResp{
		Total: count,
		List:  list,
	}, nil
}
