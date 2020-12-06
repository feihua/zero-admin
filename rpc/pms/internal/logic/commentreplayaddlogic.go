package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CommentReplayAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentReplayAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentReplayAddLogic {
	return &CommentReplayAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentReplayAddLogic) CommentReplayAdd(in *pms.CommentReplayAddReq) (*pms.CommentReplayAddResp, error) {
	// todo: add your logic here and delete this line

	return &pms.CommentReplayAddResp{}, nil
}
