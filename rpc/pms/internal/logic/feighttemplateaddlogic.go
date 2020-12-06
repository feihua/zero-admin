package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FeightTemplateAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeightTemplateAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeightTemplateAddLogic {
	return &FeightTemplateAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeightTemplateAddLogic) FeightTemplateAdd(in *pms.FeightTemplateAddReq) (*pms.FeightTemplateAddResp, error) {
	// todo: add your logic here and delete this line

	return &pms.FeightTemplateAddResp{}, nil
}
