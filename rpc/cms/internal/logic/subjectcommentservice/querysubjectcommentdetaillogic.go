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

// QuerySubjectCommentDetailLogic 查询专题评论详情
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
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

// QuerySubjectCommentDetail 查询专题评论详情
func (l *QuerySubjectCommentDetailLogic) QuerySubjectCommentDetail(in *cmsclient.QuerySubjectCommentDetailReq) (*cmsclient.QuerySubjectCommentDetailResp, error) {
	item, err := query.CmsSubjectComment.WithContext(l.ctx).Where(query.CmsSubjectComment.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询专题评论详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询专题评论详情失败")
	}

	data := &cmsclient.QuerySubjectCommentDetailResp{
		Id:             item.ID,                              // 编号
		SubjectId:      item.SubjectID,                       // 关联专题id
		MemberNickName: item.MemberNickName,                  // 关联会员昵称
		MemberIcon:     item.MemberIcon,                      // 会员头像
		Content:        item.Content,                         // 评论内容
		CreateTime:     time_util.TimeToStr(item.CreateTime), // 创建时间
		ShowStatus:     item.ShowStatus,                      // 是否显示，0->不显示；1->显示
	}

	return data, nil
}
