package brandservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// BrandListByIdsLogic 商品品牌
/*
Author: LiuFeiHua
Date: 2024/5/13 16:31
*/
type BrandListByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBrandListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandListByIdsLogic {
	return &BrandListByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// BrandListByIds 根据id列表查询品牌列表
func (l *BrandListByIdsLogic) BrandListByIds(in *pmsclient.BrandListByIdsReq) (*pmsclient.BrandListResp, error) {
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

	logc.Infof(l.ctx, "查询商品品牌列表信息,参数：%+v,响应：%+v", in, list)
	return &pmsclient.BrandListResp{
		Total: 0,
		List:  list,
	}, nil
}
