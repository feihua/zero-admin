package logic

import (
	"context"
	"go-zero-admin/rpc/oms/omsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CompayAddressAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompayAddressAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) CompayAddressAddLogic {
	return CompayAddressAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompayAddressAddLogic) CompayAddressAdd(req types.AddCompayAddressReq) (*types.AddCompayAddressResp, error) {
	_, err := l.svcCtx.Oms.CompanyAddressAdd(l.ctx, &omsclient.CompanyAddressAddReq{})

	if err != nil {
		return nil, err
	}

	return &types.AddCompayAddressResp{}, nil
}
