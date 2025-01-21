package deptservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateDeptLogic 更新部门信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:01
*/
type UpdateDeptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeptLogic {
	return &UpdateDeptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateDept 更新部门信息
// 1.根据部门id查询部门是否已存在
// 2.根据部门parentId查询部门是否已存在
// 3.根据部门名称查询部门是否已存在
// 4.查询是否有下级部门
// 5.修改下级部门祖级
// 6.部门存在时,则直接更新部门
// 7.如果该部门是启用状态，则启用该部门的所有上级部门
func (l *UpdateDeptLogic) UpdateDept(in *sysclient.UpdateDeptReq) (*sysclient.UpdateDeptResp, error) {

	dept := query.SysDept
	q := dept.WithContext(l.ctx)

	if in.ParentId == in.Id {
		return nil, errors.New("上级部门不能为当前部门")
	}

	// 1.根据部门id查询部门是否已存在
	oldDept, err := q.Where(dept.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据部门id：%d,查询部门信息失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("更新部门失败"))
	}

	if oldDept == nil {
		return nil, errors.New("更新部门失败,部门不存在")
	}

	// 2.根据部门parentId查询部门是否已存在
	parentDept, err := q.Where(dept.ID.Eq(in.ParentId)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据部门id：%d,查询部门信息失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("更新部门失败"))
	}

	if parentDept == nil {
		return nil, errors.New("更新部门失败,上级部门不存在")
	}

	// 3.根据部门名称查询部门是否已存在
	deptName := in.DeptName
	count, err := q.Where(dept.DeptName.Eq(deptName), dept.ParentID.Eq(in.ParentId)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据部门名称：%s,查询部门信息失败,异常:%s", deptName, err.Error())
		return nil, errors.New(fmt.Sprintf("查询部门信息失败"))
	}

	if count > 0 {
		logc.Errorf(l.ctx, "部门信息已存在：%+v", in)
		return nil, errors.New(fmt.Sprintf("部门：%s,已存在", deptName))
	}

	// 4.查询是否有下级部门
	sql := "select count(*) from sys_dept where status = 1 and del_flag = 1 and find_in_set(?, 'parentIds')"
	err = l.svcCtx.DB.Raw(sql, in.Id).Count(&count).Error
	if err != nil {
		logc.Errorf(l.ctx, "根据部门id查询是否有下级部门失败,异常:%s", err.Error())
		return nil, errors.New(fmt.Sprintf("更新部门失败"))
	}

	if count > 0 && in.DeptStatus == 0 {
		return nil, errors.New(fmt.Sprintf("该部门包含未停用的子部门"))
	}

	sql = "select * from sys_dept where find_in_set(?, 'parentIds')"
	list := make([]model.SysDept, 10)
	err = l.svcCtx.DB.Model(&model.SysDept{}).Raw(sql, in.Id).Scan(list).Error
	if err != nil {
		logc.Errorf(l.ctx, "根据部门id查询是否有下级部门失败,异常:%s", err.Error())
		return nil, errors.New(fmt.Sprintf("更新部门失败"))
	}

	// 5.修改下级部门祖级
	parentIds := strings.Replace(strings.Trim(fmt.Sprint(in.ParentIds), "[]"), " ", ",", -1) // 上级机构IDs，一级机构为0
	if count > 0 {
		for _, dept1 := range list {
			parentIdStr := strings.Replace(dept1.ParentIds, oldDept.ParentIds, parentIds, -1)
			_, err = q.Where(dept.ID.Eq(dept1.ID)).Update(dept.ParentIds, parentIdStr)
			if err != nil {
				logc.Errorf(l.ctx, "修改下级部门祖级失败,异常:%s", err.Error())
				return nil, errors.New(fmt.Sprintf("更新部门失败"))
			}
		}
	}

	sysDept := &model.SysDept{
		ID:         in.Id,         // 编号
		DeptName:   in.DeptName,   // 部门名称
		DeptStatus: in.DeptStatus, // 部门状态
		DeptSort:   in.DeptSort,   // 部门排序
		ParentID:   in.ParentId,   // 上级机构ID，一级机构为0
		Leader:     in.Leader,     // 负责人
		Phone:      in.Phone,      // 电话号码
		Email:      in.Email,      // 邮箱
		Remark:     in.Remark,     // 备注信息
		UpdateBy:   in.UpdateBy,   // 更新者
		ParentIds:  parentIds,     // 上级机构IDs，一级机构为0
	}

	// 6.部门存在时,则直接更新部门
	_, err = q.Where(dept.ID.Eq(in.Id)).Updates(sysDept)

	if err != nil {
		logc.Errorf(l.ctx, "更新部门信息失败,参数:%+v,异常:%s", dept, err.Error())
		return nil, errors.New(fmt.Sprintf("更新部门信息失败"))
	}

	// 7.如果该部门是启用状态，则启用该部门的所有上级部门
	if in.DeptStatus == 1 {
		_, err = q.Where(dept.ID.In(in.ParentIds...)).Update(dept.DeptStatus, in.DeptStatus)
		if err != nil {
			logc.Errorf(l.ctx, "修改上级部门状态失败,异常:%s", err.Error())
			return nil, errors.New(fmt.Sprintf("更新部门失败"))
		}

	}
	return &sysclient.UpdateDeptResp{}, nil
}
