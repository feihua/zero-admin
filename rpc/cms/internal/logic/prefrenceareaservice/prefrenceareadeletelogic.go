package prefrenceareaservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrefrenceAreaDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPrefrenceAreaDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrefrenceAreaDeleteLogic {
	return &PrefrenceAreaDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PrefrenceAreaDeleteLogic) PrefrenceAreaDelete(in *cmsclient.PrefrenceAreaDeleteReq) (*cmsclient.PrefrenceAreaDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.PrefrenceAreaDeleteResp{}, nil
}
