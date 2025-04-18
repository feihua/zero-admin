package homerecommendsubjectservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateHomeRecommendSubjectStatusLogic 更新首页推荐专题状态
/*
Author: LiuFeiHua
Date: 2024/6/12 17:59
*/
type UpdateHomeRecommendSubjectStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHomeRecommendSubjectStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomeRecommendSubjectStatusLogic {
	return &UpdateHomeRecommendSubjectStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateHomeRecommendSubjectStatus 更新首页推荐专题状态
func (l *UpdateHomeRecommendSubjectStatusLogic) UpdateHomeRecommendSubjectStatus(in *smsclient.UpdateHomeRecommendSubjectStatusReq) (*smsclient.UpdateHomeRecommendSubjectStatusResp, error) {
	q := query.SmsHomeRecommendSubject
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.RecommendStatus, in.RecommendStatus)

	if err != nil {
		logc.Errorf(l.ctx, "更新首页推荐专题状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新首页推荐专题状态失败")
	}

	return &smsclient.UpdateHomeRecommendSubjectStatusResp{}, nil
}
