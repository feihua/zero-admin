package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *CouponProductCategoryRelationListLogic) CouponProductCategoryRelationList(in *sms.CouponProductCategoryRelationListReq) (*sms.CouponProductCategoryRelationListResp, error) {
	all, _ := l.svcCtx.SmsCouponProductCategoryRelationModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	var list []*sms.CouponProductCategoryRelationListData
	for _, item := range *all {

		list = append(list, &sms.CouponProductCategoryRelationListData{
			Id:                  item.Id,
			CouponId:            item.CouponId,
			ProductCategoryId:   item.ProductCategoryId,
			ProductCategoryName: item.ProductCategoryName,
			ParentCategoryName:  item.ParentCategoryName,
		})
	}

	fmt.Println(list)
	return &sms.CouponProductCategoryRelationListResp{
		Total: 10,
		List:  list,
	}, nil
}
