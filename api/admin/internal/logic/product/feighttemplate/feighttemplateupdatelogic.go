package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeightTemplateUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeightTemplateUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) FeightTemplateUpdateLogic {
	return FeightTemplateUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeightTemplateUpdateLogic) FeightTemplateUpdate(req types.UpdateFeightTemplateReq) (*types.UpdateFeightTemplateResp, error) {
	_, err := l.svcCtx.FeightTemplateService.FeightTemplateUpdate(l.ctx, &pmsclient.FeightTemplateUpdateReq{
		Id:             req.Id,
		Name:           req.Name,
		ChargeType:     req.ChargeType,
		FirstWeight:    req.FirstWeight,
		FirstFee:       req.FirstFee,
		ContinueWeight: req.ContinueWeight,
		ContinmeFee:    req.FirstFee,
		Dest:           req.Dest,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新运费模版信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新运费模版失败")
	}

	return &types.UpdateFeightTemplateResp{
		Code:    "000000",
		Message: "更新运费模版成功",
	}, nil
}
