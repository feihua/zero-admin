package preferredareaservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePreferredAreaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePreferredAreaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePreferredAreaLogic {
	return &UpdatePreferredAreaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新优选专区
func (l *UpdatePreferredAreaLogic) UpdatePreferredArea(in *cmsclient.UpdatePreferredAreaReq) (*cmsclient.UpdatePreferredAreaResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.UpdatePreferredAreaResp{}, nil
}
