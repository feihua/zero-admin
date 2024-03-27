package homerecommendsubjectservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/model/smsmodel"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeRecommendSubjectUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeRecommendSubjectUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendSubjectUpdateLogic {
	return &HomeRecommendSubjectUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeRecommendSubjectUpdateLogic) HomeRecommendSubjectUpdate(in *smsclient.HomeRecommendSubjectUpdateReq) (*smsclient.HomeRecommendSubjectUpdateResp, error) {
	err := l.svcCtx.SmsHomeRecommendSubjectModel.Update(l.ctx, &smsmodel.SmsHomeRecommendSubject{
		Id:              in.Id,
		SubjectId:       in.SubjectId,
		SubjectName:     in.SubjectName,
		RecommendStatus: in.RecommendStatus,
		Sort:            in.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &smsclient.HomeRecommendSubjectUpdateResp{}, nil
}
