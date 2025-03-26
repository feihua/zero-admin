package helpcategoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateHelpCategoryStatusLogic 更新帮助分类
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type UpdateHelpCategoryStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHelpCategoryStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHelpCategoryStatusLogic {
	return &UpdateHelpCategoryStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateHelpCategoryStatus 更新帮助分类状态
func (l *UpdateHelpCategoryStatusLogic) UpdateHelpCategoryStatus(in *cmsclient.UpdateHelpCategoryStatusReq) (*cmsclient.UpdateHelpCategoryStatusResp, error) {
	q := query.CmsHelpCategory

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.ShowStatus, in.ShowStatus)

	if err != nil {
		logc.Errorf(l.ctx, "更新帮助分类状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新帮助分类状态失败")
	}

	return &cmsclient.UpdateHelpCategoryStatusResp{}, nil
}
