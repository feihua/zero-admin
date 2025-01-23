package helpservicelogic

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

// UpdateHelpLogic 更新帮助
/*
Author: LiuFeiHua
Date: 2025/01/23 15:23:59
*/
type UpdateHelpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHelpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHelpLogic {
	return &UpdateHelpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateHelp 更新帮助
func (l *UpdateHelpLogic) UpdateHelp(in *cmsclient.UpdateHelpReq) (*cmsclient.UpdateHelpResp, error) {
	q := query.CmsHelp.WithContext(l.ctx)

	// 1.根据帮助id查询帮助是否已存在
	_, err := q.Where(query.CmsHelp.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据帮助id：%d,查询帮助失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询帮助失败"))
	}

	item := &model.CmsHelp{
		ID:         in.Id,         // 主键ID
		CategoryID: in.CategoryId, // 分类ID
		Icon:       in.Icon,       // 图标
		Title:      in.Title,      // 标题
		ShowStatus: in.ShowStatus, // 显示状态：0->不显示；1->显示
		ReadCount:  in.ReadCount,  // 阅读量
		Content:    in.Content,    // 内容
		UpdateBy:   in.UpdateBy,   // 更新者
	}

	// 2.帮助存在时,则直接更新帮助
	_, err = q.Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新帮助失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新帮助失败")
	}

	logc.Infof(l.ctx, "更新帮助成功,参数：%+v", in)
	return &cmsclient.UpdateHelpResp{}, nil
}
