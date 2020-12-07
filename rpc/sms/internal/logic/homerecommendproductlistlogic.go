package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
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
	all, _ := l.svcCtx.SmsHomeRecommendProductModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

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

	fmt.Println(list)
	return &sms.HomeRecommendProductListResp{
		Total: 10,
		List:  list,
	}, nil
}
