package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeptAddLogic {
	return DeptAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptAddLogic) DeptAdd(req types.AddDeptReq) (*types.AddDeptResp, error) {
	_, err := l.svcCtx.Sys.DeptAdd(l.ctx, &sysclient.DeptAddReq{
		Name:     req.Name,
		ParentId: req.ParentId,
		OrderNum: req.OrderNum,
		//todo 从token里面拿
		CreateBy: "admin",
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加机构信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加机构失败")
	}

	return &types.AddDeptResp{
		Code:    "000000",
		Message: "添加机构成功",
	}, nil
}
