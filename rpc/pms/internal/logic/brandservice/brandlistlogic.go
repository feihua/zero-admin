package brandservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// BrandListLogic
/*
Author: LiuFeiHua
Date: 2024/4/7 17:37
*/
type BrandListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandListLogic {
	return &BrandListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// BrandList 查询品牌列表
func (l *BrandListLogic) BrandList(in *pmsclient.BrandListReq) (*pmsclient.BrandListResp, error) {
	q := query.PmsBrand.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(query.PmsBrand.Name.Like("%" + in.Name + "%"))
	}
	if in.FactoryStatus != 2 {
		q = q.Where(query.PmsBrand.FactoryStatus.Eq(in.FactoryStatus))
	}
	if in.ShowStatus != 2 {
		q = q.Where(query.PmsBrand.ShowStatus.Eq(in.ShowStatus))
	}

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		in, _ := json.Marshal(in)
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
		Total: count,
		List:  list,
	}, nil
}
