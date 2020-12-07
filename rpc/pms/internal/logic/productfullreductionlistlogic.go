package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductFullReductionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductFullReductionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductFullReductionListLogic {
	return &ProductFullReductionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductFullReductionListLogic) ProductFullReductionList(in *pms.ProductFullReductionListReq) (*pms.ProductFullReductionListResp, error) {
	all, _ := l.svcCtx.PmsProductFullReductionModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	var list []*pms.ProductFullReductionListData
	for _, item := range *all {

		list = append(list, &pms.ProductFullReductionListData{
			Id:          item.Id,
			ProductId:   item.ProductId,
			FullPrice:   int64(item.FullPrice),
			ReducePrice: int64(item.ReducePrice),
		})
	}

	fmt.Println(list)
	return &pms.ProductFullReductionListResp{
		Total: 10,
		List:  list,
	}, nil
}
