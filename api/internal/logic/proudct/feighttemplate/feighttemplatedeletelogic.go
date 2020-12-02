package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FeightTemplateDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeightTemplateDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) FeightTemplateDeleteLogic {
	return FeightTemplateDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeightTemplateDeleteLogic) FeightTemplateDelete(req types.DeleteFeightTemplateReq) (*types.DeleteFeightTemplateResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeleteFeightTemplateResp{}, nil
}
