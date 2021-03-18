package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

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

func (l *HomeNewProductListLogic) HomeNewProductList(in *sms.HomeNewProductListReq) (*sms.HomeNewProductListResp, error) {
	all, _ := l.svcCtx.SmsHomeNewProductModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.SmsHomeNewProductModel.Count()

	var list []*sms.HomeNewProductListData
	for _, item := range *all {

		list = append(list, &sms.HomeNewProductListData{
			Id:              item.Id,
			ProductId:       item.ProductId,
			ProductName:     item.ProductName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
		})
	}

	fmt.Println(list)
	return &sms.HomeNewProductListResp{
		Total: count,
		List:  list,
	}, nil
}
