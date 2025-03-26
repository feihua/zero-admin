package topiccategoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryTopicCategoryDetailLogic 查询话题分类详情
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type QueryTopicCategoryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryTopicCategoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTopicCategoryDetailLogic {
	return &QueryTopicCategoryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryTopicCategoryDetail 查询话题分类详情
func (l *QueryTopicCategoryDetailLogic) QueryTopicCategoryDetail(in *cmsclient.QueryTopicCategoryDetailReq) (*cmsclient.QueryTopicCategoryDetailResp, error) {
	item, err := query.CmsTopicCategory.WithContext(l.ctx).Where(query.CmsTopicCategory.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询话题分类详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询话题分类详情失败")
	}

	data := &cmsclient.QueryTopicCategoryDetailResp{
		Id:           item.ID,                                 // 主键ID
		Name:         item.Name,                               // 分类名称
		Icon:         item.Icon,                               // 分类图标
		SubjectCount: item.SubjectCount,                       // 专题数量
		ShowStatus:   item.ShowStatus,                         // 显示状态：0->不显示；1->显示
		Sort:         item.Sort,                               // 排序
		CreateBy:     item.CreateBy,                           // 创建者
		CreateTime:   time_util.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:     item.UpdateBy,                           // 更新者
		UpdateTime:   time_util.TimeToString(item.UpdateTime), // 更新时间
	}

	return data, nil
}
