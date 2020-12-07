package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CommentDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentDeleteLogic {
	return &CommentDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentDeleteLogic) CommentDelete(in *pms.CommentDeleteReq) (*pms.CommentDeleteResp, error) {
	err := l.svcCtx.PmsCommentModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.CommentDeleteResp{}, nil
}
