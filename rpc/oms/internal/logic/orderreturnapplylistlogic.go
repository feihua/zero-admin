package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *OrderReturnApplyListLogic) OrderReturnApplyList(in *oms.OrderReturnApplyListReq) (*oms.OrderReturnApplyListResp, error) {
	all, err := l.svcCtx.OmsOrderReturnApplyModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.OmsOrderReturnApplyModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询退货申请列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*oms.OrderReturnApplyListData
	for _, item := range *all {

		list = append(list, &oms.OrderReturnApplyListData{
			Id:               item.Id,
			OrderId:          item.OrderId,
			CompanyAddressId: item.CompanyAddressId,
			ProductId:        item.ProductId,
			OrderSn:          item.OrderSn,
			CreateTime:       item.CreateTime.Format("2006-01-02 15:04:05"),
			MemberUsername:   item.MemberUsername,
			ReturnAmount:     int64(item.ReturnAmount),
			ReturnName:       item.ReturnName,
			ReturnPhone:      item.ReturnPhone,
			Status:           item.Status,
			HandleTime:       item.HandleTime.Format("2006-01-02 15:04:05"),
			ProductPic:       item.ProductPic,
			ProductName:      item.ProductName,
			ProductBrand:     item.ProductBrand,
			ProductAttr:      item.ProductAttr,
			ProductCount:     item.ProductCount,
			ProductPrice:     int64(item.ProductPrice),
			ProductRealPrice: int64(item.ProductRealPrice),
			Reason:           item.Reason,
			Description:      item.Description,
			ProofPics:        item.ProofPics,
			HandleNote:       item.HandleNote,
			HandleMan:        item.HandleMan,
			ReceiveMan:       item.ReceiveMan,
			ReceiveTime:      item.ReceiveTime.Format("2006-01-02 15:04:05"),
			ReceiveNote:      item.ReceiveNote,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询退货申请列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &oms.OrderReturnApplyListResp{
		Total: count,
		List:  list,
	}, nil
}
