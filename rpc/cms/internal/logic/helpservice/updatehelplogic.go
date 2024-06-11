package helpservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHelpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHelpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHelpLogic {
	return &UpdateHelpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新帮助表
func (l *UpdateHelpLogic) UpdateHelp(in *cmsclient.UpdateHelpReq) (*cmsclient.UpdateHelpResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.UpdateHelpResp{}, nil
}
