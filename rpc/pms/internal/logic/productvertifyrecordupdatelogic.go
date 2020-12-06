package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductVertifyRecordUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductVertifyRecordUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductVertifyRecordUpdateLogic {
	return &ProductVertifyRecordUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductVertifyRecordUpdateLogic) ProductVertifyRecordUpdate(in *pms.ProductVertifyRecordUpdateReq) (*pms.ProductVertifyRecordUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &pms.ProductVertifyRecordUpdateResp{}, nil
}
