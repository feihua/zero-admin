package homerecommendsubjectservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateRecommendSubjectStatusLogic 专题推荐
/*
Author: LiuFeiHua
Date: 2024/5/14 9:35
*/
type UpdateRecommendSubjectStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRecommendSubjectStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRecommendSubjectStatusLogic {
	return &UpdateRecommendSubjectStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateRecommendSubjectStatus 批量修改专题推荐状态
func (l *UpdateRecommendSubjectStatusLogic) UpdateRecommendSubjectStatus(in *smsclient.UpdateRecommendSubjectStatusReq) (*smsclient.UpdateRecommendSubjectStatusResp, error) {
	q := query.SmsHomeRecommendSubject
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.RecommendStatus, in.RecommendStatus)

	if err != nil {
		return nil, err
	}

	return &smsclient.UpdateRecommendSubjectStatusResp{}, nil
}
