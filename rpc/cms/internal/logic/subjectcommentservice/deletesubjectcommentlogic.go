package subjectcommentservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSubjectCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSubjectCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSubjectCommentLogic {
	return &DeleteSubjectCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除专题评论表
func (l *DeleteSubjectCommentLogic) DeleteSubjectComment(in *cmsclient.DeleteSubjectCommentReq) (*cmsclient.DeleteSubjectCommentResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.DeleteSubjectCommentResp{}, nil
}
