package topicservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/model"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddTopicLogic 添加话题
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type AddTopicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTopicLogic {
	return &AddTopicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddTopic 添加话题
func (l *AddTopicLogic) AddTopic(in *cmsclient.AddTopicReq) (*cmsclient.AddTopicResp, error) {
	q := query.CmsTopic

	item := &model.CmsTopic{
		CategoryID: in.CategoryId, // 关联分类id
		Name:       in.Name,       // 话题名称
		// StartTime:      in.StartTime,      // 话题开始时间
		// EndTime:        in.EndTime,        // 话题结束时间
		AttendCount:    in.AttendCount,    // 参与人数
		AttentionCount: in.AttentionCount, // 关注人数
		ReadCount:      in.ReadCount,      // 阅读数
		AwardName:      in.AwardName,      // 奖品名称
		AttendType:     in.AttendType,     // 参与方式
		Content:        in.Content,        // 话题内容
		CreateBy:       in.CreateBy,       // 创建者
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加话题失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加话题失败")
	}

	return &cmsclient.AddTopicResp{}, nil
}
