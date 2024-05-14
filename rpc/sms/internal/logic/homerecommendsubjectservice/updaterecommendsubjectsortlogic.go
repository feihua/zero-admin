package homerecommendsubjectservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateRecommendSubjectSortLogic 专题推荐
/*
Author: LiuFeiHua
Date: 2024/5/14 9:34
*/
type UpdateRecommendSubjectSortLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRecommendSubjectSortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRecommendSubjectSortLogic {
	return &UpdateRecommendSubjectSortLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateRecommendSubjectSort 修改专题推荐排序
func (l *UpdateRecommendSubjectSortLogic) UpdateRecommendSubjectSort(in *smsclient.UpdateRecommendSubjectSortReq) (*smsclient.UpdateRecommendSubjectSortResp, error) {
	q := query.SmsHomeRecommendSubject
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Update(q.Sort, in.Sort)

	if err != nil {
		return nil, err
	}

	return &smsclient.UpdateRecommendSubjectSortResp{}, nil
}
