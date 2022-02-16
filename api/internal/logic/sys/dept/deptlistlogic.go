package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"
	"zero-admin/rpc/sys/sysclient"
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
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询机构列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询机构失败")
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
		Message: "查询机构成功",
		Data:    list,
		Success: true,
		Total:   resp.Total,
	}, nil
}
