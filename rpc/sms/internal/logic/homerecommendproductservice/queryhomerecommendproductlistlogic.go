package homerecommendproductservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryHomeRecommendProductListLogic 查询人气推荐商品表列表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:56
*/
type QueryHomeRecommendProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHomeRecommendProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHomeRecommendProductListLogic {
	return &QueryHomeRecommendProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryHomeRecommendProductList 查询人气推荐商品表列表
func (l *QueryHomeRecommendProductListLogic) QueryHomeRecommendProductList(in *smsclient.QueryHomeRecommendProductListReq) (*smsclient.QueryHomeRecommendProductListResp, error) {
	q := query.SmsHomeRecommendProduct.WithContext(l.ctx)
	if len(in.ProductName) > 0 {
		q = q.Where(query.SmsHomeRecommendProduct.ProductName.Like("%" + in.ProductName + "%"))
	}

	if in.RecommendStatus != 2 {
		q = q.Where(query.SmsHomeRecommendProduct.RecommendStatus.Eq(in.RecommendStatus))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询人气商品推荐列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*smsclient.HomeRecommendProductListData
	for _, item := range result {

		list = append(list, &smsclient.HomeRecommendProductListData{
			Id:              item.ID,
			ProductId:       item.ProductID,
			ProductName:     item.ProductName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
		})
	}

	return &smsclient.QueryHomeRecommendProductListResp{
		Total: count,
		List:  list,
	}, nil

}
