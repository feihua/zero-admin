package homebrandservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryHomeBrandListLogic 查询首页推荐品牌表列表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:53
*/
type QueryHomeBrandListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHomeBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHomeBrandListLogic {
	return &QueryHomeBrandListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryHomeBrandList 查询首页推荐品牌表列表
func (l *QueryHomeBrandListLogic) QueryHomeBrandList(in *smsclient.QueryHomeBrandListReq) (*smsclient.QueryHomeBrandListResp, error) {
	q := query.SmsHomeBrand.WithContext(l.ctx)
	if len(in.BrandName) > 0 {
		q = q.Where(query.SmsHomeBrand.BrandName.Like("%" + in.BrandName + "%"))
	}

	if in.RecommendStatus != 2 {
		q = q.Where(query.SmsHomeBrand.RecommendStatus.Eq(in.RecommendStatus))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询首页品牌列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*smsclient.HomeBrandListData
	for _, item := range result {

		list = append(list, &smsclient.HomeBrandListData{
			Id:              item.ID,
			BrandId:         item.BrandID,
			BrandName:       item.BrandName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
		})
	}

	return &smsclient.QueryHomeBrandListResp{
		Total: count,
		List:  list,
	}, nil

}
