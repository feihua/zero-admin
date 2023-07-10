package homenewproductservicelogic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *HomeNewProductListLogic) HomeNewProductList(in *smsclient.HomeNewProductListReq) (*smsclient.HomeNewProductListResp, error) {
	all, err := l.svcCtx.SmsHomeNewProductModel.FindAll(l.ctx, in)
	count, _ := l.svcCtx.SmsHomeNewProductModel.Count(l.ctx, in)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询首页新鲜好物列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*smsclient.HomeNewProductListData
	for _, item := range *all {

		list = append(list, &smsclient.HomeNewProductListData{
			Id:              item.Id,
			ProductId:       item.ProductId,
			ProductName:     item.ProductName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询首页新鲜好物列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &smsclient.HomeNewProductListResp{
		Total: count,
		List:  list,
	}, nil
}
