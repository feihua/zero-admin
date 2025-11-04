package orderreturnservicelogic

import (
	"context"
	"errors"

	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOrderReturnListLogic 查询退货/售后列表
/*
Author: LiuFeiHua
Date: 2025/06/30 15:49:28
*/
type QueryOrderReturnListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderReturnListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderReturnListLogic {
	return &QueryOrderReturnListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOrderReturnList 查询退货/售后列表
func (l *QueryOrderReturnListLogic) QueryOrderReturnList(in *omsclient.QueryOrderReturnListReq) (*omsclient.QueryOrderReturnListResp, error) {
	orderReturn := query.OmsOrderReturn
	q := orderReturn.WithContext(l.ctx)
	if in.OrderId != 0 {
		q = q.Where(orderReturn.OrderID.Eq(in.OrderId))
	}
	if len(in.ReturnNo) > 0 {
		q = q.Where(orderReturn.ReturnNo.Like("%" + in.ReturnNo + "%"))
	}
	if in.MemberId != 0 {
		q = q.Where(orderReturn.MemberID.Eq(in.MemberId))
	}
	if in.Status != 2 {
		q = q.Where(orderReturn.Status.Eq(in.Status))
	}
	if in.Type != 0 {
		q = q.Where(orderReturn.Type.Eq(in.Type))
	}

	if len(in.ReturnName) > 0 {
		q = q.Where(orderReturn.ReturnName.Like("%" + in.ReturnName + "%"))
	}
	if len(in.ReturnPhone) > 0 {
		q = q.Where(orderReturn.ReturnPhone.Like("%" + in.ReturnPhone + "%"))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询退货/售后列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询退货/售后列表失败")
	}

	var list []*omsclient.OrderReturnData

	for _, item := range result {
		list = append(list, &omsclient.OrderReturnData{
			Id:             item.ID,                                  // 主键ID
			OrderId:        item.OrderID,                             // 关联订单ID
			ReturnNo:       item.ReturnNo,                            // 退货单号
			MemberId:       item.MemberID,                            // 会员ID
			Status:         item.Status,                              // 退货状态（0待审核 1审核通过 2已收货 3已退款 4已拒绝 5已关闭）
			Type:           item.Type,                                // 售后类型（0退货退款 1仅退款 2换货）
			Reason:         item.Reason,                              // 退货原因
			Description:    item.Description,                         // 问题描述
			ProofPic:       item.ProofPic,                            // 凭证图片，逗号分隔
			RefundAmount:   float32(item.RefundAmount),               // 退款金额
			ReturnName:     item.ReturnName,                          // 退货人姓名
			ReturnPhone:    item.ReturnPhone,                         // 退货人电话
			CompanyAddress: item.CompanyAddress,                      // 退货收货地址
			CreateTime:     time_util.TimeToStr(item.CreateTime),     // 申请时间
			HandleTime:     time_util.TimeToString(item.HandleTime),  // 处理时间
			ReceiveTime:    time_util.TimeToString(item.ReceiveTime), // 收货时间
			RefundTime:     time_util.TimeToString(item.RefundTime),  // 退款时间
			CloseTime:      time_util.TimeToString(item.CloseTime),   // 关闭时间
			Remark:         item.Remark,                              // 备注

		})
	}

	return &omsclient.QueryOrderReturnListResp{
		Total: count,
		List:  list,
	}, nil
}
