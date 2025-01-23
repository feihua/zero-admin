package subjectcommentservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteSubjectCommentLogic 删除专题评论
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
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

// DeleteSubjectComment 删除专题评论
func (l *DeleteSubjectCommentLogic) DeleteSubjectComment(in *cmsclient.DeleteSubjectCommentReq) (*cmsclient.DeleteSubjectCommentResp, error) {
	q := query.CmsSubjectComment

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除专题评论失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除专题评论失败")
	}

	logc.Infof(l.ctx, "删除专题评论成功,参数：%+v", in)
	return &cmsclient.DeleteSubjectCommentResp{}, nil
}
