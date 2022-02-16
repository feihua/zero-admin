package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *DeptDeleteLogic) DeptDelete(req types.DeleteDeptReq) (*types.DeleteDeptResp, error) {
	_, err := l.svcCtx.Sys.DeptDelete(l.ctx, &sysclient.DeptDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据deptId: %d,删除部门异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除机构失败")
	}

	return &types.DeleteDeptResp{
		Code:    "000000",
		Message: "删除机构成功",
	}, nil
}
