package helpcategoryservicelogic

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

// AddHelpCategoryLogic 添加帮助分类
/*
Author: LiuFeiHua
Date: 2025/01/23 15:23:59
*/
type AddHelpCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddHelpCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHelpCategoryLogic {
	return &AddHelpCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddHelpCategory 添加帮助分类
func (l *AddHelpCategoryLogic) AddHelpCategory(in *cmsclient.AddHelpCategoryReq) (*cmsclient.AddHelpCategoryResp, error) {
	q := query.CmsHelpCategory

	item := &model.CmsHelpCategory{
		Name:       in.Name,       // 分类名称
		Icon:       in.Icon,       // 分类图标
		HelpCount:  in.HelpCount,  // 专题数量
		ShowStatus: in.ShowStatus, // 显示状态：0->不显示；1->显示
		Sort:       in.Sort,       // 排序
		CreateBy:   in.CreateBy,   // 创建者
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加帮助分类失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加帮助分类失败")
	}

	logc.Infof(l.ctx, "添加帮助分类成功,参数：%+v", in)
	return &cmsclient.AddHelpCategoryResp{}, nil
}
