package helpservicelogic

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

// AddHelpLogic 添加帮助
/*
Author: LiuFeiHua
Date: 2025/01/23 15:23:59
*/
type AddHelpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddHelpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHelpLogic {
	return &AddHelpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddHelp 添加帮助
func (l *AddHelpLogic) AddHelp(in *cmsclient.AddHelpReq) (*cmsclient.AddHelpResp, error) {
	q := query.CmsHelp

	item := &model.CmsHelp{
		CategoryID: in.CategoryId, // 分类ID
		Icon:       in.Icon,       // 图标
		Title:      in.Title,      // 标题
		ShowStatus: in.ShowStatus, // 显示状态：0->不显示；1->显示
		ReadCount:  in.ReadCount,  // 阅读量
		Content:    in.Content,    // 内容
		CreateBy:   in.CreateBy,   // 创建者
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加帮助失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加帮助失败")
	}

	logc.Infof(l.ctx, "添加帮助成功,参数：%+v", in)
	return &cmsclient.AddHelpResp{}, nil
}
