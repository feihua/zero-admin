package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

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
	resp, err := l.svcCtx.FeightTemplateService.FeightTemplateList(l.ctx, &pmsclient.FeightTemplateListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询运费模版列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询运费模版失败")
	}

	var list []*types.ListFeightTemplateData

	for _, item := range resp.List {
		list = append(list, &types.ListFeightTemplateData{
			Id:             item.Id,
			Name:           item.Name,
			ChargeType:     item.ChargeType,
			FirstWeight:    item.FirstWeight,
			FirstFee:       item.FirstFee,
			ContinueWeight: item.ContinueWeight,
			ContinmeFee:    item.FirstFee,
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
