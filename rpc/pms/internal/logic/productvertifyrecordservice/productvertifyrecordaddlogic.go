package productvertifyrecordservicelogic

import (
	"context"
	"time"
	"zero-admin/rpc/model/pmsmodel"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/rpc/pms/internal/svc"

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

func (l *ProductVertifyRecordAddLogic) ProductVertifyRecordAdd(in *pmsclient.ProductVertifyRecordAddReq) (*pmsclient.ProductVertifyRecordAddResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	_, err := l.svcCtx.PmsProductVertifyRecordModel.Insert(l.ctx, &pmsmodel.PmsProductVertifyRecord{
		ProductId:  in.ProductId,
		CreateTime: CreateTime,
		VertifyMan: in.VertifyMan,
		Status:     in.Status,
		Detail:     in.Detail,
	})
	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductVertifyRecordAddResp{}, nil
}
