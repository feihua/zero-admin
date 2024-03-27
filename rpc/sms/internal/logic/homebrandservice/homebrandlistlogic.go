package homebrandservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeBrandListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeBrandListLogic {
	return &HomeBrandListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeBrandListLogic) HomeBrandList(in *smsclient.HomeBrandListReq) (*smsclient.HomeBrandListResp, error) {
	all, err := l.svcCtx.SmsHomeBrandModel.FindAll(l.ctx, in)
	count, _ := l.svcCtx.SmsHomeBrandModel.Count(l.ctx, in)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询首页品牌列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*smsclient.HomeBrandListData
	for _, item := range *all {

		list = append(list, &smsclient.HomeBrandListData{
			Id:              item.Id,
			BrandId:         item.BrandId,
			BrandName:       item.BrandName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询首页品牌列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &smsclient.HomeBrandListResp{
		Total: count,
		List:  list,
	}, nil
}
