package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return &types.AddFeightTemplateResp{}, nil
}
