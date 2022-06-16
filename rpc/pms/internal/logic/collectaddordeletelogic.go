package logic

import (
	"context"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectAddOrDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCollectAddOrDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectAddOrDeleteLogic {
	return &CollectAddOrDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CollectAddOrDeleteLogic) CollectAddOrDelete(in *pmsclient.CollectAddOrDeleteReq) (*pmsclient.CollectAddOrDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &pmsclient.CollectAddOrDeleteResp{}, nil
}
