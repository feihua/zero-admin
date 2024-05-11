package preferredareaservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PreferredAreaUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPreferredAreaUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PreferredAreaUpdateLogic {
	return &PreferredAreaUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PreferredAreaUpdateLogic) PreferredAreaUpdate(in *cmsclient.PreferredAreaUpdateReq) (*cmsclient.PreferredAreaUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.PreferredAreaUpdateResp{}, nil
}
