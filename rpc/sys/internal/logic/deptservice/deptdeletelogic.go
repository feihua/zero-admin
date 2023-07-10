package deptservicelogic

import (
	"context"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptDeleteLogic {
	return &DeptDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptDeleteLogic) DeptDelete(in *sysclient.DeptDeleteReq) (*sysclient.DeptDeleteResp, error) {
	err := l.svcCtx.DeptModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &sysclient.DeptDeleteResp{}, nil
}
