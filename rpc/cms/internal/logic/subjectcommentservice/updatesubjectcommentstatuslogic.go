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

// UpdateSubjectCommentStatusLogic 更新专题评论
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
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

// UpdateSubjectCommentStatus 更新专题评论状态
func (l *UpdateSubjectCommentStatusLogic) UpdateSubjectCommentStatus(in *cmsclient.UpdateSubjectCommentStatusReq) (*cmsclient.UpdateSubjectCommentStatusResp, error) {
	q := query.CmsSubjectComment

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.ShowStatus, in.ShowStatus)

	if err != nil {
		logc.Errorf(l.ctx, "更新专题评论状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新专题评论状态失败")
	}

	return &cmsclient.UpdateSubjectCommentStatusResp{}, nil
}
