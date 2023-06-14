package logic

import (
	"context"
	"database/sql"
	"time"
	"zero-admin/rpc/model/omsmodel"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderReturnApplyUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnApplyUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnApplyUpdateLogic {
	return &OrderReturnApplyUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderReturnApplyUpdateLogic) OrderReturnApplyUpdate(in *oms.OrderReturnApplyUpdateReq) (*oms.OrderReturnApplyUpdateResp, error) {
	returnApply, err := l.svcCtx.OmsOrderReturnApplyModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	// 退货中处理退货中的就要设置公司地址
	if in.Status == 1 {
		returnApply.CompanyAddressId = in.CompanyAddressId
	}

	err = l.svcCtx.OmsOrderReturnApplyModel.Update(l.ctx, &omsmodel.OmsOrderReturnApply{
		Id:               in.Id,
		OrderId:          returnApply.OrderId,
		CompanyAddressId: returnApply.CompanyAddressId,
		ProductId:        returnApply.ProductId,
		OrderSn:          returnApply.OrderSn,
		CreateTime:       returnApply.CreateTime,
		MemberUsername:   returnApply.MemberUsername,
		ReturnAmount:     returnApply.ReturnAmount,
		ReturnName:       returnApply.ReturnName,
		ReturnPhone:      returnApply.ReturnPhone,
		Status:           in.Status,
		HandleTime:       time.Now(),
		ProductPic:       returnApply.ProductPic,
		ProductName:      returnApply.ProductName,
		ProductBrand:     returnApply.ProductBrand,
		ProductAttr:      returnApply.ProductAttr,
		ProductCount:     returnApply.ProductCount,
		ProductPrice:     returnApply.ProductPrice,
		ProductRealPrice: returnApply.ProductRealPrice,
		Reason:           returnApply.Reason,
		Description:      returnApply.Description,
		ProofPics:        returnApply.ProofPics,
		HandleNote:       in.HandleNote,
		HandleMan:        in.HandleMan,
		ReceiveMan:       in.ReceiveMan,
		ReceiveTime:      time.Now(),
		ReceiveNote:      sql.NullString{String: in.ReceiveNote, Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return &oms.OrderReturnApplyUpdateResp{}, nil
}
