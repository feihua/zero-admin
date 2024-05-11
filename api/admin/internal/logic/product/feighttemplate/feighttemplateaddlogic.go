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

type FeightTemplateAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeightTemplateAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) FeightTemplateAddLogic {
	return FeightTemplateAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeightTemplateAddLogic) FeightTemplateAdd(req types.AddFeightTemplateReq) (*types.AddFeightTemplateResp, error) {
	_, err := l.svcCtx.FeightTemplateService.FeightTemplateAdd(l.ctx, &pmsclient.FeightTemplateAddReq{
		Name:           req.Name,
		ChargeType:     req.ChargeType,
		FirstWeight:    req.FirstWeight,
		FirstFee:       req.FirstFee,
		ContinueWeight: req.ContinueWeight,
		ContinmeFee:    req.FirstFee,
		Dest:           req.Dest,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加运费模板信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加运费模版失败")
	}

	return &types.AddFeightTemplateResp{
		Code:    "000000",
		Message: "添加运费模版成功",
	}, nil
}
