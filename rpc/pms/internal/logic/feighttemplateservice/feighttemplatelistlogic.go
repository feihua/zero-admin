package feighttemplateservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeightTemplateListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeightTemplateListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeightTemplateListLogic {
	return &FeightTemplateListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeightTemplateListLogic) FeightTemplateList(in *pmsclient.FeightTemplateListReq) (*pmsclient.FeightTemplateListResp, error) {
	q := query.PmsFeightTemplate.WithContext(l.ctx)

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询运费模板列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.FeightTemplateListData
	for _, item := range result {

		list = append(list, &pmsclient.FeightTemplateListData{
			Id:             item.ID,
			Name:           item.Name,
			ChargeType:     item.ChargeType,
			FirstWeight:    item.FirstWeight,
			FirstFee:       item.FirstFee,
			ContinueWeight: item.ContinueWeight,
			ContinmeFee:    item.ContinueFee,
			Dest:           item.Dest,
		})
	}

	logc.Infof(l.ctx, "查询运费模板列表信息,参数：%+v,响应：%+v", in, list)
	return &pmsclient.FeightTemplateListResp{
		Total: count,
		List:  list,
	}, nil
}
