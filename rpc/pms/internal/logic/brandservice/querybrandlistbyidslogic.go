package brandservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryBrandListByIdsLogic 根据ids查询品牌信息
/*
Author: LiuFeiHua
Date: 2024/6/12 16:29
*/
type QueryBrandListByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryBrandListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryBrandListByIdsLogic {
	return &QueryBrandListByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryBrandListByIds 根据ids查询品牌信息
func (l *QueryBrandListByIdsLogic) QueryBrandListByIds(in *pmsclient.QueryBrandListByIdsReq) (*pmsclient.QueryBrandListResp, error) {
	result, err := query.PmsBrand.WithContext(l.ctx).Where(query.PmsBrand.ID.In(in.Ids...)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询商品品牌列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.BrandListData
	for _, item := range result {

		list = append(list, &pmsclient.BrandListData{
			Id:                  item.ID,
			Name:                item.Name,
			FirstLetter:         item.FirstLetter,
			Sort:                item.Sort,
			FactoryStatus:       item.FactoryStatus,
			ShowStatus:          item.ShowStatus,
			ProductCount:        item.ProductCount,
			ProductCommentCount: item.ProductCommentCount,
			Logo:                item.Logo,
			BigPic:              item.BigPic,
			BrandStory:          item.BrandStory,
		})
	}

	return &pmsclient.QueryBrandListResp{
		Total: 0,
		List:  list,
	}, nil
}
