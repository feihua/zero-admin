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

// AddTopicCategoryLogic 添加话题分类
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type AddTopicCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTopicCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTopicCategoryLogic {
	return &AddTopicCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddTopicCategory 添加话题分类
func (l *AddTopicCategoryLogic) AddTopicCategory(req *types.AddTopicCategoryReq) (resp *types.AddTopicCategoryResp, err error) {
	createBy := l.ctx.Value("userName").(string)

	_, err = l.svcCtx.TopicCategoryService.AddTopicCategory(l.ctx, &cmsclient.AddTopicCategoryReq{
		Name:       req.Name,       // 分类名称
		Icon:       req.Icon,       // 分类图标
		ShowStatus: req.ShowStatus, // 显示状态：0->不显示；1->显示
		Sort:       req.Sort,       // 排序
		CreateBy:   createBy,       // 创建者

	})

	if err != nil {
		logc.Errorf(l.ctx, "添加话题分类失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddTopicCategoryResp{
		Code:    "000000",
		Message: "添加话题分类成功",
	}, nil
}
