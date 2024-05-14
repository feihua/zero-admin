package homenewproductservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeNewProductListLogic 首页新品
/*
Author: LiuFeiHua
Date: 2024/5/14 9:28
*/
type HomeNewProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeNewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeNewProductListLogic {
	return &HomeNewProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// HomeNewProductList 查询首页新品
func (l *HomeNewProductListLogic) HomeNewProductList(in *smsclient.HomeNewProductListReq) (*smsclient.HomeNewProductListResp, error) {
	q := query.SmsHomeNewProduct.WithContext(l.ctx)
	if len(in.ProductName) > 0 {
		q = q.Where(query.SmsHomeNewProduct.ProductName.Like("%" + in.ProductName + "%"))
	}

	if in.RecommendStatus != 2 {
		q = q.Where(query.SmsHomeNewProduct.RecommendStatus.Eq(in.RecommendStatus))
	}

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询首页新鲜好物列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*smsclient.HomeNewProductListData
	for _, item := range result {

		list = append(list, &smsclient.HomeNewProductListData{
			Id:              item.ID,
			ProductId:       item.ProductID,
			ProductName:     item.ProductName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
		})
	}

	logc.Infof(l.ctx, "查询首页新鲜好物列表信息,参数：%+v,响应：%+v", in, list)
	return &smsclient.HomeNewProductListResp{
		Total: count,
		List:  list,
	}, nil
}
