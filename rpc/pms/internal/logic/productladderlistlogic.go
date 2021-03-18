package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductLadderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLadderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLadderListLogic {
	return &ProductLadderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductLadderListLogic) ProductLadderList(in *pms.ProductLadderListReq) (*pms.ProductLadderListResp, error) {
	all, _ := l.svcCtx.PmsProductLadderModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.PmsProductLadderModel.Count()

	var list []*pms.ProductLadderListData
	for _, item := range *all {

		list = append(list, &pms.ProductLadderListData{
			Id:        item.Id,
			ProductId: item.ProductId,
			Count:     item.Count,
			Discount:  int64(item.Discount),
			Price:     int64(item.Price),
		})
	}

	fmt.Println(list)
	return &pms.ProductLadderListResp{
		Total: count,
		List:  list,
	}, nil
}
