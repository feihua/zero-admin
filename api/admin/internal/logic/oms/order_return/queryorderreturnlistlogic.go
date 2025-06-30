package order_return

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	types "github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOrderReturnListLogic 查询退货/售后主列表
/*
Author: LiuFeiHua
Date: 2025/07/01 10:06:44
*/
type QueryOrderReturnListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryOrderReturnListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderReturnListLogic {
	return &QueryOrderReturnListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryOrderReturnList 查询退货/售后主列表
func (l *QueryOrderReturnListLogic) QueryOrderReturnList(req *types.QueryOrderReturnListReq) (resp *types.QueryOrderReturnListResp, err error) {
	result, err := l.svcCtx.OrderReturnService.QueryOrderReturnList(l.ctx, &omsclient.QueryOrderReturnListReq{
		PageNum:     req.Current,     // 当前页
		PageSize:    req.PageSize,    // 每页条数
		OrderId:     req.OrderId,     // 关联订单ID
		ReturnNo:    req.ReturnNo,    // 退货单号
		MemberId:    req.MemberId,    // 会员ID
		Status:      req.Status,      // 退货状态（0待审核 1审核通过 2已收货 3已退款 4已拒绝 5已关闭）
		Type:        req.Type,        // 售后类型（0退货退款 1仅退款 2换货）
		ReturnName:  req.ReturnName,  // 退货人姓名
		ReturnPhone: req.ReturnPhone, // 退货人电话
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字退货/售后主列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryOrderReturnListData

	for _, detail := range result.List {
		list = append(list, &types.QueryOrderReturnListData{
			Id:             detail.Id,                    // 主键ID
			OrderId:        detail.OrderId,               // 关联订单ID
			ReturnNo:       detail.ReturnNo,              // 退货单号
			MemberId:       detail.MemberId,              // 会员ID
			Status:         detail.Status,                // 退货状态（0待审核 1审核通过 2已收货 3已退款 4已拒绝 5已关闭）
			Type:           detail.Type,                  // 售后类型（0退货退款 1仅退款 2换货）
			Reason:         detail.Reason,                // 退货原因
			Description:    detail.Description,           // 问题描述
			ProofPic:       detail.ProofPic,              // 凭证图片，逗号分隔
			RefundAmount:   float64(detail.RefundAmount), // 退款金额
			ReturnName:     detail.ReturnName,            // 退货人姓名
			ReturnPhone:    detail.ReturnPhone,           // 退货人电话
			CompanyAddress: detail.CompanyAddress,        // 退货收货地址
			CreateTime:     detail.CreateTime,            // 申请时间
			HandleTime:     detail.HandleTime,            // 处理时间
			HandleNote:     detail.HandleNote,            // 处理备注
			HandleMan:      detail.HandleMan,             // 处理人员
			ReceiveTime:    detail.ReceiveTime,           // 收货时间
			ReceiveNote:    detail.ReceiveNote,           // 收货备注
			ReceiveMan:     detail.ReceiveMan,            // 收货人
			RefundTime:     detail.RefundTime,            // 退款时间
			CloseTime:      detail.CloseTime,             // 关闭时间
			Remark:         detail.Remark,                // 备注

		})
	}

	return &types.QueryOrderReturnListResp{
		Code:     "000000",
		Message:  "查询退货/售后主列表成功",
		Data:     list,
		Current:  req.Current,
		PageSize: req.PageSize,
		Total:    result.Total,
		Success:  true,
	}, nil
}
