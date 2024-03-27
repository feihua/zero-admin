package feighttemplateservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/model/pmsmodel"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeightTemplateAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeightTemplateAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeightTemplateAddLogic {
	return &FeightTemplateAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeightTemplateAddLogic) FeightTemplateAdd(in *pmsclient.FeightTemplateAddReq) (*pmsclient.FeightTemplateAddResp, error) {
	_, err := l.svcCtx.PmsFeightTemplateModel.Insert(l.ctx, &pmsmodel.PmsFeightTemplate{
		Name:           in.Name,
		ChargeType:     in.ChargeType,
		FirstWeight:    float64(in.FirstWeight),
		FirstFee:       float64(in.FirstFee),
		ContinueWeight: float64(in.ContinueWeight),
		ContinmeFee:    float64(in.ContinmeFee),
		Dest:           in.Dest,
	})
	if err != nil {
		return nil, err
	}

	return &pmsclient.FeightTemplateAddResp{}, nil
}
