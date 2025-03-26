package subjectcommentservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QuerySubjectCommentListLogic 查询专题评论列表
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
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

// QuerySubjectCommentList 查询专题评论列表
func (l *QuerySubjectCommentListLogic) QuerySubjectCommentList(in *cmsclient.QuerySubjectCommentListReq) (*cmsclient.QuerySubjectCommentListResp, error) {
	subjectComment := query.CmsSubjectComment
	q := subjectComment.WithContext(l.ctx)
	if in.SubjectId != 2 {
		q = q.Where(subjectComment.SubjectID.Eq(in.SubjectId))
	}
	if len(in.MemberNickName) > 0 {
		q = q.Where(subjectComment.MemberNickName.Like("%" + in.MemberNickName + "%"))
	}
	if len(in.MemberIcon) > 0 {
		q = q.Where(subjectComment.MemberIcon.Like("%" + in.MemberIcon + "%"))
	}
	if len(in.Content) > 0 {
		q = q.Where(subjectComment.Content.Like("%" + in.Content + "%"))
	}
	if in.ShowStatus != 2 {
		q = q.Where(subjectComment.ShowStatus.Eq(in.ShowStatus))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询专题评论列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询专题评论列表失败")
	}

	var list []*cmsclient.SubjectCommentListData

	for _, item := range result {
		list = append(list, &cmsclient.SubjectCommentListData{
			Id:             item.ID,                              // 编号
			SubjectId:      item.SubjectID,                       // 关联专题id
			MemberNickName: item.MemberNickName,                  // 关联会员昵称
			MemberIcon:     item.MemberIcon,                      // 会员头像
			Content:        item.Content,                         // 评论内容
			CreateTime:     time_util.TimeToStr(item.CreateTime), // 创建时间
			ShowStatus:     item.ShowStatus,                      // 是否显示，0->不显示；1->显示

		})
	}

	return &cmsclient.QuerySubjectCommentListResp{
		Total: count,
		List:  list,
	}, nil
}
