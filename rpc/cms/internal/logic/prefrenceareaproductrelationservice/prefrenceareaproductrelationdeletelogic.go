package prefrenceareaproductrelationservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrefrenceAreaProductRelationDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPrefrenceAreaProductRelationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrefrenceAreaProductRelationDeleteLogic {
	return &PrefrenceAreaProductRelationDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PrefrenceAreaProductRelationDeleteLogic) PrefrenceAreaProductRelationDelete(in *cmsclient.PrefrenceAreaProductRelationDeleteReq) (*cmsclient.PrefrenceAreaProductRelationDeleteResp, error) {
	err := l.svcCtx.CmsPrefrenceAreaProductRelationModel.DeleteByProductId(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}

	return &cmsclient.PrefrenceAreaProductRelationDeleteResp{}, nil
}
