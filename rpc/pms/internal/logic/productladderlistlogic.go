package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductLadderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLadderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLadderListLogic {
	return &ProductLadderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductLadderListLogic) ProductLadderList(in *pms.ProductLadderListReq) (*pms.ProductLadderListResp, error) {
	// todo: add your logic here and delete this line

	return &pms.ProductLadderListResp{}, nil
}
