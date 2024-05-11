package orderreturnapplyservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderReturnApplyListLogic 退货申请
/*
Author: LiuFeiHua
Date: 2024/5/8 9:32
*/
type OrderReturnApplyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnApplyListLogic {
	return &OrderReturnApplyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderReturnApplyList 查询退货申请列表
func (l *OrderReturnApplyListLogic) OrderReturnApplyList(in *omsclient.OrderReturnApplyListReq) (*omsclient.OrderReturnApplyListResp, error) {
	q := query.OmsOrderReturnApply.WithContext(l.ctx)
	if len(in.OrderSn) > 0 {
		q = q.Where(query.OmsOrderReturnApply.OrderSn.Like("%" + in.OrderSn + "%"))
	}
	if len(in.MemberUsername) > 0 {
		q = q.Where(query.OmsOrderReturnApply.OrderSn.Like("%" + in.MemberUsername + "%"))
	}

	if in.Status != 4 {
		q = q.Where(query.OmsOrderReturnApply.Status.Eq(in.Status))
	}
	//if len(in.CreateTime) > 0 {
	//	q.Where(query.OmsOrderReturnApply.CreateTime.Eq(in.CreateTime))
	//}
	//if len(in.HandleTime) > 0 {
	//	q.Where(query.OmsOrderReturnApply.HandleTime.(in.HandleTime))
	//}

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		in, _ := json.Marshal(in)
		logc.Errorf(l.ctx, "查询退货申请列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*omsclient.OrderReturnApplyListData
	for _, item := range result {

		list = append(list, &omsclient.OrderReturnApplyListData{
			Id:               item.ID,
			OrderId:          item.OrderID,
			CompanyAddressId: item.CompanyAddressID,
			ProductId:        item.ProductID,
			OrderSn:          item.OrderSn,
			CreateTime:       item.CreateTime.Format("2006-01-02 15:04:05"),
			MemberUsername:   item.MemberUsername,
			ReturnAmount:     item.ReturnAmount,
			ReturnName:       item.ReturnName,
			ReturnPhone:      item.ReturnPhone,
			Status:           item.Status,
			HandleTime:       item.HandleTime.Format("2006-01-02 15:04:05"),
			ProductPic:       item.ProductPic,
			ProductName:      item.ProductName,
			ProductBrand:     item.ProductBrand,
			ProductAttr:      item.ProductAttr,
			ProductCount:     item.ProductCount,
			ProductPrice:     item.ProductPrice,
			ProductRealPrice: item.ProductRealPrice,
			Reason:           item.Reason,
			Description:      item.Description,
			ProofPics:        item.ProofPics,
			HandleNote:       item.HandleNote,
			HandleMan:        item.HandleMan,
			ReceiveMan:       item.ReceiveMan,
			ReceiveTime:      item.ReceiveTime.Format("2006-01-02 15:04:05"),
			ReceiveNote:      *item.ReceiveNote,
		})
	}

	logc.Infof(l.ctx, "查询退货申请列表信息,参数：%+v,响应：%+v", in, list)
	return &omsclient.OrderReturnApplyListResp{
		Total: count,
		List:  list,
	}, nil
}
