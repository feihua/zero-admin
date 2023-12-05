package couponservicelogic

import (
	"context"
	"encoding/json"

	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// CouponFindByProductIdAndProductCategoryIdLogic
/*
Author: LiuFeiHua
Date: 2023/12/5 11:29
*/
type CouponFindByProductIdAndProductCategoryIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponFindByProductIdAndProductCategoryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponFindByProductIdAndProductCategoryIdLogic {
	return &CouponFindByProductIdAndProductCategoryIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CouponFindByProductIdAndProductCategoryId 根据商品Id和分类id查询可用的优惠券(app)
func (l *CouponFindByProductIdAndProductCategoryIdLogic) CouponFindByProductIdAndProductCategoryId(in *smsclient.CouponFindByProductIdAndProductCategoryIdReq) (*smsclient.CouponFindByProductIdAndProductCategoryIdResp, error) {
	all, err := l.svcCtx.SmsCouponModel.CouponFindByProductIdAndProductCategoryId(l.ctx, in.ProductId, in.ProductCategoryId)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询商品优惠券列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*smsclient.CouponListData
	for _, coupon := range *all {

		list = append(list, &smsclient.CouponListData{
			Id:           coupon.Id,
			Type:         coupon.Type,
			Name:         coupon.Name,
			Platform:     coupon.Platform,
			Count:        coupon.Count,
			Amount:       coupon.Amount,
			PerLimit:     coupon.PerLimit,
			MinPoint:     coupon.MinPoint,
			StartTime:    coupon.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:      coupon.EndTime.Format("2006-01-02 15:04:05"),
			UseType:      coupon.UseType,
			Note:         coupon.Note,
			PublishCount: coupon.PublishCount,
			UseCount:     coupon.UseCount,
			ReceiveCount: coupon.ReceiveCount,
			EnableTime:   coupon.EnableTime.Format("2006-01-02 15:04:05"),
			Code:         coupon.Code,
			MemberLevel:  coupon.MemberLevel,
		})

		reqStr, _ := json.Marshal(in)
		listStr, _ := json.Marshal(list)
		logx.WithContext(l.ctx).Infof("查询商品优惠券列表信息,参数：%s,响应：%s", reqStr, listStr)
	}

	return &smsclient.CouponFindByProductIdAndProductCategoryIdResp{
		Total: 0,
		List:  list,
	}, nil

}
