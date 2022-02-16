package logic

import (
	"context"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

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

func (l *DeptDeleteLogic) DeptDelete(in *sys.DeptDeleteReq) (*sys.DeptDeleteResp, error) {
	err := l.svcCtx.DeptModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sys.DeptDeleteResp{}, nil
}
