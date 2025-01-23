package deptservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strconv"
	"strings"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

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
	q := query.SysDept.WithContext(l.ctx)

	result, err := q.Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询部门列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询部门列表失败")
	}

	var list = make([]*sysclient.DeptListData, 0, len(result))

	for _, dept := range result {
		list = append(list, &sysclient.DeptListData{
			Id:         dept.ID,                                 // 编号
			DeptName:   dept.DeptName,                           // 部门名称
			DeptStatus: dept.DeptStatus,                         // 部门状态
			DeptSort:   dept.DeptSort,                           // 部门排序
			ParentId:   dept.ParentID,                           // 上级机构ID，一级机构为0
			Leader:     dept.Leader,                             // 负责人
			Phone:      dept.Phone,                              // 电话号码
			Email:      dept.Email,                              // 邮箱
			Remark:     dept.Remark,                             // 备注
			IsDeleted:  dept.IsDeleted,                          // 是否删除  0：否  1：是
			ParentIds:  GetParentIds(dept.ParentIds),            // 上级机构IDs，一级机构为0
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

// GetParentIds 获取父级ids
func GetParentIds(parentIdsStr string) []int64 {
	var parentIds []int64
	if len(parentIdsStr) > 0 {
		for _, i := range strings.Split(parentIdsStr, ",") {
			p, _ := strconv.ParseInt(i, 10, 64)
			parentIds = append(parentIds, p)
		}
	}
	return parentIds
}
