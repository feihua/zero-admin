package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *CouponProductRelationListLogic) CouponProductRelationList(in *sms.CouponProductRelationListReq) (*sms.CouponProductRelationListResp, error) {
	all, _ := l.svcCtx.SmsCouponProductRelationModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.SmsCouponProductRelationModel.Count()

	var list []*sms.CouponProductRelationListData
	for _, item := range *all {

		list = append(list, &sms.CouponProductRelationListData{
			Id:          item.Id,
			CouponId:    item.CouponId,
			ProductId:   item.ProductId,
			ProductName: item.ProductName,
			ProductSn:   item.ProductSn,
		})
	}

	fmt.Println(list)
	return &sms.CouponProductRelationListResp{
		Total: count,
		List:  list,
	}, nil
}
