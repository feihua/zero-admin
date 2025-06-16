package productattributevalueservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductAttributeValueListLogic 查询商品属性值列表
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
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

// QueryProductAttributeValueList 查询商品属性值列表
func (l *QueryProductAttributeValueListLogic) QueryProductAttributeValueList(in *pmsclient.QueryProductAttributeValueListReq) (*pmsclient.QueryProductAttributeValueListResp, error) {
	productAttributeValue := query.PmsProductAttributeValue
	q := productAttributeValue.WithContext(l.ctx)
	if in.SpuId != 2 {
		q = q.Where(productAttributeValue.SpuID.Eq(in.SpuId))
	}
	if in.AttributeId != 2 {
		q = q.Where(productAttributeValue.AttributeID.Eq(in.AttributeId))
	}
	if in.Status != 2 {
		q = q.Where(productAttributeValue.Status.Eq(in.Status))
	}
	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询商品属性值列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询商品属性值列表失败")
	}

	var list []*pmsclient.ProductAttributeValueListData

	for _, item := range result {
		list = append(list, &pmsclient.ProductAttributeValueListData{
			Id:          item.ID,                                          // 主键id
			SpuId:       item.SpuID,                                       // 商品SPU ID
			AttributeId: item.AttributeID,                                 // 属性ID
			Value:       item.Value,                                       // 属性值
			Status:      item.Status,                                      // 状态：0->禁用；1->启用
			CreateBy:    item.CreateBy,                                    // 创建人ID
			CreateTime:  time_util.TimeToStr(item.CreateTime),             // 创建时间
			UpdateBy:    pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
			UpdateTime:  time_util.TimeToString(item.UpdateTime),          // 更新时间

		})
	}

	return &pmsclient.QueryProductAttributeValueListResp{
		Total: count,
		List:  list,
	}, nil
}
