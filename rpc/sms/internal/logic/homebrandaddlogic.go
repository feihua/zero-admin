package logic

import (
	"context"
	"go-zero-admin/rpc/model/smsmodel"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeBrandAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeBrandAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeBrandAddLogic {
	return &HomeBrandAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeBrandAddLogic) HomeBrandAdd(in *sms.HomeBrandAddReq) (*sms.HomeBrandAddResp, error) {
	_, err := l.svcCtx.SmsHomeBrandModel.Insert(smsmodel.SmsHomeBrand{
		BrandId:         in.BrandId,
		BrandName:       in.BrandName,
		RecommendStatus: in.RecommendStatus,
		Sort:            in.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &sms.HomeBrandAddResp{}, nil
}
