package logic

import (
	"context"

	"go-zero-admin/rpc/sys/internal/svc"
	"go-zero-admin/rpc/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type DictDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictDeleteLogic {
	return &DictDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictDeleteLogic) DictDelete(in *sys.DictDeleteReq) (*sys.DictDeleteResp, error) {
	err := l.svcCtx.DictModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sys.DictDeleteResp{}, nil
}
