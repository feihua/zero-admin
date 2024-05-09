package couponproductcategoryrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponProductCategoryRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponProductCategoryRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponProductCategoryRelationListLogic {
	return &CouponProductCategoryRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponProductCategoryRelationListLogic) CouponProductCategoryRelationList(in *smsclient.CouponProductCategoryRelationListReq) (*smsclient.CouponProductCategoryRelationListResp, error) {

	result, err := query.SmsCouponProductCategoryRelation.WithContext(l.ctx).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询优惠券与产品分类关糸列表信息失败,参数：%+v,异常:%s", in, err.Error)
		return nil, err
	}

	var list []*smsclient.CouponProductCategoryRelationListData
	for _, item := range result {

		list = append(list, &smsclient.CouponProductCategoryRelationListData{
			Id:                  item.ID,
			CouponId:            item.CouponID,
			ProductCategoryId:   item.ProductCategoryID,
			ProductCategoryName: item.ProductCategoryName,
			ParentCategoryName:  item.ParentCategoryName,
		})
	}

	logc.Infof(l.ctx, "查询优惠券与产品分类关糸列表信息,参数：%+v,响应：%+v", in, list)
	return &smsclient.CouponProductCategoryRelationListResp{
		Total: 0,
		List:  list,
	}, nil
}
