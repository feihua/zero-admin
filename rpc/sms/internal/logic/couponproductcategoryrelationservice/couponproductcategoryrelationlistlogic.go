package couponproductcategoryrelationservicelogic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/smsclient"

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
	all, err := l.svcCtx.SmsCouponProductCategoryRelationModel.FindAll(l.ctx, 1, 100)
	count, _ := l.svcCtx.SmsCouponProductCategoryRelationModel.Count(l.ctx)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询优惠券与产品关糸列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*smsclient.CouponProductCategoryRelationListData
	for _, item := range *all {

		list = append(list, &smsclient.CouponProductCategoryRelationListData{
			Id:                  item.Id,
			CouponId:            item.CouponId,
			ProductCategoryId:   item.ProductCategoryId,
			ProductCategoryName: item.ProductCategoryName,
			ParentCategoryName:  item.ParentCategoryName,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询优惠券与产品关糸列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &smsclient.CouponProductCategoryRelationListResp{
		Total: count,
		List:  list,
	}, nil
}
