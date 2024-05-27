package deptservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeptUpdateLogic 更新部门信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:01
*/
type DeptUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptUpdateLogic {
	return &DeptUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeptUpdate 更新部门信息
// 1.根据部门id查询部门是否已存在
// 2.如果部门不存在,则直接返回
// 3.部门存在时,则直接更新部门
func (l *DeptUpdateLogic) DeptUpdate(in *sysclient.DeptUpdateReq) (*sysclient.DeptUpdateResp, error) {
	q := query.SysDept.WithContext(l.ctx)

	// 1.根据部门id查询部门是否已存在
	count, err := q.Where(query.SysDept.ID.Eq(in.Id)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据部门id：%d,查询部门信息失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询部门信息失败"))
	}

	// 2.如果部门不存在,则直接返回
	if count == 0 {
		logc.Errorf(l.ctx, "部门信息不存在: %+v", in)
		return nil, errors.New(fmt.Sprintf("部门: %s,不存在", in.DeptName))
	}

	dept := make(map[string]interface{})
	dept["DeptName"] = in.DeptName
	dept["ParentID"] = in.ParentId
	dept["OrderNum"] = in.OrderNum
	dept["UpdateBy"] = in.UpdateBy
	dept["DelFlag"] = in.DelFlag
	dept["ParentIds"] = strings.Replace(strings.Trim(fmt.Sprint(in.ParentIds), "[]"), " ", ",", -1)

	// 3.部门存在时,则直接更新部门
	_, err = q.Where(query.SysDept.ID.Eq(in.Id)).Updates(dept)

	if err != nil {
		logc.Errorf(l.ctx, "更新部门信息失败,参数:%+v,异常:%s", dept, err.Error())
		return nil, errors.New(fmt.Sprintf("更新部门信息失败"))
	}

	return &sysclient.DeptUpdateResp{}, nil
}
