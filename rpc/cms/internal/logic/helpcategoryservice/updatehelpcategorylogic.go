package helpcategoryservicelogic

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

// UpdateHelpCategoryLogic 更新帮助分类
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type UpdateHelpCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHelpCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHelpCategoryLogic {
	return &UpdateHelpCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateHelpCategory 更新帮助分类
func (l *UpdateHelpCategoryLogic) UpdateHelpCategory(in *cmsclient.UpdateHelpCategoryReq) (*cmsclient.UpdateHelpCategoryResp, error) {
	q := query.CmsHelpCategory.WithContext(l.ctx)

	// 1.根据帮助分类id查询帮助分类是否已存在
	_, err := q.Where(query.CmsHelpCategory.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据帮助分类id：%d,查询帮助分类失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询帮助分类失败"))
	}

	item := &model.CmsHelpCategory{
		ID:         in.Id,         // 主键ID
		Name:       in.Name,       // 分类名称
		Icon:       in.Icon,       // 分类图标
		HelpCount:  in.HelpCount,  // 专题数量
		ShowStatus: in.ShowStatus, // 显示状态：0->不显示；1->显示
		Sort:       in.Sort,       // 排序
		UpdateBy:   in.UpdateBy,   // 更新者
	}

	// 2.帮助分类存在时,则直接更新帮助分类
	_, err = q.Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新帮助分类失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新帮助分类失败")
	}

	logc.Infof(l.ctx, "更新帮助分类成功,参数：%+v", in)
	return &cmsclient.UpdateHelpCategoryResp{}, nil
}
