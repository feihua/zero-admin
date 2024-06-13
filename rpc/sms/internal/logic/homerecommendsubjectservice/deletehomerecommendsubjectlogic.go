package homerecommendsubjectservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteHomeRecommendSubjectLogic 删除首页推荐专题表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:58
*/
type DeleteHomeRecommendSubjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteHomeRecommendSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHomeRecommendSubjectLogic {
	return &DeleteHomeRecommendSubjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteHomeRecommendSubject 删除首页推荐专题表
func (l *DeleteHomeRecommendSubjectLogic) DeleteHomeRecommendSubject(in *smsclient.DeleteHomeRecommendSubjectReq) (*smsclient.DeleteHomeRecommendSubjectResp, error) {
	q := query.SmsHomeRecommendSubject
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &smsclient.DeleteHomeRecommendSubjectResp{}, nil
}
