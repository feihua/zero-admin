package topic_comment

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryTopicCommentDetailLogic 查询话题评论详情
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type QueryTopicCommentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryTopicCommentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTopicCommentDetailLogic {
	return &QueryTopicCommentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryTopicCommentDetail 查询话题评论详情
func (l *QueryTopicCommentDetailLogic) QueryTopicCommentDetail(req *types.QueryTopicCommentDetailReq) (resp *types.QueryTopicCommentDetailResp, err error) {

	detail, err := l.svcCtx.TopicCommentService.QueryTopicCommentDetail(l.ctx, &cmsclient.QueryTopicCommentDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询话题评论详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryTopicCommentDetailData{
		Id:             detail.Id,             // 主键id
		MemberNickName: detail.MemberNickName, // 评论人员昵称
		TopicId:        detail.TopicId,        // 专题id
		MemberIcon:     detail.MemberIcon,     // 评论人员头像
		Content:        detail.Content,        // 评论内容
		CreateTime:     detail.CreateTime,     // 评论时间
		ShowStatus:     detail.ShowStatus,     // 是否显示，0->不显示；1->显示
	}
	return &types.QueryTopicCommentDetailResp{
		Code:    "000000",
		Message: "查询话题评论成功",
		Data:    data,
	}, nil
}
