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

// UpdateTopicCategoryStatusLogic 更新话题分类状态状态
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type UpdateTopicCategoryStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTopicCategoryStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicCategoryStatusLogic {
	return &UpdateTopicCategoryStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateTopicCategoryStatus 更新话题分类状态
func (l *UpdateTopicCategoryStatusLogic) UpdateTopicCategoryStatus(req *types.UpdateTopicCategoryStatusReq) (resp *types.UpdateTopicCategoryStatusResp, err error) {
	_, err = l.svcCtx.TopicCategoryService.UpdateTopicCategoryStatus(l.ctx, &cmsclient.UpdateTopicCategoryStatusReq{
		Ids:        req.Ids,        // 主键id
		ShowStatus: req.ShowStatus, // 显示状态：0->不显示；1->显示
		UpdateBy:   l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新话题分类状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateTopicCategoryStatusResp{
		Code:    "000000",
		Message: "更新话题分类状态成功",
	}, nil
}
