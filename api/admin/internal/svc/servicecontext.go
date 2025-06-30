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
	"github.com/feihua/zero-admin/rpc/oms/client/orderoperationlogservice"
	"github.com/feihua/zero-admin/rpc/oms/client/orderreturnreasonservice"
	"github.com/feihua/zero-admin/rpc/oms/client/orderreturnservice"
	"github.com/feihua/zero-admin/rpc/oms/client/orderservice"
	"github.com/feihua/zero-admin/rpc/oms/client/ordersettingservice"
	"github.com/feihua/zero-admin/rpc/pms/client/commentreplayservice"
	"github.com/feihua/zero-admin/rpc/pms/client/commentservice"
	"github.com/feihua/zero-admin/rpc/pms/client/feighttemplateservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productattributegroupservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productattributeservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productattributevalueservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productbrandservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productcategoryattributerelationservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productcategoryservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productfullreductionservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productladderservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productoperatelogservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productskuservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productspecservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productspecvalueservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productspuservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productvertifyrecordservice"
	"github.com/feihua/zero-admin/rpc/sms/client/couponrecordservice"
	"github.com/feihua/zero-admin/rpc/sms/client/couponscopeservice"
	"github.com/feihua/zero-admin/rpc/sms/client/coupontypeservice"
	"github.com/feihua/zero-admin/rpc/sms/client/seckillactivityservice"
	"github.com/feihua/zero-admin/rpc/sms/client/seckillproductservice"
	"github.com/feihua/zero-admin/rpc/sms/client/seckillreservationservice"
	"github.com/feihua/zero-admin/rpc/sms/client/seckillsessionservice"

	"github.com/feihua/zero-admin/rpc/sms/client/couponservice"

	"github.com/feihua/zero-admin/rpc/sms/client/homeadvertiseservice"

	"github.com/feihua/zero-admin/rpc/sys/client/deptservice"
	"github.com/feihua/zero-admin/rpc/sys/client/dictitemservice"
	"github.com/feihua/zero-admin/rpc/sys/client/dicttypeservice"
	"github.com/feihua/zero-admin/rpc/sys/client/loginlogservice"
	"github.com/feihua/zero-admin/rpc/sys/client/menuservice"
	"github.com/feihua/zero-admin/rpc/sys/client/operatelogservice"
	"github.com/feihua/zero-admin/rpc/sys/client/postservice"
	"github.com/feihua/zero-admin/rpc/sys/client/roleservice"
	"github.com/feihua/zero-admin/rpc/sys/client/userservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberaddressservice"
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
	"github.com/feihua/zero-admin/rpc/ums/client/membersignlogservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberstatisticsinfoservice"
	"github.com/feihua/zero-admin/rpc/ums/client/membertagrelationservice"
	"github.com/feihua/zero-admin/rpc/ums/client/membertagservice"
	"github.com/feihua/zero-admin/rpc/ums/client/membertaskrelationservice"
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
	MemberGrowthLogService               membergrowthlogservice.MemberGrowthLogService
	MemberPointsLogService               memberpointslogservice.MemberPointsLogService
	MemberConsumeSettingService          memberconsumesettingservice.MemberConsumeSettingService
	MemberLevelService                   memberlevelservice.MemberLevelService
	MemberLoginLogService                memberloginlogservice.MemberLoginLogService
	MemberTagRelationService             membertagrelationservice.MemberTagRelationService
	MemberProductCategoryRelationService memberproductcategoryrelationservice.MemberProductCategoryRelationService
	MemberProductCollectionService       memberproductcollectionservice.MemberProductCollectionService
	MemberReadHistoryService             memberreadhistoryservice.MemberReadHistoryService
	MemberAddressService                 memberaddressservice.MemberAddressService
	MemberRuleSettingService             memberrulesettingservice.MemberRuleSettingService
	MemberInfoService                    memberinfoservice.MemberInfoService
	MemberStatisticsInfoService          memberstatisticsinfoservice.MemberStatisticsInfoService
	MemberTagService                     membertagservice.MemberTagService
	MemberTaskService                    membertaskservice.MemberTaskService
	MemberTaskRelationService            membertaskrelationservice.MemberTaskRelationService
	MemberSignLogService                 membersignlogservice.MemberSignLogService
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
	ProductBrandService                     productbrandservice.ProductBrandService
	CommentReplayService                    commentreplayservice.CommentReplayService
	CommentService                          commentservice.CommentService
	FeightTemplateService                   feighttemplateservice.FeightTemplateService
	ProductAttributeGroupService            productattributegroupservice.ProductAttributeGroupService
	ProductAttributeService                 productattributeservice.ProductAttributeService
	ProductAttributeValueService            productattributevalueservice.ProductAttributeValueService
	ProductCategoryAttributeRelationService productcategoryattributerelationservice.ProductCategoryAttributeRelationService
	ProductCategoryService                  productcategoryservice.ProductCategoryService
	ProductFullReductionService             productfullreductionservice.ProductFullReductionService
	ProductLadderService                    productladderservice.ProductLadderService
	ProductOperateLogService                productoperatelogservice.ProductOperateLogService
	ProductSpuService                       productspuservice.ProductSpuService
	ProductVertifyRecordService             productvertifyrecordservice.ProductVertifyRecordService
	ProductSkuService                       productskuservice.ProductSkuService
	ProductSpecService                      productspecservice.ProductSpecService
	ProductSpecValueService                 productspecvalueservice.ProductSpecValueService
	// 订单相关
	CartItemService          cartitemservice.CartItemService
	CompanyAddressService    companyaddressservice.CompanyAddressService
	OrderOperationLogService orderoperationlogservice.OrderOperationLogService
	OrderReturnService       orderreturnservice.OrderReturnService
	OrderReturnReasonService orderreturnreasonservice.OrderReturnReasonService
	OrderService             orderservice.OrderService
	OrderSettingService      ordersettingservice.OrderSettingService
	// 营销相关
	CouponRecordService       couponrecordservice.CouponRecordService
	CouponScopeService        couponscopeservice.CouponScopeService
	CouponService             couponservice.CouponService
	CouponTypeService         coupontypeservice.CouponTypeService
	HomeAdvertiseService      homeadvertiseservice.HomeAdvertiseService
	SeckillActivityService    seckillactivityservice.SeckillActivityService
	SeckillProductService     seckillproductservice.SeckillProductService
	SeckillReservationService seckillreservationservice.SeckillReservationService
	SeckillSessionService     seckillsessionservice.SeckillSessionService
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
		MemberGrowthLogService:               membergrowthlogservice.NewMemberGrowthLogService(umsClient),
		MemberPointsLogService:               memberpointslogservice.NewMemberPointsLogService(umsClient),
		MemberConsumeSettingService:          memberconsumesettingservice.NewMemberConsumeSettingService(umsClient),
		MemberLevelService:                   memberlevelservice.NewMemberLevelService(umsClient),
		MemberLoginLogService:                memberloginlogservice.NewMemberLoginLogService(umsClient),
		MemberTagRelationService:             membertagrelationservice.NewMemberTagRelationService(umsClient),
		MemberProductCategoryRelationService: memberproductcategoryrelationservice.NewMemberProductCategoryRelationService(umsClient),
		MemberProductCollectionService:       memberproductcollectionservice.NewMemberProductCollectionService(umsClient),
		MemberReadHistoryService:             memberreadhistoryservice.NewMemberReadHistoryService(umsClient),
		MemberAddressService:                 memberaddressservice.NewMemberAddressService(umsClient),
		MemberRuleSettingService:             memberrulesettingservice.NewMemberRuleSettingService(umsClient),
		MemberInfoService:                    memberinfoservice.NewMemberInfoService(umsClient),
		MemberStatisticsInfoService:          memberstatisticsinfoservice.NewMemberStatisticsInfoService(umsClient),
		MemberTagService:                     membertagservice.NewMemberTagService(umsClient),
		MemberTaskService:                    membertaskservice.NewMemberTaskService(umsClient),
		MemberTaskRelationService:            membertaskrelationservice.NewMemberTaskRelationService(umsClient),
		MemberSignLogService:                 membersignlogservice.NewMemberSignLogService(umsClient),

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

		ProductBrandService:                     productbrandservice.NewProductBrandService(pmsClient),
		CommentReplayService:                    commentreplayservice.NewCommentReplayService(pmsClient),
		CommentService:                          commentservice.NewCommentService(pmsClient),
		FeightTemplateService:                   feighttemplateservice.NewFeightTemplateService(pmsClient),
		ProductAttributeGroupService:            productattributegroupservice.NewProductAttributeGroupService(pmsClient),
		ProductAttributeService:                 productattributeservice.NewProductAttributeService(pmsClient),
		ProductAttributeValueService:            productattributevalueservice.NewProductAttributeValueService(pmsClient),
		ProductCategoryAttributeRelationService: productcategoryattributerelationservice.NewProductCategoryAttributeRelationService(pmsClient),
		ProductCategoryService:                  productcategoryservice.NewProductCategoryService(pmsClient),
		ProductFullReductionService:             productfullreductionservice.NewProductFullReductionService(pmsClient),
		ProductLadderService:                    productladderservice.NewProductLadderService(pmsClient),
		ProductOperateLogService:                productoperatelogservice.NewProductOperateLogService(pmsClient),
		ProductSpuService:                       productspuservice.NewProductSpuService(pmsClient),
		ProductVertifyRecordService:             productvertifyrecordservice.NewProductVertifyRecordService(pmsClient),
		ProductSkuService:                       productskuservice.NewProductSkuService(pmsClient),
		ProductSpecService:                      productspecservice.NewProductSpecService(pmsClient),
		ProductSpecValueService:                 productspecvalueservice.NewProductSpecValueService(pmsClient),

		CartItemService:          cartitemservice.NewCartItemService(omsClient),
		CompanyAddressService:    companyaddressservice.NewCompanyAddressService(omsClient),
		OrderOperationLogService: orderoperationlogservice.NewOrderOperationLogService(omsClient),
		OrderReturnService:       orderreturnservice.NewOrderReturnService(omsClient),
		OrderReturnReasonService: orderreturnreasonservice.NewOrderReturnReasonService(omsClient),
		OrderService:             orderservice.NewOrderService(omsClient),
		OrderSettingService:      ordersettingservice.NewOrderSettingService(omsClient),

		CouponRecordService:       couponrecordservice.NewCouponRecordService(smsClient),
		CouponScopeService:        couponscopeservice.NewCouponScopeService(smsClient),
		CouponService:             couponservice.NewCouponService(smsClient),
		CouponTypeService:         coupontypeservice.NewCouponTypeService(smsClient),
		HomeAdvertiseService:      homeadvertiseservice.NewHomeAdvertiseService(smsClient),
		SeckillActivityService:    seckillactivityservice.NewSeckillActivityService(smsClient),
		SeckillProductService:     seckillproductservice.NewSeckillProductService(smsClient),
		SeckillReservationService: seckillreservationservice.NewSeckillReservationService(smsClient),
		SeckillSessionService:     seckillsessionservice.NewSeckillSessionService(smsClient),

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
