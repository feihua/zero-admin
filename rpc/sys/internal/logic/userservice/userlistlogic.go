package userservicelogic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sysclient"
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

	all, err := l.svcCtx.UserModel.FindAll(l.ctx, in)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logc.Errorf(l.ctx, "查询用户列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	count, _ := l.svcCtx.UserModel.Count(l.ctx, in)

	var list []*sysclient.UserListData
	for _, user := range *all {
		list = append(list, &sysclient.UserListData{
			Id:             user.Id,
			Name:           user.Name,
			NickName:       user.NickName.String,
			Avatar:         user.Avatar.String,
			Email:          user.Email.String,
			Mobile:         user.Mobile.String,
			DeptId:         user.DeptId,
			Status:         user.Status,
			CreateBy:       user.CreateBy,
			CreateTime:     user.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   user.UpdateBy.String,
			LastUpdateTime: user.UpdateTime.Time.Format("2006-01-02 15:04:05"),
			DelFlag:        user.DelFlag,
			JobId:          user.JobId,
			RoleId:         user.RoleId,
			RoleName:       user.RoleName,
			JobName:        user.JobName,
			DeptName:       user.DeptName,
		})
	}

	logc.Infof(l.ctx, "查询用户列表信息,参数：%+v,响应：%+v", in, list)
	return &sysclient.UserListResp{
		Total: count,
		List:  list,
	}, nil
}
