package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CompayAddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompayAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) CompayAddressListLogic {
	return CompayAddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompayAddressListLogic) CompayAddressList(req types.ListCompayAddressReq) (*types.ListCompayAddressResp, error) {
	// todo: add your logic here and delete this line

	return &types.ListCompayAddressResp{}, nil
}
