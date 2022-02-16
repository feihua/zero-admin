package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/pmsmodel"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductVertifyRecordUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductVertifyRecordUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductVertifyRecordUpdateLogic {
	return &ProductVertifyRecordUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductVertifyRecordUpdateLogic) ProductVertifyRecordUpdate(in *pms.ProductVertifyRecordUpdateReq) (*pms.ProductVertifyRecordUpdateResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	err := l.svcCtx.PmsProductVertifyRecordModel.Update(pmsmodel.PmsProductVertifyRecord{
		Id:         in.Id,
		ProductId:  in.ProductId,
		CreateTime: CreateTime,
		VertifyMan: in.VertifyMan,
		Status:     in.Status,
		Detail:     in.Detail,
	})
	if err != nil {
		return nil, err
	}

	return &pms.ProductVertifyRecordUpdateResp{}, nil
}
