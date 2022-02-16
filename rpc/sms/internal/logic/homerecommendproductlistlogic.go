package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/sms"

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

func (l *HomeRecommendProductListLogic) HomeRecommendProductList(in *sms.HomeRecommendProductListReq) (*sms.HomeRecommendProductListResp, error) {
	all, err := l.svcCtx.SmsHomeRecommendProductModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.SmsHomeRecommendProductModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询人气商品推荐列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*sms.HomeRecommendProductListData
	for _, item := range *all {

		list = append(list, &sms.HomeRecommendProductListData{
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
	return &sms.HomeRecommendProductListResp{
		Total: count,
		List:  list,
	}, nil
}
