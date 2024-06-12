package orderreturnapplyservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOrderReturnApplyListLogic 查询订单退货申请列表
/*
Author: LiuFeiHua
Date: 2024/6/12 10:15
*/
type QueryOrderReturnApplyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderReturnApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderReturnApplyListLogic {
	return &QueryOrderReturnApplyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOrderReturnApplyList 查询订单退货申请列表
func (l *QueryOrderReturnApplyListLogic) QueryOrderReturnApplyList(in *omsclient.QueryOrderReturnApplyListReq) (*omsclient.QueryOrderReturnApplyListResp, error) {
	q := query.OmsOrderReturnApply.WithContext(l.ctx)
	if len(in.OrderSn) > 0 {
		q = q.Where(query.OmsOrderReturnApply.OrderSn.Like("%" + in.OrderSn + "%"))
	}
	if len(in.MemberUsername) > 0 {
		q = q.Where(query.OmsOrderReturnApply.MemberUsername.Like("%" + in.MemberUsername + "%"))
	}

	if in.Status != 4 {
		q = q.Where(query.OmsOrderReturnApply.Status.Eq(in.Status))
	}
	if len(in.HandleMan) > 0 {
		q = q.Where(query.OmsOrderReturnApply.HandleMan.Like("%" + in.HandleMan + "%"))
	}
	if len(in.ReturnName) > 0 {
		q = q.Where(query.OmsOrderReturnApply.ReturnName.Like("%" + in.ReturnName + "%"))
	}
	if len(in.ReturnPhone) > 0 {
		q = q.Where(query.OmsOrderReturnApply.ReturnPhone.Like("%" + in.ReturnPhone + "%"))
	}

	//if len(in.CreateTime) > 0 {
	//		where = where + fmt.Sprintf(" AND date_format(create_time,'%%Y-%%m-%%d') = '%s'", strings.Split(in.CreateTime, " ")[0])
	//	}
	//	if len(in.HandleTime) > 0 {
	//		where = where + fmt.Sprintf(" AND date_format(handle_time,'%%Y-%%m-%%d') = '%s'", strings.Split(in.HandleTime, " ")[0])
	//	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
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

	return &omsclient.QueryOrderReturnApplyListResp{
		Total: count,
		List:  list,
	}, nil

}
