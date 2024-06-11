package preferredareaservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePreferredAreaStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePreferredAreaStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePreferredAreaStatusLogic {
	return &UpdatePreferredAreaStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新优选专区状态
func (l *UpdatePreferredAreaStatusLogic) UpdatePreferredAreaStatus(in *cmsclient.UpdatePreferredAreaStatusReq) (*cmsclient.UpdatePreferredAreaStatusResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.UpdatePreferredAreaStatusResp{}, nil
}
