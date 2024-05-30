package userservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/logic/common"
	deptservicelogic "github.com/feihua/zero-admin/rpc/sys/internal/logic/deptservice"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

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
func (l *QueryDeptAndPostListLogic) QueryDeptAndPostList(in *sysclient.QueryDeptAndPostListReq) (*sysclient.QueryDeptAndPostListResp, error) {

	//1.查询所有部门
	var deptListData []*sysclient.DeptListData
	deptList, _ := query.SysDept.WithContext(l.ctx).Find()
	for _, item := range deptList {
		deptListData = append(deptListData, &sysclient.DeptListData{
			CreateBy:   item.CreateBy,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
			DeptName:   item.DeptName,
			DeptSort:   item.DeptSort,
			DeptStatus: item.DeptStatus,
			Email:      item.Email,
			Id:         item.ID,
			Leader:     item.Leader,
			ParentId:   item.ParentID,
			ParentIds:  deptservicelogic.GetParentIds(item.ParentIds),
			Phone:      item.Phone,
			Remark:     item.Remark,
			UpdateBy:   item.UpdateBy,
			UpdateTime: common.TimeToString(item.UpdateTime),
		})
	}

	//2.查询所有岗位
	var postListData []*sysclient.PostListData
	postList, _ := query.SysPost.WithContext(l.ctx).Find()
	for _, item := range postList {
		postListData = append(postListData, &sysclient.PostListData{
			CreateBy:   item.CreateBy,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
			Id:         item.ID,
			PostCode:   item.PostCode,
			PostName:   item.PostName,
			PostSort:   item.PostSort,
			PostStatus: item.PostStatus,
			Remark:     item.Remark,
			UpdateBy:   item.UpdateBy,
			UpdateTime: common.TimeToString(item.UpdateTime),
		})
	}

	data := &sysclient.QueryDeptAndPostListResp{
		DeptListData: deptListData,
		PostListData: postListData,
	}

	logc.Infof(l.ctx, "查询所有部门和岗位,参数：%+v,响应：%+v", in, data)
	return data, nil
}
