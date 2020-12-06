package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductLadderDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLadderDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLadderDeleteLogic {
	return &ProductLadderDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductLadderDeleteLogic) ProductLadderDelete(in *pms.ProductLadderDeleteReq) (*pms.ProductLadderDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &pms.ProductLadderDeleteResp{}, nil
}
