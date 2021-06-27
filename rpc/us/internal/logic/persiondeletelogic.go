package logic

import (
	"context"
	"go-zero-admin/rpc/us/internal/svc"
	"go-zero-admin/rpc/us/us"

	"github.com/tal-tech/go-zero/core/logx"
)

type PersionDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPersionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PersionDeleteLogic {
	return &PersionDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PersionDeleteLogic) PersionDelete(in *us.PersionDeleteReq) (*us.PersionDeleteResp, error) {
	// todo: add your logic here and delete this line

	err := l.svcCtx.UsRolesModel.Delete(in.Id)

	if err != nil {
		return &us.PersionDeleteResp{
			Result: false,
		}, err
	}

	return &us.PersionDeleteResp{
		Result: true,
	}, nil
}
