package deptservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateDeptStatusLogic 更新部门信息表状态
/*
Author: LiuFeiHua
Date: 2024/5/30 9:59
*/
type UpdateDeptStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDeptStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeptStatusLogic {
	return &UpdateDeptStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateDeptStatus 更新部门信息表状态
func (l *UpdateDeptStatusLogic) UpdateDeptStatus(in *sysclient.UpdateDeptStatusReq) (*sysclient.UpdateDeptStatusResp, error) {
	q := query.SysDept

	if in.DeptStatus == 1 {
		for _, id := range in.Ids {
			dept, err := query.SysDept.WithContext(l.ctx).Where(q.ID.Eq(id)).First()

			if err != nil {
				logc.Errorf(l.ctx, "更新部门信息表状态失败,参数:%+v,异常:%s", in, err.Error())
				return nil, errors.New("更新部门信息表状态失败")
			}

			if dept == nil {
				return nil, errors.New("更新部门信息表状态失败,部门不存在")
			}

			parentIds := GetParentIds(dept.ParentIds)
			_, err = query.SysDept.WithContext(l.ctx).Where(q.ID.In(parentIds...)).Update(q.DeptStatus, in.DeptStatus)
			if err != nil {
				logc.Errorf(l.ctx, "修改上级部门状态失败,异常:%s", err.Error())
				return nil, errors.New(fmt.Sprintf("更新部门信息表状态失败"))
			}
		}

	}

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.DeptStatus, in.DeptStatus)

	if err != nil {
		logc.Errorf(l.ctx, "根据部门ids：%+v,更新部门信息表状态失败,异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("更新部门信息表状态失败"))
	}

	return &sysclient.UpdateDeptStatusResp{}, nil
}
