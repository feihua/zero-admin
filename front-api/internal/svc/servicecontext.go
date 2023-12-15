package svc

import (
	"github.com/smartwalle/alipay/v3"
	"github.com/zeromicro/go-zero/zrpc"
	"zero-admin/front-api/internal/config"
	"zero-admin/rpc/cms/client/prefrenceareaproductrelationservice"
	"zero-admin/rpc/cms/client/prefrenceareaservice"
	"zero-admin/rpc/cms/client/subjectproductrelationservice"
	"zero-admin/rpc/cms/client/subjectservice"
	"zero-admin/rpc/oms/client/cartitemservice"
	"zero-admin/rpc/oms/client/companyaddressservice"
	"zero-admin/rpc/oms/client/orderitemservice"
	"zero-admin/rpc/oms/client/orderoperatehistorservice"
	"zero-admin/rpc/oms/client/orderreturnapplyservice"
	"zero-admin/rpc/oms/client/orderreturnreasonservice"
	"zero-admin/rpc/oms/client/orderservice"
	"zero-admin/rpc/oms/client/ordersettingservice"
	"zero-admin/rpc/pms/client/albumpicservice"
	"zero-admin/rpc/pms/client/albumservice"
	"zero-admin/rpc/pms/client/brandservice"
	"zero-admin/rpc/pms/client/commentreplayservice"
	"zero-admin/rpc/pms/client/commentservice"
	"zero-admin/rpc/pms/client/feighttemplateservice"
	"zero-admin/rpc/pms/client/memberpriceservice"
	"zero-admin/rpc/pms/client/productattributecategoryservice"
	"zero-admin/rpc/pms/client/productattributeservice"
	"zero-admin/rpc/pms/client/productattributevalueservice"
	"zero-admin/rpc/pms/client/productcategoryattributerelationservice"
	"zero-admin/rpc/pms/client/productcategoryservice"
	"zero-admin/rpc/pms/client/productfullreductionservice"
	"zero-admin/rpc/pms/client/productladderservice"
	"zero-admin/rpc/pms/client/productoperatelogservice"
	"zero-admin/rpc/pms/client/productservice"
	"zero-admin/rpc/pms/client/productvertifyrecordservice"
	"zero-admin/rpc/pms/client/skustockservice"
	"zero-admin/rpc/sms/client/couponhistoryservice"
	"zero-admin/rpc/sms/client/couponproductcategoryrelationservice"
	"zero-admin/rpc/sms/client/couponproductrelationservice"
	"zero-admin/rpc/sms/client/couponservice"
	"zero-admin/rpc/sms/client/flashpromotionlogservice"
	"zero-admin/rpc/sms/client/flashpromotionproductrelationservice"
	"zero-admin/rpc/sms/client/flashpromotionservice"
	"zero-admin/rpc/sms/client/flashpromotionsessionservice"
	"zero-admin/rpc/sms/client/homeadvertiseservice"
	"zero-admin/rpc/sms/client/homebrandservice"
	"zero-admin/rpc/sms/client/homenewproductservice"
	"zero-admin/rpc/sms/client/homerecommendproductservice"
	"zero-admin/rpc/sms/client/homerecommendsubjectservice"
	"zero-admin/rpc/sys/client/configservice"
	"zero-admin/rpc/sys/client/deptservice"
	"zero-admin/rpc/sys/client/dictservice"
	"zero-admin/rpc/sys/client/jobservice"
	"zero-admin/rpc/sys/client/loginlogservice"
	"zero-admin/rpc/sys/client/menuservice"
	"zero-admin/rpc/sys/client/roleservice"
	"zero-admin/rpc/sys/client/syslogservice"
	"zero-admin/rpc/sys/client/userservice"
	"zero-admin/rpc/ums/client/growthchangehistoryservice"
	"zero-admin/rpc/ums/client/integrationchangehistoryservice"
	"zero-admin/rpc/ums/client/integrationconsumesettingservice"
	"zero-admin/rpc/ums/client/memberattentionservice"
	"zero-admin/rpc/ums/client/memberlevelservice"
	"zero-admin/rpc/ums/client/memberloginlogservice"
	"zero-admin/rpc/ums/client/membermembertagrelationservice"
	"zero-admin/rpc/ums/client/memberproductcategoryrelationservice"
	"zero-admin/rpc/ums/client/memberproductcollectionservice"
	"zero-admin/rpc/ums/client/memberreadhistoryservice"
	"zero-admin/rpc/ums/client/memberreceiveaddressservice"
	"zero-admin/rpc/ums/client/memberrulesettingservice"
	"zero-admin/rpc/ums/client/memberservice"
	"zero-admin/rpc/ums/client/memberstatisticsinfoservice"
	"zero-admin/rpc/ums/client/membertagservice"
	"zero-admin/rpc/ums/client/membertaskservice"
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
	ConfigService   configservice.ConfigService
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
	PrefrenceAreaService                prefrenceareaservice.PrefrenceAreaService
	PrefrenceAreaProductRelationService prefrenceareaproductrelationservice.PrefrenceAreaProductRelationService

	AlipayClient *alipay.Client
}

const (
	//kAppId      = "2021000122672643"
	//kPrivateKey = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCG/iWUj0IScEpQWeOxEdC8muaPNZdvmnXTeMXTLARMbOIlg3LsED1Nt4VqjZ/K2J+3jQGcGMtCGbIOTlJpBnjpM5n/IA+oiUmhooH6ftBJcDEpbixDAk9pvrW/LXE52V3+MMrpKKEwJMt+YOUKf//m4svx5aXgxzb7QlMk1F7TaYjxxbULcb2eZKhRfhfJwQYrt8DP8jRQq1Jo5m52tnvMrKQjyVJwJy44YTsEMrXRaPWZ0+XKMjUWTbnqdVECt9r2z/5M3r6x2j0qxiQAYBAGgeD/iazDtvhXCrrgb4nG7yEjOKhuVSN1mRItdmqW4TferOGCJZanxCFcdzrQkf2BAgMBAAECggEAcrbOIKyMrTaXMCjzAKnvBBduDgywn7pWnlpnYchp7rgohVBq/IfgUIa/7YhkXfAv6b79uzSmpYlIcjfEeFNztFiRaOhJ5iKkW6LJaaESRxX78QUav+barTXPJKLtMQeyhCvagsBwGYVrF/4nJQEY6Y+ZV/qbN6SS6Hm4RffijSxJKJv6CzUjkajxnDzJBz3ySR0z5L8Q0abRkb+PS8cgUO9VbiGcD4leQV3wXM1e3BFDBsQ7POVe2EAYWcGbDVHqV8hFM186zs2zsgMp4VzUODuPuBPaVbKnZcml3Z5aEIamJH52YtES3R5WvWWQ2We6xwmoGXmuGENHejtdqB21AQKBgQDE0Tu6yaP76lWxlq15TCeQORzVlWJ7zhcvoxEj4BTHTggs1r9dRLqLai9Cd2+cciirh+VeOFIUKSfX/yF+1iTYPO0qXVXtPVBkc8iVrRU1bDO7UCAWAMz36DteOSYsPL6roJrudmLuke1AQQz6I5ydRlupZpDYFsuxVNs2XhBtyQKBgQCvlbtCr8LEVpPP5Cye/KhBT3q8vORijFyz/U0Tq4IK4DRjKR2MSSrueR5Osq+qcAmqi7Fdn89XtjHQE1QFABLbHPE9F3f79M1/pndHABkOGXOHIxq8K+X9AshSRK1MoQp1pJLm2E4KiMxBSzaIvneug3ukPqeqm44SKG9uBuQN+QKBgQCE1peSzY+xcoseDo3NJZo6XGHawjWzS/kYPN5PsWk0z7Ty1opYYA/sEuIM4WHiXKaYh2NHAYpccx6iSV+JJO2/SPfltRNOyShedEs4wpZi9UHBNiZB046D8ClJwhbCmskyO3b2Zc8GKFXSHVWt6qVE/XzWTBSM1G3spVJDUp+SCQKBgQCpsNJmU4iuyWFW1BTPnixZ2h8rUn6CQ1bAWHfqH6GxMxdOEglNb9T+3a0Nr6EX3elpmlHSwsTW5uzjRBq6LmUKv8DhItJBfUgxKscxpgWQ28YL/0AyRVajG9JPt7GoUibSpTeXw8pAYg7Mt4y/wRvXW5jdlfPibS1znQJ72ksCuQKBgHOTEVs6M42vjXL/ijZPR1RBsY7rROKSsklpPjbyiHGrQeeJMWh4n1SNAVIrlaQmj3B8gCaCAFEJX702ajOEfnPrnDTO3V0w9rSP9S1PGoRr86JTOT8g3A+0d7Vp2BGnHZ/9s9Ed3cZzTev2TBZHdKeoTlYQfYaD7n/kttpI0Xx2"
	kAppId      = "2021003185602741"
	kPrivateKey = "123456"

	kServerPort = "9989"
	// TODO 设置回调地址域名
	kServerDomain = ""
)

func NewServiceContext(c config.Config) *ServiceContext {
	client, err := alipay.New(kAppId, kPrivateKey, true)
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

		ConfigService:   configservice.NewConfigService(sysClient),
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
		PrefrenceAreaService:                prefrenceareaservice.NewPrefrenceAreaService(cmsClient),
		PrefrenceAreaProductRelationService: prefrenceareaproductrelationservice.NewPrefrenceAreaProductRelationService(cmsClient),

		AlipayClient: client,
	}
}
