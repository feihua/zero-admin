package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCommentListLogic {
	return ProductCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCommentListLogic) ProductCommentList(req types.ListProductCommentReq) (*types.ListProductCommentResp, error) {
	// todo: add your logic here and delete this line

	return &types.ListProductCommentResp{}, nil
}
