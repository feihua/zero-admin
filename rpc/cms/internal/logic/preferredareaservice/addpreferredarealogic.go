package preferredareaservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPreferredAreaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPreferredAreaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPreferredAreaLogic {
	return &AddPreferredAreaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加优选专区
func (l *AddPreferredAreaLogic) AddPreferredArea(in *cmsclient.AddPreferredAreaReq) (*cmsclient.AddPreferredAreaResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.AddPreferredAreaResp{}, nil
}
