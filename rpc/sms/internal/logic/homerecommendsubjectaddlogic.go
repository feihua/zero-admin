package logic

import (
	"context"
	"zero-admin/rpc/model/smsmodel"

	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/sms"

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

func (l *HomeRecommendSubjectAddLogic) HomeRecommendSubjectAdd(in *sms.HomeRecommendSubjectAddReq) (*sms.HomeRecommendSubjectAddResp, error) {
	_, err := l.svcCtx.SmsHomeRecommendSubjectModel.Insert(smsmodel.SmsHomeRecommendSubject{
		SubjectId:       in.SubjectId,
		SubjectName:     in.SubjectName,
		RecommendStatus: in.RecommendStatus,
		Sort:            in.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &sms.HomeRecommendSubjectAddResp{}, nil
}
