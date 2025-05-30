package productvertifyrecordservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductVertifyRecordLogic 删除商品审核记录
/*
Author: LiuFeiHua
Date: 2024/6/12 17:13
*/
type DeleteProductVertifyRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductVertifyRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductVertifyRecordLogic {
	return &DeleteProductVertifyRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductVertifyRecord 删除商品审核记录
func (l *DeleteProductVertifyRecordLogic) DeleteProductVertifyRecord(in *pmsclient.DeleteProductVertifyRecordReq) (*pmsclient.DeleteProductVertifyRecordResp, error) {
	q := query.PmsProductVertifyRecord
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除商品审核记录失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除商品审核记录失败")
	}

	return &pmsclient.DeleteProductVertifyRecordResp{}, nil
}
