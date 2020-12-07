package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductVertifyRecordListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductVertifyRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductVertifyRecordListLogic {
	return &ProductVertifyRecordListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductVertifyRecordListLogic) ProductVertifyRecordList(in *pms.ProductVertifyRecordListReq) (*pms.ProductVertifyRecordListResp, error) {
	all, _ := l.svcCtx.PmsProductVertifyRecordModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	var list []*pms.ProductVertifyRecordListData
	for _, item := range *all {

		list = append(list, &pms.ProductVertifyRecordListData{
			Id:         item.Id,
			ProductId:  item.ProductId,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
			VertifyMan: item.VertifyMan,
			Status:     item.Status,
			Detail:     item.Detail,
		})
	}

	fmt.Println(list)
	return &pms.ProductVertifyRecordListResp{
		Total: 10,
		List:  list,
	}, nil
}
