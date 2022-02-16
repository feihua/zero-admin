package logic

import (
	"context"
	"zero-admin/rpc/model/pmsmodel"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

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

func (l *FeightTemplateAddLogic) FeightTemplateAdd(in *pms.FeightTemplateAddReq) (*pms.FeightTemplateAddResp, error) {
	_, err := l.svcCtx.PmsFeightTemplateModel.Insert(pmsmodel.PmsFeightTemplate{
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

	return &pms.FeightTemplateAddResp{}, nil
}
