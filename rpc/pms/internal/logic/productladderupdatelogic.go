package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductLadderUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLadderUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLadderUpdateLogic {
	return &ProductLadderUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductLadderUpdateLogic) ProductLadderUpdate(in *pms.ProductLadderUpdateReq) (*pms.ProductLadderUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &pms.ProductLadderUpdateResp{}, nil
}
