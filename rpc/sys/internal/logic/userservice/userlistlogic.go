package userservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/logic/common"
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
	if len(in.UserName) > 0 {
		q = q.Where(query.SysUser.UserName.Like("%" + in.UserName + "%"))
	}
	if len(in.Mobile) > 0 {
		q = q.Where(query.SysUser.Mobile.Like("%" + in.Mobile + "%"))
	}

	if in.UserStatus != 2 {
		q = q.Where(query.SysUser.UserStatus.Eq(in.UserStatus))
	}

	offset := (in.Current - 1) * in.PageSize
	result, count, err := q.FindByPage(int(offset), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询用户列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*sysclient.UserListData
	for _, user := range result {
		list = append(list, &sysclient.UserListData{
			Id:         user.ID,
			UserName:   user.UserName,
			NickName:   user.NickName,
			Avatar:     user.Avatar,
			Email:      user.Email,
			Mobile:     user.Mobile,
			DeptId:     user.DeptID,
			UserStatus: user.UserStatus,
			CreateBy:   user.CreateBy,
			CreateTime: user.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateBy:   user.UpdateBy,
			UpdateTime: common.TimeToString(user.UpdateTime),
			DelFlag:    user.DelFlag,
			JobId:      user.JobID,
		})
	}

	logc.Infof(l.ctx, "查询用户列表信息,参数：%+v,响应：%+v", in, list)
	return &sysclient.UserListResp{
		Total: count,
		List:  list,
	}, nil
}
