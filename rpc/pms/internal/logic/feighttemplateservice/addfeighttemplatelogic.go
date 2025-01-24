package feighttemplateservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddFeightTemplateLogic 添加运费模版
/*
Author: LiuFeiHua
Date: 2024/6/12 16:38
*/
type AddFeightTemplateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFeightTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFeightTemplateLogic {
	return &AddFeightTemplateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddFeightTemplate 添加运费模版
func (l *AddFeightTemplateLogic) AddFeightTemplate(in *pmsclient.AddFeightTemplateReq) (*pmsclient.AddFeightTemplateResp, error) {
	err := query.PmsFeightTemplate.WithContext(l.ctx).Create(&model.PmsFeightTemplate{
		Name:           in.Name,           // 运费模版名称
		ChargeType:     in.ChargeType,     // 计费类型:0->按重量；1->按件数
		FirstWeight:    in.FirstWeight,    // 首重kg
		FirstFee:       in.FirstFee,       // 首费（元）
		ContinueWeight: in.ContinueWeight, // 续重kg
		ContinueFee:    in.ContinueFee,    // 续费（元）
		Dest:           in.Dest,           // 目的地（省、市）
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.AddFeightTemplateResp{}, nil
}
