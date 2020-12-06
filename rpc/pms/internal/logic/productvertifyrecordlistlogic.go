package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductVertifyRecordListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductVertifyRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductVertifyRecordListLogic {
	return &ProductVertifyRecordListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductVertifyRecordListLogic) ProductVertifyRecordList(in *pms.ProductVertifyRecordListReq) (*pms.ProductVertifyRecordListResp, error) {
	// todo: add your logic here and delete this line

	return &pms.ProductVertifyRecordListResp{}, nil
}
