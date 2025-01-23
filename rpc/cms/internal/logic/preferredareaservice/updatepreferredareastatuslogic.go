package preferredareaservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdatePreferredAreaStatusLogic 更新优选专区
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type UpdatePreferredAreaStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePreferredAreaStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePreferredAreaStatusLogic {
	return &UpdatePreferredAreaStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdatePreferredAreaStatus 更新优选专区状态
func (l *UpdatePreferredAreaStatusLogic) UpdatePreferredAreaStatus(in *cmsclient.UpdatePreferredAreaStatusReq) (*cmsclient.UpdatePreferredAreaStatusResp, error) {
	q := query.CmsPreferredArea

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.ShowStatus, in.ShowStatus)

	if err != nil {
		logc.Errorf(l.ctx, "更新优选专区状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新优选专区状态失败")
	}

	logc.Infof(l.ctx, "更新优选专区状态成功,参数：%+v", in)
	return &cmsclient.UpdatePreferredAreaStatusResp{}, nil
}
