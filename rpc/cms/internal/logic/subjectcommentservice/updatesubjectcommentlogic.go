package subjectcommentservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSubjectCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSubjectCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubjectCommentLogic {
	return &UpdateSubjectCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新专题评论表
func (l *UpdateSubjectCommentLogic) UpdateSubjectComment(in *cmsclient.UpdateSubjectCommentReq) (*cmsclient.UpdateSubjectCommentResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.UpdateSubjectCommentResp{}, nil
}
