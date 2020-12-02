package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CompayAddressUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompayAddressUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) CompayAddressUpdateLogic {
	return CompayAddressUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompayAddressUpdateLogic) CompayAddressUpdate(req types.UpdateCompayAddressReq) (*types.UpdateCompayAddressResp, error) {
	// todo: add your logic here and delete this line

	return &types.UpdateCompayAddressResp{}, nil
}
