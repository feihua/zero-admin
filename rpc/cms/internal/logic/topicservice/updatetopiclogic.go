package topicservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/model"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateTopicLogic 更新话题
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type UpdateTopicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicLogic {
	return &UpdateTopicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateTopic 更新话题
func (l *UpdateTopicLogic) UpdateTopic(in *cmsclient.UpdateTopicReq) (*cmsclient.UpdateTopicResp, error) {
	q := query.CmsTopic.WithContext(l.ctx)

	// 1.根据话题id查询话题是否已存在
	_, err := q.Where(query.CmsTopic.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据话题id：%d,查询话题失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询话题失败"))
	}

	item := &model.CmsTopic{
		ID:         in.Id,         // 主键ID
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
		UpdateBy:       in.UpdateBy,       // 更新者
	}

	// 2.话题存在时,则直接更新话题
	_, err = q.Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新话题失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新话题失败")
	}

	return &cmsclient.UpdateTopicResp{}, nil
}
