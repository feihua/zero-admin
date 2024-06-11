package subjectcommentservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySubjectCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySubjectCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubjectCommentListLogic {
	return &QuerySubjectCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询专题评论表列表
func (l *QuerySubjectCommentListLogic) QuerySubjectCommentList(in *cmsclient.QuerySubjectCommentListReq) (*cmsclient.QuerySubjectCommentListResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.QuerySubjectCommentListResp{}, nil
}
