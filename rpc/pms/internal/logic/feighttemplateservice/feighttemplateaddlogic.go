package feighttemplateservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// FeightTemplateAddLogic 运费模板
/*
Author: LiuFeiHua
Date: 2024/5/8 10:44
*/
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

// FeightTemplateAdd 添加运费模板
func (l *FeightTemplateAddLogic) FeightTemplateAdd(in *pmsclient.FeightTemplateAddReq) (*pmsclient.FeightTemplateAddResp, error) {
	err := query.PmsFeightTemplate.WithContext(l.ctx).Create(&model.PmsFeightTemplate{
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
