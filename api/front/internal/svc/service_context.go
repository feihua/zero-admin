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
	"github.com/feihua/zero-admin/rpc/oms/client/orderoperatehistoryservice"
	"github.com/feihua/zero-admin/rpc/oms/client/orderreturnapplyservice"
	"github.com/feihua/zero-admin/rpc/oms/client/orderreturnreasonservice"
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
	"github.com/feihua/zero-admin/rpc/sms/client/couponservice"
	"github.com/feihua/zero-admin/rpc/sms/client/coupontypeservice"
	"github.com/feihua/zero-admin/rpc/sms/client/homeadvertiseservice"
	"github.com/feihua/zero-admin/rpc/sms/client/seckillactivityservice"
	"github.com/feihua/zero-admin/rpc/sms/client/seckillproductservice"
	"github.com/feihua/zero-admin/rpc/sms/client/seckillreservationservice"
	"github.com/feihua/zero-admin/rpc/sms/client/seckillsessionservice"
	"github.com/feihua/zero-admin/rpc/sys/client/deptservice"
	"github.com/feihua/zero-admin/rpc/sys/client/loginlogservice"
	"github.com/feihua/zero-admin/rpc/sys/client/menuservice"
	"github.com/feihua/zero-admin/rpc/sys/client/roleservice"
	"github.com/feihua/zero-admin/rpc/sys/client/userservice"
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

	// 系统相关
	DeptService     deptservice.DeptService
	LoginLogService loginlogservice.LoginLogService
	MenuService     menuservice.MenuService
	RoleService     roleservice.RoleService
	UserService     userservice.UserService
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
	CartItemService            cartitemservice.CartItemService
	CompanyAddressService      companyaddressservice.CompanyAddressService
	OrderItemService           orderitemservice.OrderItemService
	OrderOperateHistoryService orderoperatehistoryservice.OrderOperateHistoryService
	OrderReturnApplyService    orderreturnapplyservice.OrderReturnApplyService
	OrderReturnReasonService   orderreturnreasonservice.OrderReturnReasonService
	OrderService               orderservice.OrderService
	OrderSettingService        ordersettingservice.OrderSettingService
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

	AlipayClient *alipay.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化支付宝客户端
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

		DeptService:     deptservice.NewDeptService(sysClient),
		LoginLogService: loginlogservice.NewLoginLogService(sysClient),
		MenuService:     menuservice.NewMenuService(sysClient),
		RoleService:     roleservice.NewRoleService(sysClient),
		UserService:     userservice.NewUserService(sysClient),

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

		CartItemService:            cartitemservice.NewCartItemService(omsClient),
		CompanyAddressService:      companyaddressservice.NewCompanyAddressService(omsClient),
		OrderItemService:           orderitemservice.NewOrderItemService(omsClient),
		OrderOperateHistoryService: orderoperatehistoryservice.NewOrderOperateHistoryService(omsClient),
		OrderReturnApplyService:    orderreturnapplyservice.NewOrderReturnApplyService(omsClient),
		OrderReturnReasonService:   orderreturnreasonservice.NewOrderReturnReasonService(omsClient),
		OrderService:               orderservice.NewOrderService(omsClient),
		OrderSettingService:        ordersettingservice.NewOrderSettingService(omsClient),

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

		AlipayClient: client,
	}
}
