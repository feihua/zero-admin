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
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// AddDeptLogic 添加部门
/*
Author: LiuFeiHua
Date: 2023/12/18 16:59
*/
type AddDeptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDeptLogic {
	return &AddDeptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddDept 添加部门
// 1.根据部门名称查询部门是否已存在
// 2.如果部门已存在,则直接返回
// 3.如果父节点不为正常状态,则不允许新增子节点
// 4.部门不存在时,则直接添加部门
func (l *AddDeptLogic) AddDept(in *sysclient.AddDeptReq) (*sysclient.AddDeptResp, error) {
	q := query.SysDept.WithContext(l.ctx)

	// 1.根据部门名称查询部门是否已存在
	name := in.DeptName
	count, err := q.Where(query.SysDept.DeptName.Eq(name), query.SysDept.ParentID.Eq(in.ParentId)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据部门名称：%s,查询部门失败,异常:%s", name, err.Error())
		return nil, errors.New(fmt.Sprintf("添加部门失败"))
	}

	// 2.如果部门已存在,则直接返回
	if count > 0 {
		logc.Errorf(l.ctx, "部门名称已存在：%+v", in)
		return nil, errors.New(fmt.Sprintf("部门名称：%s,已存在", name))
	}

	// 3.如果父节点不为正常状态,则不允许新增子节点
	parentDept, err := q.Where(query.SysDept.ID.Eq(in.ParentId)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "上级部门不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("添加部门失败,上级部门不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询上级部门异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("添加部门失败,查询上级部门异常")
	}

	if parentDept.Status != 1 {
		return nil, errors.New(fmt.Sprintf("添加部门失败,：%s,停用，不允许新增", parentDept.DeptName))
	}

	ancestors := fmt.Sprintf("%s,%d", parentDept.Ancestors, parentDept.ID)
	// 4.部门不存在时,则直接添加部门
	dept := &model.SysDept{
		ParentID:  in.ParentId, // 上级部门id
		Ancestors: ancestors,   // 祖级列表
		DeptName:  in.DeptName, // 部门名称
		Sort:      in.Sort,     // 显示顺序
		Leader:    in.Leader,   // 负责人
		Phone:     in.Phone,    // 联系电话
		Email:     in.Email,    // 邮箱
		Status:    in.Status,   // 部门状态（0：停用，1:正常）
		DelFlag:   1,           // 删除标志（0代表删除 1代表存在）
		Remark:    in.Remark,   // 备注信息
		CreateBy:  in.CreateBy, // 创建者
	}

	err = q.Create(dept)

	if err != nil {
		logc.Errorf(l.ctx, "添加部门失败,参数:%+v,异常:%s", dept, err.Error())
		return nil, errors.New("添加部门失败")
	}

	return &sysclient.AddDeptResp{}, nil
}
