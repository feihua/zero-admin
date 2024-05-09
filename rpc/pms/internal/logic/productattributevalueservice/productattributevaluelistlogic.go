package productattributevalueservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAttributeValueListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeValueListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeValueListLogic {
	return &ProductAttributeValueListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeValueListLogic) ProductAttributeValueList(in *pmsclient.ProductAttributeValueListReq) (*pmsclient.ProductAttributeValueListResp, error) {
	q := query.PmsProductAttributeValue
	result, err := q.WithContext(l.ctx).Where(q.ProductID.Eq(in.ProductId)).Find()

	if err != nil {
		in, _ := json.Marshal(in)
		logc.Errorf(l.ctx, "查询产品参数列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.ProductAttributeValueListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductAttributeValueListData{
			Id:                 item.ID,
			ProductId:          item.ProductID,
			ProductAttributeId: item.ProductAttributeID,
			Value:              *item.Value,
		})
	}

	logc.Infof(l.ctx, "查询产品参数信息,参数：%+v,响应：%+v", in, list)
	return &pmsclient.ProductAttributeValueListResp{
		Total: 0,
		List:  list,
	}, nil
}
