package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CouponHistoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponHistoryListLogic {
	return &CouponHistoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponHistoryListLogic) CouponHistoryList(in *sms.CouponHistoryListReq) (*sms.CouponHistoryListResp, error) {
	all, _ := l.svcCtx.SmsCouponHistoryModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	var list []*sms.CouponHistoryListData
	for _, item := range *all {

		list = append(list, &sms.CouponHistoryListData{
			Id:             item.Id,
			CouponId:       item.CouponId,
			MemberId:       item.MemberId,
			CouponCode:     item.CouponCode,
			MemberNickname: item.MemberNickname,
			GetType:        item.GetType,
			CreateTime:     item.CreateTime.Format("2006-01-02 15:04:05"),
			UseStatus:      item.UseStatus,
			UseTime:        item.UseTime.Format("2006-01-02 15:04:05"),
			OrderId:        item.OrderId,
			OrderSn:        item.OrderSn,
		})
	}

	fmt.Println(list)
	return &sms.CouponHistoryListResp{
		Total: 10,
		List:  list,
	}, nil
}
