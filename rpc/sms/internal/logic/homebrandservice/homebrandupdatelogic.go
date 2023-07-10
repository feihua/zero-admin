package homebrandservicelogic

import (
	"context"
	"zero-admin/rpc/model/smsmodel"
	"zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
	"zero-admin/rpc/sms/internal/svc"
)

type HomeBrandUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeBrandUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeBrandUpdateLogic {
	return &HomeBrandUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeBrandUpdateLogic) HomeBrandUpdate(in *smsclient.HomeBrandUpdateReq) (*smsclient.HomeBrandUpdateResp, error) {
	err := l.svcCtx.SmsHomeBrandModel.Update(l.ctx, &smsmodel.SmsHomeBrand{
		Id:              in.Id,
		BrandId:         in.BrandId,
		BrandName:       in.BrandName,
		RecommendStatus: in.RecommendStatus,
		Sort:            in.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &smsclient.HomeBrandUpdateResp{}, nil
}
