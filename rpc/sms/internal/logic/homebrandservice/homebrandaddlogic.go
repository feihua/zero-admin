package homebrandservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/model/smsmodel"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *HomeBrandAddLogic) HomeBrandAdd(in *smsclient.HomeBrandAddReq) (*smsclient.HomeBrandAddResp, error) {
	for _, data := range in.BrandAddData {
		homeBrand, _ := l.svcCtx.SmsHomeBrandModel.FindOneByBrandId(l.ctx, data.BrandId)
		//存在的时候不添加
		if homeBrand == nil {
			_, err := l.svcCtx.SmsHomeBrandModel.Insert(l.ctx, &smsmodel.SmsHomeBrand{
				BrandId:         data.BrandId,
				BrandName:       data.BrandName,
				RecommendStatus: data.RecommendStatus,
				Sort:            data.Sort,
			})
			if err != nil {
				return nil, err
			}
		}

	}

	return &smsclient.HomeBrandAddResp{}, nil
}
