package preferredareaservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PreferredAreaDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPreferredAreaDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PreferredAreaDeleteLogic {
	return &PreferredAreaDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PreferredAreaDeleteLogic) PreferredAreaDelete(in *cmsclient.PreferredAreaDeleteReq) (*cmsclient.PreferredAreaDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.PreferredAreaDeleteResp{}, nil
}
