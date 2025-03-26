package feighttemplateservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateFeightTemplateLogic 更新运费模版
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:05
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
	q := query.PmsFeightTemplate.WithContext(l.ctx)

	// 1.根据运费模版id查询运费模版是否已存在
	_, err := q.Where(query.PmsFeightTemplate.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据运费模版id：%d,查询运费模版失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询运费模版失败"))
	}

	item := &model.PmsFeightTemplate{
		ID:             in.Id,             //
		Name:           in.Name,           // 运费模版名称
		ChargeType:     in.ChargeType,     // 计费类型:0->按重量；1->按件数
		FirstWeight:    in.FirstWeight,    // 首重kg
		FirstFee:       in.FirstFee,       // 首费（元）
		ContinueWeight: in.ContinueWeight, // 续重kg
		ContinueFee:    in.ContinueFee,    // 续费（元）
		Dest:           in.Dest,           // 目的地（省、市）
	}

	// 2.运费模版存在时,则直接更新运费模版
	_, err = q.Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新运费模版失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新运费模版失败")
	}

	return &pmsclient.UpdateFeightTemplateResp{}, nil
}
