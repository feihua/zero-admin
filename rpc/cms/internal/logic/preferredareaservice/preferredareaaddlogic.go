package preferredareaservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PreferredAreaAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPreferredAreaAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PreferredAreaAddLogic {
	return &PreferredAreaAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 商品优选
func (l *PreferredAreaAddLogic) PreferredAreaAdd(in *cmsclient.PreferredAreaAddReq) (*cmsclient.PreferredAreaAddResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.PreferredAreaAddResp{}, nil
}
