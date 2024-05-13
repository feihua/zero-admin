package homebrandservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeBrandListLogic 首页品牌信息
/*
Author: LiuFeiHua
Date: 2024/5/13 15:58
*/
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

// HomeBrandList 查询首页品牌信息
func (l *HomeBrandListLogic) HomeBrandList(in *smsclient.HomeBrandListReq) (*smsclient.HomeBrandListResp, error) {
	q := query.SmsHomeBrand.WithContext(l.ctx)
	if len(in.BrandName) > 0 {
		q = q.Where(query.SmsHomeBrand.BrandName.Like("%" + in.BrandName + "%"))
	}

	if in.RecommendStatus != 2 {
		q = q.Where(query.SmsHomeBrand.RecommendStatus.Eq(in.RecommendStatus))
	}

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		in, _ := json.Marshal(in)
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

	logc.Infof(l.ctx, "查询首页品牌列表信息,参数：%+v,响应：%+v", in, list)
	return &smsclient.HomeBrandListResp{
		Total: count,
		List:  list,
	}, nil
}
