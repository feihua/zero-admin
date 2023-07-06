package logic

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"
	"zero-admin/rpc/cms/cmsclient"
	"zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductListLogic {
	return ProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductListLogic) ProductList(req types.ListProductReq) (*types.ListProductResp, error) {
	resp, err := l.svcCtx.Pms.ProductList(l.ctx, &pmsclient.ProductListReq{
		Current:           req.Current,
		PageSize:          req.PageSize,
		Name:              req.Name,
		VerifyStatus:      req.VerifyStatus,
		ProductCategoryId: req.ProductCategoryId,
		PublishStatus:     req.PublishStatus,
		DeleteStatus:      req.DeleteStatus,
		BrandId:           req.BrandId,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询商品信息列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询商品信息失败")
	}

	var list []*types.ListtProductData

	for _, item := range resp.List {
		var productCategoryIdArray []int64
		for _, s := range strings.Split(item.ProductCategoryIdArray, ",") {
			id, _ := strconv.ParseInt(s, 10, 64)
			productCategoryIdArray = append(productCategoryIdArray, id)
		}

		//查询专题关联
		subjectRelationList, _ := l.svcCtx.Cms.SubjectProductRelationList(l.ctx, &cmsclient.SubjectProductRelationListReq{
			ProductId: item.Id,
		})

		//查询优选关联
		areaRelationList, _ := l.svcCtx.Cms.PrefrenceAreaProductRelationList(l.ctx, &cmsclient.PrefrenceAreaProductRelationListReq{
			ProductId: item.Id,
		})

		list = append(list, &types.ListtProductData{
			Id:                               item.Id,
			BrandId:                          item.BrandId,
			ProductCategoryId:                item.ProductCategoryId,
			ProductCategoryIdArray:           productCategoryIdArray,
			FeightTemplateId:                 item.FeightTemplateId,
			ProductAttributeCategoryId:       item.ProductAttributeCategoryId,
			Name:                             item.Name,
			Pic:                              item.Pic,
			ProductSn:                        item.ProductSn,
			DeleteStatus:                     item.DeleteStatus,
			PublishStatus:                    item.PublishStatus,
			NewStatus:                        item.NewStatus,
			RecommandStatus:                  item.RecommandStatus,
			VerifyStatus:                     item.VerifyStatus,
			Sort:                             item.Sort,
			Sale:                             item.Sale,
			Price:                            item.Price,
			PromotionPrice:                   item.PromotionPrice,
			GiftGrowth:                       item.GiftGrowth,
			GiftPoint:                        item.GiftPoint,
			UsePointLimit:                    item.UsePointLimit,
			SubTitle:                         item.SubTitle,
			Description:                      item.Description,
			OriginalPrice:                    item.OriginalPrice,
			Stock:                            item.Stock,
			LowStock:                         item.LowStock,
			Unit:                             item.Unit,
			Weight:                           item.Weight,
			PreviewStatus:                    item.PreviewStatus,
			ServiceIds:                       item.ServiceIds,
			Keywords:                         item.Keywords,
			Note:                             item.Note,
			AlbumPics:                        item.AlbumPics,
			DetailTitle:                      item.DetailTitle,
			DetailDesc:                       item.DetailDesc,
			DetailHtml:                       item.DetailHtml,
			DetailMobileHtml:                 item.DetailMobileHtml,
			PromotionStartTime:               item.PromotionStartTime,
			PromotionEndTime:                 item.PromotionEndTime,
			PromotionPerLimit:                item.PromotionPerLimit,
			PromotionType:                    item.PromotionType,
			BrandName:                        item.BrandName,
			ProductCategoryName:              item.ProductCategoryName,
			SubjectProductRelationList:       subjectRelationList.SubjectId,
			PrefrenceAreaProductRelationList: areaRelationList.PrefrenceAreaId,
		})
	}

	return &types.ListProductResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询商品信息成功",
	}, nil
}
