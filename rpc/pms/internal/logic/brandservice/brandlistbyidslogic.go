package brandservicelogic

import (
	"context"
	"encoding/json"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *BrandListByIdsLogic) BrandListByIds(in *pmsclient.BrandListByIdsReq) (*pmsclient.BrandListResp, error) {
	all, err := l.svcCtx.PmsBrandModel.FindAllByIds(l.ctx, in.Ids)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询商品品牌列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*pmsclient.BrandListData
	for _, item := range *all {

		list = append(list, &pmsclient.BrandListData{
			Id:                  item.Id,
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

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询商品品牌列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &pmsclient.BrandListResp{
		Total: 0,
		List:  list,
	}, nil
}
