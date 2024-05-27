package deptservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/logic/common"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strconv"
	"strings"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeptListLogic 查询部门信息列表
/*
Author: LiuFeiHua
Date: 2023/12/18 17:00
*/
type DeptListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptListLogic {
	return &DeptListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeptList 查询部门信息列表
func (l *DeptListLogic) DeptList(in *sysclient.DeptListReq) (*sysclient.DeptListResp, error) {
	q := query.SysDept.WithContext(l.ctx)

	result, err := q.Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询部门列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询部门列表信息失败")
	}

	var list []*sysclient.DeptListData

	for _, dept := range result {
		list = append(list, &sysclient.DeptListData{
			Id:         dept.ID,
			DeptName:   dept.DeptName,
			ParentId:   dept.ParentID,
			OrderNum:   dept.OrderNum,
			CreateBy:   dept.CreateBy,
			CreateTime: dept.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateBy:   dept.UpdateBy,
			UpdateTime: common.TimeToString(dept.UpdateTime),
			DelFlag:    dept.DelFlag,
			ParentIds:  getParentIds(dept.ParentIds),
		})
	}

	logc.Infof(l.ctx, "查询部门列表信息,参数：%+v,响应：%+v", in, list)
	return &sysclient.DeptListResp{
		List: list,
	}, nil

}

// 获取父级ids
func getParentIds(parentIdsStr string) []int64 {
	var parentIds []int64
	if len(parentIdsStr) > 0 {
		for _, i := range strings.Split(parentIdsStr, ",") {
			p, _ := strconv.ParseInt(i, 10, 64)
			parentIds = append(parentIds, p)
		}
	}
	return parentIds
}
