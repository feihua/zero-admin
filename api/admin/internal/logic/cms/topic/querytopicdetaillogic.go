package topic

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

// QueryTopicDetailLogic 查询话题详情
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type QueryTopicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryTopicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTopicDetailLogic {
	return &QueryTopicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryTopicDetail 查询话题详情
func (l *QueryTopicDetailLogic) QueryTopicDetail(req *types.QueryTopicDetailReq) (resp *types.QueryTopicDetailResp, err error) {

	detail, err := l.svcCtx.TopicService.QueryTopicDetail(l.ctx, &cmsclient.QueryTopicDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询话题详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryTopicDetailData{
		Id:             detail.Id,             // 主键id
		CategoryId:     detail.CategoryId,     // 关联分类id
		Name:           detail.Name,           // 话题名称
		StartTime:      detail.StartTime,      // 话题开始时间
		EndTime:        detail.EndTime,        // 话题结束时间
		AttendCount:    detail.AttendCount,    // 参与人数
		AttentionCount: detail.AttentionCount, // 关注人数
		ReadCount:      detail.ReadCount,      // 阅读数
		AwardName:      detail.AwardName,      // 奖品名称
		AttendType:     detail.AttendType,     // 参与方式
		Content:        detail.Content,        // 话题内容
		CreateBy:       detail.CreateBy,       // 创建者
		CreateTime:     detail.CreateTime,     // 创建时间
		UpdateBy:       detail.UpdateBy,       // 更新者
		UpdateTime:     detail.UpdateTime,     // 更新时间
	}
	return &types.QueryTopicDetailResp{
		Code:    "000000",
		Message: "查询话题成功",
		Data:    data,
	}, nil
}
