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

// DeleteTopicLogic 删除话题
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type DeleteTopicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTopicLogic {
	return &DeleteTopicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteTopic 删除话题
func (l *DeleteTopicLogic) DeleteTopic(in *cmsclient.DeleteTopicReq) (*cmsclient.DeleteTopicResp, error) {
	q := query.CmsTopic

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除话题失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除话题失败")
	}

	logc.Infof(l.ctx, "删除话题成功,参数：%+v", in)
	return &cmsclient.DeleteTopicResp{}, nil
}
