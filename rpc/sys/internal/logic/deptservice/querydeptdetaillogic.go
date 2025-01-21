package deptservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
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
		Id:         dept.ID,                                       // 编号
		DeptName:   dept.DeptName,                                 // 部门名称
		DeptStatus: dept.DeptStatus,                               // 部门状态
		DeptSort:   dept.DeptSort,                                 // 部门排序
		ParentId:   dept.ParentID,                                 // 上级机构ID，一级机构为0
		Leader:     dept.Leader,                                   // 负责人
		Phone:      dept.Phone,                                    // 电话号码
		Email:      dept.Email,                                    // 邮箱
		Remark:     dept.Remark,                                   // 备注信息
		IsDeleted:  dept.IsDeleted,                                // 是否删除  0：否  1：是
		ParentIds:  GetParentIds(dept.ParentIds),                  // 上级机构IDs，一级机构为0
		CreateBy:   dept.CreateBy,                                 // 创建者
		CreateTime: dept.CreateTime.Format("2006-01-02 15:04:05"), // 创建时间
		UpdateBy:   dept.UpdateBy,                                 // 更新者
		UpdateTime: time_util.TimeToString(dept.UpdateTime),       // 更新时间
	}

	logc.Infof(l.ctx, "查询部门信息表详情,参数：%+v,响应：%+v", in, data)

	return data, nil

}
