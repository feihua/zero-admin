package logic

import (
	"context"
	"encoding/json"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/pms/pmsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
	_, err := l.svcCtx.Pms.FeightTemplateUpdate(l.ctx, &pmsclient.FeightTemplateUpdateReq{
		Id:             req.Id,
		Name:           req.Name,
		ChargeType:     req.ChargeType,
		FirstWeight:    int64(req.FirstWeight),
		FirstFee:       int64(req.FirstFee),
		ContinueWeight: int64(req.ContinueWeight),
		ContinmeFee:    int64(req.FirstFee),
		Dest:           req.Dest,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.Errorf("更新运费模版参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新运费模版失败")
	}

	return &types.UpdateFeightTemplateResp{
		Code:    "000000",
		Message: "更新运费模版成功",
	}, nil
}
