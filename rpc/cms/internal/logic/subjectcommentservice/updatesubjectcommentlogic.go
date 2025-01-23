package subjectcommentservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/model"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateSubjectCommentLogic 更新专题评论
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
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

// UpdateSubjectComment 更新专题评论
func (l *UpdateSubjectCommentLogic) UpdateSubjectComment(in *cmsclient.UpdateSubjectCommentReq) (*cmsclient.UpdateSubjectCommentResp, error) {
	q := query.CmsSubjectComment.WithContext(l.ctx)

	// 1.根据专题评论id查询专题评论是否已存在
	_, err := q.Where(query.CmsSubjectComment.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据专题评论id：%d,查询专题评论失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询专题评论失败"))
	}

	item := &model.CmsSubjectComment{
		ID:             in.Id,             // 编号
		SubjectID:      in.SubjectId,      // 关联专题id
		MemberNickName: in.MemberNickName, // 关联会员昵称
		MemberIcon:     in.MemberIcon,     // 会员头像
		Content:        in.Content,        // 评论内容
		ShowStatus:     in.ShowStatus,     // 是否显示，0->不显示；1->显示
	}

	// 2.专题评论存在时,则直接更新专题评论
	_, err = q.Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新专题评论失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新专题评论失败")
	}

	logc.Infof(l.ctx, "更新专题评论成功,参数：%+v", in)
	return &cmsclient.UpdateSubjectCommentResp{}, nil
}
