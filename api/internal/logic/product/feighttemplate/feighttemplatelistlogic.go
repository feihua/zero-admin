package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询运费模版列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询运费模版失败")
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
		Message:  "查询运费模版成功",
	}, nil
}
