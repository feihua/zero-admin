package topicservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryTopicDetailLogic 查询话题详情
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type QueryTopicDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryTopicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTopicDetailLogic {
	return &QueryTopicDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryTopicDetail 查询话题详情
func (l *QueryTopicDetailLogic) QueryTopicDetail(in *cmsclient.QueryTopicDetailReq) (*cmsclient.QueryTopicDetailResp, error) {
	item, err := query.CmsTopic.WithContext(l.ctx).Where(query.CmsTopic.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询话题详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询话题详情失败")
	}

	data := &cmsclient.QueryTopicDetailResp{
		Id:             item.ID,                                      // 主键ID
		CategoryId:     item.CategoryID,                              // 关联分类id
		Name:           item.Name,                                    // 话题名称
		StartTime:      item.StartTime.Format("2006-01-02 15:04:05"), // 话题开始时间
		EndTime:        item.EndTime.Format("2006-01-02 15:04:05"),   // 话题结束时间
		AttendCount:    item.AttendCount,                             // 参与人数
		AttentionCount: item.AttentionCount,                          // 关注人数
		ReadCount:      item.ReadCount,                               // 阅读数
		AwardName:      item.AwardName,                               // 奖品名称
		AttendType:     item.AttendType,                              // 参与方式
		Content:        item.Content,                                 // 话题内容
		CreateBy:       item.CreateBy,                                // 创建者
		CreateTime:     time_util.TimeToStr(item.CreateTime),         // 创建时间
		UpdateBy:       item.UpdateBy,                                // 更新者
		UpdateTime:     time_util.TimeToString(item.UpdateTime),      // 更新时间
	}

	logc.Infof(l.ctx, "查询话题详情,参数：%+v,响应：%+v", in, data)
	return data, nil
}
