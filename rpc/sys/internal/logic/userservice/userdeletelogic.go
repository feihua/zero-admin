package userservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// UserDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 14:20
*/
type UserDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserDelete 删除用户(id为1是系统预留超级管理员用户,不能删除)
func (l *UserDeleteLogic) UserDelete(in *sysclient.UserDeleteReq) (*sysclient.UserDeleteResp, error) {

	//判断是否包含超级管理员
	var flag = false
	for id := range in.Ids {
		if id == 1 {
			flag = true
			break
		}
	}

	//id为1是系统预留超级管理员用户,不能删除
	if flag {
		logc.Errorf(l.ctx, "删除用户异常,参数userId:%+v,异常:%s", in.Ids, "id为1是系统预留超级管理员用户,不能删除")
		return nil, errors.New("删除用户异常")
	}

	err := l.svcCtx.UserModel.DeleteByIds(l.ctx, in.Ids)

	//删除用户与角色的关联
	for _, id := range in.Ids {
		_ = l.svcCtx.UserRoleModel.DeleteByUserId(l.ctx, id)
	}

	if err != nil {
		logc.Errorf(l.ctx, "删除用户异常,参数userId:%+v,异常:%s", in.Ids, err.Error())
		return nil, errors.New("删除用户异常")
	}
	return &sysclient.UserDeleteResp{}, nil
}
