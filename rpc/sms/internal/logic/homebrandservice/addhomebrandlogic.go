package homebrandservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddHomeBrandLogic 添加首页推荐品牌表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:52
*/
type AddHomeBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddHomeBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHomeBrandLogic {
	return &AddHomeBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddHomeBrand 添加首页推荐品牌表
func (l *AddHomeBrandLogic) AddHomeBrand(in *smsclient.AddHomeBrandReq) (*smsclient.AddHomeBrandResp, error) {
	q := query.SmsHomeBrand

	var brands []*model.SmsHomeBrand
	for _, item := range in.BrandAddData {
		homeBrand, _ := q.WithContext(l.ctx).Where(q.BrandID.Eq(item.BrandId)).First()
		//存在的时候不添加
		if homeBrand == nil {
			brands = append(brands, &model.SmsHomeBrand{
				BrandID:         item.BrandId,
				BrandName:       item.BrandName,
				RecommendStatus: item.RecommendStatus,
				Sort:            item.Sort,
			})
		}

	}

	if len(brands) == 0 {
		return &smsclient.AddHomeBrandResp{}, nil
	}

	err := q.WithContext(l.ctx).CreateInBatches(brands, len(brands))
	if err != nil {
		return nil, err
	}

	return &smsclient.AddHomeBrandResp{}, nil
}
