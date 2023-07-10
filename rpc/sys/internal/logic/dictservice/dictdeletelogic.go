package dictservicelogic

import (
	"context"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *DictDeleteLogic) DictDelete(in *sysclient.DictDeleteReq) (*sysclient.DictDeleteResp, error) {
	err := l.svcCtx.DictModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &sysclient.DictDeleteResp{}, nil
}
