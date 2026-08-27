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

// UpdateTopicStatusLogic 更新话题状态状态
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type UpdateTopicStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTopicStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicStatusLogic {
	return &UpdateTopicStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateTopicStatus 更新话题状态
func (l *UpdateTopicStatusLogic) UpdateTopicStatus(req *types.UpdateTopicStatusReq) (resp *types.UpdateTopicStatusResp, err error) {
	_, err = l.svcCtx.TopicService.UpdateTopicStatus(l.ctx, &cmsclient.UpdateTopicStatusReq{
		Ids:      req.Ids, // 主键id
		UpdateBy: l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新话题状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateTopicStatusResp{
		Code:    "000000",
		Message: "更新话题状态成功",
	}, nil
}
