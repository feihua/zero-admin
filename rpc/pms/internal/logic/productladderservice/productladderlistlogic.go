package productladderservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

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

func (l *ProductLadderListLogic) ProductLadderList(in *pmsclient.ProductLadderListReq) (*pmsclient.ProductLadderListResp, error) {
	all, err := l.svcCtx.PmsProductLadderModel.FindAll(l.ctx, in.ProductId)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询产品阶梯价格列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*pmsclient.ProductLadderListData
	for _, item := range *all {

		list = append(list, &pmsclient.ProductLadderListData{
			Id:        item.Id,
			ProductId: item.ProductId,
			Count:     item.Count,
			Discount:  float32(item.Discount),
			Price:     float32(item.Price),
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询产品阶梯价格列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &pmsclient.ProductLadderListResp{
		Total: 0,
		List:  list,
	}, nil
}
