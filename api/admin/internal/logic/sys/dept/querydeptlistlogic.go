package dept

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
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

// QueryDeptList 部门列表
func (l *QueryDeptListLogic) QueryDeptList(req *types.ListDeptReq) (*types.ListDeptResp, error) {
	resp, err := l.svcCtx.DeptService.DeptList(l.ctx, &sysclient.DeptListReq{
		DeptName: req.Name,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询机构列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询机构失败")
	}

	var list []*types.ListDeptData

	for _, dept := range resp.List {
		list = append(list, &types.ListDeptData{
			Id:         dept.Id,
			DeptName:   dept.DeptName,
			ParentId:   dept.ParentId,
			OrderNum:   dept.OrderNum,
			CreateBy:   dept.CreateBy,
			CreateTime: dept.CreateTime,
			UpdateBy:   dept.UpdateBy,
			UpdateTime: dept.UpdateTime,
			DelFlag:    dept.DelFlag,
			ParentIds:  dept.ParentIds,
		})
	}

	return &types.ListDeptResp{
		Code:    "000000",
		Message: "查询机构成功",
		Data:    list,
		Success: true,
		Total:   resp.Total,
	}, nil
}
