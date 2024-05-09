package roleservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// RoleDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:55
*/
type RoleDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RoleDelete 删除角色(id为1的是系统预留超级管理员角色,不能删除)
func (l *RoleDeleteLogic) RoleDelete(in *sysclient.RoleDeleteReq) (*sysclient.RoleDeleteResp, error) {

	var flag = false
	for _, id := range in.Ids {
		if id == 1 {
			flag = true
			break
		}
	}

	if flag {
		logc.Errorf(l.ctx, "删除角色信息失败,参数:%+v,异常:%s", in, "不能删除超级管理员角色")
		return nil, errors.New("不能删除超级管理员角色")
	}

	q := query.SysRole
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}
	return &sysclient.RoleDeleteResp{}, nil
}
