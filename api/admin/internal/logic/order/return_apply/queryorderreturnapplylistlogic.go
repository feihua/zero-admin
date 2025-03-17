package return_apply

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOrderReturnApplyListLogic 查询退货申请
/*
Author: LiuFeiHua
Date: 2024/6/15 11:38
*/
type QueryOrderReturnApplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryOrderReturnApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderReturnApplyListLogic {
	return &QueryOrderReturnApplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryOrderReturnApplyList 查询退货申请
func (l *QueryOrderReturnApplyListLogic) QueryOrderReturnApplyList(req *types.QueryOrderReturnApplyListReq) (resp *types.QueryOrderReturnApplyListResp, err error) {
	result, err := l.svcCtx.OrderReturnApplyService.QueryOrderReturnApplyList(l.ctx, &omsclient.QueryOrderReturnApplyListReq{
		PageNum:        req.Current,
		PageSize:       req.PageSize,
		OrderSn:        strings.TrimSpace(req.OrderSn),
		MemberUsername: strings.TrimSpace(req.MemberUsername),
		HandleTime:     req.HandleTime,
		CreateTime:     req.CreateTime,
		Status:         req.Status,
		HandleMan:      strings.TrimSpace(req.HandleMan),
		ReturnName:     strings.TrimSpace(req.ReturnName),
		ReturnPhone:    strings.TrimSpace(req.ReturnPhone),
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询退货申请列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询退货申请列表失败")
	}

	var list []*types.QueryOrderReturnApplyListData

	for _, item := range result.List {
		list = append(list, &types.QueryOrderReturnApplyListData{
			Id:               item.Id,               //
			OrderId:          item.OrderId,          // 订单id
			CompanyAddressId: item.CompanyAddressId, // 收货地址表id
			ProductId:        item.ProductId,        // 退货商品id
			OrderSn:          item.OrderSn,          // 订单编号
			CreateTime:       item.CreateTime,       // 申请时间
			MemberUsername:   item.MemberUsername,   // 会员用户名
			ReturnAmount:     item.ReturnAmount,     // 退款金额
			ReturnName:       item.ReturnName,       // 退货人姓名
			ReturnPhone:      item.ReturnPhone,      // 退货人电话
			Status:           item.Status,           // 申请状态：0->待处理；1->退货中；2->已完成；3->已拒绝
			HandleTime:       item.HandleTime,       // 处理时间
			ProductPic:       item.ProductPic,       // 商品图片
			ProductName:      item.ProductName,      // 商品名称
			ProductBrand:     item.ProductBrand,     // 商品品牌
			ProductAttr:      item.ProductAttr,      // 商品销售属性：颜色：红色；尺码：xl;
			ProductCount:     item.ProductCount,     // 退货数量
			ProductPrice:     item.ProductPrice,     // 商品单价
			ProductRealPrice: item.ProductRealPrice, // 商品实际支付单价
			Reason:           item.Reason,           // 原因
			Description:      item.Description,      // 描述
			ProofPics:        item.ProofPics,        // 凭证图片，以逗号隔开
			HandleNote:       item.HandleNote,       // 处理备注
			HandleMan:        item.HandleMan,        // 处理人员
			ReceiveMan:       item.ReceiveMan,       // 收货人
			ReceiveTime:      item.ReceiveTime,      // 收货时间
			ReceiveNote:      item.ReceiveNote,      // 收货备注
		})
	}

	return &types.QueryOrderReturnApplyListResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询退货申请列表成功",
	}, nil
}
