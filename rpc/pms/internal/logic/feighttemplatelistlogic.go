package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

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

func (l *FeightTemplateListLogic) FeightTemplateList(in *pms.FeightTemplateListReq) (*pms.FeightTemplateListResp, error) {
	all, err := l.svcCtx.PmsFeightTemplateModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.PmsFeightTemplateModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询运费模板列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*pms.FeightTemplateListData
	for _, item := range *all {

		list = append(list, &pms.FeightTemplateListData{
			Id:             item.Id,
			Name:           item.Name,
			ChargeType:     item.ChargeType,
			FirstWeight:    int64(item.FirstWeight),
			FirstFee:       int64(item.FirstFee),
			ContinueWeight: int64(item.ContinueWeight),
			ContinmeFee:    int64(item.ContinmeFee),
			Dest:           item.Dest,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询运费模板列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &pms.FeightTemplateListResp{
		Total: count,
		List:  list,
	}, nil
}
