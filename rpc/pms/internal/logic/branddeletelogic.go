package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type BrandDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBrandDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandDeleteLogic {
	return &BrandDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BrandDeleteLogic) BrandDelete(in *pms.BrandDeleteReq) (*pms.BrandDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &pms.BrandDeleteResp{}, nil
}
