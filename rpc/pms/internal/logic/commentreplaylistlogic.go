package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CommentReplayListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentReplayListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentReplayListLogic {
	return &CommentReplayListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentReplayListLogic) CommentReplayList(in *pms.CommentReplayListReq) (*pms.CommentReplayListResp, error) {
	// todo: add your logic here and delete this line

	return &pms.CommentReplayListResp{}, nil
}
