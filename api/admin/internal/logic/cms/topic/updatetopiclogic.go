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

// UpdateTopicLogic 更新话题
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type UpdateTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicLogic {
	return &UpdateTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateTopic 更新话题
func (l *UpdateTopicLogic) UpdateTopic(req *types.UpdateTopicReq) (resp *types.UpdateTopicResp, err error) {

	updateBy := l.ctx.Value("userName").(string)
	_, err = l.svcCtx.TopicService.UpdateTopic(l.ctx, &cmsclient.UpdateTopicReq{
		Id:         req.Id,         // 主键id
		CategoryId: req.CategoryId, // 关联分类id
		Name:       req.Name,       // 话题名称
		StartTime:  req.StartTime,  // 话题开始时间
		EndTime:    req.EndTime,    // 话题结束时间
		AwardName:  req.AwardName,  // 奖品名称
		AttendType: req.AttendType, // 参与方式
		Content:    req.Content,    // 话题内容
		UpdateBy:   updateBy,       // 更新者

	})

	if err != nil {
		logc.Errorf(l.ctx, "更新话题失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateTopicResp{
		Code:    "000000",
		Message: "更新话题成功",
	}, nil
}
