package couponproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponProductRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponProductRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponProductRelationListLogic {
	return &CouponProductRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponProductRelationListLogic) CouponProductRelationList(in *smsclient.CouponProductRelationListReq) (*smsclient.CouponProductRelationListResp, error) {
	result, err := query.SmsCouponProductRelation.WithContext(l.ctx).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询优惠券与产品关糸列表信息失败,参数：%+v,异常:%s", in, err.Error)
		return nil, err
	}

	var list []*smsclient.CouponProductRelationListData
	for _, item := range result {

		list = append(list, &smsclient.CouponProductRelationListData{
			Id:          item.ID,
			CouponId:    item.CouponID,
			ProductId:   item.ProductID,
			ProductName: item.ProductName,
			ProductSn:   item.ProductSn,
		})
	}

	logc.Infof(l.ctx, "查询优惠券与产品关糸列表信息,参数：%+v,响应：%+v", in, list)
	return &smsclient.CouponProductRelationListResp{
		Total: 0,
		List:  list,
	}, nil
}
