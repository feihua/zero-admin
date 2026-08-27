package topiccategoryservicelogic

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

// UpdateTopicCategoryLogic 更新话题分类
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type UpdateTopicCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTopicCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicCategoryLogic {
	return &UpdateTopicCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateTopicCategory 更新话题分类
func (l *UpdateTopicCategoryLogic) UpdateTopicCategory(in *cmsclient.UpdateTopicCategoryReq) (*cmsclient.UpdateTopicCategoryResp, error) {
	q := query.CmsTopicCategory

	// 1.根据话题分类id查询话题分类是否已存在
	detail, err := q.WithContext(l.ctx).Where(query.CmsTopicCategory.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据话题分类id：%d,查询话题分类失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("更新话题分类失败"))
	}

	count, err := q.WithContext(l.ctx).Where(q.Name.Eq(in.Name), q.ID.Neq(in.Id)).Count()

	if err != nil {
		return nil, errors.New("查询话题分类失败")
	}

	if count > 0 {
		return nil, errors.New("更新话题分类失败,分类名称已存在")
	}

	item := &model.CmsTopicCategory{
		ID:           in.Id,               // 主键ID
		Name:         in.Name,             // 分类名称
		Icon:         in.Icon,             // 分类图标
		SubjectCount: detail.SubjectCount, // 专题数量
		ShowStatus:   in.ShowStatus,       // 显示状态：0->不显示；1->显示
		Sort:         in.Sort,             // 排序
		UpdateBy:     in.UpdateBy,         // 更新者
	}

	// 2.话题分类存在时,则直接更新话题分类
	_, err = q.WithContext(l.ctx).Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新话题分类失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新话题分类失败")
	}

	return &cmsclient.UpdateTopicCategoryResp{}, nil
}
