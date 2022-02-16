package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
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
	all, err := l.svcCtx.PmsProductLadderModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.PmsProductLadderModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询产品阶梯价格列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

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

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询产品阶梯价格列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &pms.ProductLadderListResp{
		Total: count,
		List:  list,
	}, nil
}
