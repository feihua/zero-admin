package svc

import (
	"github.com/feihua/zero-admin/api/front/internal/config"
	"github.com/feihua/zero-admin/rpc/cms/client/preferredareaproductrelationservice"
	"github.com/feihua/zero-admin/rpc/cms/client/preferredareaservice"
	"github.com/feihua/zero-admin/rpc/cms/client/subjectproductrelationservice"
	"github.com/feihua/zero-admin/rpc/cms/client/subjectservice"
	"github.com/feihua/zero-admin/rpc/oms/client/cartitemservice"
	"github.com/feihua/zero-admin/rpc/oms/client/companyaddressservice"
	"github.com/feihua/zero-admin/rpc/oms/client/orderitemservice"
	"github.com/feihua/zero-admin/rpc/oms/client/orderoperatehistorservice"
	"github.com/feihua/zero-admin/rpc/oms/client/orderreturnapplyservice"
	"github.com/feihua/zero-admin/rpc/oms/client/orderreturnreasonservice"
	"github.com/feihua/zero-admin/rpc/oms/client/orderservice"
	"github.com/feihua/zero-admin/rpc/oms/client/ordersettingservice"
	"github.com/feihua/zero-admin/rpc/pms/client/albumpicservice"
	"github.com/feihua/zero-admin/rpc/pms/client/albumservice"
	"github.com/feihua/zero-admin/rpc/pms/client/brandservice"
	"github.com/feihua/zero-admin/rpc/pms/client/commentreplayservice"
	"github.com/feihua/zero-admin/rpc/pms/client/commentservice"
	"github.com/feihua/zero-admin/rpc/pms/client/feighttemplateservice"
	"github.com/feihua/zero-admin/rpc/pms/client/memberpriceservice"
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
	"github.com/feihua/zero-admin/rpc/sms/client/couponproductcategoryrelationservice"
	"github.com/feihua/zero-admin/rpc/sms/client/couponproductrelationservice"
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
	"github.com/feihua/zero-admin/rpc/sys/client/dictservice"
	"github.com/feihua/zero-admin/rpc/sys/client/jobservice"
	"github.com/feihua/zero-admin/rpc/sys/client/loginlogservice"
	"github.com/feihua/zero-admin/rpc/sys/client/menuservice"
	"github.com/feihua/zero-admin/rpc/sys/client/roleservice"
	"github.com/feihua/zero-admin/rpc/sys/client/syslogservice"
	"github.com/feihua/zero-admin/rpc/sys/client/userservice"
	"github.com/feihua/zero-admin/rpc/ums/client/growthchangehistoryservice"
	"github.com/feihua/zero-admin/rpc/ums/client/integrationchangehistoryservice"
	"github.com/feihua/zero-admin/rpc/ums/client/integrationconsumesettingservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberattentionservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberlevelservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberloginlogservice"
	"github.com/feihua/zero-admin/rpc/ums/client/membermembertagrelationservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberproductcategoryrelationservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberproductcollectionservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberreadhistoryservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberreceiveaddressservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberrulesettingservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberstatisticsinfoservice"
	"github.com/feihua/zero-admin/rpc/ums/client/membertagservice"
	"github.com/feihua/zero-admin/rpc/ums/client/membertaskservice"
	"github.com/smartwalle/alipay/v3"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	//会员相关
	GrowthChangeHistoryService           growthchangehistoryservice.GrowthChangeHistoryService
	IntegrationChangeHistoryService      integrationchangehistoryservice.IntegrationChangeHistoryService
	IntegrationConsumeSettingService     integrationconsumesettingservice.IntegrationConsumeSettingService
	MemberLevelService                   memberlevelservice.MemberLevelService
	MemberLoginLogService                memberloginlogservice.MemberLoginLogService
	MemberMemberTagRelationService       membermembertagrelationservice.MemberMemberTagRelationService
	MemberProductCategoryRelationService memberproductcategoryrelationservice.MemberProductCategoryRelationService
	MemberProductCollectionService       memberproductcollectionservice.MemberProductCollectionService
	MemberReadHistoryService             memberreadhistoryservice.MemberReadHistoryService
	MemberReceiveAddressService          memberreceiveaddressservice.MemberReceiveAddressService
	MemberRuleSettingService             memberrulesettingservice.MemberRuleSettingService
	MemberService                        memberservice.MemberService
	MemberStatisticsInfoService          memberstatisticsinfoservice.MemberStatisticsInfoService
	MemberTagService                     membertagservice.MemberTagService
	MemberTaskService                    membertaskservice.MemberTaskService
	MemberAttentionService               memberattentionservice.MemberAttentionService

	//系统相关
	DeptService     deptservice.DeptService
	DictService     dictservice.DictService
	JobService      jobservice.JobService
	LoginLogService loginlogservice.LoginLogService
	SysLogService   syslogservice.SysLogService
	MenuService     menuservice.MenuService
	RoleService     roleservice.RoleService
	UserService     userservice.UserService
	//商品相关
	AlbumPicService                         albumpicservice.AlbumPicService
	AlbumService                            albumservice.AlbumService
	BrandService                            brandservice.BrandService
	CommentReplayService                    commentreplayservice.CommentReplayService
	CommentService                          commentservice.CommentService
	FeightTemplateService                   feighttemplateservice.FeightTemplateService
	MemberPriceService                      memberpriceservice.MemberPriceService
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
	//订单相关
	CartItemService            cartitemservice.CartItemService
	CompanyAddressService      companyaddressservice.CompanyAddressService
	OrderItemService           orderitemservice.OrderItemService
	OrderOperateHistoryService orderoperatehistorservice.OrderOperateHistorService
	OrderReturnApplyService    orderreturnapplyservice.OrderReturnApplyService
	OrderReturnReasonService   orderreturnreasonservice.OrderReturnReasonService
	OrderService               orderservice.OrderService
	OrderSettingService        ordersettingservice.OrderSettingService
	//营销相关
	CouponHistoryService                 couponhistoryservice.CouponHistoryService
	CouponProductCategoryRelationService couponproductcategoryrelationservice.CouponProductCategoryRelationService
	CouponProductRelationService         couponproductrelationservice.CouponProductRelationService
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
	//内容相关
	SubjectService                      subjectservice.SubjectService
	SubjectProductRelationService       subjectproductrelationservice.SubjectProductRelationService
	PreferredAreaService                preferredareaservice.PreferredAreaService
	PreferredAreaProductRelationService preferredareaproductrelationservice.PreferredAreaProductRelationService

	AlipayClient *alipay.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	//初始化支付宝客户端
	client, err := alipay.New(c.Alipay.AppId, c.Alipay.PrivateKey, c.Alipay.IsProduction)
	if err != nil {
		print("初始化支付宝失败")
	}
	umsClient := zrpc.MustNewClient(c.UmsRpc)
	sysClient := zrpc.MustNewClient(c.SysRpc)
	pmsClient := zrpc.MustNewClient(c.PmsRpc)
	omsClient := zrpc.MustNewClient(c.OmsRpc)
	smsClient := zrpc.MustNewClient(c.SmsRpc)
	cmsClient := zrpc.MustNewClient(c.CmsRpc)
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
		MemberReceiveAddressService:          memberreceiveaddressservice.NewMemberReceiveAddressService(umsClient),
		MemberRuleSettingService:             memberrulesettingservice.NewMemberRuleSettingService(umsClient),
		MemberService:                        memberservice.NewMemberService(umsClient),
		MemberStatisticsInfoService:          memberstatisticsinfoservice.NewMemberStatisticsInfoService(umsClient),
		MemberTagService:                     membertagservice.NewMemberTagService(umsClient),
		MemberTaskService:                    membertaskservice.NewMemberTaskService(umsClient),
		MemberAttentionService:               memberattentionservice.NewMemberAttentionService(umsClient),

		DeptService:     deptservice.NewDeptService(sysClient),
		DictService:     dictservice.NewDictService(sysClient),
		JobService:      jobservice.NewJobService(sysClient),
		LoginLogService: loginlogservice.NewLoginLogService(sysClient),
		SysLogService:   syslogservice.NewSysLogService(sysClient),
		MenuService:     menuservice.NewMenuService(sysClient),
		RoleService:     roleservice.NewRoleService(sysClient),
		UserService:     userservice.NewUserService(sysClient),

		AlbumPicService:                         albumpicservice.NewAlbumPicService(pmsClient),
		AlbumService:                            albumservice.NewAlbumService(pmsClient),
		BrandService:                            brandservice.NewBrandService(pmsClient),
		CommentReplayService:                    commentreplayservice.NewCommentReplayService(pmsClient),
		CommentService:                          commentservice.NewCommentService(pmsClient),
		FeightTemplateService:                   feighttemplateservice.NewFeightTemplateService(pmsClient),
		MemberPriceService:                      memberpriceservice.NewMemberPriceService(pmsClient),
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
		OrderOperateHistoryService: orderoperatehistorservice.NewOrderOperateHistorService(omsClient),
		OrderReturnApplyService:    orderreturnapplyservice.NewOrderReturnApplyService(omsClient),
		OrderReturnReasonService:   orderreturnreasonservice.NewOrderReturnReasonService(omsClient),
		OrderService:               orderservice.NewOrderService(omsClient),
		OrderSettingService:        ordersettingservice.NewOrderSettingService(omsClient),

		CouponHistoryService:                 couponhistoryservice.NewCouponHistoryService(smsClient),
		CouponProductCategoryRelationService: couponproductcategoryrelationservice.NewCouponProductCategoryRelationService(smsClient),
		CouponProductRelationService:         couponproductrelationservice.NewCouponProductRelationService(smsClient),
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
