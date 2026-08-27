package topic_category

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

// QueryTopicCategoryDetailLogic 查询话题分类详情
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type QueryTopicCategoryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryTopicCategoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTopicCategoryDetailLogic {
	return &QueryTopicCategoryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryTopicCategoryDetail 查询话题分类详情
func (l *QueryTopicCategoryDetailLogic) QueryTopicCategoryDetail(req *types.QueryTopicCategoryDetailReq) (resp *types.QueryTopicCategoryDetailResp, err error) {

	detail, err := l.svcCtx.TopicCategoryService.QueryTopicCategoryDetail(l.ctx, &cmsclient.QueryTopicCategoryDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询话题分类详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryTopicCategoryDetailData{
		Id:           detail.Id,           // 主键id
		Name:         detail.Name,         // 分类名称
		Icon:         detail.Icon,         // 分类图标
		SubjectCount: detail.SubjectCount, // 专题数量
		ShowStatus:   detail.ShowStatus,   // 显示状态：0->不显示；1->显示
		Sort:         detail.Sort,         // 排序
		CreateBy:     detail.CreateBy,     // 创建者
		CreateTime:   detail.CreateTime,   // 创建时间
		UpdateBy:     detail.UpdateBy,     // 更新者
		UpdateTime:   detail.UpdateTime,   // 更新时间
	}
	return &types.QueryTopicCategoryDetailResp{
		Code:    "000000",
		Message: "查询话题分类成功",
		Data:    data,
	}, nil
}
