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

// QueryTopicListLogic 查询话题列表
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type QueryTopicListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryTopicListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTopicListLogic {
	return &QueryTopicListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryTopicList 查询话题列表
func (l *QueryTopicListLogic) QueryTopicList(req *types.QueryTopicListReq) (resp *types.QueryTopicListResp, err error) {
	result, err := l.svcCtx.TopicService.QueryTopicList(l.ctx, &cmsclient.QueryTopicListReq{
		PageNum:    int32(req.Current),
		PageSize:   int32(req.PageSize),
		CategoryId: req.CategoryId, // 关联分类id
		Name:       req.Name,       // 话题名称
		StartTime:  req.StartTime,  // 话题开始时间
		EndTime:    req.EndTime,    // 话题结束时间
		AwardName:  req.AwardName,  // 奖品名称
		AttendType: req.AttendType, // 参与方式

	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字话题列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryTopicListData

	for _, item := range result.List {
		list = append(list, &types.QueryTopicListData{
			Id:             item.Id,             // 主键id
			CategoryId:     item.CategoryId,     // 关联分类id
			Name:           item.Name,           // 话题名称
			StartTime:      item.StartTime,      // 话题开始时间
			EndTime:        item.EndTime,        // 话题结束时间
			AttendCount:    item.AttendCount,    // 参与人数
			AttentionCount: item.AttentionCount, // 关注人数
			ReadCount:      item.ReadCount,      // 阅读数
			AwardName:      item.AwardName,      // 奖品名称
			AttendType:     item.AttendType,     // 参与方式
			Content:        item.Content,        // 话题内容
			CreateBy:       item.CreateBy,       // 创建者
			CreateTime:     item.CreateTime,     // 创建时间
			UpdateBy:       item.UpdateBy,       // 更新者
			UpdateTime:     item.UpdateTime,     // 更新时间
		})
	}

	return &types.QueryTopicListResp{
		Code:     "000000",
		Message:  "查询话题列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
