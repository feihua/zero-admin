package logic

import (
	"context"
	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"
	"go-zero-admin/rpc/oms/omsclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type CompayAddressDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompayAddressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) CompayAddressDeleteLogic {
	return CompayAddressDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompayAddressDeleteLogic) CompayAddressDelete(req types.DeleteCompayAddressReq) (*types.DeleteCompayAddressResp, error) {
	_, _ = l.svcCtx.Oms.CompanyAddressDelete(l.ctx, &omsclient.CompanyAddressDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteCompayAddressResp{}, nil
}
