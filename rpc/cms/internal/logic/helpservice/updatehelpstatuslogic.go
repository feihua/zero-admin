package helpservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHelpStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHelpStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHelpStatusLogic {
	return &UpdateHelpStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新帮助表状态
func (l *UpdateHelpStatusLogic) UpdateHelpStatus(in *cmsclient.UpdateHelpStatusReq) (*cmsclient.UpdateHelpStatusResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.UpdateHelpStatusResp{}, nil
}
