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

// QueryTopicListLogic 查询话题列表
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type QueryTopicListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryTopicListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTopicListLogic {
	return &QueryTopicListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryTopicList 查询话题列表
func (l *QueryTopicListLogic) QueryTopicList(in *cmsclient.QueryTopicListReq) (*cmsclient.QueryTopicListResp, error) {
	topic := query.CmsTopic
	q := topic.WithContext(l.ctx)
	if in.CategoryId != 2 {
		q = q.Where(topic.CategoryID.Eq(in.CategoryId))
	}
	if len(in.Name) > 0 {
		q = q.Where(topic.Name.Like("%" + in.Name + "%"))
	}
	// if in.StartTime != 2 {
	// 	q = q.Where(topic.StartTime.Eq(in.StartTime))
	// }
	// if in.EndTime != 2 {
	// 	q = q.Where(topic.EndTime.Eq(in.EndTime))
	// }
	if in.AttendCount != 2 {
		q = q.Where(topic.AttendCount.Eq(in.AttendCount))
	}
	if in.AttentionCount != 2 {
		q = q.Where(topic.AttentionCount.Eq(in.AttentionCount))
	}
	if in.ReadCount != 2 {
		q = q.Where(topic.ReadCount.Eq(in.ReadCount))
	}
	if len(in.AwardName) > 0 {
		q = q.Where(topic.AwardName.Like("%" + in.AwardName + "%"))
	}
	if len(in.AttendType) > 0 {
		q = q.Where(topic.AttendType.Like("%" + in.AttendType + "%"))
	}
	if len(in.Content) > 0 {
		q = q.Where(topic.Content.Like("%" + in.Content + "%"))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询话题列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询话题列表失败")
	}

	var list []*cmsclient.TopicListData

	for _, item := range result {
		list = append(list, &cmsclient.TopicListData{
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

		})
	}

	logc.Infof(l.ctx, "查询话题列表信息,参数：%+v,响应：%+v", in, list)

	return &cmsclient.QueryTopicListResp{
		Total: count,
		List:  list,
	}, nil
}
