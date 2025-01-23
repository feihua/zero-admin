package subjectcategoryservicelogic

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

	// 1.根据专题分类id查询专题分类是否已存在
	_, err := q.Where(query.CmsSubjectCategory.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据专题分类id：%d,查询专题分类失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询专题分类失败"))
	}

	item := &model.CmsSubjectCategory{
		ID:           in.Id,           // 主键ID
		Name:         in.Name,         // 专题分类名称
		Icon:         in.Icon,         // 分类图标
		SubjectCount: in.SubjectCount, // 专题数量
		ShowStatus:   in.ShowStatus,   // 显示状态：0->不显示；1->显示
		Sort:         in.Sort,         // 排序
		UpdateBy:     in.UpdateBy,     // 更新者
	}

	// 2.专题分类存在时,则直接更新专题分类
	_, err = q.Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新专题分类失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新专题分类失败")
	}

	logc.Infof(l.ctx, "更新专题分类成功,参数：%+v", in)
	return &cmsclient.UpdateSubjectCategoryResp{}, nil
}
