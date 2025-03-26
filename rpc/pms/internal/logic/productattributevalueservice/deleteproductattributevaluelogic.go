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

// DeleteProductAttributeValueLogic 删除存储产品参数信息
/*
Author: LiuFeiHua
Date: 2024/6/12 16:49
*/
type DeleteProductAttributeValueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductAttributeValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductAttributeValueLogic {
	return &DeleteProductAttributeValueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductAttributeValue 删除存储产品参数信息
func (l *DeleteProductAttributeValueLogic) DeleteProductAttributeValue(in *pmsclient.DeleteProductAttributeValueReq) (*pmsclient.DeleteProductAttributeValueResp, error) {
	q := query.PmsProductAttributeValue
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除存储产品参数信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除存储产品参数信息失败")
	}

	return &pmsclient.DeleteProductAttributeValueResp{}, nil
}
