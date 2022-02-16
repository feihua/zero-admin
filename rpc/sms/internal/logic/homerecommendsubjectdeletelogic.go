package logic

import (
	"context"

	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeRecommendSubjectDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeRecommendSubjectDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendSubjectDeleteLogic {
	return &HomeRecommendSubjectDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeRecommendSubjectDeleteLogic) HomeRecommendSubjectDelete(in *sms.HomeRecommendSubjectDeleteReq) (*sms.HomeRecommendSubjectDeleteResp, error) {
	err := l.svcCtx.SmsHomeRecommendSubjectModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sms.HomeRecommendSubjectDeleteResp{}, nil
}
