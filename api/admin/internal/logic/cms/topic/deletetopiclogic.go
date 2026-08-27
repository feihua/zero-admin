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

// DeleteTopicLogic 删除话题
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type DeleteTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTopicLogic {
	return &DeleteTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteTopic 删除话题
func (l *DeleteTopicLogic) DeleteTopic(req *types.DeleteTopicReq) (resp *types.DeleteTopicResp, err error) {
	_, err = l.svcCtx.TopicService.DeleteTopic(l.ctx, &cmsclient.DeleteTopicReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除话题失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.DeleteTopicResp{
		Code:    "000000",
		Message: "删除话题成功",
	}, nil
}
