package subjectcommentservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSubjectCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSubjectCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSubjectCommentLogic {
	return &AddSubjectCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加专题评论表
func (l *AddSubjectCommentLogic) AddSubjectComment(in *cmsclient.AddSubjectCommentReq) (*cmsclient.AddSubjectCommentResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.AddSubjectCommentResp{}, nil
}
