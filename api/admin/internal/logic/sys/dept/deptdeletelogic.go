package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeptDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:16
*/
type DeptDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeptDeleteLogic {
	return DeptDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeptDelete 删除部门信息
func (l *DeptDeleteLogic) DeptDelete(req types.DeleteDeptReq) (*types.DeleteDeptResp, error) {
	_, err := l.svcCtx.DeptService.DeptDelete(l.ctx, &sysclient.DeptDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据deptId: %d,删除部门异常:%s", req.Ids, err.Error())
		return nil, errorx.NewDefaultError("删除机构失败")
	}

	return &types.DeleteDeptResp{
		Code:    "000000",
		Message: "删除机构成功",
	}, nil
}
