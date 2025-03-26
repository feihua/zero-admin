package topicservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateTopicStatusLogic 更新话题
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type UpdateTopicStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTopicStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicStatusLogic {
	return &UpdateTopicStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateTopicStatus 更新话题状态
func (l *UpdateTopicStatusLogic) UpdateTopicStatus(in *cmsclient.UpdateTopicStatusReq) (*cmsclient.UpdateTopicStatusResp, error) {
	q := query.CmsTopic

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.StartTime, 1)

	if err != nil {
		logc.Errorf(l.ctx, "更新话题状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新话题状态失败")
	}

	return &cmsclient.UpdateTopicStatusResp{}, nil
}
