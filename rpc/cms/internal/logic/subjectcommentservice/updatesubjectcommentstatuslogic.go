package subjectcommentservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSubjectCommentStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSubjectCommentStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubjectCommentStatusLogic {
	return &UpdateSubjectCommentStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新专题评论表状态
func (l *UpdateSubjectCommentStatusLogic) UpdateSubjectCommentStatus(in *cmsclient.UpdateSubjectCommentStatusReq) (*cmsclient.UpdateSubjectCommentStatusResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.UpdateSubjectCommentStatusResp{}, nil
}
