package topiccategoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateTopicCategoryStatusLogic 更新话题分类
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type UpdateTopicCategoryStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTopicCategoryStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicCategoryStatusLogic {
	return &UpdateTopicCategoryStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateTopicCategoryStatus 更新话题分类状态
func (l *UpdateTopicCategoryStatusLogic) UpdateTopicCategoryStatus(in *cmsclient.UpdateTopicCategoryStatusReq) (*cmsclient.UpdateTopicCategoryStatusResp, error) {
	q := query.CmsTopicCategory

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.ShowStatus, in.ShowStatus)

	if err != nil {
		logc.Errorf(l.ctx, "更新话题分类状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新话题分类状态失败")
	}

	logc.Infof(l.ctx, "更新话题分类状态成功,参数：%+v", in)
	return &cmsclient.UpdateTopicCategoryStatusResp{}, nil
}
