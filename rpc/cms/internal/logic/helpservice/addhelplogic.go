package helpservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddHelpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddHelpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHelpLogic {
	return &AddHelpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加帮助表
func (l *AddHelpLogic) AddHelp(in *cmsclient.AddHelpReq) (*cmsclient.AddHelpResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.AddHelpResp{}, nil
}
