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

// QueryHelpDetailLogic 查询帮助详情
/*
Author: LiuFeiHua
Date: 2025/01/23 15:23:59
*/
type QueryHelpDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHelpDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHelpDetailLogic {
	return &QueryHelpDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryHelpDetail 查询帮助详情
func (l *QueryHelpDetailLogic) QueryHelpDetail(in *cmsclient.QueryHelpDetailReq) (*cmsclient.QueryHelpDetailResp, error) {
	item, err := query.CmsHelp.WithContext(l.ctx).Where(query.CmsHelp.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询帮助详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询帮助详情失败")
	}

	data := &cmsclient.QueryHelpDetailResp{
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
	}

	return data, nil
}
