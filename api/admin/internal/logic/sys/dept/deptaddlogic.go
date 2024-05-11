package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeptAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:16
*/
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

// DeptAdd 添加部门信息
func (l *DeptAddLogic) DeptAdd(req types.AddDeptReq) (*types.AddDeptResp, error) {
	_, err := l.svcCtx.DeptService.DeptAdd(l.ctx, &sysclient.DeptAddReq{
		DeptName:  req.DeptName,
		ParentId:  req.ParentId,
		OrderNum:  req.OrderNum,
		CreateBy:  l.ctx.Value("userName").(string),
		ParentIds: req.ParentIds,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加机构信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加机构失败")
	}

	return &types.AddDeptResp{
		Code:    "000000",
		Message: "添加机构成功",
	}, nil
}
