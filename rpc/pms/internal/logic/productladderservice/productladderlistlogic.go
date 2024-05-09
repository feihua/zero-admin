package productladderservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

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
	result, err := query.PmsProductLadder.WithContext(l.ctx).Where(query.PmsProductLadder.ProductID.Eq(in.ProductId)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询产品阶梯价格列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.ProductLadderListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductLadderListData{
			Id:        item.ID,
			ProductId: item.ProductID,
			Count:     item.Count,
			Discount:  float32(item.Discount),
			Price:     float32(item.Price),
		})
	}

	logc.Infof(l.ctx, "查询产品阶梯价格列表信息,参数：%+v,响应：%+v", in, list)
	return &pmsclient.ProductLadderListResp{
		Total: 0,
		List:  list,
	}, nil
}
