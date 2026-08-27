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

// QueryTopicCommentListLogic 查询话题评论列表
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type QueryTopicCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryTopicCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTopicCommentListLogic {
	return &QueryTopicCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryTopicCommentList 查询话题评论列表
func (l *QueryTopicCommentListLogic) QueryTopicCommentList(req *types.QueryTopicCommentListReq) (resp *types.QueryTopicCommentListResp, err error) {
	result, err := l.svcCtx.TopicCommentService.QueryTopicCommentList(l.ctx, &cmsclient.QueryTopicCommentListReq{
		PageNum:        int32(req.Current),
		PageSize:       int32(req.PageSize),
		MemberNickName: req.MemberNickName, // 评论人员昵称
		TopicId:        req.TopicId,        // 专题id
		ShowStatus:     req.ShowStatus,     // 是否显示，0->不显示；1->显示

	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字话题评论列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryTopicCommentListData

	for _, item := range result.List {
		list = append(list, &types.QueryTopicCommentListData{
			Id:             item.Id,             // 主键id
			MemberNickName: item.MemberNickName, // 评论人员昵称
			TopicId:        item.TopicId,        // 专题id
			MemberIcon:     item.MemberIcon,     // 评论人员头像
			Content:        item.Content,        // 评论内容
			CreateTime:     item.CreateTime,     // 评论时间
			ShowStatus:     item.ShowStatus,     // 是否显示，0->不显示；1->显示
		})
	}

	return &types.QueryTopicCommentListResp{
		Code:     "000000",
		Message:  "查询话题评论列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
