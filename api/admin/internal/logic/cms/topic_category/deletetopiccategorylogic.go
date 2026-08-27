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

// DeleteTopicCategoryLogic 删除话题分类
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type DeleteTopicCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTopicCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTopicCategoryLogic {
	return &DeleteTopicCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteTopicCategory 删除话题分类
func (l *DeleteTopicCategoryLogic) DeleteTopicCategory(req *types.DeleteTopicCategoryReq) (resp *types.DeleteTopicCategoryResp, err error) {
	_, err = l.svcCtx.TopicCategoryService.DeleteTopicCategory(l.ctx, &cmsclient.DeleteTopicCategoryReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除话题分类失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.DeleteTopicCategoryResp{
		Code:    "000000",
		Message: "删除话题分类成功",
	}, nil
}
