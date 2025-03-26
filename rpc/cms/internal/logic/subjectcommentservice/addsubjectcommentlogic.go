package subjectcommentservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/model"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddSubjectCommentLogic 添加专题评论
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
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

// AddSubjectComment 添加专题评论
func (l *AddSubjectCommentLogic) AddSubjectComment(in *cmsclient.AddSubjectCommentReq) (*cmsclient.AddSubjectCommentResp, error) {
	q := query.CmsSubjectComment

	item := &model.CmsSubjectComment{
		SubjectID:      in.SubjectId,      // 关联专题id
		MemberNickName: in.MemberNickName, // 关联会员昵称
		MemberIcon:     in.MemberIcon,     // 会员头像
		Content:        in.Content,        // 评论内容
		ShowStatus:     in.ShowStatus,     // 是否显示，0->不显示；1->显示
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加专题评论失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加专题评论失败")
	}

	return &cmsclient.AddSubjectCommentResp{}, nil
}
