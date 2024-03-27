package prefrenceareaproductrelationservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrefrenceAreaProductRelationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPrefrenceAreaProductRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrefrenceAreaProductRelationUpdateLogic {
	return &PrefrenceAreaProductRelationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PrefrenceAreaProductRelationUpdateLogic) PrefrenceAreaProductRelationUpdate(in *cmsclient.PrefrenceAreaProductRelationUpdateReq) (*cmsclient.PrefrenceAreaProductRelationUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.PrefrenceAreaProductRelationUpdateResp{}, nil
}
