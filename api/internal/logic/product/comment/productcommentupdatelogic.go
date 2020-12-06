package logic

import (
	"context"
	"go-zero-admin/rpc/pms/pmsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductCommentUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCommentUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCommentUpdateLogic {
	return ProductCommentUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCommentUpdateLogic) ProductCommentUpdate(req types.UpdateProductCommentReq) (*types.UpdateProductCommentResp, error) {
	_, err := l.svcCtx.Pms.CommentUpdate(l.ctx, &pmsclient.CommentUpdateReq{})

	if err != nil {
		return nil, err
	}

	return &types.UpdateProductCommentResp{}, nil
}
