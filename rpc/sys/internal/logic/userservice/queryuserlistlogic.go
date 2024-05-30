package userservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/logic/common"
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

	logc.Infof(l.ctx, "查询用户列表信息,参数：%+v,响应：%+v", in, list)
	return &sysclient.QueryUserListResp{
		Total: count,
		List:  list,
	}, nil
}
