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

type ProductAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductAddLogic {
	return ProductAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductAdd 添加商品
//逻辑步骤：
//1.创建商品基本信息
//2.会员价格
//3.阶梯价格
//4.满减价格
//5.添加sku库存信息
//6.添加商品参数,添加自定义商品规格
//7.关联专题
//8.关联优选
//注意: 步骤1到6是在商品(rpc)中操作,步骤7到8是在cms模块中操作
func (l *ProductAddLogic) ProductAdd(req types.AddProductReq) (*types.AddProductResp, error) {
	//添加商品成功后返回商品id,然后关联专题和关联优选(因为这二个表在cms模块中,需要通过rpc去调用)
	result, err := l.svcCtx.Pms.ProductAdd(l.ctx, &pmsclient.ProductAddReq{
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
		logx.WithContext(l.ctx).Errorf("添加商品信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加商品信息失败")
	}

	//商品id
	productId := result.ProductId
	//7.关联专题
	addSubjectProductRelation(req, l, productId)

	//8.关联优选
	addPrefrenceAreaProductRelation(req, l, productId)

	return &types.AddProductResp{
		Code:    "000000",
		Message: "添加商品信息成功",
	}, nil
}

// 添加专题关联
func addSubjectProductRelation(req types.AddProductReq, l *ProductAddLogic, productId int32) {
	for _, item := range req.SubjectProductRelationList {
		_, _ = l.svcCtx.Cms.SubjectProductRelationAdd(l.ctx, &cmsclient.SubjectProductRelationAddReq{
			SubjectId: item.SubjectID,
			ProductId: int64(productId),
		})
	}
}

// 添加优选商品关联
func addPrefrenceAreaProductRelation(req types.AddProductReq, l *ProductAddLogic, productId int32) {
	for _, item := range req.PrefrenceAreaProductRelationList {
		_, _ = l.svcCtx.Cms.PrefrenceAreaProductRelationAdd(l.ctx, &cmsclient.PrefrenceAreaProductRelationAddReq{
			PrefrenceAreaId: item.PrefrenceAreaID,
			ProductId:       int64(productId),
		})
	}
}
