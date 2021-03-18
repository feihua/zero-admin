package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberPriceListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberPriceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberPriceListLogic {
	return &MemberPriceListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberPriceListLogic) MemberPriceList(in *pms.MemberPriceListReq) (*pms.MemberPriceListResp, error) {
	all, _ := l.svcCtx.PmsMemberPriceModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.PmsMemberPriceModel.Count()

	var list []*pms.MemberPriceListData
	for _, item := range *all {

		list = append(list, &pms.MemberPriceListData{
			Id:              item.Id,
			ProductId:       item.ProductId,
			MemberLevelId:   item.MemberLevelId,
			MemberPrice:     int64(item.MemberPrice),
			MemberLevelName: item.MemberLevelName,
		})
	}

	fmt.Println(list)
	return &pms.MemberPriceListResp{
		Total: count,
		List:  list,
	}, nil
}
