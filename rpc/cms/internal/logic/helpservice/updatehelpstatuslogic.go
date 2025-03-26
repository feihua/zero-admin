package helpservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateHelpStatusLogic 更新帮助
/*
Author: LiuFeiHua
Date: 2025/01/23 15:23:59
*/
type UpdateHelpStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHelpStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHelpStatusLogic {
	return &UpdateHelpStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateHelpStatus 更新帮助状态
func (l *UpdateHelpStatusLogic) UpdateHelpStatus(in *cmsclient.UpdateHelpStatusReq) (*cmsclient.UpdateHelpStatusResp, error) {
	q := query.CmsHelp

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.ShowStatus, in.ShowStatus)

	if err != nil {
		logc.Errorf(l.ctx, "更新帮助状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新帮助状态失败")
	}

	return &cmsclient.UpdateHelpStatusResp{}, nil
}
