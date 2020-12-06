package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductVertifyRecordAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductVertifyRecordAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductVertifyRecordAddLogic {
	return &ProductVertifyRecordAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductVertifyRecordAddLogic) ProductVertifyRecordAdd(in *pms.ProductVertifyRecordAddReq) (*pms.ProductVertifyRecordAddResp, error) {
	// todo: add your logic here and delete this line

	return &pms.ProductVertifyRecordAddResp{}, nil
}
