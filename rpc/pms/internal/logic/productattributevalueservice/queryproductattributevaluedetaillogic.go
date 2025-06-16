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
	"gorm.io/gorm"
)

// QueryProductAttributeValueDetailLogic 查询商品属性值详情
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type QueryProductAttributeValueDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductAttributeValueDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductAttributeValueDetailLogic {
	return &QueryProductAttributeValueDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductAttributeValueDetail 查询商品属性值详情
func (l *QueryProductAttributeValueDetailLogic) QueryProductAttributeValueDetail(in *pmsclient.QueryProductAttributeValueDetailReq) (*pmsclient.QueryProductAttributeValueDetailResp, error) {
	item, err := query.PmsProductAttributeValue.WithContext(l.ctx).Where(query.PmsProductAttributeValue.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "商品属性值不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("商品属性值不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询商品属性值异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询商品属性值异常")
	}

	data := &pmsclient.QueryProductAttributeValueDetailResp{
		Id:          item.ID,                                          // 主键id
		SpuId:       item.SpuID,                                       // 商品SPU ID
		AttributeId: item.AttributeID,                                 // 属性ID
		Value:       item.Value,                                       // 属性值
		Status:      item.Status,                                      // 状态：0->禁用；1->启用
		CreateBy:    item.CreateBy,                                    // 创建人ID
		CreateTime:  time_util.TimeToStr(item.CreateTime),             // 创建时间
		UpdateBy:    pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
		UpdateTime:  time_util.TimeToString(item.UpdateTime),          // 更新时间
	}

	return data, nil
}
