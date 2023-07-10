package homerecommendsubjectservicelogic

import (
	"context"
	"zero-admin/rpc/model/smsmodel"
	"zero-admin/rpc/sms/smsclient"

	"zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeRecommendSubjectAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeRecommendSubjectAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendSubjectAddLogic {
	return &HomeRecommendSubjectAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeRecommendSubjectAddLogic) HomeRecommendSubjectAdd(in *smsclient.HomeRecommendSubjectAddReq) (*smsclient.HomeRecommendSubjectAddResp, error) {
	for _, data := range in.RecommendSubjectAddData {
		homeBrand, _ := l.svcCtx.SmsHomeRecommendSubjectModel.FindOneBySubjectId(l.ctx, data.SubjectId)
		if homeBrand == nil {
			_, err := l.svcCtx.SmsHomeRecommendSubjectModel.Insert(l.ctx, &smsmodel.SmsHomeRecommendSubject{
				SubjectId:       data.SubjectId,
				SubjectName:     data.SubjectName,
				RecommendStatus: data.RecommendStatus,
				Sort:            data.Sort,
			})
			if err != nil {
				return nil, err
			}
		}
	}

	return &smsclient.HomeRecommendSubjectAddResp{}, nil
}
