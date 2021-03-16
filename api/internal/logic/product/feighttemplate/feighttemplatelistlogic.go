package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/pmsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FeightTemplateListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeightTemplateListLogic(ctx context.Context, svcCtx *svc.ServiceContext) FeightTemplateListLogic {
	return FeightTemplateListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeightTemplateListLogic) FeightTemplateList(req types.ListFeightTemplateReq) (*types.ListFeightTemplateResp, error) {
	resp, err := l.svcCtx.Pms.FeightTemplateList(l.ctx, &pmsclient.FeightTemplateListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtFeightTemplateData

	for _, item := range resp.List {
		list = append(list, &types.ListtFeightTemplateData{
			Id:             item.Id,
			Name:           item.Name,
			ChargeType:     item.ChargeType,
			FirstWeight:    float64(item.FirstWeight),
			FirstFee:       float64(item.FirstFee),
			ContinueWeight: float64(item.ContinueWeight),
			ContinmeFee:    float64(item.FirstFee),
			Dest:           item.Dest,
		})
	}

	return &types.ListFeightTemplateResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "",
	}, nil
}
