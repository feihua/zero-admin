package deptservicelogic

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

// QueryDeptDetailLogic 查询部门信息表详情
/*
Author: LiuFeiHua
Date: 2024/5/30 10:13
*/
type QueryDeptDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDeptDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDeptDetailLogic {
	return &QueryDeptDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryDeptDetail 查询部门信息表详情
func (l *QueryDeptDetailLogic) QueryDeptDetail(in *sysclient.QueryDeptDetailReq) (*sysclient.QueryDeptDetailResp, error) {
	dept, err := query.SysDept.WithContext(l.ctx).Where(query.SysDept.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询部门信息表详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询部门信息表详情失败")
	}

	data := &sysclient.QueryDeptDetailResp{
		CreateBy:   dept.CreateBy,
		CreateTime: dept.CreateTime.Format("2006-01-02 15:04:05"),
		DeptName:   dept.DeptName,
		DeptSort:   dept.DeptSort,
		DeptStatus: dept.DeptStatus,
		Email:      dept.Email,
		Id:         dept.ID,
		Leader:     dept.Leader,
		ParentId:   dept.ParentID,
		ParentIds:  GetParentIds(dept.ParentIds),
		Phone:      dept.Phone,
		Remark:     dept.Remark,
		UpdateBy:   dept.UpdateBy,
		UpdateTime: common.TimeToString(dept.UpdateTime),
	}

	logc.Infof(l.ctx, "查询部门信息表详情,参数：%+v,响应：%+v", in, data)

	return data, nil

}
