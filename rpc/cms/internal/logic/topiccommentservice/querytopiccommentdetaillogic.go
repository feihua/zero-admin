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

// QueryTopicCommentDetailLogic 查询专题评论详情
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type QueryTopicCommentDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryTopicCommentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTopicCommentDetailLogic {
	return &QueryTopicCommentDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryTopicCommentDetail 查询专题评论详情
func (l *QueryTopicCommentDetailLogic) QueryTopicCommentDetail(in *cmsclient.QueryTopicCommentDetailReq) (*cmsclient.QueryTopicCommentDetailResp, error) {
	item, err := query.CmsTopicComment.WithContext(l.ctx).Where(query.CmsTopicComment.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询专题评论详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询专题评论详情失败")
	}

	data := &cmsclient.QueryTopicCommentDetailResp{
		Id:             item.ID,                              // 主键ID
		MemberNickName: item.MemberNickName,                  // 评论人员昵称
		TopicId:        item.TopicID,                         // 专题ID
		MemberIcon:     item.MemberIcon,                      // 评论人员头像
		Content:        item.Content,                         // 评论内容
		CreateTime:     time_util.TimeToStr(item.CreateTime), // 评论时间
		ShowStatus:     item.ShowStatus,                      // 是否显示，0->不显示；1->显示
	}

	return data, nil
}
