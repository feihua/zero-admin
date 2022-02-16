package logic

import (
	"context"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAttributeDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeDeleteLogic {
	return &ProductAttributeDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeDeleteLogic) ProductAttributeDelete(in *pms.ProductAttributeDeleteReq) (*pms.ProductAttributeDeleteResp, error) {
	err := l.svcCtx.PmsProductAttributeModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.ProductAttributeDeleteResp{}, nil
}
