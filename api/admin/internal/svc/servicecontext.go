package svc

import (
	"github.com/feihua/zero-admin/api/admin/internal/config"
	"github.com/feihua/zero-admin/api/admin/internal/middleware"
	"github.com/feihua/zero-admin/rpc/cms/client/preferredareaproductrelationservice"
	"github.com/feihua/zero-admin/rpc/cms/client/preferredareaservice"
	"github.com/feihua/zero-admin/rpc/cms/client/subjectcategoryservice"
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
	"github.com/feihua/zero-admin/rpc/pms/client/brandservice"
	"github.com/feihua/zero-admin/rpc/pms/client/commentreplayservice"
	"github.com/feihua/zero-admin/rpc/pms/client/commentservice"
	"github.com/feihua/zero-admin/rpc/pms/client/feighttemplateservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productattributecategoryservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productattributeservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productattributevalueservice"
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
	"github.com/feihua/zero-admin/rpc/sys/client/deptservice"
	"github.com/feihua/zero-admin/rpc/sys/client/dictitemservice"
	"github.com/feihua/zero-admin/rpc/sys/client/dicttypeservice"
	"github.com/feihua/zero-admin/rpc/sys/client/loginlogservice"
	"github.com/feihua/zero-admin/rpc/sys/client/menuservice"
	"github.com/feihua/zero-admin/rpc/sys/client/operatelogservice"
	"github.com/feihua/zero-admin/rpc/sys/client/postservice"
	"github.com/feihua/zero-admin/rpc/sys/client/roleservice"
	"github.com/feihua/zero-admin/rpc/sys/client/userservice"
	"github.com/feihua/zero-admin/rpc/ums/client/growthchangehistoryservice"
	"github.com/feihua/zero-admin/rpc/ums/client/integrationchangehistoryservice"
	"github.com/feihua/zero-admin/rpc/ums/client/integrationconsumesettingservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberaddressservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberinfoservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberlevelservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberloginlogservice"
	"github.com/feihua/zero-admin/rpc/ums/client/membermembertagrelationservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberproductcategoryrelationservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberproductcollectionservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberreadhistoryservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberrulesettingservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberstatisticsinfoservice"
	"github.com/feihua/zero-admin/rpc/ums/client/membertagservice"
	"github.com/feihua/zero-admin/rpc/ums/client/membertaskservice"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	CheckUrl rest.Middleware
	AddLog   rest.Middleware
	// 会员相关
	GrowthChangeHistoryService           growthchangehistoryservice.GrowthChangeHistoryService
	IntegrationChangeHistoryService      integrationchangehistoryservice.IntegrationChangeHistoryService
	IntegrationConsumeSettingService     integrationconsumesettingservice.IntegrationConsumeSettingService
	MemberLevelService                   memberlevelservice.MemberLevelService
	MemberLoginLogService                memberloginlogservice.MemberLoginLogService
	MemberMemberTagRelationService       membermembertagrelationservice.MemberMemberTagRelationService
	MemberProductCategoryRelationService memberproductcategoryrelationservice.MemberProductCategoryRelationService
	MemberProductCollectionService       memberproductcollectionservice.MemberProductCollectionService
	MemberReadHistoryService             memberreadhistoryservice.MemberReadHistoryService
	MemberAddressService                 memberaddressservice.MemberAddressService
	MemberRuleSettingService             memberrulesettingservice.MemberRuleSettingService
	MemberInfoService                    memberinfoservice.MemberInfoService
	MemberStatisticsInfoService          memberstatisticsinfoservice.MemberStatisticsInfoService
	MemberTagService                     membertagservice.MemberTagService
	MemberTaskService                    membertaskservice.MemberTaskService
	// 系统相关
	DeptService       deptservice.DeptService
	DictTypeService   dicttypeservice.DictTypeService
	DictItemService   dictitemservice.DictItemService
	PostService       postservice.PostService
	LoginLogService   loginlogservice.LoginLogService
	Operatelogservice operatelogservice.OperateLogService
	MenuService       menuservice.MenuService
	RoleService       roleservice.RoleService
	UserService       userservice.UserService
	// 商品相关
	BrandService                            brandservice.BrandService
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
	SubjectCategoryService              subjectcategoryservice.SubjectCategoryService
	Redis                               *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	newRedis := redis.New(c.Redis.Address, redisConfig(c))
	umsClient := zrpc.MustNewClient(c.UmsRpc)
	sysClient := zrpc.MustNewClient(c.SysRpc)
	pmsClient := zrpc.MustNewClient(c.PmsRpc)
	omsClient := zrpc.MustNewClient(c.OmsRpc)
	smsClient := zrpc.MustNewClient(c.SmsRpc)
	cmsClient := zrpc.MustNewClient(c.CmsRpc)
	operateLogService := operatelogservice.NewOperateLogService(sysClient)
	return &ServiceContext{
		Config:                               c,
		GrowthChangeHistoryService:           growthchangehistoryservice.NewGrowthChangeHistoryService(umsClient),
		IntegrationChangeHistoryService:      integrationchangehistoryservice.NewIntegrationChangeHistoryService(umsClient),
		IntegrationConsumeSettingService:     integrationconsumesettingservice.NewIntegrationConsumeSettingService(umsClient),
		MemberLevelService:                   memberlevelservice.NewMemberLevelService(umsClient),
		MemberLoginLogService:                memberloginlogservice.NewMemberLoginLogService(umsClient),
		MemberMemberTagRelationService:       membermembertagrelationservice.NewMemberMemberTagRelationService(umsClient),
		MemberProductCategoryRelationService: memberproductcategoryrelationservice.NewMemberProductCategoryRelationService(umsClient),
		MemberProductCollectionService:       memberproductcollectionservice.NewMemberProductCollectionService(umsClient),
		MemberReadHistoryService:             memberreadhistoryservice.NewMemberReadHistoryService(umsClient),
		MemberAddressService:                 memberaddressservice.NewMemberAddressService(umsClient),
		MemberRuleSettingService:             memberrulesettingservice.NewMemberRuleSettingService(umsClient),
		MemberInfoService:                    memberinfoservice.NewMemberInfoService(umsClient),
		MemberStatisticsInfoService:          memberstatisticsinfoservice.NewMemberStatisticsInfoService(umsClient),
		MemberTagService:                     membertagservice.NewMemberTagService(umsClient),
		MemberTaskService:                    membertaskservice.NewMemberTaskService(umsClient),

		DeptService:       deptservice.NewDeptService(sysClient),
		DictTypeService:   dicttypeservice.NewDictTypeService(sysClient),
		DictItemService:   dictitemservice.NewDictItemService(sysClient),
		PostService:       postservice.NewPostService(sysClient),
		LoginLogService:   loginlogservice.NewLoginLogService(sysClient),
		Operatelogservice: operateLogService,
		MenuService:       menuservice.NewMenuService(sysClient),
		RoleService:       roleservice.NewRoleService(sysClient),
		UserService:       userservice.NewUserService(sysClient),
		CheckUrl:          middleware.NewCheckUrlMiddleware(newRedis, c.Auth.ExcludeUrl).Handle,
		AddLog:            middleware.NewAddLogMiddleware(operateLogService).Handle,

		BrandService:                            brandservice.NewBrandService(pmsClient),
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
		SubjectCategoryService:              subjectcategoryservice.NewSubjectCategoryService(cmsClient),
		Redis:                               newRedis,
	}
}

func redisConfig(c config.Config) redis.Option {
	return func(r *redis.Redis) {
		r.Type = redis.NodeType
		r.Pass = c.Redis.Pass
	}
}
