package subjectcategoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/model"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

// UpdateSubjectCategoryLogic 更新专题分类
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type UpdateSubjectCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSubjectCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubjectCategoryLogic {
	return &UpdateSubjectCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateSubjectCategory 更新专题分类
func (l *UpdateSubjectCategoryLogic) UpdateSubjectCategory(in *cmsclient.UpdateSubjectCategoryReq) (*cmsclient.UpdateSubjectCategoryResp, error) {
	q := query.CmsSubjectCategory.WithContext(l.ctx)

	s, err := q.Where(query.CmsSubjectCategory.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errors.New("专题分类不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询专题分类异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询专题分类异常")
	}

	now := time.Now()
	item := &model.CmsSubjectCategory{
		ID:           in.Id,           // 主键ID
		Name:         in.Name,         // 专题分类名称
		Icon:         in.Icon,         // 分类图标
		SubjectCount: in.SubjectCount, // 专题数量
		ShowStatus:   in.ShowStatus,   // 显示状态：0->不显示；1->显示
		Sort:         in.Sort,         // 排序
		CreateBy:     s.CreateBy,      // 创建者
		CreateTime:   s.CreateTime,    // 创建时间
		UpdateBy:     in.UpdateBy,     // 更新者
		UpdateTime:   &now,            // 更新时间
	}

	// 2.专题分类存在时,则直接更新专题分类
	err = l.svcCtx.DB.Model(&model.CmsSubjectCategory{}).WithContext(l.ctx).Where(query.CmsSubjectCategory.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新专题分类失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新专题分类失败")
	}

	return &cmsclient.UpdateSubjectCategoryResp{}, nil
}
