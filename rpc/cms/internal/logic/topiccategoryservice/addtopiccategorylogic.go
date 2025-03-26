package topiccategoryservicelogic

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

// AddTopicCategoryLogic 添加话题分类
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type AddTopicCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTopicCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTopicCategoryLogic {
	return &AddTopicCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddTopicCategory 添加话题分类
func (l *AddTopicCategoryLogic) AddTopicCategory(in *cmsclient.AddTopicCategoryReq) (*cmsclient.AddTopicCategoryResp, error) {
	q := query.CmsTopicCategory

	item := &model.CmsTopicCategory{
		Name:         in.Name,         // 分类名称
		Icon:         in.Icon,         // 分类图标
		SubjectCount: in.SubjectCount, // 专题数量
		ShowStatus:   in.ShowStatus,   // 显示状态：0->不显示；1->显示
		Sort:         in.Sort,         // 排序
		CreateBy:     in.CreateBy,     // 创建者
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加话题分类失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加话题分类失败")
	}

	return &cmsclient.AddTopicCategoryResp{}, nil
}
