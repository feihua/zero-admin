package prefrenceareaservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrefrenceAreaUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPrefrenceAreaUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrefrenceAreaUpdateLogic {
	return &PrefrenceAreaUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PrefrenceAreaUpdateLogic) PrefrenceAreaUpdate(in *cmsclient.PrefrenceAreaUpdateReq) (*cmsclient.PrefrenceAreaUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.PrefrenceAreaUpdateResp{}, nil
}
