package orderreturnapplyservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/pkg/time_util"
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

	// if len(in.CreateTime) > 0 {
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
			Id:               item.ID,                                  //
			OrderId:          item.OrderID,                             // 订单id
			CompanyAddressId: item.CompanyAddressID,                    // 收货地址表id
			ProductId:        item.ProductID,                           // 退货商品id
			OrderSn:          item.OrderSn,                             // 订单编号
			CreateTime:       time_util.TimeToStr(item.CreateTime),     // 申请时间
			MemberUsername:   item.MemberUsername,                      // 会员用户名
			ReturnAmount:     item.ReturnAmount,                        // 退款金额
			ReturnName:       item.ReturnName,                          // 退货人姓名
			ReturnPhone:      item.ReturnPhone,                         // 退货人电话
			Status:           item.Status,                              // 申请状态：0->待处理；1->退货中；2->已完成；3->已拒绝
			HandleTime:       time_util.TimeToString(item.HandleTime),  // 处理时间
			ProductPic:       item.ProductPic,                          // 商品图片
			ProductName:      item.ProductName,                         // 商品名称
			ProductBrand:     item.ProductBrand,                        // 商品品牌
			ProductAttr:      item.ProductAttr,                         // 商品销售属性：颜色：红色；尺码：xl;
			ProductCount:     item.ProductCount,                        // 退货数量
			ProductPrice:     item.ProductPrice,                        // 商品单价
			ProductRealPrice: item.ProductRealPrice,                    // 商品实际支付单价
			Reason:           item.Reason,                              // 原因
			Description:      item.Description,                         // 描述
			ProofPics:        item.ProofPics,                           // 凭证图片，以逗号隔开
			HandleNote:       item.HandleNote,                          // 处理备注
			HandleMan:        item.HandleMan,                           // 处理人员
			ReceiveMan:       item.ReceiveMan,                          // 收货人
			ReceiveTime:      time_util.TimeToString(item.ReceiveTime), // 收货时间
			ReceiveNote:      item.ReceiveNote,                         // 收货备注
		})
	}

	return &omsclient.QueryOrderReturnApplyListResp{
		Total: count,
		List:  list,
	}, nil

}
