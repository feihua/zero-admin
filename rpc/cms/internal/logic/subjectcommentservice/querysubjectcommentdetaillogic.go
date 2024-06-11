package subjectcommentservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySubjectCommentDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySubjectCommentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubjectCommentDetailLogic {
	return &QuerySubjectCommentDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询专题评论表详情
func (l *QuerySubjectCommentDetailLogic) QuerySubjectCommentDetail(in *cmsclient.QuerySubjectCommentDetailReq) (*cmsclient.QuerySubjectCommentDetailResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.QuerySubjectCommentDetailResp{}, nil
}
