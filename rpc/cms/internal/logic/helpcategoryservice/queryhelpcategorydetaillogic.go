package helpcategoryservicelogic

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

// QueryHelpCategoryDetailLogic 查询帮助分类详情
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type QueryHelpCategoryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHelpCategoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHelpCategoryDetailLogic {
	return &QueryHelpCategoryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryHelpCategoryDetail 查询帮助分类详情
func (l *QueryHelpCategoryDetailLogic) QueryHelpCategoryDetail(in *cmsclient.QueryHelpCategoryDetailReq) (*cmsclient.QueryHelpCategoryDetailResp, error) {
	item, err := query.CmsHelpCategory.WithContext(l.ctx).Where(query.CmsHelpCategory.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询帮助分类详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询帮助分类详情失败")
	}

	data := &cmsclient.QueryHelpCategoryDetailResp{
		Id:         item.ID,                                 // 主键ID
		Name:       item.Name,                               // 分类名称
		Icon:       item.Icon,                               // 分类图标
		HelpCount:  item.HelpCount,                          // 专题数量
		ShowStatus: item.ShowStatus,                         // 显示状态：0->不显示；1->显示
		Sort:       item.Sort,                               // 排序
		CreateBy:   item.CreateBy,                           // 创建者
		CreateTime: time_util.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:   item.UpdateBy,                           // 更新者
		UpdateTime: time_util.TimeToString(item.UpdateTime), // 更新时间
	}

	return data, nil
}
