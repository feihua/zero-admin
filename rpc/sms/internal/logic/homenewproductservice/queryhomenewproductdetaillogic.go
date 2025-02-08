package homenewproductservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryHomeNewProductDetailLogic 查询新鲜好物表详情
/*
Author: 刘飞华
Date: 2025/02/07 10:11:43
*/
type QueryHomeNewProductDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHomeNewProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHomeNewProductDetailLogic {
	return &QueryHomeNewProductDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryHomeNewProductDetail 查询新鲜好物表详情
func (l *QueryHomeNewProductDetailLogic) QueryHomeNewProductDetail(in *smsclient.QueryHomeNewProductDetailReq) (*smsclient.QueryHomeNewProductDetailResp, error) {
	item, err := query.SmsHomeNewProduct.WithContext(l.ctx).Where(query.SmsHomeNewProduct.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询新鲜好物表详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询新鲜好物表详情失败")
	}

	data := &smsclient.QueryHomeNewProductDetailResp{
		Id:              item.ID,              //
		ProductId:       item.ProductID,       // 商品id
		ProductName:     item.ProductName,     // 商品名称
		RecommendStatus: item.RecommendStatus, // 推荐状态：0->不推荐;1->推荐
		Sort:            item.Sort,            // 排序

	}

	logc.Infof(l.ctx, "查询新鲜好物表详情,参数：%+v,响应：%+v", in, data)
	return data, nil
}
