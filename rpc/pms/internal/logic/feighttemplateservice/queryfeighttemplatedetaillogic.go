package feighttemplateservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryFeightTemplateDetailLogic 查询运费模版详情
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:05
*/
type QueryFeightTemplateDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryFeightTemplateDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFeightTemplateDetailLogic {
	return &QueryFeightTemplateDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryFeightTemplateDetail 查询运费模版详情
func (l *QueryFeightTemplateDetailLogic) QueryFeightTemplateDetail(in *pmsclient.QueryFeightTemplateDetailReq) (*pmsclient.QueryFeightTemplateDetailResp, error) {
	item, err := query.PmsFeightTemplate.WithContext(l.ctx).Where(query.PmsFeightTemplate.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询运费模版详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询运费模版详情失败")
	}

	data := &pmsclient.QueryFeightTemplateDetailResp{
		Id:             item.ID,             //
		Name:           item.Name,           // 运费模版名称
		ChargeType:     item.ChargeType,     // 计费类型:0->按重量；1->按件数
		FirstWeight:    item.FirstWeight,    // 首重kg
		FirstFee:       item.FirstFee,       // 首费（元）
		ContinueWeight: item.ContinueWeight, // 续重kg
		ContinueFee:    item.ContinueFee,    // 续费（元）
		Dest:           item.Dest,           // 目的地（省、市）
	}

	return data, nil
}
