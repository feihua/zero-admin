package feighttemplateservicelogic

import (
	"context"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeightTemplateDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeightTemplateDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeightTemplateDeleteLogic {
	return &FeightTemplateDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeightTemplateDeleteLogic) FeightTemplateDelete(in *pmsclient.FeightTemplateDeleteReq) (*pmsclient.FeightTemplateDeleteResp, error) {
	err := l.svcCtx.PmsFeightTemplateModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &pmsclient.FeightTemplateDeleteResp{}, nil
}
