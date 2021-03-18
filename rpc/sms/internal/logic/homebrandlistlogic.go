package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeBrandListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeBrandListLogic {
	return &HomeBrandListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeBrandListLogic) HomeBrandList(in *sms.HomeBrandListReq) (*sms.HomeBrandListResp, error) {
	all, _ := l.svcCtx.SmsHomeBrandModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.SmsHomeBrandModel.Count()

	var list []*sms.HomeBrandListData
	for _, item := range *all {

		list = append(list, &sms.HomeBrandListData{
			Id:              item.Id,
			BrandId:         item.BrandId,
			BrandName:       item.BrandName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
		})
	}

	fmt.Println(list)
	return &sms.HomeBrandListResp{
		Total: count,
		List:  list,
	}, nil
}
