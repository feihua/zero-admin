package deptservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

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
// 1.判断部门是否存在
// 2.查询是否有下级部门
func (l *UpdateDeptStatusLogic) UpdateDeptStatus(in *sysclient.UpdateDeptStatusReq) (*sysclient.UpdateDeptStatusResp, error) {
	status := in.DeptStatus // 1.判断状态是否正确

	q := query.SysDept

	for _, id := range in.Ids {
		dept, err := query.SysDept.WithContext(l.ctx).Where(q.ID.Eq(id)).First()

		// 1.判断部门是否存在
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			logc.Errorf(l.ctx, "部门不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
			return nil, errors.New("部门不存在")
		case err != nil:
			logc.Errorf(l.ctx, "查询部门异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
			return nil, errors.New("查询部门异常")
		}

		// 2.查询是否有下级部门
		var count int64
		sql := "select count(*) from sys_dept where status = 1 and del_flag = 1 and find_in_set(?, 'parentIds')"
		err = l.svcCtx.DB.Raw(sql, id).Count(&count).Error
		if err != nil {
			logc.Errorf(l.ctx, "根据部门id查询是否有下级部门失败,异常:%s", err.Error())
			return nil, errors.New(fmt.Sprintf("更新部门失败"))
		}

		if count > 0 && in.DeptStatus == 0 {
			return nil, errors.New(fmt.Sprintf("该部门包含未停用的子部门"))
		}

		if status == 1 {
			parentIds := GetParentIds(dept.ParentIds)
			_, err = query.SysDept.WithContext(l.ctx).Where(q.ID.In(parentIds...)).Update(q.DeptStatus, status)
			if err != nil {
				logc.Errorf(l.ctx, "修改上级部门状态失败,异常:%s", err.Error())
				return nil, errors.New(fmt.Sprintf("更新部门信息表状态失败"))
			}
		}

	}

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.DeptStatus, status)

	if err != nil {
		logc.Errorf(l.ctx, "根据部门ids：%+v,更新部门信息表状态失败,异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("更新部门信息表状态失败"))
	}

	return &sysclient.UpdateDeptStatusResp{}, nil
}
