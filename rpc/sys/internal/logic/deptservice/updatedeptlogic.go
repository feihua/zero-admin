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
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"

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
// 8.如果该部门是禁用状态，则禁用该部门的所有下级部门
func (l *UpdateDeptLogic) UpdateDept(in *sysclient.UpdateDeptReq) (*sysclient.UpdateDeptResp, error) {

	dept := query.SysDept
	q := dept.WithContext(l.ctx)

	if in.ParentId == in.Id {
		return nil, errors.New("上级部门不能为当前部门")
	}

	// 1.根据部门id查询部门是否已存在
	oldDept, err := q.Where(dept.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "更新部门失败,部门不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("更新部门失败,部门不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询部门异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询部门异常")
	}

	// 2.根据部门parentId查询部门是否已存在
	parentDept, err := q.Where(dept.ID.Eq(in.ParentId)).First()

	// 1.判断上级部门是否存在
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "上级部门不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("上级部门不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询上级部门异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询上级部门异常")
	}

	ancestors := fmt.Sprintf("%s,%d", parentDept.Ancestors, parentDept.ID)
	// 3.根据部门名称查询部门是否已存在
	deptName := in.DeptName
	count, err := q.Where(dept.ID.Neq(in.Id), dept.DeptName.Eq(deptName), dept.ParentID.Eq(parentDept.ID)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据部门名称：%s,查询部门信息失败,异常:%s", deptName, err.Error())
		return nil, errors.New(fmt.Sprintf("查询部门信息失败"))
	}

	if count > 0 {
		logc.Errorf(l.ctx, "部门信息已存在：%+v", in)
		return nil, errors.New(fmt.Sprintf("部门：%s,已存在", deptName))
	}

	// 4.查询是否有下级部门
	sql := "select count(*) from sys_dept where status = 1 and del_flag = 1 and find_in_set(?, 'ancestors')"
	err = l.svcCtx.DB.Raw(sql, in.Id).Count(&count).Error
	if err != nil {
		logc.Errorf(l.ctx, "根据部门id查询是否有下级部门失败,异常:%s", err.Error())
		return nil, errors.New(fmt.Sprintf("更新部门失败"))
	}

	if count > 0 && in.Status == 0 {
		return nil, errors.New(fmt.Sprintf("该部门包含未停用的子部门"))
	}

	sql = "select * from sys_dept where find_in_set(?, 'ancestors')"
	list := make([]model.SysDept, 10)
	err = l.svcCtx.DB.Model(&model.SysDept{}).Raw(sql, in.Id).Scan(list).Error
	if err != nil {
		logc.Errorf(l.ctx, "根据部门id查询是否有下级部门失败,异常:%s", err.Error())
		return nil, errors.New(fmt.Sprintf("更新部门失败"))
	}

	// 5.修改下级部门祖级
	if count > 0 {
		for _, dept1 := range list {
			newAncestors := strings.Replace(dept1.Ancestors, oldDept.Ancestors, ancestors, -1)
			_, err = q.Where(dept.ID.Eq(dept1.ID)).Update(dept.Ancestors, newAncestors)
			if err != nil {
				logc.Errorf(l.ctx, "修改下级部门祖级失败,异常:%s", err.Error())
				return nil, errors.New(fmt.Sprintf("更新部门失败"))
			}
		}
	}
	now := time.Now()
	sysDept := &model.SysDept{
		ID:         in.Id,              // 部门id
		ParentID:   in.ParentId,        // 上级部门id
		Ancestors:  ancestors,          // 祖级列表
		DeptName:   in.DeptName,        // 部门名称
		Sort:       in.Sort,            // 显示顺序
		Leader:     in.Leader,          // 负责人
		Phone:      in.Phone,           // 联系电话
		Email:      in.Email,           // 邮箱
		Status:     in.Status,          // 部门状态（0：停用，1:正常）
		DelFlag:    oldDept.DelFlag,    // 删除标志（0代表删除 1代表存在）
		Remark:     in.Remark,          // 备注信息
		CreateBy:   oldDept.CreateBy,   // 创建者
		CreateTime: oldDept.CreateTime, // 创建时间
		UpdateBy:   in.UpdateBy,        // 更新者
		UpdateTime: &now,               // 更新时间
	}

	// 6.部门存在时,则直接更新部门
	err = l.svcCtx.DB.Model(&model.SysDept{}).WithContext(l.ctx).Where(dept.ID.Eq(in.Id)).Save(sysDept).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新部门信息失败,参数:%+v,异常:%s", dept, err.Error())
		return nil, errors.New(fmt.Sprintf("更新部门信息失败"))
	}

	// 7.如果该部门是启用状态，则启用该部门的所有上级部门
	if in.Status == 1 {
		_, err = q.Where(dept.ID.In(GetParentIds(ancestors)...)).Update(dept.Status, in.Status)
		if err != nil {
			logc.Errorf(l.ctx, "修改上级部门状态失败,异常:%s", err.Error())
			return nil, errors.New(fmt.Sprintf("更新部门失败"))
		}

	}

	key := l.svcCtx.RedisKey + "dept"
	filed := strconv.FormatInt(in.Id, 10)
	_, _ = l.svcCtx.Redis.HdelCtx(l.ctx, key, filed)
	return &sysclient.UpdateDeptResp{}, nil
}

func GetParentIds(ancestors string) []int64 {
	var parentIds []int64
	if len(ancestors) > 0 {
		for _, i := range strings.Split(ancestors, ",") {
			p, _ := strconv.ParseInt(i, 10, 64)
			parentIds = append(parentIds, p)
		}
	}
	return parentIds
}
