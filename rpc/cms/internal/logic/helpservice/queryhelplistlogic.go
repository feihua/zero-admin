package helpservicelogic

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

// QueryHelpListLogic 查询帮助列表
/*
Author: LiuFeiHua
Date: 2025/01/23 15:23:59
*/
type QueryHelpListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHelpListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHelpListLogic {
	return &QueryHelpListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryHelpList 查询帮助列表
func (l *QueryHelpListLogic) QueryHelpList(in *cmsclient.QueryHelpListReq) (*cmsclient.QueryHelpListResp, error) {
	help := query.CmsHelp
	q := help.WithContext(l.ctx)
	if in.CategoryId != 2 {
		q = q.Where(help.CategoryID.Eq(in.CategoryId))
	}
	if len(in.Icon) > 0 {
		q = q.Where(help.Icon.Like("%" + in.Icon + "%"))
	}
	if len(in.Title) > 0 {
		q = q.Where(help.Title.Like("%" + in.Title + "%"))
	}
	if in.ShowStatus != 2 {
		q = q.Where(help.ShowStatus.Eq(in.ShowStatus))
	}
	if in.ReadCount != 2 {
		q = q.Where(help.ReadCount.Eq(in.ReadCount))
	}
	if len(in.Content) > 0 {
		q = q.Where(help.Content.Like("%" + in.Content + "%"))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询帮助列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询帮助列表失败")
	}

	var list []*cmsclient.HelpListData

	for _, item := range result {
		list = append(list, &cmsclient.HelpListData{
			Id:         item.ID,                                 // 主键ID
			CategoryId: item.CategoryID,                         // 分类ID
			Icon:       item.Icon,                               // 图标
			Title:      item.Title,                              // 标题
			ShowStatus: item.ShowStatus,                         // 显示状态：0->不显示；1->显示
			ReadCount:  item.ReadCount,                          // 阅读量
			Content:    item.Content,                            // 内容
			CreateBy:   item.CreateBy,                           // 创建者
			CreateTime: time_util.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:   item.UpdateBy,                           // 更新者
			UpdateTime: time_util.TimeToString(item.UpdateTime), // 更新时间

		})
	}

	return &cmsclient.QueryHelpListResp{
		Total: count,
		List:  list,
	}, nil
}
