package logic

import (
	"context"
	"go-zero-admin/rpc/model/pmsmodel"
	"time"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
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
