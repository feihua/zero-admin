package logic

import (
	"context"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCollectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectListLogic {
	return &CollectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CollectListLogic) CollectList(in *pmsclient.CollectListReq) (*pmsclient.CollectListResp, error) {
	// todo: add your logic here and delete this line

	return &pmsclient.CollectListResp{}, nil
}
