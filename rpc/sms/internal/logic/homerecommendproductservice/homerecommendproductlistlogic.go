package homerecommendproductservicelogic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/smsclient"
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
	all, err := l.svcCtx.SmsHomeRecommendProductModel.FindAll(l.ctx, in)
	count, _ := l.svcCtx.SmsHomeRecommendProductModel.Count(l.ctx, in)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询人气商品推荐列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*smsclient.HomeRecommendProductListData
	for _, item := range *all {

		list = append(list, &smsclient.HomeRecommendProductListData{
			Id:              item.Id,
			ProductId:       item.ProductId,
			ProductName:     item.ProductName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询人气商品推荐列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &smsclient.HomeRecommendProductListResp{
		Total: count,
		List:  list,
	}, nil
}
