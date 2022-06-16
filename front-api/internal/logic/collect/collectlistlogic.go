package collect

import (
	"context"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) CollectListLogic {
	return CollectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollectListLogic) CollectList(req types.CollectListReq) (resp *types.CollectListResp, err error) {
	l.svcCtx.Pms.CollectList(l.ctx, &pmsclient.CollectListReq{
		Id: 0,
	})

	return &types.CollectListResp{
		Errno:  0,
		Errmsg: "",
	}, nil
}
