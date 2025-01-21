package dept

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

// QueryDeptListLogic 部门列表
/*
Author: LiuFeiHua
Date: 2023/12/18 17:16
*/
type QueryDeptListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryDeptListLogic {
	return QueryDeptListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryDeptList 查询部门列表
func (l *QueryDeptListLogic) QueryDeptList(req *types.QueryDeptListReq) (*types.QueryDeptListResp, error) {
	resp, err := l.svcCtx.DeptService.QueryDeptList(l.ctx, &sysclient.QueryDeptListReq{})

	if err != nil {
		logc.Errorf(l.ctx, "查询部门列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryDeptListData

	for _, dept := range resp.List {
		list = append(list, &types.QueryDeptListData{
			Id:         dept.Id,         // 编号
			DeptName:   dept.DeptName,   // 部门名称
			DeptStatus: dept.DeptStatus, // 部门状态
			DeptSort:   dept.DeptSort,   // 部门排序
			ParentId:   dept.ParentId,   // 上级机构ID，一级机构为0
			Leader:     dept.Leader,     // 负责人
			Phone:      dept.Phone,      // 电话号码
			Email:      dept.Email,      // 邮箱
			Remark:     dept.Remark,     // 备注信息
			IsDeleted:  dept.IsDeleted,  // 是否删除  0：否  1：是
			ParentIds:  dept.ParentIds,  // 上级机构IDs，一级机构为0
			CreateBy:   dept.CreateBy,   // 创建者
			CreateTime: dept.CreateTime, // 创建时间
			UpdateBy:   dept.UpdateBy,   // 更新者
			UpdateTime: dept.UpdateTime, // 更新时间
		})
	}

	return &types.QueryDeptListResp{
		Code:    "000000",
		Message: "查询部门成功",
		Data:    list,
		Success: true,
	}, nil
}
