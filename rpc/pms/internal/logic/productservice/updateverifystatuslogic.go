package productservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateVerifyStatusLogic
/*
Author: LiuFeiHua
Date: 2024/5/15 16:12
*/
type UpdateVerifyStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateVerifyStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVerifyStatusLogic {
	return &UpdateVerifyStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateVerifyStatus 批量修改审核状态
func (l *UpdateVerifyStatusLogic) UpdateVerifyStatus(in *pmsclient.UpdateProductStatusReq) (*pmsclient.UpdateProductStatusResp, error) {
	q := query.PmsProduct
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.VerifyStatus, in.Status)

	if err != nil {
		return nil, err
	}

	//修改完审核状态后插入审核记录
	var list []*model.PmsProductVertifyRecord
	for _, id := range in.Ids {
		list = append(list, &model.PmsProductVertifyRecord{
			ProductID:  id,
			CreateTime: time.Now(),
			VertifyMan: in.VertifyMan,
			Status:     in.Status,
			Detail:     in.Detail,
		})
	}
	err = query.PmsProductVertifyRecord.WithContext(l.ctx).CreateInBatches(list, len(list))
	if err != nil {
		return nil, err
	}

	return &pmsclient.UpdateProductStatusResp{}, nil
}
