package feighttemplateservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateFeightTemplateLogic 更新运费模版
/*
Author: LiuFeiHua
Date: 2024/6/12 16:41
*/
type UpdateFeightTemplateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFeightTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFeightTemplateLogic {
	return &UpdateFeightTemplateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateFeightTemplate 更新运费模版
func (l *UpdateFeightTemplateLogic) UpdateFeightTemplate(in *pmsclient.UpdateFeightTemplateReq) (*pmsclient.UpdateFeightTemplateResp, error) {
	q := query.PmsFeightTemplate
	_, err := q.WithContext(l.ctx).Updates(&model.PmsFeightTemplate{
		ID:             in.Id,
		Name:           in.Name,
		ChargeType:     in.ChargeType,
		FirstWeight:    in.FirstWeight,
		FirstFee:       in.FirstFee,
		ContinueWeight: in.ContinueWeight,
		ContinueFee:    in.ContinueFee,
		Dest:           in.Dest,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.UpdateFeightTemplateResp{}, nil
}
