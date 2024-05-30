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
			CreateBy:   dept.CreateBy,
			CreateTime: dept.CreateTime,
			DeptName:   dept.DeptName,
			DeptSort:   dept.DeptSort,
			DeptStatus: dept.DeptStatus,
			Email:      dept.Email,
			Id:         dept.Id,
			Leader:     dept.Leader,
			ParentId:   dept.ParentId,
			ParentIds:  dept.ParentIds,
			Phone:      dept.Phone,
			Remark:     dept.Remark,
			UpdateBy:   dept.UpdateBy,
			UpdateTime: dept.UpdateTime,
		})
	}

	return &types.QueryDeptListResp{
		Code:    "000000",
		Message: "查询部门成功",
		Data:    list,
		Success: true,
	}, nil
}
