package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/pmsmodel"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductVertifyRecordAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductVertifyRecordAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductVertifyRecordAddLogic {
	return &ProductVertifyRecordAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductVertifyRecordAddLogic) ProductVertifyRecordAdd(in *pms.ProductVertifyRecordAddReq) (*pms.ProductVertifyRecordAddResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	_, err := l.svcCtx.PmsProductVertifyRecordModel.Insert(pmsmodel.PmsProductVertifyRecord{
		ProductId:  in.ProductId,
		CreateTime: CreateTime,
		VertifyMan: in.VertifyMan,
		Status:     in.Status,
		Detail:     in.Detail,
	})
	if err != nil {
		return nil, err
	}

	return &pms.ProductVertifyRecordAddResp{}, nil
}
