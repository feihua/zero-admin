package helpcategoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHelpCategoryStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHelpCategoryStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHelpCategoryStatusLogic {
	return &UpdateHelpCategoryStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新帮助分类表状态
func (l *UpdateHelpCategoryStatusLogic) UpdateHelpCategoryStatus(in *cmsclient.UpdateHelpCategoryStatusReq) (*cmsclient.UpdateHelpCategoryStatusResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.UpdateHelpCategoryStatusResp{}, nil
}
