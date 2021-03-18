package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

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

func (l *BrandListLogic) BrandList(in *pms.BrandListReq) (*pms.BrandListResp, error) {
	all, _ := l.svcCtx.PmsBrandModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.PmsBrandModel.Count()

	var list []*pms.BrandListData
	for _, item := range *all {

		list = append(list, &pms.BrandListData{
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

	fmt.Println(list)
	return &pms.BrandListResp{
		Total: count,
		List:  list,
	}, nil
}
