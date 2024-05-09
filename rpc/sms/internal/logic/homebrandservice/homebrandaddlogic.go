package homebrandservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
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
	q := query.SmsHomeBrand
	for _, data := range in.BrandAddData {
		homeBrand, _ := q.WithContext(l.ctx).Where(q.BrandID.Eq(data.BrandId)).First()
		//存在的时候不添加
		if homeBrand == nil {
			err := q.WithContext(l.ctx).Create(&model.SmsHomeBrand{
				BrandID:         data.BrandId,
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
