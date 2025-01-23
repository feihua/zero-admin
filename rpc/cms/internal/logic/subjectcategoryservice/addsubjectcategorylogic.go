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
)

// AddSubjectCategoryLogic 添加专题分类
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type AddSubjectCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSubjectCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSubjectCategoryLogic {
	return &AddSubjectCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddSubjectCategory 添加专题分类
func (l *AddSubjectCategoryLogic) AddSubjectCategory(in *cmsclient.AddSubjectCategoryReq) (*cmsclient.AddSubjectCategoryResp, error) {
	q := query.CmsSubjectCategory

	item := &model.CmsSubjectCategory{
		Name:         in.Name,         // 专题分类名称
		Icon:         in.Icon,         // 分类图标
		SubjectCount: in.SubjectCount, // 专题数量
		ShowStatus:   in.ShowStatus,   // 显示状态：0->不显示；1->显示
		Sort:         in.Sort,         // 排序
		CreateBy:     in.CreateBy,     // 创建者
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加专题分类失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加专题分类失败")
	}

	logc.Infof(l.ctx, "添加专题分类成功,参数：%+v", in)
	return &cmsclient.AddSubjectCategoryResp{}, nil
}
