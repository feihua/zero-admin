package logic

import (
	"context"
	"go-zero-admin/rpc/pms/pmsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductCommentAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCommentAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCommentAddLogic {
	return ProductCommentAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCommentAddLogic) ProductCommentAdd(req types.AddProductCommentReq) (*types.AddProductCommentResp, error) {
	_, err := l.svcCtx.Pms.CommentAdd(l.ctx, &pmsclient.CommentAddReq{})

	if err != nil {
		return nil, err
	}

	return &types.AddProductCommentResp{}, nil
}
