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

// AddTopicLogic 添加话题
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type AddTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTopicLogic {
	return &AddTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddTopic 添加话题
func (l *AddTopicLogic) AddTopic(req *types.AddTopicReq) (resp *types.AddTopicResp, err error) {
	createBy := l.ctx.Value("userName").(string)

	_, err = l.svcCtx.TopicService.AddTopic(l.ctx, &cmsclient.AddTopicReq{
		CategoryId: req.CategoryId, // 关联分类id
		Name:       req.Name,       // 话题名称
		StartTime:  req.StartTime,  // 话题开始时间
		EndTime:    req.EndTime,    // 话题结束时间
		AwardName:  req.AwardName,  // 奖品名称
		AttendType: req.AttendType, // 参与方式
		Content:    req.Content,    // 话题内容
		CreateBy:   createBy,       // 创建者

	})

	if err != nil {
		logc.Errorf(l.ctx, "添加话题失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddTopicResp{
		Code:    "000000",
		Message: "添加话题成功",
	}, nil
}
