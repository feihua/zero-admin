package homebrandservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeBrandAddLogic 首页品牌信息
/*
Author: LiuFeiHua
Date: 2024/5/13 15:57
*/
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

// HomeBrandAdd 添加首页品牌信息
func (l *HomeBrandAddLogic) HomeBrandAdd(in *smsclient.HomeBrandAddReq) (*smsclient.HomeBrandAddResp, error) {
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
		return &smsclient.HomeBrandAddResp{}, nil
	}

	err := q.WithContext(l.ctx).CreateInBatches(brands, len(brands))
	if err != nil {
		return nil, err
	}

	return &smsclient.HomeBrandAddResp{}, nil
}
