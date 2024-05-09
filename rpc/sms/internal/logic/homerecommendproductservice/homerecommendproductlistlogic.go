package homerecommendproductservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type HomeRecommendProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeRecommendProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendProductListLogic {
	return &HomeRecommendProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeRecommendProductListLogic) HomeRecommendProductList(in *smsclient.HomeRecommendProductListReq) (*smsclient.HomeRecommendProductListResp, error) {
	q := query.SmsHomeRecommendProduct.WithContext(l.ctx)
	if len(in.ProductName) > 0 {
		q = q.Where(query.SmsHomeRecommendProduct.ProductName.Like("%" + in.ProductName + "%"))
	}

	if in.RecommendStatus != 2 {
		q = q.Where(query.SmsHomeRecommendProduct.RecommendStatus.Eq(in.RecommendStatus))
	}

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

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

	logc.Infof(l.ctx, "查询人气商品推荐列表信息,参数：%+v,响应：%+v", in, list)
	return &smsclient.HomeRecommendProductListResp{
		Total: count,
		List:  list,
	}, nil
}
