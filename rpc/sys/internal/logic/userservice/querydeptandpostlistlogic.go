package userservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryDeptAndPostListLogic 查询所有部门和岗位
/*
Author: LiuFeiHua
Date: 2024/5/30 9:22
*/
type QueryDeptAndPostListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDeptAndPostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDeptAndPostListLogic {
	return &QueryDeptAndPostListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryDeptAndPostList 查询所有部门和岗位
// 1.查询所有部门
// 2.查询所有岗位
func (l *QueryDeptAndPostListLogic) QueryDeptAndPostList(in *sysclient.QueryDeptAndPostListReq) (*sysclient.QueryDeptAndPostListResp, error) {

	// 1.查询所有部门
	deptList, _ := query.SysDept.WithContext(l.ctx).Find()
	var deptListData = make([]*sysclient.DeptData, 0, len(deptList))

	for _, dept := range deptList {
		deptListData = append(deptListData, &sysclient.DeptData{
			Id:         dept.ID,         // 编号
			DeptName:   dept.DeptName,   // 部门名称
			DeptStatus: dept.DeptStatus, // 部门状态
			DeptSort:   dept.DeptSort,   // 部门排序
			ParentId:   dept.ParentID,   // 上级机构ID，一级机构为0
			Leader:     dept.Leader,     // 负责人
			Phone:      dept.Phone,      // 电话号码
			Email:      dept.Email,      // 邮箱
			Remark:     dept.Remark,     // 备注信息
		})
	}

	// 2.查询所有岗位
	postList, _ := query.SysPost.WithContext(l.ctx).Find()
	var postListData = make([]*sysclient.PostData, 0, len(postList))

	for _, job := range postList {
		postListData = append(postListData, &sysclient.PostData{
			Id:         job.ID,         // 岗位id
			PostName:   job.PostName,   // 岗位名称
			PostCode:   job.PostCode,   // 岗位编码
			PostStatus: job.PostStatus, // 岗位状态
			PostSort:   job.PostSort,   // 岗位排序
			Remark:     job.Remark,     // 备注信息
		})
	}

	data := &sysclient.QueryDeptAndPostListResp{
		DeptListData: deptListData,
		PostListData: postListData,
	}

	return data, nil
}
