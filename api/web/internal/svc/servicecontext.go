package svc

import (
	"github.com/feihua/zero-admin/api/web/internal/config"
	"github.com/feihua/zero-admin/rpc/cms/client/preferredareaproductrelationservice"
	"github.com/feihua/zero-admin/rpc/cms/client/preferredareaservice"
	"github.com/feihua/zero-admin/rpc/cms/client/subjectproductrelationservice"
	"github.com/feihua/zero-admin/rpc/cms/client/subjectservice"
	"github.com/feihua/zero-admin/rpc/oms/client/cartitemservice"
	"github.com/feihua/zero-admin/rpc/oms/client/companyaddressservice"
	"github.com/feihua/zero-admin/rpc/oms/client/orderitemservice"
	"github.com/feihua/zero-admin/rpc/oms/client/orderoperatehistoryservice"
	"github.com/feihua/zero-admin/rpc/oms/client/orderreturnapplyservice"
	"github.com/feihua/zero-admin/rpc/oms/client/orderreturnreasonservice"
	"github.com/feihua/zero-admin/rpc/oms/client/orderservice"
	"github.com/feihua/zero-admin/rpc/oms/client/ordersettingservice"
	"github.com/feihua/zero-admin/rpc/pms/client/commentreplayservice"
	"github.com/feihua/zero-admin/rpc/pms/client/commentservice"
	"github.com/feihua/zero-admin/rpc/pms/client/feighttemplateservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productattributecategoryservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productattributeservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productattributevalueservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productbrandservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productcategoryattributerelationservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productcategoryservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productfullreductionservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productladderservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productoperatelogservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productvertifyrecordservice"
	"github.com/feihua/zero-admin/rpc/pms/client/skustockservice"
	"github.com/feihua/zero-admin/rpc/sms/client/couponhistoryservice"
	"github.com/feihua/zero-admin/rpc/sms/client/couponservice"
	"github.com/feihua/zero-admin/rpc/sms/client/flashpromotionlogservice"
	"github.com/feihua/zero-admin/rpc/sms/client/flashpromotionproductrelationservice"
	"github.com/feihua/zero-admin/rpc/sms/client/flashpromotionservice"
	"github.com/feihua/zero-admin/rpc/sms/client/flashpromotionsessionservice"
	"github.com/feihua/zero-admin/rpc/sms/client/homeadvertiseservice"
	"github.com/feihua/zero-admin/rpc/sms/client/homebrandservice"
	"github.com/feihua/zero-admin/rpc/sms/client/homenewproductservice"
	"github.com/feihua/zero-admin/rpc/sms/client/homerecommendproductservice"
	"github.com/feihua/zero-admin/rpc/sms/client/homerecommendsubjectservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberaddressservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberbrandattentionservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberconsumesettingservice"
	"github.com/feihua/zero-admin/rpc/ums/client/membergrowthlogservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberinfoservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberlevelservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberloginlogservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberpointslogservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberproductcategoryrelationservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberproductcollectionservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberreadhistoryservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberrulesettingservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberstatisticsinfoservice"
	"github.com/feihua/zero-admin/rpc/ums/client/membertagrelationservice"
	"github.com/feihua/zero-admin/rpc/ums/client/membertagservice"
	"github.com/feihua/zero-admin/rpc/ums/client/membertaskrelationservice"
	"github.com/feihua/zero-admin/rpc/ums/client/membertaskservice"
	"github.com/smartwalle/alipay/v3"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	// 会员相关
	MemberGrowthLogService               membergrowthlogservice.MemberGrowthLogService
	MemberPointsLogService               memberpointslogservice.MemberPointsLogService
	MemberConsumeSettingService          memberconsumesettingservice.MemberConsumeSettingService
	MemberLevelService                   memberlevelservice.MemberLevelService
	MemberLoginLogService                memberloginlogservice.MemberLoginLogService
	MemberProductCategoryRelationService memberproductcategoryrelationservice.MemberProductCategoryRelationService
	MemberProductCollectionService       memberproductcollectionservice.MemberProductCollectionService
	MemberReadHistoryService             memberreadhistoryservice.MemberReadHistoryService
	MemberAddressService                 memberaddressservice.MemberAddressService
	MemberRuleSettingService             memberrulesettingservice.MemberRuleSettingService
	MemberService                        memberinfoservice.MemberInfoService
	MemberStatisticsInfoService          memberstatisticsinfoservice.MemberStatisticsInfoService
	MemberTagService                     membertagservice.MemberTagService
	MemberTagRelationService             membertagrelationservice.MemberTagRelationService
	MemberTaskService                    membertaskservice.MemberTaskService
	MemberTaskRelationService            membertaskrelationservice.MemberTaskRelationService
	MemberBrandAttentionService          memberbrandattentionservice.MemberBrandAttentionService

	// 商品相关
	ProductBrandService                     productbrandservice.ProductBrandService
	CommentReplayService                    commentreplayservice.CommentReplayService
	CommentService                          commentservice.CommentService
	FeightTemplateService                   feighttemplateservice.FeightTemplateService
	ProductAttributeCategoryService         productattributecategoryservice.ProductAttributeCategoryService
	ProductAttributeService                 productattributeservice.ProductAttributeService
	ProductAttributeValueService            productattributevalueservice.ProductAttributeValueService
	ProductCategoryAttributeRelationService productcategoryattributerelationservice.ProductCategoryAttributeRelationService
	ProductCategoryService                  productcategoryservice.ProductCategoryService
	ProductFullReductionService             productfullreductionservice.ProductFullReductionService
	ProductLadderService                    productladderservice.ProductLadderService
	ProductOperateLogService                productoperatelogservice.ProductOperateLogService
	ProductService                          productservice.ProductService
	ProductVertifyRecordService             productvertifyrecordservice.ProductVertifyRecordService
	SkuStockService                         skustockservice.SkuStockService
	// 订单相关
	CartItemService            cartitemservice.CartItemService
	CompanyAddressService      companyaddressservice.CompanyAddressService
	OrderItemService           orderitemservice.OrderItemService
	OrderOperateHistoryService orderoperatehistoryservice.OrderOperateHistoryService
	OrderReturnApplyService    orderreturnapplyservice.OrderReturnApplyService
	OrderReturnReasonService   orderreturnreasonservice.OrderReturnReasonService
	OrderService               orderservice.OrderService
	OrderSettingService        ordersettingservice.OrderSettingService
	// 营销相关
	CouponHistoryService                 couponhistoryservice.CouponHistoryService
	CouponService                        couponservice.CouponService
	FlashPromotionLogService             flashpromotionlogservice.FlashPromotionLogService
	FlashPromotionProductRelationService flashpromotionproductrelationservice.FlashPromotionProductRelationService
	FlashPromotionService                flashpromotionservice.FlashPromotionService
	FlashPromotionSessionService         flashpromotionsessionservice.FlashPromotionSessionService
	HomeAdvertiseService                 homeadvertiseservice.HomeAdvertiseService
	HomeBrandService                     homebrandservice.HomeBrandService
	HomeNewProductService                homenewproductservice.HomeNewProductService
	HomeRecommendProductService          homerecommendproductservice.HomeRecommendProductService
	HomeRecommendSubjectService          homerecommendsubjectservice.HomeRecommendSubjectService
	// 内容相关
	SubjectService                      subjectservice.SubjectService
	SubjectProductRelationService       subjectproductrelationservice.SubjectProductRelationService
	PreferredAreaService                preferredareaservice.PreferredAreaService
	PreferredAreaProductRelationService preferredareaproductrelationservice.PreferredAreaProductRelationService

	AlipayClient *alipay.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化支付宝客户端
	client, err := alipay.New(c.Alipay.AppId, c.Alipay.PrivateKey, c.Alipay.IsProduction)
	if err != nil {
		print("初始化支付宝失败")
	}
	umsClient := zrpc.MustNewClient(c.UmsRpc)
	pmsClient := zrpc.MustNewClient(c.PmsRpc)
	omsClient := zrpc.MustNewClient(c.OmsRpc)
	smsClient := zrpc.MustNewClient(c.SmsRpc)
	cmsClient := zrpc.MustNewClient(c.CmsRpc)
	return &ServiceContext{
		Config:                               c,
		MemberGrowthLogService:               membergrowthlogservice.NewMemberGrowthLogService(umsClient),
		MemberPointsLogService:               memberpointslogservice.NewMemberPointsLogService(umsClient),
		MemberConsumeSettingService:          memberconsumesettingservice.NewMemberConsumeSettingService(umsClient),
		MemberLevelService:                   memberlevelservice.NewMemberLevelService(umsClient),
		MemberLoginLogService:                memberloginlogservice.NewMemberLoginLogService(umsClient),
		MemberProductCategoryRelationService: memberproductcategoryrelationservice.NewMemberProductCategoryRelationService(umsClient),
		MemberProductCollectionService:       memberproductcollectionservice.NewMemberProductCollectionService(umsClient),
		MemberReadHistoryService:             memberreadhistoryservice.NewMemberReadHistoryService(umsClient),
		MemberAddressService:                 memberaddressservice.NewMemberAddressService(umsClient),
		MemberRuleSettingService:             memberrulesettingservice.NewMemberRuleSettingService(umsClient),
		MemberService:                        memberinfoservice.NewMemberInfoService(umsClient),
		MemberStatisticsInfoService:          memberstatisticsinfoservice.NewMemberStatisticsInfoService(umsClient),
		MemberTagService:                     membertagservice.NewMemberTagService(umsClient),
		MemberTagRelationService:             membertagrelationservice.NewMemberTagRelationService(umsClient),
		MemberTaskService:                    membertaskservice.NewMemberTaskService(umsClient),
		MemberTaskRelationService:            membertaskrelationservice.NewMemberTaskRelationService(umsClient),
		MemberBrandAttentionService:          memberbrandattentionservice.NewMemberBrandAttentionService(umsClient),

		ProductBrandService:                     productbrandservice.NewProductBrandService(pmsClient),
		CommentReplayService:                    commentreplayservice.NewCommentReplayService(pmsClient),
		CommentService:                          commentservice.NewCommentService(pmsClient),
		FeightTemplateService:                   feighttemplateservice.NewFeightTemplateService(pmsClient),
		ProductAttributeCategoryService:         productattributecategoryservice.NewProductAttributeCategoryService(pmsClient),
		ProductAttributeService:                 productattributeservice.NewProductAttributeService(pmsClient),
		ProductAttributeValueService:            productattributevalueservice.NewProductAttributeValueService(pmsClient),
		ProductCategoryAttributeRelationService: productcategoryattributerelationservice.NewProductCategoryAttributeRelationService(pmsClient),
		ProductCategoryService:                  productcategoryservice.NewProductCategoryService(pmsClient),
		ProductFullReductionService:             productfullreductionservice.NewProductFullReductionService(pmsClient),
		ProductLadderService:                    productladderservice.NewProductLadderService(pmsClient),
		ProductOperateLogService:                productoperatelogservice.NewProductOperateLogService(pmsClient),
		ProductService:                          productservice.NewProductService(pmsClient),
		ProductVertifyRecordService:             productvertifyrecordservice.NewProductVertifyRecordService(pmsClient),
		SkuStockService:                         skustockservice.NewSkuStockService(pmsClient),

		CartItemService:            cartitemservice.NewCartItemService(omsClient),
		CompanyAddressService:      companyaddressservice.NewCompanyAddressService(omsClient),
		OrderItemService:           orderitemservice.NewOrderItemService(omsClient),
		OrderOperateHistoryService: orderoperatehistoryservice.NewOrderOperateHistoryService(omsClient),
		OrderReturnApplyService:    orderreturnapplyservice.NewOrderReturnApplyService(omsClient),
		OrderReturnReasonService:   orderreturnreasonservice.NewOrderReturnReasonService(omsClient),
		OrderService:               orderservice.NewOrderService(omsClient),
		OrderSettingService:        ordersettingservice.NewOrderSettingService(omsClient),

		CouponHistoryService:                 couponhistoryservice.NewCouponHistoryService(smsClient),
		CouponService:                        couponservice.NewCouponService(smsClient),
		FlashPromotionLogService:             flashpromotionlogservice.NewFlashPromotionLogService(smsClient),
		FlashPromotionProductRelationService: flashpromotionproductrelationservice.NewFlashPromotionProductRelationService(smsClient),
		FlashPromotionService:                flashpromotionservice.NewFlashPromotionService(smsClient),
		FlashPromotionSessionService:         flashpromotionsessionservice.NewFlashPromotionSessionService(smsClient),
		HomeAdvertiseService:                 homeadvertiseservice.NewHomeAdvertiseService(smsClient),
		HomeBrandService:                     homebrandservice.NewHomeBrandService(smsClient),
		HomeNewProductService:                homenewproductservice.NewHomeNewProductService(smsClient),
		HomeRecommendProductService:          homerecommendproductservice.NewHomeRecommendProductService(smsClient),
		HomeRecommendSubjectService:          homerecommendsubjectservice.NewHomeRecommendSubjectService(smsClient),

		SubjectService:                      subjectservice.NewSubjectService(cmsClient),
		SubjectProductRelationService:       subjectproductrelationservice.NewSubjectProductRelationService(cmsClient),
		PreferredAreaService:                preferredareaservice.NewPreferredAreaService(cmsClient),
		PreferredAreaProductRelationService: preferredareaproductrelationservice.NewPreferredAreaProductRelationService(cmsClient),

		AlipayClient: client,
	}
}
