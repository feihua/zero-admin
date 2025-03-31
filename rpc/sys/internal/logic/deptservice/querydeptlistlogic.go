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

// QueryDeptListLogic 查询部门列表
/*
Author: LiuFeiHua
Date: 2023/12/18 17:00
*/
type QueryDeptListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDeptListLogic {
	return &QueryDeptListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryDeptList 查询部门列表
func (l *QueryDeptListLogic) QueryDeptList(in *sysclient.QueryDeptListReq) (*sysclient.QueryDeptListResp, error) {

	result, err := query.SysDept.WithContext(l.ctx).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询部门列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询部门列表失败")
	}

	var list = make([]*sysclient.DeptListData, 0, len(result))

	for _, dept := range result {
		list = append(list, &sysclient.DeptListData{
			Id:         dept.ID,                                 // 部门id
			ParentId:   dept.ParentID,                           // 上级部门id
			Ancestors:  dept.Ancestors,                          // 祖级列表
			DeptName:   dept.DeptName,                           // 部门名称
			Sort:       dept.Sort,                               // 显示顺序
			Leader:     dept.Leader,                             // 负责人
			Phone:      dept.Phone,                              // 联系电话
			Email:      dept.Email,                              // 邮箱
			Status:     dept.Status,                             // 部门状态（0：停用，1:正常）
			DelFlag:    dept.DelFlag,                            // 删除标志（0代表删除 1代表存在）
			Remark:     dept.Remark,                             // 备注信息
			CreateBy:   dept.CreateBy,                           // 创建者
			CreateTime: time_util.TimeToStr(dept.CreateTime),    // 创建时间
			UpdateBy:   dept.UpdateBy,                           // 更新者
			UpdateTime: time_util.TimeToString(dept.UpdateTime), // 更新时间
		})
	}

	return &sysclient.QueryDeptListResp{
		List: list,
	}, nil

}
