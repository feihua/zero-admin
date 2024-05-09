package feighttemplateservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// FeightTemplateDeleteLogic 运费模板
/*
Author: LiuFeiHua
Date: 2024/5/8 9:51
*/
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

// FeightTemplateDelete 删除运费模板
func (l *FeightTemplateDeleteLogic) FeightTemplateDelete(in *pmsclient.FeightTemplateDeleteReq) (*pmsclient.FeightTemplateDeleteResp, error) {
	q := query.PmsFeightTemplate
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.FeightTemplateDeleteResp{}, nil
}
