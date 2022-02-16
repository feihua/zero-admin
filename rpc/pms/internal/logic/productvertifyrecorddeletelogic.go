package logic

import (
	"context"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductVertifyRecordDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductVertifyRecordDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductVertifyRecordDeleteLogic {
	return &ProductVertifyRecordDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductVertifyRecordDeleteLogic) ProductVertifyRecordDelete(in *pms.ProductVertifyRecordDeleteReq) (*pms.ProductVertifyRecordDeleteResp, error) {
	err := l.svcCtx.PmsProductVertifyRecordModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.ProductVertifyRecordDeleteResp{}, nil
}
