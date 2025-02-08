package homebrandservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryHomeBrandDetailLogic 查询首页推荐品牌表详情
/*
Author: 刘飞华
Date: 2025/02/07 10:11:43
*/
type QueryHomeBrandDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHomeBrandDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHomeBrandDetailLogic {
	return &QueryHomeBrandDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryHomeBrandDetail 查询首页推荐品牌表详情
func (l *QueryHomeBrandDetailLogic) QueryHomeBrandDetail(in *smsclient.QueryHomeBrandDetailReq) (*smsclient.QueryHomeBrandDetailResp, error) {
	item, err := query.SmsHomeBrand.WithContext(l.ctx).Where(query.SmsHomeBrand.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询首页推荐品牌表详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询首页推荐品牌表详情失败")
	}

	data := &smsclient.QueryHomeBrandDetailResp{
		Id:              item.ID,              //
		BrandId:         item.BrandID,         // 商品品牌id
		BrandName:       item.BrandName,       // 商品品牌名称
		RecommendStatus: item.RecommendStatus, // 推荐状态：0->不推荐;1->推荐
		Sort:            item.Sort,            // 排序

	}

	logc.Infof(l.ctx, "查询首页推荐品牌表详情,参数：%+v,响应：%+v", in, data)
	return data, nil
}
