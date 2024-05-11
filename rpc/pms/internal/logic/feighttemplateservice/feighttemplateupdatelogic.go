package feighttemplateservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// FeightTemplateUpdateLogic 运费模板
/*
Author: LiuFeiHua
Date: 2024/5/8 10:46
*/
type FeightTemplateUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeightTemplateUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeightTemplateUpdateLogic {
	return &FeightTemplateUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FeightTemplateUpdate 更新运费模板
func (l *FeightTemplateUpdateLogic) FeightTemplateUpdate(in *pmsclient.FeightTemplateUpdateReq) (*pmsclient.FeightTemplateUpdateResp, error) {
	q := query.PmsFeightTemplate
	_, err := q.WithContext(l.ctx).Updates(&model.PmsFeightTemplate{
		ID:             in.Id,
		Name:           in.Name,
		ChargeType:     in.ChargeType,
		FirstWeight:    in.FirstWeight,
		FirstFee:       in.FirstFee,
		ContinueWeight: in.ContinueWeight,
		ContinueFee:    in.ContinmeFee,
		Dest:           in.Dest,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.FeightTemplateUpdateResp{}, nil
}
