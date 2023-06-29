package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/cms/cmsclient"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductUpdateLogic {
	return ProductUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductUpdate 更新商品
//逻辑步骤：
//1.更新商品基本信息
//2.会员价格
//3.阶梯价格
//4.满减价格
//5.更新sku库存信息
//6.更新商品参数,添加自定义商品规格
//7.关联专题
//8.关联优选
//注意: 步骤1到6是在商品(rpc)中操作,步骤7到8是在cms模块中操作
func (l *ProductUpdateLogic) ProductUpdate(req types.UpdateProductReq) (*types.UpdateProductResp, error) {
	productId := req.Id
	_, err := l.svcCtx.Pms.ProductUpdate(l.ctx, &pmsclient.ProductUpdateReq{
		Id:                         productId,
		BrandId:                    req.BrandId,
		ProductCategoryId:          req.ProductCategoryId,
		FeightTemplateId:           req.FeightTemplateId,
		ProductAttributeCategoryId: req.ProductAttributeCategoryId,
		Name:                       req.Name,
		Pic:                        req.Pic,
		ProductSn:                  req.ProductSn,
		DeleteStatus:               req.DeleteStatus,
		PublishStatus:              req.PublishStatus,
		NewStatus:                  req.NewStatus,
		RecommandStatus:            req.RecommandStatus,
		VerifyStatus:               req.VerifyStatus,
		Sort:                       req.Sort,
		Sale:                       req.Sale,
		Price:                      req.Price,
		PromotionPrice:             req.PromotionPrice,
		GiftGrowth:                 req.GiftGrowth,
		GiftPoint:                  req.GiftPoint,
		UsePointLimit:              req.UsePointLimit,
		SubTitle:                   req.SubTitle,
		Description:                req.Description,
		OriginalPrice:              req.OriginalPrice,
		Stock:                      req.Stock,
		LowStock:                   req.LowStock,
		Unit:                       req.Unit,
		Weight:                     req.Weight,
		PreviewStatus:              req.PreviewStatus,
		ServiceIds:                 req.ServiceIds,
		Keywords:                   req.Keywords,
		Note:                       req.Note,
		AlbumPics:                  req.AlbumPics,
		DetailTitle:                req.DetailTitle,
		DetailDesc:                 req.DetailDesc,
		DetailHtml:                 req.DetailHtml,
		DetailMobileHtml:           req.DetailMobileHtml,
		PromotionStartTime:         req.PromotionStartTime,
		PromotionEndTime:           req.PromotionEndTime,
		PromotionPerLimit:          req.PromotionPerLimit,
		PromotionType:              req.PromotionType,
		BrandName:                  req.BrandName,
		ProductCategoryName:        req.ProductCategoryName,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新商品信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新商品信息失败")
	}

	//7.关联专题
	updateSubjectProductRelation(req, l, productId)

	//8.关联优选
	updatePrefrenceAreaProductRelation(req, l, productId)

	return &types.UpdateProductResp{
		Code:    "000000",
		Message: "更新商品信息成功",
	}, nil
}

// 更新专题关联
func updateSubjectProductRelation(req types.UpdateProductReq, l *ProductUpdateLogic, productId int64) {
	//1.先删除专题的关联
	_, _ = l.svcCtx.Cms.SubjectProductRelationDelete(l.ctx, &cmsclient.SubjectProductRelationDeleteReq{Id: productId})

	//2.重新添加专题的关联
	for _, item := range req.SubjectProductRelationList {
		_, _ = l.svcCtx.Cms.SubjectProductRelationAdd(l.ctx, &cmsclient.SubjectProductRelationAddReq{
			SubjectId: item.SubjectID,
			ProductId: productId,
		})
	}
}

// 更新优选商品关联
func updatePrefrenceAreaProductRelation(req types.UpdateProductReq, l *ProductUpdateLogic, productId int64) {
	//1.先删除优选商品的关联
	_, _ = l.svcCtx.Cms.PrefrenceAreaProductRelationDelete(l.ctx, &cmsclient.PrefrenceAreaProductRelationDeleteReq{Id: productId})

	//2.重新添加优选商品的关联
	for _, item := range req.PrefrenceAreaProductRelationList {
		_, _ = l.svcCtx.Cms.PrefrenceAreaProductRelationAdd(l.ctx, &cmsclient.PrefrenceAreaProductRelationAddReq{
			PrefrenceAreaId: item.PrefrenceAreaID,
			ProductId:       productId,
		})
	}
}
