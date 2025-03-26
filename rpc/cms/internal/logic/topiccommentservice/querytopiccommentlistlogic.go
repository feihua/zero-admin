package topiccommentservicelogic

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

// QueryTopicCommentListLogic 查询专题评论列表
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type QueryTopicCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryTopicCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTopicCommentListLogic {
	return &QueryTopicCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryTopicCommentList 查询专题评论列表
func (l *QueryTopicCommentListLogic) QueryTopicCommentList(in *cmsclient.QueryTopicCommentListReq) (*cmsclient.QueryTopicCommentListResp, error) {
	topicComment := query.CmsTopicComment
	q := topicComment.WithContext(l.ctx)
	if len(in.MemberNickName) > 0 {
		q = q.Where(topicComment.MemberNickName.Like("%" + in.MemberNickName + "%"))
	}
	if in.TopicId != 2 {
		q = q.Where(topicComment.TopicID.Eq(in.TopicId))
	}
	if len(in.MemberIcon) > 0 {
		q = q.Where(topicComment.MemberIcon.Like("%" + in.MemberIcon + "%"))
	}
	if len(in.Content) > 0 {
		q = q.Where(topicComment.Content.Like("%" + in.Content + "%"))
	}
	if in.ShowStatus != 2 {
		q = q.Where(topicComment.ShowStatus.Eq(in.ShowStatus))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询专题评论列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询专题评论列表失败")
	}

	var list []*cmsclient.TopicCommentListData

	for _, item := range result {
		list = append(list, &cmsclient.TopicCommentListData{
			Id:             item.ID,                              // 主键ID
			MemberNickName: item.MemberNickName,                  // 评论人员昵称
			TopicId:        item.TopicID,                         // 专题ID
			MemberIcon:     item.MemberIcon,                      // 评论人员头像
			Content:        item.Content,                         // 评论内容
			CreateTime:     time_util.TimeToStr(item.CreateTime), // 评论时间
			ShowStatus:     item.ShowStatus,                      // 是否显示，0->不显示；1->显示

		})
	}

	return &cmsclient.QueryTopicCommentListResp{
		Total: count,
		List:  list,
	}, nil
}
