package productattributevalueservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductAttributeValueListLogic 查询存储产品参数信息列表
/*
Author: LiuFeiHua
Date: 2024/6/12 16:50
*/
type QueryProductAttributeValueListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductAttributeValueListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductAttributeValueListLogic {
	return &QueryProductAttributeValueListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductAttributeValueList 查询存储产品参数信息列表
func (l *QueryProductAttributeValueListLogic) QueryProductAttributeValueList(in *pmsclient.QueryProductAttributeValueListReq) (*pmsclient.QueryProductAttributeValueListResp, error) {
	q := query.PmsProductAttributeValue
	result, err := q.WithContext(l.ctx).Where(q.ProductID.Eq(in.ProductId)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询存储产品参数信息列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询存储产品参数信息列表失败")
	}

	var list []*pmsclient.ProductAttributeValueListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductAttributeValueListData{
			Id:                 item.ID,                 //
			ProductId:          item.ProductID,          // 商品id
			ProductAttributeId: item.ProductAttributeID, // 商品属性id
			Value:              item.Value,              // 手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开
		})
	}

	return &pmsclient.QueryProductAttributeValueListResp{
		Total: 0,
		List:  list,
	}, nil

}
