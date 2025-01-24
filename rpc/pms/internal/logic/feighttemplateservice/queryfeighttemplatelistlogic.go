package feighttemplateservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryFeightTemplateListLogic 查询运费模版列表
/*
Author: LiuFeiHua
Date: 2024/6/12 16:41
*/
type QueryFeightTemplateListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryFeightTemplateListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFeightTemplateListLogic {
	return &QueryFeightTemplateListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryFeightTemplateList 查询运费模版列表
func (l *QueryFeightTemplateListLogic) QueryFeightTemplateList(in *pmsclient.QueryFeightTemplateListReq) (*pmsclient.QueryFeightTemplateListResp, error) {
	q := query.PmsFeightTemplate.WithContext(l.ctx)

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询运费模板列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.FeightTemplateListData
	for _, item := range result {

		list = append(list, &pmsclient.FeightTemplateListData{
			Id:             item.ID,             //
			Name:           item.Name,           // 运费模版名称
			ChargeType:     item.ChargeType,     // 计费类型:0->按重量；1->按件数
			FirstWeight:    item.FirstWeight,    // 首重kg
			FirstFee:       item.FirstFee,       // 首费（元）
			ContinueWeight: item.ContinueWeight, // 续重kg
			ContinueFee:    item.ContinueFee,    // 续费（元）
			Dest:           item.Dest,           // 目的地（省、市）
		})
	}

	return &pmsclient.QueryFeightTemplateListResp{
		Total: count,
		List:  list,
	}, nil
}
