package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
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
	all, _ := l.svcCtx.PmsFeightTemplateModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.PmsFeightTemplateModel.Count()

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

	fmt.Println(list)
	return &pms.FeightTemplateListResp{
		Total: count,
		List:  list,
	}, nil
}
