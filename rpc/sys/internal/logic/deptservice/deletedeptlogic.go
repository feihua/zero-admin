package deptservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteDeptLogic 删除部门信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:00
*/
type DeleteDeptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDeptLogic {
	return &DeleteDeptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteDept 删除部门信息
func (l *DeleteDeptLogic) DeleteDept(in *sysclient.DeleteDeptReq) (*sysclient.DeleteDeptResp, error) {
	q := query.SysDept

	id := in.Id
	if id == 1 {
		return nil, errors.New("顶级部门,不允许删除")
	}

	count, err := q.WithContext(l.ctx).Where(q.ParentID.Eq(id)).Count()
	if err != nil {
		logc.Errorf(l.ctx, "根据父部门id查询下级部门数量失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除部门信息失败")
	}
	if count > 0 {
		return nil, errors.New("存在下级部门,不允许删除")
	}

	count, err = query.SysUser.WithContext(l.ctx).Where(query.SysUser.DeptID.Eq(id)).Count()
	if err != nil {
		logc.Errorf(l.ctx, "查询部门是否存在用户失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询部门是否存在用户失败")
	}
	if count > 0 {
		return nil, errors.New("部门存在用户,不允许删除")
	}

	_, err = q.WithContext(l.ctx).Where(q.ID.Eq(id)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除部门信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除部门信息失败")
	}

	return &sysclient.DeleteDeptResp{}, nil
}
