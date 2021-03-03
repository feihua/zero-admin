package logic

import (
	"context"
	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"
	"go-zero-admin/rpc/sys/sysclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeptListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeptListLogic {
	return DeptListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptListLogic) DeptList(req types.ListDeptReq) (*types.ListDeptResp, error) {
	resp, err := l.svcCtx.Sys.DeptList(l.ctx, &sysclient.DeptListReq{
		Name:     req.Name,
		CreateBy: req.CreateBy,
	})

	if err != nil {
		return nil, err
	}

	var list []*types.ListDeptData

	for _, dept := range resp.List {
		list = append(list, &types.ListDeptData{
			Id:             dept.Id,
			Name:           dept.Name,
			ParentId:       dept.ParentId,
			OrderNum:       dept.OrderNum,
			CreateBy:       dept.CreateBy,
			CreateTime:     dept.CreateTime,
			LastUpdateBy:   dept.LastUpdateBy,
			LastUpdateTime: dept.LastUpdateTime,
			DelFlag:        dept.DelFlag,
		})
	}

	return &types.ListDeptResp{
		Code:    "000000",
		Message: "",
		Data:    list,
		Success: true,
		Total:   resp.Total,
	}, nil
}
