package helpservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteHelpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteHelpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHelpLogic {
	return &DeleteHelpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除帮助表
func (l *DeleteHelpLogic) DeleteHelp(in *cmsclient.DeleteHelpReq) (*cmsclient.DeleteHelpResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.DeleteHelpResp{}, nil
}
